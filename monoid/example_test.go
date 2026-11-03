package monoid_test

import (
	"fmt"
	"maps"

	"go.unmango.dev/fp/monoid"
)

func ExampleFold() {
	stock := map[string]int{"mango": 3, "papaya": 2}

	fmt.Println(monoid.Fold(monoid.Sum[int](), maps.Values(stock)))
	// Output: 5
}
