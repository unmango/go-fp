package either_test

import (
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"go.unmango.dev/fp/either"
)

var _ = Describe("Either", func() {
	It("should be Left with the zero value as a zero value", func() {
		var e either.Either[int, string]

		Expect(e.IsLeft()).To(BeTrue())
		Expect(e).To(Equal(either.Left[string](0)))
	})

	Describe("Left", func() {
		It("should hold the left value", func() {
			e := either.Left[string](69)

			l, ok := e.Left()
			Expect(ok).To(BeTrue())
			Expect(l).To(Equal(69))

			_, ok = e.Right()
			Expect(ok).To(BeFalse())
		})
	})

	Describe("Right", func() {
		It("should hold the right value", func() {
			e := either.Right[int]("mango")

			r, ok := e.Right()
			Expect(ok).To(BeTrue())
			Expect(r).To(Equal("mango"))
			Expect(e.IsRight()).To(BeTrue())

			_, ok = e.Left()
			Expect(ok).To(BeFalse())
		})
	})

	It("should be comparable", func() {
		Expect(either.Left[string](1) == either.Left[string](1)).To(BeTrue())
		Expect(either.Left[string](0) == either.Right[int]("")).To(BeFalse())
		Expect(either.Right[int]("a") == either.Right[int]("b")).To(BeFalse())
	})

	Describe("MapLeft", func() {
		It("should map a left value", func() {
			e := either.Left[string](69).MapLeft(strconv.Itoa)

			Expect(e).To(Equal(either.Left[string]("69")))
		})

		It("should keep a right value", func() {
			e := either.Right[int]("mango").MapLeft(func(int) bool {
				Fail("map called")
				return false
			})

			Expect(e).To(Equal(either.Right[bool]("mango")))
		})
	})

	Describe("MapRight", func() {
		It("should map a right value", func() {
			e := either.Right[string](69).MapRight(strconv.Itoa)

			Expect(e).To(Equal(either.Right[string]("69")))
		})

		It("should keep a left value", func() {
			e := either.Left[int]("mango").MapRight(func(int) bool {
				Fail("map called")
				return false
			})

			Expect(e).To(Equal(either.Left[bool]("mango")))
		})
	})

	Describe("BindLeft", func() {
		It("should bind a left value", func() {
			e := either.Left[string](69).BindLeft(func(x int) either.Either[int, string] {
				return either.Right[int](strconv.Itoa(x))
			})

			Expect(e).To(Equal(either.Right[int]("69")))
		})

		It("should keep a right value", func() {
			e := either.Right[int]("mango").BindLeft(func(int) either.Either[bool, string] {
				Fail("bind called")
				return either.Left[string](false)
			})

			Expect(e).To(Equal(either.Right[bool]("mango")))
		})
	})

	Describe("BindRight", func() {
		It("should bind a right value", func() {
			e := either.Right[string](69).BindRight(func(x int) either.Either[string, string] {
				return either.Left[string](strconv.Itoa(x))
			})

			Expect(e).To(Equal(either.Left[string]("69")))
		})

		It("should keep a left value", func() {
			e := either.Left[int]("mango").BindRight(func(int) either.Either[string, bool] {
				Fail("bind called")
				return either.Right[string](false)
			})

			Expect(e).To(Equal(either.Left[bool]("mango")))
		})
	})

	Describe("Swap", func() {
		It("should turn Left into Right", func() {
			Expect(either.Left[string](69).Swap()).To(Equal(either.Right[string](69)))
		})

		It("should turn Right into Left", func() {
			Expect(either.Right[int]("mango").Swap()).To(Equal(either.Left[int]("mango")))
		})
	})

	Describe("Fold", func() {
		onLeft := func(x int) string { return "left " + strconv.Itoa(x) }
		onRight := func(s string) string { return "right " + s }

		It("should fold a left value", func() {
			Expect(either.Left[string](69).Fold(onLeft, onRight)).To(Equal("left 69"))
		})

		It("should fold a right value", func() {
			Expect(either.Right[int]("mango").Fold(onLeft, onRight)).To(Equal("right mango"))
		})
	})

	Describe("String", func() {
		It("should format Left", func() {
			Expect(either.Left[string](69).String()).To(Equal("Left(69)"))
		})

		It("should format Right", func() {
			Expect(either.Right[int]("mango").String()).To(Equal("Right(mango)"))
		})
	})
})
