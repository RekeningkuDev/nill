package nill

// Int8 is a wrapper for int type.
type Int8 Type[int8]

// NewInt8 creates a new Int8.
func NewInt8(value int8) Int8 {
	return Int8{Valid: true, V: value}
}

// Add adds value to Int8.
func (v *Int8) Add(value int8) {
	v.V += value
}

// Sub subtracts value from Int8.
func (v *Int8) Sub(value int8) {
	v.V -= value
}

// Mul multiplies Int8 by value.
func (v *Int8) Mul(value int8) {
	v.V *= value
}

// Div divides Int8 by value.
func (v *Int8) Div(value int8) {
	v.V /= value
}

// Mod calculates the remainder of Int8 divided by value.
func (v *Int8) Mod(value int8) {
	v.V %= value
}

// Inc increments Int8 by 1.
func (v *Int8) Inc() {
	v.V++
}

// Dec decrements Int8 by 1.
func (v *Int8) Dec() {
	v.V--
}

// Zero returns true if Int8 is zero.
func (v *Int8) Zero() bool {
	return v.V == 0
}

// Positive returns true if Int8 is positive.
func (v *Int8) Positive() bool {
	return v.V > 0
}

// Negative returns true if Int8 is negative.
func (v *Int8) Negative() bool {
	return v.V < 0
}

// Even returns true if Int8 is even.
func (v *Int8) Even() bool {
	return v.V%2 == 0
}

// Odd returns true if Int8 is odd.
func (v *Int8) Odd() bool {
	return v.V%2 != 0
}

// GreaterThan returns true if Int8 is greater than value.
func (v *Int8) GreaterThan(value int8) bool {
	return v.V > value
}

// GreaterThanOrEqual returns true if Int8 is greater than or equal to value.
func (v *Int8) GreaterThanOrEqual(value int8) bool {
	return v.V >= value
}

// LessThan returns true if Int8 is less than value.
func (v *Int8) LessThan(value int8) bool {
	return v.V < value
}

// LessThanOrEqual returns true if Int8 is less than or equal to value.
func (v *Int8) LessThanOrEqual(value int8) bool {
	return v.V <= value
}

// Equal returns true if Int8 is equal to value.
func (v *Int8) Equal(value int8) bool {
	return v.V == value
}

// NotEqual returns true if Int8 is not equal to value.
func (v *Int8) NotEqual(value int8) bool {
	return v.V != value
}

// DivisibleBy returns true if Int8 is divisible by value.
func (v *Int8) DivisibleBy(value int8) bool {
	return v.V%value == 0
}

// Prime returns true if Int8 is a prime number.
func (v *Int8) Prime() bool {
	if v.V <= 1 {
		return false
	}

	for i := int8(2); i < v.V; i++ {
		if v.V%i == 0 {
			return false
		}
	}

	return true
}
