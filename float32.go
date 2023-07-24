package nill

// Float32 is a wrapper for float32 type.
type Float32 Type[float32]

// NewFloat32 creates a new Float32.
func NewFloat32(value float32) Float32 {
	return Float32{Valid: true, Value: value}
}

// Add adds value to Float32.
func (v *Float32) Add(value float32) {
	v.Value += value
}

// Sub subtracts value from Float32.
func (v *Float32) Sub(value float32) {
	v.Value -= value
}

// Mul multiplies Float32 by value.
func (v *Float32) Mul(value float32) {
	v.Value *= value
}

// Div divides Float32 by value.
func (v *Float32) Div(value float32) {
	v.Value /= value
}

// Inc calculates the remainder of Float32 divided by value.
func (v *Float32) Inc() {
	v.Value++
}

// Dec decrements Float32 by 1.
func (v *Float32) Dec() {
	v.Value--
}

// Zero returns true if Float32 is zero.
func (v *Float32) Zero() bool {
	return v.Value == 0
}

// Positive returns true if Float32 is positive.
func (v *Float32) Positive() bool {
	return v.Value > 0
}

// Negative returns true if Float32 is negative.
func (v *Float32) Negative() bool {
	return v.Value < 0
}
