package reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReader(t *testing.T) {
	t.Run("should input a string from the IO interface", func(t *testing.T) {
		assert.IsType(t, "", Reader())
	})
}
