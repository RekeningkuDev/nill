package nill

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BarTest struct {
	Text Type[string] `json:"text"`
}

type FooTest struct {
	Bar Type[BarTest] `json:"bar"`
}

func TestType_UnmarshalJSON(t *testing.T) {
	t.Run("should unmarshal json", func(t *testing.T) {
		var foo FooTest
		jsonStr := `{"bar":{"text":"test"}}`
		if err := json.Unmarshal([]byte(jsonStr), &foo); err != nil {
			assert.Error(t, err)
		}

		assert.True(t, foo.Bar.Valid)
		assert.True(t, foo.Bar.V.Text.Valid)
		assert.Equal(t, foo.Bar.V.Text.V, "test")
	})

	t.Run("should unmarshal json Baz with null value", func(t *testing.T) {
		var foo FooTest
		jsonStr := `{"bar":{"text":null}}`

		if err := json.Unmarshal([]byte(jsonStr), &foo); err != nil {
			assert.Error(t, err)
		}

		assert.True(t, foo.Bar.Valid)
		assert.False(t, foo.Bar.V.Text.Valid)
		assert.Equal(t, foo.Bar.V.Text.V, "")
	})

	t.Run("should unmarshal json BarTest with empty value", func(t *testing.T) {
		var foo FooTest
		jsonStr := `{"bar":{}}`

		if err := json.Unmarshal([]byte(jsonStr), &foo); err != nil {
			assert.Error(t, err)
		}

		assert.True(t, foo.Bar.Valid)
		assert.False(t, foo.Bar.V.Text.Valid)
		assert.Equal(t, foo.Bar.V.Text.V, "")
	})

	t.Run("should unmarshal json with empty value", func(t *testing.T) {
		var foo FooTest
		jsonStr := `{}`

		if err := json.Unmarshal([]byte(jsonStr), &foo); err != nil {
			assert.Error(t, err)
		}

		assert.False(t, foo.Bar.Valid)
		assert.False(t, foo.Bar.V.Text.Valid)
		assert.Equal(t, foo.Bar.V.Text.V, "")
	})
}

func TestType_MarshalJSON(t *testing.T) {
	t.Run("should marshal json", func(t *testing.T) {
		foo := FooTest{
			Bar: Type[BarTest]{
				Valid: true,
				V: BarTest{
					Text: Type[string]{
						Valid: true,
						V:     "test",
					},
				},
			},
		}

		jsonStr, err := json.Marshal(foo)
		if err != nil {
			assert.Error(t, err)
		}

		assert.Equal(t, string(jsonStr), `{"bar":{"text":"test"}}`)
	})

	t.Run("should marshal json Text with null value", func(t *testing.T) {
		foo := FooTest{
			Bar: Type[BarTest]{
				Valid: true,
				V: BarTest{
					Text: Type[string]{
						Valid: false,
						V:     "",
					},
				},
			},
		}

		jsonStr, err := json.Marshal(foo)
		if err != nil {
			assert.Error(t, err)
		}

		assert.Equal(t, string(jsonStr), `{"bar":{"text":null}}`)
	})

	t.Run("should marshal json BarTest with null value", func(t *testing.T) {
		foo := FooTest{
			Bar: Type[BarTest]{
				Valid: false,
			},
		}

		jsonStr, err := json.Marshal(foo)
		if err != nil {
			assert.Error(t, err)
		}

		assert.Equal(t, string(jsonStr), `{"bar":null}`)
	})

}
