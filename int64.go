package nill

// Int64 is a wrapper for int type.
type Int64 Type[int64]

// NewInt64 creates a new Int64.
func NewInt64(value int64) Int64 {
	return Int64{Valid: true, V: value}
}

// Add adds value to Int64.
func (v *Int64) Add(value int64) {
	v.V += value
}

// Sub subtracts value from Int64.
func (v *Int64) Sub(value int64) {
	v.V -= value
}

// Mul multiplies Int64 by value.
func (v *Int64) Mul(value int64) {
	v.V *= value
}

// Div divides Int64 by value.
func (v *Int64) Div(value int64) {
	v.V /= value
}

// Mod calculates the remainder of Int64 divided by value.
func (v *Int64) Mod(value int64) {
	v.V %= value
}

// Inc increments Int64 by 1.
func (v *Int64) Inc() {
	v.V++
}

// Dec decrements Int64 by 1.
func (v *Int64) Dec() {
	v.V--
}

// Zero returns true if Int64 is zero.
func (v *Int64) Zero() bool {
	return v.V == 0
}

// Positive returns true if Int64 is positive.
func (v *Int64) Positive() bool {
	return v.V > 0
}

// Negative returns true if Int64 is negative.
func (v *Int64) Negative() bool {
	return v.V < 0
}

// Even returns true if Int64 is even.
func (v *Int64) Even() bool {
	return v.V%2 == 0
}

// Odd returns true if Int64 is odd.
func (v *Int64) Odd() bool {
	return v.V%2 != 0
}

// GreaterThan returns true if Int64 is greater than value.
func (v *Int64) GreaterThan(value int64) bool {
	return v.V > value
}

// GreaterThanOrEqual returns true if Int64 is greater than or equal to value.
func (v *Int64) GreaterThanOrEqual(value int64) bool {
	return v.V >= value
}

// LessThan returns true if Int64 is less than value.
func (v *Int64) LessThan(value int64) bool {
	return v.V < value
}

// LessThanOrEqual returns true if Int64 is less than or equal to value.
func (v *Int64) LessThanOrEqual(value int64) bool {
	return v.V <= value
}

// Equal returns true if Int64 is equal to value.
func (v *Int64) Equal(value int64) bool {
	return v.V == value
}

// NotEqual returns true if Int64 is not equal to value.
func (v *Int64) NotEqual(value int64) bool {
	return v.V != value
}

// DivisibleBy returns true if Int64 is divisible by value.
func (v *Int64) DivisibleBy(value int64) bool {
	return v.V%value == 0
}

// Prime returns true if Int64 is a prime number.
func (v *Int64) Prime() bool {
	if v.V <= 1 {
		return false
	}

	for i := int64(2); i < v.V; i++ {
		if v.V%i == 0 {
			return false
		}
	}

	return true
}
