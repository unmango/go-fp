package monoid_test

import (
	"slices"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"go.unmango.dev/fp/monoid"
	"go.unmango.dev/fp/semigroup"
)

type count int

func (a count) Combine(b count) count { return a + b }

// laws checks associativity and identity on a few sample values.
func laws[T any](m monoid.Monoid[T], a, b, c T) {
	GinkgoHelper()
	Expect(m.Combine(m.Combine(a, b), c)).To(Equal(m.Combine(a, m.Combine(b, c))), "associativity")
	Expect(m.Combine(m.Empty(), a)).To(Equal(a), "left identity")
	Expect(m.Combine(a, m.Empty())).To(Equal(a), "right identity")
}

var _ = Describe("Monoid", func() {
	Describe("Lift", func() {
		It("should combine with the function and use the identity", func() {
			m := monoid.Lift(1, func(a, b int) int { return a * b })

			Expect(m.Empty()).To(Equal(1))
			Expect(m.Combine(2, 3)).To(Equal(6))
		})
	})

	Describe("From", func() {
		It("should wrap a semigroup", func() {
			m := monoid.From("", semigroup.Lift(func(a, b string) string { return a + b }))

			Expect(m.Combine("a", "b")).To(Equal("ab"))
			Expect(m.Empty()).To(BeEmpty())
		})
	})

	Describe("Method", func() {
		It("should combine with the type's Combine method", func() {
			m := monoid.Method[count](0)

			laws(m, 1, 2, 3)
		})
	})

	Describe("instances", func() {
		It("Sum should satisfy the laws", func() {
			laws(monoid.Sum[int](), 1, 2, 3)
			Expect(monoid.Sum[float64]().Combine(1.5, 2)).To(Equal(3.5))
		})

		It("Product should satisfy the laws", func() {
			laws(monoid.Product[int](), 2, 3, 4)
		})

		It("String should satisfy the laws", func() {
			laws(monoid.String(), "a", "b", "c")
		})

		It("Slice should satisfy the laws", func() {
			laws(monoid.Slice[int](), []int{1}, []int{2, 3}, []int{4})
		})

		It("Slice should not alias its inputs", func() {
			m := monoid.Slice[int]()
			a := make([]int, 1, 4)

			x := m.Combine(a, []int{1})
			y := m.Combine(a, []int{2})

			Expect(x).To(Equal([]int{0, 1}))
			Expect(y).To(Equal([]int{0, 2}))
		})

		It("Slice should not alias a when b is empty", func() {
			a := []int{1, 2}

			x := monoid.Slice[int]().Combine(a, nil)
			x[0] = 9

			Expect(a).To(Equal([]int{1, 2}))
		})

		It("Slice should keep a non-nil empty right operand under the identity", func() {
			m := monoid.Slice[int]()

			Expect(m.Combine(m.Empty(), []int{})).To(Equal([]int{}))
			Expect(m.Combine(m.Empty(), m.Empty())).To(BeNil())
		})

		It("All should satisfy the laws", func() {
			laws(monoid.All(), true, false, true)
			Expect(monoid.All().Empty()).To(BeTrue())
		})

		It("Any should satisfy the laws", func() {
			laws(monoid.Any(), false, true, false)
			Expect(monoid.Any().Empty()).To(BeFalse())
		})
	})

	Describe("Fold", func() {
		It("should combine every value", func() {
			Expect(monoid.Fold(monoid.Sum[int](), slices.Values([]int{1, 2, 3}))).To(Equal(6))
		})

		It("should return the identity for an empty sequence", func() {
			Expect(monoid.Fold(monoid.Product[int](), slices.Values([]int{}))).To(Equal(1))
		})
	})

	Describe("FoldMap", func() {
		It("should map then combine every value", func() {
			s := monoid.FoldMap(monoid.String(), slices.Values([]int{1, 2, 3}), strconv.Itoa)

			Expect(s).To(Equal("123"))
		})
	})
})
