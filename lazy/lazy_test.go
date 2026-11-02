package lazy_test

import (
	"strconv"
	"sync"
	"sync/atomic"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"go.unmango.dev/fp/lazy"
)

var _ = Describe("Lazy", func() {
	Describe("Of", func() {
		It("should return the value", func() {
			Expect(lazy.Of(69).Get()).To(Equal(69))
		})
	})

	Describe("Once", func() {
		It("should not call the function until forced", func() {
			var calls atomic.Int32
			_ = lazy.Once(func() int {
				calls.Add(1)
				return 69
			})

			Expect(calls.Load()).To(BeZero())
		})

		It("should call the function once across goroutines", func() {
			var calls atomic.Int32
			l := lazy.Once(func() int {
				calls.Add(1)
				return 69
			})

			var wg sync.WaitGroup
			for range 16 {
				wg.Go(func() {
					defer GinkgoRecover()
					Expect(l.Get()).To(Equal(69))
				})
			}
			wg.Wait()

			Expect(calls.Load()).To(Equal(int32(1)))
		})

		It("should accept a Lazy", func() {
			calls := 0
			l := lazy.Once(lazy.Lazy[int](func() int {
				calls++
				return 69
			}))

			l.Get()
			l.Get()

			Expect(calls).To(Equal(1))
		})
	})

	Describe("Once2", func() {
		It("should call the function once", func() {
			calls := 0
			l := lazy.Once2(func() (int, string) {
				calls++
				return 69, "mango"
			})

			l.Get()
			a, b := l.Get()

			Expect(a).To(Equal(69))
			Expect(b).To(Equal("mango"))
			Expect(calls).To(Equal(1))
		})
	})

	It("should yield the zero value for a nil Lazy", func() {
		var l lazy.Lazy[int]
		var l2 lazy.Lazy2[int, string]

		Expect(l.Get()).To(BeZero())
		Expect(l.Map(func(x int) int { return x + 1 }).Get()).To(Equal(1))

		a, b := l2.Get()
		Expect(a).To(BeZero())
		Expect(b).To(BeZero())
	})

	Describe("Map", func() {
		It("should not force the value", func() {
			l := lazy.Lazy[int](func() int {
				Fail("forced")
				return 0
			})

			_ = l.Map(strconv.Itoa)
		})

		It("should map the value when forced", func() {
			Expect(lazy.Of(69).Map(strconv.Itoa).Get()).To(Equal("69"))
		})
	})
})
