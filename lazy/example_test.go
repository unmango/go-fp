package lazy_test

import (
	"fmt"

	"go.unmango.dev/fp/lazy"
)

func ExampleOnce() {
	config := lazy.Once(func() string {
		fmt.Println("loading")
		return "mango"
	})

	fmt.Println(config.Get())
	fmt.Println(config.Get())
	// Output:
	// loading
	// mango
	// mango
}
