package nill

// Int32 is a wrapper for int type.
type Int32 Type[int32]

// NewInt32 creates a new Int32.
func NewInt32(value int32) Int32 {
	return Int32{Valid: true, Value: value}
}

// Add adds value to Int32.
func (v *Int32) Add(value int32) {
	v.Value += value
}

// Sub subtracts value from Int32.
func (v *Int32) Sub(value int32) {
	v.Value -= value
}

// Mul multiplies Int32 by value.
func (v *Int32) Mul(value int32) {
	v.Value *= value
}

// Div divides Int32 by value.
func (v *Int32) Div(value int32) {
	v.Value /= value
}

// Mod calculates the remainder of Int32 divided by value.
func (v *Int32) Mod(value int32) {
	v.Value %= value
}

// Inc increments Int32 by 1.
func (v *Int32) Inc() {
	v.Value++
}

// Dec decrements Int32 by 1.
func (v *Int32) Dec() {
	v.Value--
}

// Zero returns true if Int32 is zero.
func (v *Int32) Zero() bool {
	return v.Value == 0
}

// Positive returns true if Int32 is positive.
func (v *Int32) Positive() bool {
	return v.Value > 0
}

// Negative returns true if Int32 is negative.
func (v *Int32) Negative() bool {
	return v.Value < 0
}

// Even returns true if Int32 is even.
func (v *Int32) Even() bool {
	return v.Value%2 == 0
}

// Odd returns true if Int32 is odd.
func (v *Int32) Odd() bool {
	return v.Value%2 != 0
}

// GreaterThan returns true if Int32 is greater than value.
func (v *Int32) GreaterThan(value int32) bool {
	return v.Value > value
}

// GreaterThanOrEqual returns true if Int32 is greater than or equal to value.
func (v *Int32) GreaterThanOrEqual(value int32) bool {
	return v.Value >= value
}

// LessThan returns true if Int32 is less than value.
func (v *Int32) LessThan(value int32) bool {
	return v.Value < value
}

// LessThanOrEqual returns true if Int32 is less than or equal to value.
func (v *Int32) LessThanOrEqual(value int32) bool {
	return v.Value <= value
}

// Equal returns true if Int32 is equal to value.
func (v *Int32) Equal(value int32) bool {
	return v.Value == value
}

// NotEqual returns true if Int32 is not equal to value.
func (v *Int32) NotEqual(value int32) bool {
	return v.Value != value
}

// DivisibleBy returns true if Int32 is divisible by value.
func (v *Int32) DivisibleBy(value int32) bool {
	return v.Value%value == 0
}

// Prime returns true if Int32 is a prime number.
func (v *Int32) Prime() bool {
	if v.Value <= 1 {
		return false
	}

	for i := int32(2); i < v.Value; i++ {
		if v.Value%i == 0 {
			return false
		}
	}

	return true
}
