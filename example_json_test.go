package nill_test

import (
	"encoding/json"
	"fmt"
	"github.com/RekeningkuDev/nill"
	"log"
)

func ExampleTypeMarshal() {
	type BarTest struct {
		Text nill.Type[string] `json:"text"`
	}

	type FooTest struct {
		Bar nill.Type[BarTest] `json:"bar"`
	}

	foo := FooTest{
		Bar: nill.Type[BarTest]{
			Valid: true,
			Value: BarTest{
				Text: nill.Type[string]{
					Valid: true,
					Value: "test",
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
