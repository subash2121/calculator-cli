package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("should return an entity with operand and operator", func(t *testing.T) {
		assert.IsType(t, Expression{}, Parse("add 5"))
	})

	t.Run("should parse the read value as a operand and operator", func(t *testing.T) {
		parser := Parse("add 5")
		assert.IsType(t, "", parser.Operator)
		assert.IsType(t, 5.0, parser.Operand)
	})

	t.Run("should parse the 'subtract 5' as operator-subtract and operand-5.0", func(t *testing.T) {
		parser := Parse("subtract 5")
		assert.Equal(t, "subtract", parser.Operator)
		assert.Equal(t, 5.0, parser.Operand)
	})

	t.Run("should panic if wrong operand is provided except a number", func(t *testing.T) {
		assert.Panics(t, func() {
			Parse("subtract 5.a")
		})
	})
}

func TestConvertToFloat(t *testing.T) {
	t.Run("should convert the string 10 to 10.0", func(t *testing.T) {
		assert.Equal(t, 10.0, ConvertToFloat("10"))
	})

	t.Run("should convert the string -10 to -10.0", func(t *testing.T) {
		assert.Equal(t, -10.0, ConvertToFloat("-10"))
	})

	t.Run("should panic if string cannot be converted to a float", func(t *testing.T) {
		assert.Panics(t, func() {
			Parse("abc")
		})
	})
}
