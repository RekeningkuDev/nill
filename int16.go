package nill

// Int16 is a wrapper for int type.
type Int16 Type[int16]

// NewInt16 creates a new Int16.
func NewInt16(value int16) Int16 {
	return Int16{Valid: true, Value: value}
}

// Add adds value to Int16.
func (v *Int16) Add(value int16) {
	v.Value += value
}

// Sub subtracts value from Int16.
func (v *Int16) Sub(value int16) {
	v.Value -= value
}

// Mul multiplies Int16 by value.
func (v *Int16) Mul(value int16) {
	v.Value *= value
}

// Div divides Int16 by value.
func (v *Int16) Div(value int16) {
	v.Value /= value
}

// Mod calculates the remainder of Int16 divided by value.
func (v *Int16) Mod(value int16) {
	v.Value %= value
}

// Inc increments Int16 by 1.
func (v *Int16) Inc() {
	v.Value++
}

// Dec decrements Int16 by 1.
func (v *Int16) Dec() {
	v.Value--
}

// Zero returns true if Int16 is zero.
func (v *Int16) Zero() bool {
	return v.Value == 0
}

// Positive returns true if Int16 is positive.
func (v *Int16) Positive() bool {
	return v.Value > 0
}

// Negative returns true if Int16 is negative.
func (v *Int16) Negative() bool {
	return v.Value < 0
}

// Even returns true if Int16 is even.
func (v *Int16) Even() bool {
	return v.Value%2 == 0
}

// Odd returns true if Int16 is odd.
func (v *Int16) Odd() bool {
	return v.Value%2 != 0
}

// GreaterThan returns true if Int16 is greater than value.
func (v *Int16) GreaterThan(value int16) bool {
	return v.Value > value
}

// GreaterThanOrEqual returns true if Int16 is greater than or equal to value.
func (v *Int16) GreaterThanOrEqual(value int16) bool {
	return v.Value >= value
}

// LessThan returns true if Int16 is less than value.
func (v *Int16) LessThan(value int16) bool {
	return v.Value < value
}

// LessThanOrEqual returns true if Int16 is less than or equal to value.
func (v *Int16) LessThanOrEqual(value int16) bool {
	return v.Value <= value
}

// Equal returns true if Int16 is equal to value.
func (v *Int16) Equal(value int16) bool {
	return v.Value == value
}

// NotEqual returns true if Int16 is not equal to value.
func (v *Int16) NotEqual(value int16) bool {
	return v.Value != value
}

// DivisibleBy returns true if Int16 is divisible by value.
func (v *Int16) DivisibleBy(value int16) bool {
	return v.Value%value == 0
}

// Prime returns true if Int16 is a prime number.
func (v *Int16) Prime() bool {
	if v.Value <= 1 {
		return false
	}

	for i := int16(2); i < v.Value; i++ {
		if v.Value%i == 0 {
			return false
		}
	}

	return true
}
