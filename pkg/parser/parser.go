package parser

import (
	"strconv"
	"strings"
)

type Expression struct {
	operator string
	operand  float64
}

func Parse(inputString string) Expression {
	inputValue := Expression{}
	inputValue.operator = strings.Split(inputString, " ")[0]
	inputValue.operand = ConvertToFloat(strings.Split(inputString, " ")[1])
	return inputValue
}

func ConvertToFloat(operandString string) float64 {
	operand, err := strconv.ParseFloat(operandString, 64)
	if err != nil {
		panic("Error: Wrong Operand")
	}
	return operand
}
