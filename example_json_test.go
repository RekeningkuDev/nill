package nill_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/RekeningkuDev/nill"
)

func ExampleType() {
	type BarTest struct {
		Text nill.Type[string] `json:"text"`
	}

	type FooTest struct {
		Bar nill.Type[BarTest] `json:"bar"`
	}

	foo := FooTest{
		Bar: nill.Type[BarTest]{
			Valid: true,
			V: BarTest{
				Text: nill.Type[string]{
					Valid: true,
					V:     "test",
				},
			},
		},
	}

	jsonStr, err := json.Marshal(foo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonStr))
	// Output:
	// {"bar":{"text":"test"}}
}
