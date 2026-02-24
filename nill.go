package nill

import "database/sql"

type Type[T any] sql.Null[T]

func New[T any](v T) Type[T] {
	return Type[T]{Valid: true, V: v}
}

func (t Type[T]) IsZero() bool {
	return !t.Valid
}
