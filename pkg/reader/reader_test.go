package reader

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	t.Run("should input a string from the IO interface", func(t *testing.T) {
		input := "add 5"
		reader := strings.NewReader(input)
		assert.Equal(t, "add 5", Reader(reader))
	})
}
