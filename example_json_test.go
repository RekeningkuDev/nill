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
		Zoo nill.Type[string]  `json:"zoo,omitzero"`
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
		Zoo: nill.Type[string]{
			V: "this should be omitted",
			// Valid is false, so this field will be omitted in JSON output
			Valid: false,
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
