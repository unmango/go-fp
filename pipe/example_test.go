package pipe_test

import (
	"fmt"
	"strings"

	"go.unmango.dev/fp/pipe"
)

func ExamplePipe3() {
	shout := pipe.Pipe3(strings.TrimSpace, strings.ToUpper, func(s string) string { return s + "!" })

	fmt.Println(shout("  mango "))
	// Output: MANGO!
}
