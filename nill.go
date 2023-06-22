package nill

type Type[T comparable] struct {
	Valid bool
	Value T
}

func New[T comparable](value T) Type[T] {
	return Type[T]{Valid: true, Value: value}
}
