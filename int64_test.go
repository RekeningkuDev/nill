package nill

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt64(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		value := NewInt64(1)
		value.Add(1)
		assert.Equal(t, Int64{Valid: true, Value: 2}, value)
	})
	t.Run("sub", func(t *testing.T) {
		value := NewInt64(1)
		value.Sub(1)
		assert.Equal(t, Int64{Valid: true, Value: 0}, value)
	})
	t.Run("mul", func(t *testing.T) {
		value := NewInt64(3)
		value.Mul(5)
		assert.Equal(t, Int64{Valid: true, Value: 15}, value)
	})
	t.Run("div", func(t *testing.T) {
		value := NewInt64(10)
		value.Div(5)
		assert.Equal(t, Int64{Valid: true, Value: 2}, value)
	})
	t.Run("mod", func(t *testing.T) {
		value := NewInt64(1)
		value.Mod(5)
		assert.Equal(t, Int64{Valid: true, Value: 1}, value)
	})
	t.Run("inc", func(t *testing.T) {
		value := NewInt64(1)
		value.Inc()
		assert.Equal(t, Int64{Valid: true, Value: 2}, value)
	})
	t.Run("dec", func(t *testing.T) {
		value := NewInt64(1)
		value.Dec()
		assert.Equal(t, Int64{Valid: true, Value: 0}, value)
	})
	t.Run("zero", func(t *testing.T) {
		value := NewInt64(0)
		assert.True(t, value.Zero())
	})
	t.Run("positive", func(t *testing.T) {
		value := NewInt64(1)
		assert.True(t, value.Positive())
	})
	t.Run("negative", func(t *testing.T) {
		value := NewInt64(-1)
		assert.True(t, value.Negative())
	})
	t.Run("event", func(t *testing.T) {
		value := NewInt64(10)
		value.Even()
		assert.True(t, value.Even())
	})
	t.Run("odd", func(t *testing.T) {
		value := NewInt64(11)
		assert.True(t, value.Odd())
	})
	t.Run("greater than", func(t *testing.T) {
		value := NewInt64(11)
		assert.True(t, value.GreaterThan(10))
	})
	t.Run("greater than equal", func(t *testing.T) {
		value := NewInt64(10)
		assert.True(t, value.LessThanOrEqual(10))
	})
	t.Run("less than", func(t *testing.T) {
		value := NewInt64(10)
		assert.True(t, value.LessThan(11))
	})
	t.Run("less than equal", func(t *testing.T) {
		value := NewInt64(10)
		assert.True(t, value.LessThanOrEqual(10))
	})
	t.Run("equal", func(t *testing.T) {
		value := NewInt64(10)
		assert.True(t, value.Equal(10))
	})
	t.Run("not equal", func(t *testing.T) {
		value := NewInt64(10)
		assert.True(t, value.NotEqual(11))
	})
	t.Run("divisible by", func(t *testing.T) {
		value := NewInt64(10)
		assert.True(t, value.DivisibleBy(5))
	})
	t.Run("prime", func(t *testing.T) {
		value := NewInt64(19)
		assert.True(t, value.Prime())
	})
}
