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
	t.Run("should add 5 to 0 and result as 5.0", func(t *testing.T) {
		value := HandleAddFunction(5, calculator.NewCalculator())
		assert.Equal(t, 5.0, value)
	})
}

func TestHandleCancelFunction(t *testing.T) {
	calc := calculator.NewCalculator()

	t.Run("should reset the calculator value to 0", func(t *testing.T) {
		HandleAddFunction(5, calc)
		value := HandleCancelFunction(0, calc)
		assert.Equal(t, 0.0, value)
	})
}

func TestHandleDivideFunction(t *testing.T) {
	calc := calculator.NewCalculator()

	t.Run("should divide 10 by 2 and result 5.0", func(t *testing.T) {
		HandleAddFunction(10, calc)
		value := HandleDivideFunction(2, calc)
		assert.Equal(t, 5.0, value)
	})
}

func TestHandleMultiplyFunction(t *testing.T) {
	t.Run("should multiply 2 to 10 and result as 20", func(t *testing.T) {
		calc := calculator.NewCalculator()
		HandleAddFunction(10, calc)
		value := HandleMultiplyFunction(2, calc)
		assert.Equal(t, 20.0, value)
	})
}

func TestHandleSubtractFunction(t *testing.T) {
	t.Run("should subtract 5 to 0 and result as -5", func(t *testing.T) {
		value := HandleSubtractFunction(5, calculator.NewCalculator())
		assert.Equal(t, -5.0, value)
	})
}

func TestHandler(t *testing.T) {
	t.Run("should compute the addition of 5 with 0 and result 5", func(t *testing.T) {
		value := Handler(parser.Expression{Operator: constants.Operations("add"), Operand: 5.0}, calculator.NewCalculator())
		assert.Equal(t, 5.0, value)
	})
}
