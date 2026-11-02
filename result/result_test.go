package result_test

import (
	"errors"
	"fmt"
	"io/fs"
	"slices"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"go.unmango.dev/fp/maybe"
	"go.unmango.dev/fp/result"
)

var errBoom = errors.New("boom")

var _ = Describe("Result", func() {
	It("should be Ok with the zero value as a zero value", func() {
		var r result.Result[int]

		Expect(r.IsOk()).To(BeTrue())
		Expect(r.Get()).To(Equal(0))
	})

	Describe("Ok", func() {
		It("should hold the value", func() {
			v, err := result.Ok(69).Get()

			Expect(err).NotTo(HaveOccurred())
			Expect(v).To(Equal(69))
		})
	})

	Describe("Error", func() {
		It("should hold the error", func() {
			r := result.Error[int](errBoom)

			Expect(r.IsErr()).To(BeTrue())
			Expect(r.Err()).To(MatchError(errBoom))
		})

		It("should be Ok with the zero value for a nil error", func() {
			Expect(result.Error[int](nil)).To(Equal(result.Ok(0)))
		})
	})

	Describe("Errorf", func() {
		It("should format the error", func() {
			r := result.Errorf[int]("bad %d", 69)

			Expect(r.Err()).To(MatchError("bad 69"))
		})

		It("should wrap with %w", func() {
			r := result.Errorf[int]("open: %w", fs.ErrNotExist)

			Expect(errors.Is(r.Err(), fs.ErrNotExist)).To(BeTrue())
		})
	})

	Describe("Of", func() {
		It("should be Ok for a nil error", func() {
			Expect(result.Of(strconv.Atoi("69"))).To(Equal(result.Ok(69)))
		})

		It("should be an error and drop the value for a non-nil error", func() {
			v, err := result.Of(69, errBoom).Get()

			Expect(err).To(MatchError(errBoom))
			Expect(v).To(BeZero())
		})
	})

	Describe("FromMaybe", func() {
		It("should be Ok for Some", func() {
			Expect(result.FromMaybe(maybe.Some(69))).To(Equal(result.Ok(69)))
		})

		It("should be ErrNone for None", func() {
			r := result.FromMaybe(maybe.None[int]())

			Expect(r.Err()).To(MatchError(maybe.ErrNone))
		})
	})

	Describe("OrElse", func() {
		It("should return the value when Ok", func() {
			Expect(result.Ok(69).OrElse(420)).To(Equal(69))
		})

		It("should return the fallback when an error", func() {
			Expect(result.Error[int](errBoom).OrElse(420)).To(Equal(420))
		})
	})

	Describe("Map", func() {
		It("should map Ok", func() {
			Expect(result.Ok(69).Map(strconv.Itoa)).To(Equal(result.Ok("69")))
		})

		It("should keep the error and not call the function", func() {
			r := result.Error[int](errBoom).Map(func(int) string {
				Fail("map called")
				return ""
			})

			Expect(r.Err()).To(MatchError(errBoom))
		})
	})

	Describe("Bind", func() {
		parse := func(s string) result.Result[int] {
			return result.Of(strconv.Atoi(s))
		}

		It("should chain Ok", func() {
			r := result.Ok("69").Bind(parse)

			Expect(r).To(Equal(result.Ok(69)))
		})

		It("should return the error from the function", func() {
			r := result.Ok("mango").Bind(parse)

			Expect(r.Err()).To(MatchError(strconv.ErrSyntax))
		})

		It("should stop at the first error", func() {
			r := result.Error[string](errBoom).Bind(func(string) result.Result[int] {
				Fail("bind called")
				return result.Ok(0)
			})

			Expect(r.Err()).To(MatchError(errBoom))
		})
	})

	Describe("MapErr", func() {
		It("should map the error", func() {
			r := result.Error[int](errBoom).MapErr(func(err error) error {
				return fmt.Errorf("wrapped: %w", err)
			})

			Expect(r.Err()).To(MatchError("wrapped: boom"))
			Expect(errors.Is(r.Err(), errBoom)).To(BeTrue())
		})

		It("should not call the function for Ok", func() {
			r := result.Ok(69).MapErr(func(err error) error {
				Fail("map called")
				return err
			})

			Expect(r).To(Equal(result.Ok(69)))
		})
	})

	Describe("Maybe", func() {
		It("should be Some for Ok", func() {
			Expect(result.Ok(69).Maybe()).To(Equal(maybe.Some(69)))
		})

		It("should be None for an error", func() {
			Expect(result.Error[int](errBoom).Maybe()).To(Equal(maybe.None[int]()))
		})
	})

	Describe("Collect", func() {
		It("should collect every Ok value", func() {
			seq := slices.Values([]result.Result[int]{result.Ok(1), result.Ok(2)})

			Expect(result.Collect(seq)).To(Equal(result.Ok([]int{1, 2})))
		})

		It("should stop at the first error", func() {
			seen := 0
			seq := func(yield func(result.Result[int]) bool) {
				for _, r := range []result.Result[int]{result.Ok(1), result.Error[int](errBoom), result.Ok(3)} {
					seen++
					if !yield(r) {
						return
					}
				}
			}

			r := result.Collect(seq)

			Expect(r.Err()).To(MatchError(errBoom))
			Expect(seen).To(Equal(2))
		})

		It("should be Ok and empty for an empty sequence", func() {
			v, err := result.Collect(slices.Values([]result.Result[int]{})).Get()

			Expect(err).NotTo(HaveOccurred())
			Expect(v).To(BeEmpty())
		})
	})

	Describe("String", func() {
		It("should format Ok", func() {
			Expect(result.Ok(69).String()).To(Equal("Ok(69)"))
		})

		It("should format an error", func() {
			Expect(result.Error[int](errBoom).String()).To(Equal("Error(boom)"))
		})
	})
})
