// Package pipe composes functions left to right.
package pipe

// Pipe returns a function that applies fa and then fb.
//
//	pipe.Pipe(strings.TrimSpace, strings.ToUpper)
func Pipe[A, B, C any](fa func(A) B, fb func(B) C) func(A) C {
	return func(a A) C {
		return fb(fa(a))
	}
}

// Pipe3 returns a function that applies fa, fb and fc in order.
func Pipe3[A, B, C, D any](fa func(A) B, fb func(B) C, fc func(C) D) func(A) D {
	return func(a A) D {
		return fc(fb(fa(a)))
	}
}

// Pipe4 returns a function that applies fa through fd in order.
func Pipe4[A, B, C, D, E any](fa func(A) B, fb func(B) C, fc func(C) D, fd func(D) E) func(A) E {
	return func(a A) E {
		return fd(fc(fb(fa(a))))
	}
}

// Pipe5 returns a function that applies fa through fe in order.
func Pipe5[A, B, C, D, E, F any](fa func(A) B, fb func(B) C, fc func(C) D, fd func(D) E, fe func(E) F) func(A) F {
	return func(a A) F {
		return fe(fd(fc(fb(fa(a)))))
	}
}

// Pipe6 returns a function that applies fa through ff in order.
func Pipe6[A, B, C, D, E, F, G any](fa func(A) B, fb func(B) C, fc func(C) D, fd func(D) E, fe func(E) F, ff func(F) G) func(A) G {
	return func(a A) G {
		return ff(fe(fd(fc(fb(fa(a))))))
	}
}
