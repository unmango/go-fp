// Package lazy provides Lazy, a value computed when it is first needed.
package lazy

import "sync"

// Lazy is a deferred computation of a T.
// The zero value (nil) yields the zero value of T.
type Lazy[T any] func() T

// Lazy2 is a deferred computation of a pair of values, such as a (T, error).
// The zero value (nil) yields the zero values of T and V.
type Lazy2[T, V any] func() (T, V)

// Of returns a Lazy that always returns x.
func Of[T any](x T) Lazy[T] {
	return func() T {
		return x
	}
}

// Once returns a Lazy that calls f the first time it is forced and returns the same value after that.
// It is safe to force from several goroutines.
func Once[T any](f func() T) Lazy[T] {
	return sync.OnceValue(f)
}

// Once2 is Once for a function returning two values.
func Once2[T, V any](f func() (T, V)) Lazy2[T, V] {
	return sync.OnceValues(f)
}

// Get forces l and returns its value.
func (l Lazy[T]) Get() (v T) {
	if l == nil {
		return v
	}
	return l()
}

// Map returns a Lazy that applies f to the value of l when forced.
// Neither l nor f is called until then.
// The result is not memoized: wrap it in Once to compute it only once.
func (l Lazy[T]) Map[U any](f func(T) U) Lazy[U] {
	return func() U {
		return f(l.Get())
	}
}

// Get forces l and returns its values.
func (l Lazy2[T, V]) Get() (t T, v V) {
	if l == nil {
		return t, v
	}
	return l()
}
