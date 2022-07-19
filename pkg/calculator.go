package main

type Calculator struct {
	value float64
}

func NewCalculator() Calculator {
	return Calculator{0}
}

func (calculator *Calculator) Add(operand float64) float64 {
	calculator.value += operand
	return calculator.value
}

func (calculator Calculator) GetValue() float64 {
	return calculator.value
}
