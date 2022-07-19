package main

type Calculator struct {
	value float64
}

func NewCalculator() Calculator {
	return Calculator{0}
}

func (calculator Calculator) GetValue() float64 {
	return calculator.value
}

func (calculator *Calculator) Add(operand float64) float64 {
	calculator.value += operand
	return calculator.GetValue()
}

func (calculator *Calculator) Subtract(operand float64) float64 {
	return calculator.Add(-operand)
}

func (calculator *Calculator) Multiply(operand float64) float64 {
	calculator.value *= operand
	return calculator.GetValue()
}
