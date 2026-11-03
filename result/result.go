// Package result provides Result, a value paired with an error, as returned by most Go functions.
package result

import (
	"fmt"
	"iter"

	"go.unmango.dev/fp/maybe"
)

// Result holds either a value (Ok) or an error.
// The zero value is Ok with the zero value of T.
type Result[T any] struct {
	value T
	err   error
}

// Ok returns a successful Result holding v.
func Ok[T any](v T) Result[T] {
	return Result[T]{value: v}
}

// Error returns a failed Result holding err.
// A nil err gives Ok with the zero value of T.
func Error[T any](err error) Result[T] {
	return Result[T]{err: err}
}

// Errorf returns a failed Result holding fmt.Errorf(format, a...).
func Errorf[T any](format string, a ...any) Result[T] {
	return Error[T](fmt.Errorf(format, a...))
}

// Of converts a (value, error) pair into a Result.
// When err is not nil the value is dropped.
//
//	r := result.Of(strconv.Atoi(s))
func Of[T any](v T, err error) Result[T] {
	if err != nil {
		return Error[T](err)
	}
	return Ok(v)
}

// FromMaybe returns Ok for Some and maybe.ErrNone for None.
func FromMaybe[T any](m maybe.Maybe[T]) Result[T] {
	if v, ok := m.Get(); ok {
		return Ok(v)
	}
	return Error[T](maybe.ErrNone)
}

// Collect gathers the values of seq into a slice, stopping at the first error.
func Collect[T any](seq iter.Seq[Result[T]]) Result[[]T] {
	values := []T{}
	for r := range seq {
		if r.err != nil {
			return Error[[]T](r.err)
		}
		values = append(values, r.value)
	}
	return Ok(values)
}

// Get returns the value and error as an ordinary Go pair.
func (r Result[T]) Get() (T, error) {
	return r.value, r.err
}

// Err returns the held error, or nil for Ok.
func (r Result[T]) Err() error {
	return r.err
}

// IsOk reports whether r holds a value.
func (r Result[T]) IsOk() bool {
	return r.err == nil
}

// IsErr reports whether r holds an error.
func (r Result[T]) IsErr() bool {
	return r.err != nil
}

// OrElse returns the held value, or fallback when r holds an error.
func (r Result[T]) OrElse(fallback T) T {
	if r.err != nil {
		return fallback
	}
	return r.value
}

// Map applies f to the held value. f is not called when r holds an error.
func (r Result[T]) Map[U any](f func(T) U) Result[U] {
	if r.err != nil {
		return Error[U](r.err)
	}
	return Ok(f(r.value))
}

// Bind applies f to the held value and returns its result.
// f is not called when r holds an error.
func (r Result[T]) Bind[U any](f func(T) Result[U]) Result[U] {
	if r.err != nil {
		return Error[U](r.err)
	}
	return f(r.value)
}

// MapErr applies f to the held error. f is not called for Ok.
// If f returns nil, the result is Ok with the zero value of T.
func (r Result[T]) MapErr(f func(error) error) Result[T] {
	if r.err == nil {
		return r
	}
	return Error[T](f(r.err))
}

// Maybe returns Some for Ok and None when r holds an error, dropping the error.
func (r Result[T]) Maybe() maybe.Maybe[T] {
	return maybe.FromOk(r.value, r.err == nil)
}

// String formats r as Ok(v) or Error(err).
func (r Result[T]) String() string {
	if r.err != nil {
		return fmt.Sprintf("Error(%v)", r.err)
	}
	return fmt.Sprintf("Ok(%v)", r.value)
}
