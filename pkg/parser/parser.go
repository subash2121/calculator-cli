package parser

import (
	"calculator-cli/pkg/constants"
	"strconv"
	"strings"
)

type Expression struct {
	Operator constants.Operations
	Operand  float64
}

func Parse(inputString string) Expression {
	inputValue := Expression{}
	stringList := strings.Split(inputString, " ")
	inputValue.Operator = constants.Operations(stringList[0])
	if len(stringList) < 2 {
		return inputValue
	}
	inputValue.Operand = ConvertToFloat(strings.Split(inputString, " ")[1])
	return inputValue
}

func ConvertToFloat(operandString string) float64 {
	operand, err := strconv.ParseFloat(operandString, 64)
	if err != nil {
		panic("Error: Wrong Operand")
	}
	return operand
}
