package calculator

import "os"

type Calculator struct {
	value float64
}

type ArithmeticInterface interface {
	Add(value float64) float64
	Subtract(value float64) float64
	Multiply(value float64) float64
	Divide(value float64) float64
}

type CalculatorInterface interface {
	ArithmeticInterface
	UtilityInterface
}

type UtilityInterface interface {
	Cancel(float64) float64
	Exit(float64) float64
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (calculator *Calculator) Add(operand float64) float64 {
	calculator.value += operand
	return calculator.value
}

func (calculator *Calculator) Subtract(operand float64) float64 {
	return calculator.Add(-operand)
}

func (calculator *Calculator) Multiply(operand float64) float64 {
	calculator.value *= operand
	return calculator.value
}

func (calculator *Calculator) Divide(operand float64) float64 {
	if operand == 0 {
		panic("Zero division error")
	}
	calculator.value /= operand
	return calculator.value
}

func (calculator *Calculator) Cancel(float64) float64 {
	calculator.value = 0
	return calculator.value
}

func (calculator *Calculator) Exit(float64) float64 {
	defer func() { os.Exit(0) }()
	return 0
}
