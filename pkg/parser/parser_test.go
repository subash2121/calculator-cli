package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("should return a entity with operand and operator", func(t *testing.T) {
		assert.IsType(t, UserInput{}, Parse())
	})

}
