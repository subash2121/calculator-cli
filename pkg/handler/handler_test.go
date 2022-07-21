package handler

import (
	"calculator-cli/pkg/calculator"
	"calculator-cli/pkg/constants"
	"calculator-cli/pkg/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddHandler(t *testing.T) {
	operations := operationsMap{}

	t.Run("should add a new handler to the operations", func(t *testing.T) {
		operations.AddHandler(constants.Operations("add"), func(operand float64, calculator *calculator.Calculator) float64 {
			return calculator.Add(operand)
		})
		assert.Contains(t, operations, constants.Operations("add"))
	})

	t.Run("should panic if same handler is added", func(t *testing.T) {
		operations := operationsMap{}
		operations.AddHandler(constants.Operations("add"), func(operand float64, calculator *calculator.Calculator) float64 {
			return calculator.Add(operand)
		})
		assert.Panics(t, func() {
			operations.AddHandler(constants.Operations("add"), func(operand float64, calculator *calculator.Calculator) float64 {
				return calculator.Add(operand)
			})
		})
	})
}

func TestExecuteQuery(t *testing.T) {
	operations := operationsMap{}
	operations.AddHandler(constants.Operations("add"), func(operand float64, calculator *calculator.Calculator) float64 {
		return calculator.Add(operand)
	})

	t.Run("should execute a given handler and add 5 to get 5", func(t *testing.T) {
		value := operations.ExecuteQuery(parser.Expression{Operator: "add", Operand: 5}, calculator.NewCalculator())
		assert.Equal(t, 5.0, value)
	})
}

func TestHandleAddFunction(t *testing.T) {
	t.Run("Should add 5 to 0 and result as 5.0", func(t *testing.T) {
		value := HandleAddFunction(5, calculator.NewCalculator())
		assert.Equal(t, 5.0, value)
	})
}

func TestHandleCancelFunction(t *testing.T) {
	calc := calculator.NewCalculator()

	t.Run("Should reset the calculator value to 0", func(t *testing.T) {
		HandleAddFunction(5, calc)
		value := HandleCancelFunction(0, calc)
		assert.Equal(t, 0.0, value)
	})
}

func TestHandleDivideFunction(t *testing.T) {
	calc := calculator.NewCalculator()

	t.Run("Should divide 10 by 2 and result 5.0", func(t *testing.T) {
		HandleAddFunction(10, calc)
		value := HandleDivideFunction(2, calc)
		assert.Equal(t, 5.0, value)
	})

	t.Run("Should panic if divided by 0", func(t *testing.T) {
		assert.Panics(t, func() {
			HandleDivideFunction(0, calc)
		})
	})
}

//func TestHandleExitFunction(t *testing.T) {
//	t.Run("should exit from the program", func(t *testing.T) {
//		assert.Panics(t, func() {
//			Cancel
//		})
//		assert.Equal(t, 1.0, HandleExitFunction(0.0, calculator.NewCalculator()))
//	})
//}

func TestHandleMultiplyFunction(t *testing.T) {
	t.Run("Should add 2 to 10 and result as 20", func(t *testing.T) {
		calc := calculator.NewCalculator()
		HandleAddFunction(10, calc)
		value := HandleMultiplyFunction(2, calc)
		assert.Equal(t, 20.0, value)
	})
}

func TestHandleSubtractFunction(t *testing.T) {
	t.Run("Should subtract 5 to 0 and result as -5", func(t *testing.T) {
		value := HandleSubtractFunction(5, calculator.NewCalculator())
		assert.Equal(t, -5.0, value)
	})
}
