package maybe_test

import (
	"slices"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"go.unmango.dev/fp/maybe"
)

var _ = Describe("Maybe", func() {
	It("should be None as a zero value", func() {
		var m maybe.Maybe[int]

		Expect(m.IsNone()).To(BeTrue())
		Expect(m).To(Equal(maybe.None[int]()))
	})

	Describe("Some", func() {
		It("should hold the value", func() {
			v, ok := maybe.Some(69).Get()

			Expect(ok).To(BeTrue())
			Expect(v).To(Equal(69))
		})

		It("should hold a zero value", func() {
			m := maybe.Some(0)

			Expect(m.IsSome()).To(BeTrue())
			Expect(m).NotTo(Equal(maybe.None[int]()))
		})
	})

	Describe("None", func() {
		It("should hold nothing", func() {
			v, ok := maybe.None[int]().Get()

			Expect(ok).To(BeFalse())
			Expect(v).To(BeZero())
		})
	})

	It("should be comparable", func() {
		Expect(maybe.Some(1) == maybe.Some(1)).To(BeTrue())
		Expect(maybe.Some(1) == maybe.Some(2)).To(BeFalse())
		Expect(maybe.None[int]() == maybe.None[int]()).To(BeTrue())
		Expect(maybe.Some(0) == maybe.None[int]()).To(BeFalse())
	})

	Describe("FromPtr", func() {
		It("should be None for nil", func() {
			Expect(maybe.FromPtr[int](nil).IsNone()).To(BeTrue())
		})

		It("should copy the pointed-to value", func() {
			x := 69
			m := maybe.FromPtr(&x)
			x = 420

			Expect(m).To(Equal(maybe.Some(69)))
		})
	})

	Describe("FromOk", func() {
		It("should be Some when ok", func() {
			Expect(maybe.FromOk(69, true)).To(Equal(maybe.Some(69)))
		})

		It("should be None when not ok", func() {
			Expect(maybe.FromOk(69, false)).To(Equal(maybe.None[int]()))
		})
	})

	Describe("Ptr", func() {
		It("should be nil for None", func() {
			Expect(maybe.None[int]().Ptr()).To(BeNil())
		})

		It("should point to a copy of the value", func() {
			m := maybe.Some(69)
			p := m.Ptr()
			*p = 420

			Expect(m).To(Equal(maybe.Some(69)))
		})
	})

	Describe("OrElse", func() {
		It("should return the value when Some", func() {
			Expect(maybe.Some(69).OrElse(420)).To(Equal(69))
		})

		It("should return the fallback when None", func() {
			Expect(maybe.None[int]().OrElse(420)).To(Equal(420))
		})
	})

	Describe("OrElseGet", func() {
		It("should not call the fallback when Some", func() {
			v := maybe.Some(69).OrElseGet(func() int {
				Fail("fallback called")
				return 0
			})

			Expect(v).To(Equal(69))
		})

		It("should call the fallback when None", func() {
			Expect(maybe.None[int]().OrElseGet(func() int { return 420 })).To(Equal(420))
		})
	})

	Describe("Map", func() {
		It("should map Some", func() {
			m := maybe.Some(69).Map(strconv.Itoa)

			Expect(m).To(Equal(maybe.Some("69")))
		})

		It("should not call the function for None", func() {
			m := maybe.None[int]().Map(func(int) string {
				Fail("map called")
				return ""
			})

			Expect(m).To(Equal(maybe.None[string]()))
		})
	})

	Describe("Bind", func() {
		half := func(x int) maybe.Maybe[int] {
			if x%2 != 0 {
				return maybe.None[int]()
			}
			return maybe.Some(x / 2)
		}

		It("should chain Some", func() {
			Expect(maybe.Some(8).Bind(half).Bind(half)).To(Equal(maybe.Some(2)))
		})

		It("should stop at the first None", func() {
			Expect(maybe.Some(6).Bind(half).Bind(half)).To(Equal(maybe.None[int]()))
		})
	})

	Describe("Filter", func() {
		even := func(x int) bool { return x%2 == 0 }

		It("should keep a matching value", func() {
			Expect(maybe.Some(2).Filter(even)).To(Equal(maybe.Some(2)))
		})

		It("should drop a value that does not match", func() {
			Expect(maybe.Some(3).Filter(even)).To(Equal(maybe.None[int]()))
		})

		It("should keep None", func() {
			Expect(maybe.None[int]().Filter(even)).To(Equal(maybe.None[int]()))
		})
	})

	Describe("All", func() {
		It("should yield the value when Some", func() {
			Expect(slices.Collect(maybe.Some(69).All())).To(Equal([]int{69}))
		})

		It("should yield nothing when None", func() {
			Expect(slices.Collect(maybe.None[int]().All())).To(BeEmpty())
		})

		It("should stop when the consumer stops", func() {
			for range maybe.Some(69).All() {
				break
			}
		})
	})

	Describe("String", func() {
		It("should format Some", func() {
			Expect(maybe.Some(69).String()).To(Equal("Some(69)"))
		})

		It("should format None", func() {
			Expect(maybe.None[int]().String()).To(Equal("None"))
		})
	})
})
