package nill

func (n *Type[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.Valid = false
		return nil
	}

	var value T
	if err := JSONUnmarshal(data, &value); err != nil {
		return err
	}

	n.V = value
	n.Valid = true
	return nil
}

func (n Type[T]) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return JSONMarshal(nil)
	}

	return JSONMarshal(n.V)
}
