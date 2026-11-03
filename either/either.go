// Package either provides Either, a value that holds one of two types.
package either

import "fmt"

// Either holds a value of type L (Left) or a value of type R (Right).
// The zero value is Left with the zero value of L.
// Either values of comparable types L and R can be compared with ==.
type Either[L, R any] struct {
	left  L
	right R
	isR   bool
}

// Left returns an Either holding the left value l.
// The right type comes first so it can be given on its own:
//
//	e := either.Left[string](69) // Either[int, string]
func Left[R, L any](l L) Either[L, R] {
	return Either[L, R]{left: l}
}

// Right returns an Either holding the right value r.
//
//	e := either.Right[int]("mango") // Either[int, string]
func Right[L, R any](r R) Either[L, R] {
	return Either[L, R]{right: r, isR: true}
}

// Left returns the left value and true, or the zero value and false when e is Right.
func (e Either[L, R]) Left() (L, bool) {
	return e.left, !e.isR
}

// Right returns the right value and true, or the zero value and false when e is Left.
func (e Either[L, R]) Right() (R, bool) {
	return e.right, e.isR
}

// IsLeft reports whether e holds a left value.
func (e Either[L, R]) IsLeft() bool {
	return !e.isR
}

// IsRight reports whether e holds a right value.
func (e Either[L, R]) IsRight() bool {
	return e.isR
}

// MapLeft applies f to a left value. f is not called when e is Right.
func (e Either[L, R]) MapLeft[LB any](f func(L) LB) Either[LB, R] {
	if e.isR {
		return Right[LB](e.right)
	}
	return Left[R](f(e.left))
}

// MapRight applies f to a right value. f is not called when e is Left.
func (e Either[L, R]) MapRight[RB any](f func(R) RB) Either[L, RB] {
	if !e.isR {
		return Left[RB](e.left)
	}
	return Right[L](f(e.right))
}

// BindLeft applies f to a left value and returns its result.
// f is not called when e is Right.
func (e Either[L, R]) BindLeft[LB any](f func(L) Either[LB, R]) Either[LB, R] {
	if e.isR {
		return Right[LB](e.right)
	}
	return f(e.left)
}

// BindRight applies f to a right value and returns its result.
// f is not called when e is Left.
func (e Either[L, R]) BindRight[RB any](f func(R) Either[L, RB]) Either[L, RB] {
	if !e.isR {
		return Left[RB](e.left)
	}
	return f(e.right)
}

// Swap turns a Left into a Right and a Right into a Left.
func (e Either[L, R]) Swap() Either[R, L] {
	if e.isR {
		return Left[L](e.right)
	}
	return Right[R](e.left)
}

// Fold collapses e into a single value by calling onLeft or onRight.
func (e Either[L, R]) Fold[U any](onLeft func(L) U, onRight func(R) U) U {
	if e.isR {
		return onRight(e.right)
	}
	return onLeft(e.left)
}

// String formats e as Left(v) or Right(v).
func (e Either[L, R]) String() string {
	if e.isR {
		return fmt.Sprintf("Right(%v)", e.right)
	}
	return fmt.Sprintf("Left(%v)", e.left)
}
