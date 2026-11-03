package pipe_test

import (
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"go.unmango.dev/fp/pipe"
)

var (
	inc     = func(x int) int { return x + 1 }
	double  = func(x int) int { return x * 2 }
	itoa    = strconv.Itoa
	exclaim = func(s string) string { return s + "!" }
	upper   = strings.ToUpper
)

var _ = Describe("Pipe", func() {
	It("should compose two functions left to right", func() {
		Expect(pipe.Pipe(inc, double)(1)).To(Equal(4))
	})

	It("should infer types across steps", func() {
		var f func(int) string = pipe.Pipe(double, itoa)

		Expect(f(21)).To(Equal("42"))
	})

	It("should compose three functions", func() {
		Expect(pipe.Pipe3(inc, double, itoa)(1)).To(Equal("4"))
	})

	It("should compose four functions", func() {
		Expect(pipe.Pipe4(inc, double, itoa, exclaim)(1)).To(Equal("4!"))
	})

	It("should compose five functions", func() {
		Expect(pipe.Pipe5(inc, double, inc, itoa, exclaim)(1)).To(Equal("5!"))
	})

	It("should compose six functions", func() {
		Expect(pipe.Pipe6(inc, double, inc, itoa, exclaim, upper)(1)).To(Equal("5!"))
	})
})
