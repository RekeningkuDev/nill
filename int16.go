package nill

// Int16 is a wrapper for int type.
type Int16 Type[int16]

// NewInt16 creates a new Int16.
func NewInt16(value int16) Int16 {
	return Int16{Valid: true, V: value}
}

// Add adds value to Int16.
func (v *Int16) Add(value int16) {
	v.V += value
}

// Sub subtracts value from Int16.
func (v *Int16) Sub(value int16) {
	v.V -= value
}

// Mul multiplies Int16 by value.
func (v *Int16) Mul(value int16) {
	v.V *= value
}

// Div divides Int16 by value.
func (v *Int16) Div(value int16) {
	v.V /= value
}

// Mod calculates the remainder of Int16 divided by value.
func (v *Int16) Mod(value int16) {
	v.V %= value
}

// Inc increments Int16 by 1.
func (v *Int16) Inc() {
	v.V++
}

// Dec decrements Int16 by 1.
func (v *Int16) Dec() {
	v.V--
}

// Zero returns true if Int16 is zero.
func (v *Int16) Zero() bool {
	return v.V == 0
}

// Positive returns true if Int16 is positive.
func (v *Int16) Positive() bool {
	return v.V > 0
}

// Negative returns true if Int16 is negative.
func (v *Int16) Negative() bool {
	return v.V < 0
}

// Even returns true if Int16 is even.
func (v *Int16) Even() bool {
	return v.V%2 == 0
}

// Odd returns true if Int16 is odd.
func (v *Int16) Odd() bool {
	return v.V%2 != 0
}

// GreaterThan returns true if Int16 is greater than value.
func (v *Int16) GreaterThan(value int16) bool {
	return v.V > value
}

// GreaterThanOrEqual returns true if Int16 is greater than or equal to value.
func (v *Int16) GreaterThanOrEqual(value int16) bool {
	return v.V >= value
}

// LessThan returns true if Int16 is less than value.
func (v *Int16) LessThan(value int16) bool {
	return v.V < value
}

// LessThanOrEqual returns true if Int16 is less than or equal to value.
func (v *Int16) LessThanOrEqual(value int16) bool {
	return v.V <= value
}

// Equal returns true if Int16 is equal to value.
func (v *Int16) Equal(value int16) bool {
	return v.V == value
}

// NotEqual returns true if Int16 is not equal to value.
func (v *Int16) NotEqual(value int16) bool {
	return v.V != value
}

// DivisibleBy returns true if Int16 is divisible by value.
func (v *Int16) DivisibleBy(value int16) bool {
	return v.V%value == 0
}

// Prime returns true if Int16 is a prime number.
func (v *Int16) Prime() bool {
	if v.V <= 1 {
		return false
	}

	for i := int16(2); i < v.V; i++ {
		if v.V%i == 0 {
			return false
		}
	}

	return true
}
