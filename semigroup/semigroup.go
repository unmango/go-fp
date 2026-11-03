// Package semigroup describes values that can be combined associatively.
package semigroup

// Semigroup combines two values of T into one.
// Combine must be associative: Combine(Combine(a, b), c) == Combine(a, Combine(b, c)).
type Semigroup[T any] interface {
	Combine(a, b T) T
}

// Combinable is a type that knows how to combine itself with another value of its own type.
// It is used as a self-referential constraint:
//
//	func Concat[T semigroup.Combinable[T]](xs ...T) T
type Combinable[T Combinable[T]] interface {
	Combine(T) T
}

// Func adapts a plain function to the Semigroup interface.
type Func[T any] func(a, b T) T

// Combine calls f(a, b).
func (f Func[T]) Combine(a, b T) T {
	return f(a, b)
}

// Lift returns a Semigroup that combines with f.
func Lift[T any](f func(a, b T) T) Semigroup[T] {
	return Func[T](f)
}

// Method returns a Semigroup that combines with T's own Combine method.
func Method[T Combinable[T]]() Semigroup[T] {
	return Func[T](func(a, b T) T {
		return a.Combine(b)
	})
}
