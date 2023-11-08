package nill

type Type[T any] struct {
	Valid bool
	Value T
}

func New[T any](value T) Type[T] {
	return Type[T]{Valid: true, Value: value}
}
