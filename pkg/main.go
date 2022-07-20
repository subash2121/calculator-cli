package main

import (
	"calculator-cli/pkg/calculator"
	"calculator-cli/pkg/handler"
	"calculator-cli/pkg/parser"
	"calculator-cli/pkg/reader"
	"fmt"
	"os"
)

func main() {
	newCalculator := calculator.NewCalculator()
	for true {
		fmt.Print("> ")
		Input := reader.Reader(os.Stdin)
		ParsedObject := parser.Parse(Input)
		handler.Handler(ParsedObject, newCalculator)
	}
}
