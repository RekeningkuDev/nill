package nill

import "testing"

func TestScan(t *testing.T) {
	var foo Type[int]
	err := foo.Scan(nil)
	if err != nil {
		t.Error(err)
	}

	if foo.Valid {
		t.Error("should be invalid")
	}

	err = foo.Scan(1)
	if err != nil {
		t.Error(err)
	}

	wantFoo := Type[int]{Valid: true, V: 1}

	if foo != wantFoo {
		t.Errorf("foo %v want %v", foo, wantFoo)
	}
}

func TestValue(t *testing.T) {
	foo := Type[int64]{Valid: true, V: 1}
	v, err := foo.Value()
	if err != nil {
		t.Error("error should be nil")
	}

	vInt, ok := v.(int64)
	if !ok {
		t.Errorf("value should be int64, got %T", v)
	}

	if vInt != 1 {
		t.Errorf("value should be 1, got %d", vInt)
	}
}
