package renderer

import (
	"fmt"
	"io"
)

func RenderInput(writer io.Writer) {
	fmt.Fprintf(writer, "> ")
}

func RenderValue(writer io.Writer, value float64) {
	fmt.Fprintf(writer, "%.2f", value)
}

func RenderError(writer io.Writer, error string) {
	fmt.Fprintf(writer, "%s", error)
}

func RenderNewLine(writer io.Writer) {
	fmt.Fprintf(writer, "\n")
}
