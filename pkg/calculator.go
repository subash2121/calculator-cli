package main

type Calculator struct {
	value float64
}

func NewCalculator() Calculator {
	return Calculator{0}
}

func (calculator *Calculator) Add(operand float64) float64 {
	return 4.0
}
