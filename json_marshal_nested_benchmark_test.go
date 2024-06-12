package nill

import (
	"encoding/json"
	"testing"

	goJson "github.com/goccy/go-json"
)

type BarBenchmarkNestedMarshal struct {
	Text Type[string] `json:"text"`
}

type FooBenchmarkNestedMarshal struct {
	Bar Type[BarBenchmarkNestedMarshal] `json:"bar"`
}

func BenchmarkNestedMarshal(b *testing.B) {
	JSONMarshal = json.Marshal

	b.Run("With Value", func(b *testing.B) {
		foo := FooBenchmarkNestedMarshal{
			Bar: Type[BarBenchmarkNestedMarshal]{
				Valid: true,
				V: BarBenchmarkNestedMarshal{
					Text: Type[string]{
						Valid: true,
						V:     "test",
					},
				},
			},
		}

		for i := 0; i < b.N; i++ {
			if _, err := JSONMarshal(foo); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Null", func(b *testing.B) {
		foo := FooBenchmarkNestedMarshal{
			Bar: Type[BarBenchmarkNestedMarshal]{
				Valid: false,
			},
		}

		for i := 0; i < b.N; i++ {
			if _, err := JSONMarshal(foo); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Empty", func(b *testing.B) {
		foo := FooBenchmarkNestedMarshal{}

		for i := 0; i < b.N; i++ {
			if _, err := JSONMarshal(foo); err != nil {
				b.Error(err)
			}
		}
	})
}

func BenchmarkNestedMarshal_GoJson(b *testing.B) {
	JSONMarshal = goJson.Marshal

	b.Run("With Value", func(b *testing.B) {
		foo := FooBenchmarkNestedMarshal{
			Bar: Type[BarBenchmarkNestedMarshal]{
				Valid: true,
				V: BarBenchmarkNestedMarshal{
					Text: Type[string]{
						Valid: true,
						V:     "test",
					},
				},
			},
		}

		for i := 0; i < b.N; i++ {
			if _, err := JSONMarshal(foo); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Null", func(b *testing.B) {
		foo := FooBenchmarkNestedMarshal{
			Bar: Type[BarBenchmarkNestedMarshal]{
				Valid: false,
			},
		}

		for i := 0; i < b.N; i++ {
			if _, err := JSONMarshal(foo); err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Empty", func(b *testing.B) {
		foo := FooBenchmarkNestedMarshal{}

		for i := 0; i < b.N; i++ {
			if _, err := JSONMarshal(foo); err != nil {
				b.Error(err)
			}
		}
	})
}
