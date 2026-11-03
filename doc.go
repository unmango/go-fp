// Package fp is the root of go.unmango.dev/fp, a small functional programming library for Go.
//
// The library is split into packages by concept:
//
//   - [go.unmango.dev/fp/maybe]: a value that may be absent.
//   - [go.unmango.dev/fp/result]: a value or an error.
//   - [go.unmango.dev/fp/either]: one of two typed values.
//   - [go.unmango.dev/fp/lazy]: a value computed on first use.
//   - [go.unmango.dev/fp/pipe]: left-to-right function composition.
//   - [go.unmango.dev/fp/semigroup] and [go.unmango.dev/fp/monoid]: combining values and folding sequences.
//
// Each type converts to and from the matching ordinary Go shape,
// such as (T, bool), (T, error), *T and iter.Seq, so it can be used at the edges of existing code.
package fp
