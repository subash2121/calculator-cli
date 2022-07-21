package renderer

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

func TestRenderInput(t *testing.T) {
	defer setupStdout()()
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = fakeStdout
	RenderInput(os.Stdout)
	fakeStdout.Close()
	bytes, err := io.ReadAll(r)
	r.Close()
	t.Run("should print > ", func(t *testing.T) {
		assert.Equal(t, "> ", string(bytes))
	})
}

func TestRenderValue(t *testing.T) {
	defer setupStdout()()
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = fakeStdout
	RenderValue(os.Stdout, 2)
	fakeStdout.Close()
	bytes, err := io.ReadAll(r)
	r.Close()
	t.Run("should print the value 2", func(t *testing.T) {
		assert.Equal(t, "2.00", string(bytes))
	})
}

func TestRenderError(t *testing.T) {
	defer setupStdout()()
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = fakeStdout
	RenderError(os.Stdout, "error")
	fakeStdout.Close()
	bytes, err := io.ReadAll(r)
	r.Close()
	t.Run("should print given error", func(t *testing.T) {
		assert.Equal(t, "error", string(bytes))
	})
}

func TestRenderNewLine(t *testing.T) {
	defer setupStdout()()
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = fakeStdout
	RenderNewLine(os.Stdout)
	fakeStdout.Close()
	bytes, err := io.ReadAll(r)
	r.Close()
	t.Run("should print newline", func(t *testing.T) {
		assert.Equal(t, "\n", string(bytes))
	})
}

func setupStdout() func() {
	originalStdout := os.Stdout
	return func() {
		os.Stdout = originalStdout
	}
}
