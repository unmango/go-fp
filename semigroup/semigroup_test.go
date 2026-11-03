package semigroup_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"go.unmango.dev/fp/semigroup"
)

// max combines by keeping the larger value; it has a Combine method.
type max int

func (a max) Combine(b max) max {
	if b > a {
		return b
	}
	return a
}

var _ = Describe("Semigroup", func() {
	Describe("Lift", func() {
		It("should combine with the function", func() {
			s := semigroup.Lift(func(a, b int) int { return a + b })

			Expect(s.Combine(1, 2)).To(Equal(3))
		})
	})

	Describe("Method", func() {
		s := semigroup.Method[max]()

		It("should combine with the type's Combine method", func() {
			Expect(s.Combine(1, 2)).To(Equal(max(2)))
		})

		It("should be associative", func() {
			a, b, c := max(3), max(9), max(5)

			Expect(s.Combine(s.Combine(a, b), c)).To(Equal(s.Combine(a, s.Combine(b, c))))
		})
	})
})
