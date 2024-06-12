package nill

// Int is a wrapper for int type.
type Int Type[int]

// NewInt creates a new Int.
func NewInt(value int) Int {
	return Int{Valid: true, V: value}
}

// Add adds value to Int.
func (v *Int) Add(value int) {
	v.V += value
}

// Sub subtracts value from Int.
func (v *Int) Sub(value int) {
	v.V -= value
}

// Mul multiplies Int by value.
func (v *Int) Mul(value int) {
	v.V *= value
}

// Div divides Int by value.
func (v *Int) Div(value int) {
	v.V /= value
}

// Mod calculates the remainder of Int divided by value.
func (v *Int) Mod(value int) {
	v.V %= value
}

// Inc increments Int by 1.
func (v *Int) Inc() {
	v.V++
}

// Dec decrements Int by 1.
func (v *Int) Dec() {
	v.V--
}

// Zero returns true if Int is zero.
func (v *Int) Zero() bool {
	return v.V == 0
}

// Positive returns true if Int is positive.
func (v *Int) Positive() bool {
	return v.V > 0
}

// Negative returns true if Int is negative.
func (v *Int) Negative() bool {
	return v.V < 0
}

// Even returns true if Int is even.
func (v *Int) Even() bool {
	return v.V%2 == 0
}

// Odd returns true if Int is odd.
func (v *Int) Odd() bool {
	return v.V%2 != 0
}

// GreaterThan returns true if Int is greater than value.
func (v *Int) GreaterThan(value int) bool {
	return v.V > value
}

// GreaterThanOrEqual returns true if Int is greater than or equal to value.
func (v *Int) GreaterThanOrEqual(value int) bool {
	return v.V >= value
}

// LessThan returns true if Int is less than value.
func (v *Int) LessThan(value int) bool {
	return v.V < value
}

// LessThanOrEqual returns true if Int is less than or equal to value.
func (v *Int) LessThanOrEqual(value int) bool {
	return v.V <= value
}

// Equal returns true if Int is equal to value.
func (v *Int) Equal(value int) bool {
	return v.V == value
}

// NotEqual returns true if Int is not equal to value.
func (v *Int) NotEqual(value int) bool {
	return v.V != value
}

// DivisibleBy returns true if Int is divisible by value.
func (v *Int) DivisibleBy(value int) bool {
	return v.V%value == 0
}

// Prime returns true if Int is a prime number.
func (v *Int) Prime() bool {
	if v.V <= 1 {
		return false
	}

	for i := 2; i < v.V; i++ {
		if v.V%i == 0 {
			return false
		}
	}

	return true
}
