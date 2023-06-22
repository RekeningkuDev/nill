package nill

import (
	"encoding/json"
	goJson "github.com/goccy/go-json"
	"testing"
)

type BarBenchmarkUnmarshal struct {
	Text Type[string] `json:"text"`
}

func BenchmarkUnmarshal(b *testing.B) {
	JSONUnmarshal = json.Unmarshal
	b.Run("With Value", func(b *testing.B) {
		var foo BarBenchmarkUnmarshal
		jsonStr := `{"text":"test"}`

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Null", func(b *testing.B) {
		var foo BarBenchmarkUnmarshal
		jsonStr := `{"text":null}`

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Empty", func(b *testing.B) {
		var foo BarBenchmarkUnmarshal
		jsonStr := `{}`

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Empty Value", func(b *testing.B) {
		var foo BarBenchmarkUnmarshal
		jsonStr := `{"text":""}`

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})
}

func BenchmarkUnmarshal_GoJson(b *testing.B) {
	JSONUnmarshal = goJson.Unmarshal
	b.Run("With Value", func(b *testing.B) {
		var foo BarBenchmarkUnmarshal
		jsonStr := `{"text":"test"}`

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Null", func(b *testing.B) {
		var foo BarBenchmarkUnmarshal
		jsonStr := `{"text":null}`

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Empty", func(b *testing.B) {
		var foo BarBenchmarkUnmarshal
		jsonStr := `{}`

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Empty Value", func(b *testing.B) {
		var foo BarBenchmarkUnmarshal
		jsonStr := `{"text":""}`

		for i := 0; i < b.N; i++ {
			if err := JSONUnmarshal([]byte(jsonStr), &foo); err != nil {
				b.Error(err)
			}
		}
	})
}
