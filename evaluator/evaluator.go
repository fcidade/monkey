package evaluator

import (
	"fmt"

	"github.com/fcidade/monkey-lang/ast"
	"github.com/fcidade/monkey-lang/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node, env *object.Environment) object.Object {

	switch v := node.(type) {
	case *ast.ReturnStatement:
		value := Eval(v.ReturnValue, env)
		if isError(value) {
			return value
		}
		return &object.ReturnValue{Value: value}

	case *ast.BlockStatement:
		return evalBlockStatement(v, env)

	case *ast.IfExpression:
		return evalIfExpression(v, env)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: v.Value}

	case *ast.Program:
		return evalProgram(v, env)

	case *ast.Boolean:
		return boolean(v.Value)

	case *ast.ExpressionStatement:
		return Eval(v.Expression, env)

	case *ast.PrefixExpression:
		right := Eval(v.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(v.Operator, right)

	case *ast.InfixExpression:
		left := Eval(v.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(v.Right, env)
		if isError(right) {
			return right
		}
		return evalInfixExpression(left, v.Operator, right)

	case *ast.LetStatement:
		val := Eval(v.Value, env)
		if isError(val) {
			return val
		}
		env.Set(v.Name.Value, val)

	case *ast.Identifier:
		return evalIdentifier(env, v)

	case *ast.FunctionLiteral:
		return &object.Function{
			Parameters: v.Parameters,
			Body:       *v.Body,
			Env:        env,
		}

	case *ast.CallExpression:
		function := Eval(v.Function, env)
		if isError(function) {
			return function
		}

		args := evalExpressions(v.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		return applyFunction(function, args)
	}

	return nil
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	function, ok := fn.(*object.Function)
	if !ok {
		return newError("not a function: %s", fn.Type())
	}

	extendedEnv := extendedFunctionEnv(function, args)
	evaluated := Eval(&function.Body, extendedEnv)

	return unwrapReturnValue(evaluated)
}

func extendedFunctionEnv(function *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(function.Env)

	for paramIdx, param := range function.Parameters {
		env.Set(param.Value, args[paramIdx])
	}

	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

func evalExpressions(
	exps []ast.Expression,
	env *object.Environment,
) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func evalIdentifier(env *object.Environment, node *ast.Identifier) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}
	return newError("identifier not found: %s", node.Value)
}

func evalIfExpression(v *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(v.Condition, env)
	if isError(condition) {
		return condition
	}
	if isTruthy(condition) {
		return Eval(v.Consequence, env)
	}
	if v.Alternative != nil {
		return Eval(v.Alternative, env)
	}
	return NULL
}

func isTruthy(v object.Object) bool {
	return v != FALSE && v != NULL
}

func evalInfixExpression(left object.Object, operator string, right object.Object) object.Object {
	switch {
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(left, operator, right)
	case operator == "==":
		return boolean(left == right)
	case operator == "!=":
		return boolean(left != right)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalIntegerInfixExpression(left object.Object, operator string, right object.Object) object.Object {
	leftVal := left.(*object.Integer)
	rightVal := right.(*object.Integer)

	switch operator {
	case "+":
		return integer(leftVal.Value + rightVal.Value)
	case "-":
		return integer(leftVal.Value - rightVal.Value)
	case "*":
		return integer(leftVal.Value * rightVal.Value)
	case "/":
		return integer(leftVal.Value / rightVal.Value)
	case ">":
		return boolean(leftVal.Value > rightVal.Value)
	case "<":
		return boolean(leftVal.Value < rightVal.Value)
	case "==":
		return boolean(leftVal.Value == rightVal.Value)
	case "!=":
		return boolean(leftVal.Value != rightVal.Value)
	}
	return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
}

func integer(number int64) *object.Integer {
	return &object.Integer{Value: number}
}

func boolean(boolean bool) *object.Boolean {
	if boolean {
		return TRUE
	}
	return FALSE
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperator(right)
	case "-":
		return evalMinusOperator(right)
	default:
		return newError("unkown operator: %s%s", operator, right.Type())
	}
}

func evalMinusOperator(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return newError("unknown operator: -%s", right.Type())
	}
	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func evalBangOperator(value object.Object) object.Object {
	switch value {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var last object.Object
	for _, stmt := range program.Statements {
		last = Eval(stmt, env)
		switch last := last.(type) {
		case *object.ReturnValue:
			return last.Value
		case *object.Error:
			return last
		}
	}
	return last
}

func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var last object.Object
	for _, stmt := range block.Statements {
		last = Eval(stmt, env)
		if last.Type() == object.ERROR_OBJ || last.Type() == object.RETURN_VALUE_OBJ {
			return last
		}
	}
	return last
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}
