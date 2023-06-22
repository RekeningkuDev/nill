package null

import (
	"encoding/json"
	goJson "github.com/goccy/go-json"
	"testing"
)

type BarBenchmarkUnmarshalWithEscape struct {
	Text Type[string] `json:"text"`
}

func BenchmarkUnmarshalWithEscape(b *testing.B) {
	JSONUnmarshal = json.Unmarshal
	b.Run("With Value", func(b *testing.B) {
		var foo BarBenchmarkUnmarshalWithEscape
		jsonStr := "{\"text\":\"test\"}"

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Null", func(b *testing.B) {
		var foo BarBenchmarkUnmarshalWithEscape
		jsonStr := "{\"text\":null}"

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Empty", func(b *testing.B) {
		var foo BarBenchmarkUnmarshalWithEscape
		jsonStr := "{}"

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Empty Value", func(b *testing.B) {
		var foo BarBenchmarkUnmarshalWithEscape
		jsonStr := "{\"text\":\"\"}"

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})
}

func BenchmarkUnmarshal_GoJsonWithEscape(b *testing.B) {
	JSONUnmarshal = goJson.Unmarshal
	b.Run("With Value", func(b *testing.B) {
		var foo BarBenchmarkUnmarshalWithEscape
		jsonStr := "{\"text\":\"test\"}"

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Null", func(b *testing.B) {
		var foo BarBenchmarkUnmarshalWithEscape
		jsonStr := "{\"text\":null}"

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Empty", func(b *testing.B) {
		var foo BarBenchmarkUnmarshalWithEscape
		jsonStr := "{}"

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Empty Value", func(b *testing.B) {
		var foo BarBenchmarkUnmarshalWithEscape
		jsonStr := "{\"text\":\"\"}"

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})
}
