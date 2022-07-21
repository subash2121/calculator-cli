package handler

import (
	"calculator-cli/pkg/calculator"
	"calculator-cli/pkg/constants"
	"calculator-cli/pkg/parser"
	"calculator-cli/pkg/renderer"
	"os"
)

type operationsMap map[constants.Operations]handleFunc

type handleFunc func(float64, *calculator.Calculator)

var operations = operationsMap{}

func (operations operationsMap) AddHandler(operator constants.Operations, function handleFunc) {
	if _, exists := operations[operator]; exists {
		panic("command already exists")
	}
	operations[operator] = function
}

func (operations operationsMap) ExecuteQuery(expression parser.Expression, calculator *calculator.Calculator) {
	if evaluator, exists := operations[expression.Operator]; exists {
		evaluator(expression.Operand, calculator)
	} else {
		renderer.RenderError(os.Stdout, "Operation not supported")
	}
}

func RegisterHandler(operations *operationsMap) {
	operations.AddHandler("add", HandleAddFunction)
	operations.AddHandler("subtract", HandleSubtractFunction)
	operations.AddHandler("multiply", HandleMultiplyFunction)
	operations.AddHandler("divide", HandleDivideFunction)
	operations.AddHandler("exit", HandleExitFunction)
	operations.AddHandler("cancel", HandleCancelFunction)
}

func init() {
	RegisterHandler(&operations)
}

func Handler(expression parser.Expression, calculator *calculator.Calculator) {
	operations.ExecuteQuery(expression, calculator)
}

func HandleAddFunction(operand float64, calculator *calculator.Calculator) {
	renderer.RenderValue(os.Stdout, calculator.Add(operand))
}

func HandleSubtractFunction(operand float64, calculator *calculator.Calculator) {
	renderer.RenderValue(os.Stdout, calculator.Subtract(operand))
}

func HandleMultiplyFunction(operand float64, calculator *calculator.Calculator) {
	renderer.RenderValue(os.Stdout, calculator.Multiply(operand))
}

func HandleDivideFunction(operand float64, calculator *calculator.Calculator) {
	renderer.RenderValue(os.Stdout, calculator.Divide(operand))
}

func HandleExitFunction(operand float64, calculator *calculator.Calculator) {
	calculator.Exit(operand)
}

func HandleCancelFunction(operand float64, calculator *calculator.Calculator) {
	renderer.RenderValue(os.Stdout, calculator.Cancel(operand))
}
