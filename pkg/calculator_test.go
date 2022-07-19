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

func TestGetValue(t *testing.T) {
	t.Run("should result the value of the calculator", func(t *testing.T) {
		assert.IsType(t, 0.0, NewCalculator().GetValue())
	})

	t.Run("should result 0.0 for a new calculator", func(t *testing.T) {
		assert.Equal(t, 0.0, NewCalculator().GetValue())
	})

	t.Run("should return 5.4 as the value of calculator after adding 5.4", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(5.4)
		assert.Equal(t, 5.4, calculator.GetValue())
	})
}

func TestAdd(t *testing.T) {
	t.Run("should result as 4.0 when 4 is added with 0", func(t *testing.T) {
		calculator := NewCalculator()
		assert.Equal(t, 4.0, calculator.Add(4))
	})

	t.Run("should result as 8.0 when 4 is added with 4", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(4)
		assert.Equal(t, 8.0, calculator.Add(4))
	})

	t.Run("should result as 0.0 when -4 is added with 4", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(4)
		assert.Equal(t, 0.0, calculator.Add(-4))
	})

	t.Run("should result as -8.0 when -4 is added with -4", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(-4)
		assert.Equal(t, -8.0, calculator.Add(-4))
	})
}

func TestSubtract(t *testing.T) {
	t.Run("should result as -4.0  when 4 is subtracted from 0", func(t *testing.T) {
		calculator := NewCalculator()
		assert.Equal(t, -4.0, calculator.Subtract(4))
	})

	t.Run("should result as 0.0 when 5 is subtracted from 5", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(5)
		assert.Equal(t, 0.0, calculator.Subtract(5))
	})

	t.Run("should result as 10.0 when -5 is subtracted from 5", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(5)
		assert.Equal(t, 10.0, calculator.Subtract(-5))
	})

	t.Run("should result as 0.0 when -5 is subtracted from -5", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(-5)
		assert.Equal(t, 0.0, calculator.Subtract(-5))
	})

	t.Run("should result as -10.0 when 5 is subtracted from -5", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(-5)
		assert.Equal(t, -10.0, calculator.Subtract(5))
	})
}

func TestMultiply(t *testing.T) {
	t.Run("should result as 0.0  when 4 is multiplied with 0", func(t *testing.T) {
		calculator := NewCalculator()
		assert.Equal(t, 0.0, calculator.Multiply(4))
	})

	t.Run("should result as 4.0  when 4 is multiplied with 1", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(1)
		assert.Equal(t, 4.0, calculator.Multiply(4))
	})

	t.Run("should result as -2.0  when 2 is multiplied with -1", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(-1)
		assert.Equal(t, -2.0, calculator.Multiply(2))
	})

	t.Run("should result as 10.0  when 4 is multiplied with 2.5", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(2.5)
		assert.Equal(t, 10.0, calculator.Multiply(4))
	})

	t.Run("should result as -20.0  when 8 is multiplied with -2.5", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(-2.5)
		assert.Equal(t, -20.0, calculator.Multiply(8))
	})
}
