package nill

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat64(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		value := NewFloat64(1.5)
		value.Add(1)
		assert.Equal(t, Float64{Valid: true, V: 2.5}, value)
	})
	t.Run("sub", func(t *testing.T) {
		value := NewFloat64(1.5)
		value.Sub(1)
		assert.Equal(t, Float64{Valid: true, V: 0.5}, value)
	})
	t.Run("mul", func(t *testing.T) {
		value := NewFloat64(1.5)
		value.Mul(5)
		assert.Equal(t, Float64{Valid: true, V: 7.5}, value)
	})
	t.Run("div", func(t *testing.T) {
		value := NewFloat64(1.5)
		value.Div(5)
		assert.Equal(t, Float64{Valid: true, V: 0.3}, value)
	})
	t.Run("inc", func(t *testing.T) {
		value := NewFloat64(1.5)
		value.Inc()
		assert.Equal(t, Float64{Valid: true, V: 2.5}, value)
	})
	t.Run("dec", func(t *testing.T) {
		value := NewFloat64(1.5)
		value.Dec()
		assert.Equal(t, Float64{Valid: true, V: 0.5}, value)
	})
	t.Run("zero", func(t *testing.T) {
		value := NewFloat64(0)
		assert.True(t, value.Zero())
	})
	t.Run("positive", func(t *testing.T) {
		value := NewFloat32(1)
		assert.True(t, value.Positive())
	})
	t.Run("negative", func(t *testing.T) {
		value := NewFloat64(-1)
		assert.True(t, value.Negative())
	})
}
