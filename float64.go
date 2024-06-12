package nill

type Float64 Type[float64]

func NewFloat64(value float64) Float64 {
	return Float64{Valid: true, V: value}
}

// Add adds value to Float32.
func (v *Float64) Add(value float64) {
	v.V += value
}

// Sub subtracts value from Float32.
func (v *Float64) Sub(value float64) {
	v.V -= value
}

// Mul multiplies Float32 by value.
func (v *Float64) Mul(value float64) {
	v.V *= value
}

// Div divides Float32 by value.
func (v *Float64) Div(value float64) {
	v.V /= value
}

// Inc calculates the remainder of Float32 divided by value.
func (v *Float64) Inc() {
	v.V++
}

// Dec decrements Float32 by 1.
func (v *Float64) Dec() {
	v.V--
}

// Zero returns true if Float32 is zero.
func (v *Float64) Zero() bool {
	return v.V == 0
}

// Positive returns true if Float32 is positive.
func (v *Float64) Positive() bool {
	return v.V > 0
}

// Negative returns true if Float32 is negative.
func (v *Float64) Negative() bool {
	return v.V < 0
}
