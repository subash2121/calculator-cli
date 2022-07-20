package handler

import (
	"calculator-cli/pkg/calculator"
	"calculator-cli/pkg/constants"
	"calculator-cli/pkg/parser"
	"fmt"
)

type operationsMap map[constants.Operations]handleFunc

type handleFunc func(float64, *calculator.Calculator) float64

func (operations operationsMap) AddHandler(operator constants.Operations, function handleFunc) {
	if _, exists := operations[operator]; exists {
		panic("command already exists")
	}
	operations[operator] = function
}

func (operations operationsMap) ExecuteQuery(expression parser.Expression, calculator *calculator.Calculator) float64 {
	if evaluator, exists := operations[expression.Operator]; exists {
		return evaluator(expression.Operand, calculator)
	}
	return 0.0
}

func RegisterHandler(operations *operationsMap) {
	operations.AddHandler("add", HandleAddFunction)
	operations.AddHandler("subtract", HandleSubtractFunction)
	operations.AddHandler("multiply", HandleMultiplyFunction)
	operations.AddHandler("divide", HandleDivideFunction)
	operations.AddHandler("exit", HandleExitFunction)
	operations.AddHandler("cancel", HandleCancelFunction)
}

func Handler(expression parser.Expression, calculator *calculator.Calculator) {
	operations := operationsMap{}
	RegisterHandler(&operations)
	fmt.Println(operations.ExecuteQuery(expression, calculator))
}

func HandleAddFunction(operand float64, calculator *calculator.Calculator) float64 {
	return calculator.Add(operand)
}

func HandleSubtractFunction(operand float64, calculator *calculator.Calculator) float64 {
	return calculator.Subtract(operand)
}

func HandleMultiplyFunction(operand float64, calculator *calculator.Calculator) float64 {
	return calculator.Multiply(operand)
}

func HandleDivideFunction(operand float64, calculator *calculator.Calculator) float64 {
	return calculator.Divide(operand)
}

func HandleExitFunction(operand float64, calculator *calculator.Calculator) float64 {
	return calculator.Exit(operand)
}

func HandleCancelFunction(operand float64, calculator *calculator.Calculator) float64 {
	return calculator.Cancel(operand)
}
