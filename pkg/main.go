package main

import (
	"calculator-cli/pkg/calculator"
	"calculator-cli/pkg/handler"
	"calculator-cli/pkg/parser"
	"calculator-cli/pkg/reader"
	"calculator-cli/pkg/renderer"
	"os"
)

func main() {
	newCalculator := calculator.NewCalculator()
	for true {
		renderer.RenderInput(os.Stdout)
		Input := reader.Reader(os.Stdin)
		ParsedObject := parser.Parse(Input)
		handler.Handler(ParsedObject, newCalculator)
		renderer.RenderNewLine(os.Stdout)
	}
}
