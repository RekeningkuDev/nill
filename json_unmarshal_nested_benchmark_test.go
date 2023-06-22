package null

import (
	"encoding/json"
	goJson "github.com/goccy/go-json"
	"testing"
)

type BarBenchmarkNestedUnmarshal struct {
	Text Type[string] `json:"text"`
}

type FooBenchmarkNestedUnmarshal struct {
	Bar Type[BarBenchmarkNestedUnmarshal] `json:"bar"`
}

func BenchmarkNestedUnmarshal(b *testing.B) {
	JSONUnmarshal = json.Unmarshal
	b.Run("With Value", func(b *testing.B) {
		var foo FooBenchmarkNestedUnmarshal
		jsonStr := `{"bar":{"text":"test"}}`

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Null", func(b *testing.B) {
		var foo FooBenchmarkNestedUnmarshal
		jsonStr := `{"bar":{"text":null}}`

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Empty", func(b *testing.B) {
		var foo FooBenchmarkNestedUnmarshal
		jsonStr := `{"bar":{}}`

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Empty Value", func(b *testing.B) {
		var foo FooBenchmarkNestedUnmarshal
		jsonStr := `{"bar":{"text":""}}`

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})
}

func BenchmarkNestedUnmarshal_GoJson(b *testing.B) {
	JSONUnmarshal = goJson.Unmarshal
	b.Run("With Value", func(b *testing.B) {
		var foo FooBenchmarkNestedUnmarshal
		jsonStr := `{"bar":{"text":"test"}}`

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Null", func(b *testing.B) {
		var foo FooBenchmarkNestedUnmarshal
		jsonStr := `{"bar":{"text":null}}`

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Empty", func(b *testing.B) {
		var foo FooBenchmarkNestedUnmarshal
		jsonStr := `{"bar":{}}`

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Empty Value", func(b *testing.B) {
		var foo FooBenchmarkNestedUnmarshal
		jsonStr := `{"bar":{"text":""}}`

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})
}
