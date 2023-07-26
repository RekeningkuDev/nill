package nill

// Int is a wrapper for int type.
type Int Type[int]

// NewInt creates a new Int.
func NewInt(value int) Int {
	return Int{Valid: true, Value: value}
}

// Add adds value to Int.
func (v *Int) Add(value int) {
	v.Value += value
}

// Sub subtracts value from Int.
func (v *Int) Sub(value int) {
	v.Value -= value
}

// Mul multiplies Int by value.
func (v *Int) Mul(value int) {
	v.Value *= value
}

// Div divides Int by value.
func (v *Int) Div(value int) {
	v.Value /= value
}

// Mod calculates the remainder of Int divided by value.
func (v *Int) Mod(value int) {
	v.Value %= value
}

// Inc increments Int by 1.
func (v *Int) Inc() {
	v.Value++
}

// Dec decrements Int by 1.
func (v *Int) Dec() {
	v.Value--
}

// Zero returns true if Int is zero.
func (v *Int) Zero() bool {
	return v.Value == 0
}

// Positive returns true if Int is positive.
func (v *Int) Positive() bool {
	return v.Value > 0
}

// Negative returns true if Int is negative.
func (v *Int) Negative() bool {
	return v.Value < 0
}

// Even returns true if Int is even.
func (v *Int) Even() bool {
	return v.Value%2 == 0
}

// Odd returns true if Int is odd.
func (v *Int) Odd() bool {
	return v.Value%2 != 0
}

// GreaterThan returns true if Int is greater than value.
func (v *Int) GreaterThan(value int) bool {
	return v.Value > value
}

// GreaterThanOrEqual returns true if Int is greater than or equal to value.
func (v *Int) GreaterThanOrEqual(value int) bool {
	return v.Value >= value
}

// LessThan returns true if Int is less than value.
func (v *Int) LessThan(value int) bool {
	return v.Value < value
}

// LessThanOrEqual returns true if Int is less than or equal to value.
func (v *Int) LessThanOrEqual(value int) bool {
	return v.Value <= value
}

// Equal returns true if Int is equal to value.
func (v *Int) Equal(value int) bool {
	return v.Value == value
}

// NotEqual returns true if Int is not equal to value.
func (v *Int) NotEqual(value int) bool {
	return v.Value != value
}

// DivisibleBy returns true if Int is divisible by value.
func (v *Int) DivisibleBy(value int) bool {
	return v.Value%value == 0
}

// Prime returns true if Int is a prime number.
func (v *Int) Prime() bool {
	if v.Value <= 1 {
		return false
	}

	for i := 2; i < v.Value; i++ {
		if v.Value%i == 0 {
			return false
		}
	}

	return true
}
