// Package maybe provides Maybe, a value that is either present (Some) or absent (None).
package maybe

import (
	"errors"
	"fmt"
	"iter"
)

// ErrNone is the error used when a None is converted into a value that requires an error,
// such as a result.Result.
var ErrNone = errors.New("none")

// Maybe holds either a value (Some) or nothing (None).
// The zero value is None.
// Maybe values of a comparable type T can be compared with ==.
type Maybe[T any] struct {
	value T
	ok    bool
}

// Some returns a Maybe holding v.
func Some[T any](v T) Maybe[T] {
	return Maybe[T]{v, true}
}

// None returns a Maybe holding nothing.
func None[T any]() Maybe[T] {
	return Maybe[T]{}
}

// FromPtr returns None for a nil pointer and Some with a copy of the pointed-to value otherwise.
func FromPtr[T any](p *T) Maybe[T] {
	if p == nil {
		return None[T]()
	}
	return Some(*p)
}

// FromOk converts a comma-ok pair into a Maybe.
//
//	v, ok := m[key]
//	value := maybe.FromOk(v, ok)
func FromOk[T any](v T, ok bool) Maybe[T] {
	if !ok {
		return None[T]()
	}
	return Some(v)
}

// Get returns the held value and true, or the zero value and false for None.
func (m Maybe[T]) Get() (T, bool) {
	return m.value, m.ok
}

// IsSome reports whether m holds a value.
func (m Maybe[T]) IsSome() bool {
	return m.ok
}

// IsNone reports whether m holds nothing.
func (m Maybe[T]) IsNone() bool {
	return !m.ok
}

// Ptr returns a pointer to a copy of the held value, or nil for None.
func (m Maybe[T]) Ptr() *T {
	if !m.ok {
		return nil
	}
	return new(m.value)
}

// OrElse returns the held value, or fallback for None.
func (m Maybe[T]) OrElse(fallback T) T {
	if !m.ok {
		return fallback
	}
	return m.value
}

// OrElseGet returns the held value, or the result of calling fallback for None.
// fallback is only called for None.
func (m Maybe[T]) OrElseGet(fallback func() T) T {
	if !m.ok {
		return fallback()
	}
	return m.value
}

// Map applies f to the held value. f is not called for None.
//
//	maybe.Some(2).Map(strconv.Itoa) // Some("2")
func (m Maybe[T]) Map[U any](f func(T) U) Maybe[U] {
	if !m.ok {
		return None[U]()
	}
	return Some(f(m.value))
}

// Bind applies f to the held value and returns its result. f is not called for None.
func (m Maybe[T]) Bind[U any](f func(T) Maybe[U]) Maybe[U] {
	if !m.ok {
		return None[U]()
	}
	return f(m.value)
}

// Filter returns m if it holds a value that satisfies pred, and None otherwise.
func (m Maybe[T]) Filter(pred func(T) bool) Maybe[T] {
	if !m.ok || !pred(m.value) {
		return None[T]()
	}
	return m
}

// All returns a sequence yielding the held value once, or nothing for None.
func (m Maybe[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		if m.ok {
			yield(m.value)
		}
	}
}

// String formats m as Some(v) or None.
func (m Maybe[T]) String() string {
	if !m.ok {
		return "None"
	}
	return fmt.Sprintf("Some(%v)", m.value)
}
