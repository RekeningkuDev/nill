package nill

import (
	"testing"

	goJson "github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

type BarBenchmarkMarshal struct {
	Text Type[string] `json:"text"`
}

func BenchmarkMarshal(b *testing.B) {
	b.Run("should marshal json", func(b *testing.B) {
		foo := BarBenchmarkMarshal{
			Text: Type[string]{
				Valid: true,
				V:     "test",
			},
		}

		for i := 0; i < b.N; i++ {
			if _, err := JSONMarshal(foo); err != nil {
				assert.Error(b, err)
			}
		}
	})

	b.Run("should marshal null", func(b *testing.B) {
		foo := BarBenchmarkMarshal{
			Text: Type[string]{
				Valid: false,
			},
		}

		for i := 0; i < b.N; i++ {
			if _, err := JSONMarshal(foo); err != nil {
				assert.Error(b, err)
			}
		}
	})

	b.Run("should marshal empty", func(b *testing.B) {
		foo := BarBenchmarkMarshal{}

		for i := 0; i < b.N; i++ {
			if _, err := JSONMarshal(foo); err != nil {
				assert.Error(b, err)
			}
		}
	})
}

func BenchmarkMarshal_GoJson(b *testing.B) {
	JSONMarshal = goJson.Marshal

	b.Run("should marshal json", func(b *testing.B) {
		foo := BarBenchmarkMarshal{
			Text: Type[string]{
				Valid: true,
				V:     "test",
			},
		}

		for i := 0; i < b.N; i++ {
			if _, err := JSONMarshal(foo); err != nil {
				assert.Error(b, err)
			}
		}
	})

	b.Run("should marshal null", func(b *testing.B) {
		foo := BarBenchmarkMarshal{
			Text: Type[string]{
				Valid: false,
			},
		}

		for i := 0; i < b.N; i++ {
			if _, err := JSONMarshal(foo); err != nil {
				assert.Error(b, err)
			}
		}
	})

	b.Run("should marshal empty", func(b *testing.B) {
		foo := BarBenchmarkMarshal{}

		for i := 0; i < b.N; i++ {
			if _, err := JSONMarshal(foo); err != nil {
				assert.Error(b, err)
			}
		}
	})
}
