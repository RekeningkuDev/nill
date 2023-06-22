package nill_test

import (
	"encoding/json"
	"fmt"
	"github.com/RekeningkuDev/null"
	"log"
)

func ExampleTypeMarshal() {
	type BarTest struct {
		Text null.Type[string] `json:"text"`
	}

	type FooTest struct {
		Bar null.Type[BarTest] `json:"bar"`
	}

	foo := FooTest{
		Bar: null.Type[BarTest]{
			Valid: true,
			Value: BarTest{
				Text: null.Type[string]{
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
