package nill

import (
	"database/sql"
	"database/sql/driver"
)

// Scan implements the [sql.Scanner] interface.
func (t *Type[T]) Scan(value interface{}) error {
	var x sql.Null[T]
	err := x.Scan(value)
	if err != nil {
		return err
	}

	*t = Type[T](x)

	return nil
}

// Value implements the [driver.Valuer] interface.
func (t Type[T]) Value() (driver.Value, error) {
	sql := sql.Null[T](t)
	return sql.Value()
}
