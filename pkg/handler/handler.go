package handler

import (
	"calculator-cli/pkg/calculator"
	"calculator-cli/pkg/constants"
	"calculator-cli/pkg/parser"
	"errors"
	"fmt"
)

type operationsMap map[constants.Operations]handleFunc

type handleFunc func(float64) float64

func (operations operationsMap) AddHandler(operator constants.Operations, function handleFunc) error {
	if _, exists := operations[operator]; exists {
		return errors.New("command already exists")
	}
	operations[operator] = function
	return nil
}

func (operations operationsMap) ExecuteQuery(expression parser.Expression) float64 {
	if com, exists := operations[expression.Operator]; exists {
		return com(expression.Operand)
	}
	return 0.0
}

func Handler(expression parser.Expression, calculator *calculator.Calculator) {
	operations := operationsMap{}
	operations.AddHandler("add", calculator.Add)
	operations.AddHandler("subtract", calculator.Subtract)
	operations.AddHandler("multiply", calculator.Multiply)
	operations.AddHandler("divide", calculator.Divide)
	operations.AddHandler("exit", calculator.Exit)
	operations.AddHandler("cancel", calculator.Cancel)
	fmt.Println(operations.ExecuteQuery(expression))
}
