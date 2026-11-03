// Package monoid describes semigroups with an identity value, and folds sequences with them.
package monoid

import (
	"iter"
	"slices"

	"go.unmango.dev/fp/semigroup"
)

// Monoid is a Semigroup with an identity value.
// Empty must satisfy Combine(Empty(), x) == x and Combine(x, Empty()) == x.
type Monoid[T any] interface {
	semigroup.Semigroup[T]
	Empty() T
}

// Number is the set of built-in numeric types, excluding complex numbers.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type monoid[T any] struct {
	semigroup.Semigroup[T]
	empty T
}

func (m monoid[T]) Empty() T {
	return m.empty
}

// From returns a Monoid combining with s and using empty as the identity.
func From[T any](empty T, s semigroup.Semigroup[T]) Monoid[T] {
	return monoid[T]{s, empty}
}

// Lift returns a Monoid combining with f and using empty as the identity.
func Lift[T any](empty T, f func(a, b T) T) Monoid[T] {
	return From(empty, semigroup.Lift(f))
}

// Method returns a Monoid combining with T's own Combine method and using empty as the identity.
func Method[T semigroup.Combinable[T]](empty T) Monoid[T] {
	return From(empty, semigroup.Method[T]())
}

// Sum adds numbers, with 0 as the identity.
// Floating-point addition is only approximately associative, so folding floats in a different order can round differently.
func Sum[T Number]() Monoid[T] {
	return Lift(0, func(a, b T) T { return a + b })
}

// Product multiplies numbers, with 1 as the identity.
// Floating-point multiplication is only approximately associative, as with Sum.
func Product[T Number]() Monoid[T] {
	return Lift(1, func(a, b T) T { return a * b })
}

// String concatenates strings, with "" as the identity.
// Each Combine copies both strings, so folding many strings is quadratic; prefer strings.Join for large inputs.
func String() Monoid[string] {
	return Lift("", func(a, b string) string { return a + b })
}

// Slice appends slices, with nil as the identity.
// The result never shares a backing array with a, so each Combine copies both slices
// and folding many slices is quadratic; prefer slices.Concat for large inputs.
func Slice[T any]() Monoid[[]T] {
	return Lift(nil, func(a, b []T) []T {
		if a == nil {
			if b == nil {
				return nil
			}
			return append([]T{}, b...)
		}
		return append(slices.Clone(a), b...)
	})
}

// All combines booleans with &&, with true as the identity.
func All() Monoid[bool] {
	return Lift(true, func(a, b bool) bool { return a && b })
}

// Any combines booleans with ||, with false as the identity.
func Any() Monoid[bool] {
	return Lift(false, func(a, b bool) bool { return a || b })
}

// Fold combines every value of seq with m, starting from m.Empty().
func Fold[T any](m Monoid[T], seq iter.Seq[T]) T {
	acc := m.Empty()
	for v := range seq {
		acc = m.Combine(acc, v)
	}
	return acc
}

// FoldMap applies f to every value of seq and combines the results with m.
func FoldMap[A, T any](m Monoid[T], seq iter.Seq[A], f func(A) T) T {
	acc := m.Empty()
	for a := range seq {
		acc = m.Combine(acc, f(a))
	}
	return acc
}
