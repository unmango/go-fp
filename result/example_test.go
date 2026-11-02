package result_test

import (
	"fmt"
	"strconv"

	"go.unmango.dev/fp/result"
)

func Example() {
	double := func(s string) (int, error) {
		return result.Of(strconv.Atoi(s)).
			Map(func(x int) int { return x * 2 }).
			Get()
	}

	fmt.Println(double("21"))
	fmt.Println(double("mango"))
	// Output:
	// 42 <nil>
	// 0 strconv.Atoi: parsing "mango": invalid syntax
}
