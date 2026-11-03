package semigroup_test

import (
	"fmt"

	"go.unmango.dev/fp/semigroup"
)

func ExampleLift() {
	longest := semigroup.Lift(func(a, b string) string {
		if len(b) > len(a) {
			return b
		}
		return a
	})

	fmt.Println(longest.Combine("fig", "mango"))
	// Output: mango
}

type Highest int

func (m Highest) Combine(o Highest) Highest {
	if o > m {
		return o
	}
	return m
}

func ExampleMethod() {
	s := semigroup.Method[Highest]()

	fmt.Println(s.Combine(3, 7))
	// Output: 7
}
