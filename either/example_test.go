package either_test

import (
	"fmt"
	"strconv"

	"go.unmango.dev/fp/either"
)

func Example() {
	parse := func(s string) either.Either[int, string] {
		if n, err := strconv.Atoi(s); err == nil {
			return either.Left[string](n)
		}
		return either.Right[int](s)
	}

	describe := func(e either.Either[int, string]) string {
		return e.Fold(
			func(n int) string { return fmt.Sprint("number ", n) },
			func(s string) string { return "word " + s },
		)
	}

	fmt.Println(describe(parse("69")))
	fmt.Println(describe(parse("mango")))
	// Output:
	// number 69
	// word mango
}
