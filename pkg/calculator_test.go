package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCalculator(t *testing.T) {
	t.Run("should create a new calculator and return it", func(t *testing.T) {
		assert.IsType(t, Calculator{}, NewCalculator())
	})

	t.Run("should initialize calculator with inital zero value", func(t *testing.T) {
		assert.Equal(t, 0.0, NewCalculator().value)
	})
}

func TestAdd(t *testing.T) {
	t.Run("should result as 4 when add 4 with 0", func(t *testing.T) {
		calculator := NewCalculator()
		assert.Equal(t, 4.0, calculator.Add(4))
	})
}
