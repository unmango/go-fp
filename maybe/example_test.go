package maybe_test

import (
	"fmt"
	"strings"

	"go.unmango.dev/fp/maybe"
)

func Example() {
	lookup := func(m map[string]string, k string) maybe.Maybe[string] {
		v, ok := m[k]
		return maybe.FromOk(v, ok)
	}

	env := map[string]string{"name": "mango"}

	fmt.Println(lookup(env, "name").Map(strings.ToUpper))
	fmt.Println(lookup(env, "missing").Map(strings.ToUpper).OrElse("anonymous"))
	// Output:
	// Some(MANGO)
	// anonymous
}

func ExampleMaybe_Get() {
	if v, ok := maybe.Some(69).Get(); ok {
		fmt.Println(v)
	}
	// Output: 69
}
