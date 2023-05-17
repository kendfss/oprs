package oprs

// Ternary reduces the if-else statement to a one-liner
func Ternary[T any](pred bool, whenTrue, whenFalse T) T {
	if pred {
		return whenTrue
	}
	return whenFalse
}

func Pred[T, U any](pred func(T) bool, whenTrue, whenFalse U) func(T) U {
	return func(t T) U {
		if pred(t) {
			return whenTrue
		}
		return whenFalse
	}
}

// DropRight strips a function of its right-most return value
func DropRight[I, L, R any](fn func(I) (L, R)) func(I) L {
	return func(i I) L {
		l, _ := fn(i)
		return l
	}
}

// DropLeft strips a function of its left-most return value
func DropLeft[I, L, R any](fn func(I) (L, R)) func(I) R {
	return func(i I) R {
		_, r := fn(i)
		return r
	}
}

// Not returns the negation of a predicate
func Not[T any](pred func(T) bool) func(T) bool {
	return func(arg T) bool {
		return !pred(arg)
	}
}

// And returns the && gate for non-boolean types
func And[L, R any](one func(L) bool, two func(R) bool) func(L, R) bool {
	return func(l L, r R) bool {
		return one(l) && two(r)
	}
}

// Or returns the || gate for non-boolean types
func Or[L, R any](one func(L) bool, two func(R) bool) func(L, R) bool {
	return func(l L, r R) bool {
		return one(l) || two(r)
	}
}

// Or returns the && gate for non-boolean types
func Xor[L, R any](one func(L) bool, two func(R) bool) func(L, R) bool {
	return func(l L, r R) bool {
		return !(one(l) && two(r))
	}
}

// Both offers the AND gate
func Both[T any](one, two func(T) bool) func(T) bool {
	return func(arg T) bool {
		return one(arg) && two(arg)
	}
}

// Neither offers the !AND gate
func Neither[T any](one, two func(T) bool) func(T) bool {
	return func(arg T) bool {
		return !(one(arg) || two(arg))
	}
}

// Either offers the OR gate
func Either[T any](one, two func(T) bool) func(T) bool {
	return func(arg T) bool {
		return one(arg) || two(arg)
	}
}

// NotBoth offers the XOR gate
func NotBoth[T any](one, two func(T) bool) func(T) bool {
	return Not(Both(one, two))
}

// All AND-concatenates a sequence of boolean operators with of a shared type
// If no argument is given, the returned predicate will check if its argument's
// pointer is nil
func All[T any](preds ...func(T) bool) func(T) bool {
	switch len(preds) {
	case 0:
		return func(arg T) bool { return nil == &arg }
	case 1:
		return preds[0]
	default:
		out := preds[0]
		for _, pred := range preds[1:] {
			out = Both(out, pred)
		}
		return out
	}
}

// Any OR-concatenates a sequence of boolean operators with of a shared type
// If no argument is given, the returned predicate will check if its argument's
// pointer is nil
func Any[T any](preds ...func(T) bool) func(T) bool {
	switch len(preds) {
	case 0:
		return func(arg T) bool { return nil == &arg }
	case 1:
		return preds[0]
	default:
		out := preds[0]
		for _, pred := range preds[1:] {
			out = Either(out, pred)
		}
		return out
	}
}

// Any XOR-concatenates a sequence of boolean operators with of a shared type
// If no argument is given, the returned predicate will check if its argument's
// pointer is nil
func One[T any](preds ...func(T) bool) func(T) bool {
	switch len(preds) {
	case 0:
		return func(arg T) bool { return nil == &arg }
	case 1:
		return preds[0]
	default:
		out := preds[0]
		for _, pred := range preds[1:] {
			out = NotBoth(out, pred)
		}
		return out
	}
}
