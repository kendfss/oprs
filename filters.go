package oprs

import "github.com/kendfss/rules"

// IsEven returns true iff the argument is an even number
func IsEven[T rules.Integer](a T) bool {
	return a%2 == 0
}

// IsOdd returns true iff the argument is an odd number
func IsOdd[T rules.Integer](a T) bool {
	return a%2 == 1
}

// IsTrue checks if a bool is true, for declarative testing
func IsTrue(b bool) bool {
	return b == true
}

// IsFalse checks if a bool is false, for declarative testing
func IsFalse(b bool) bool {
	return b == false
}

// Is creates an equivalence predicate for a value
func Is[T comparable](val T) func(T) bool {
	return func(arg T) bool {
		return arg == val
	}
}

// Isnt creates an anti-equivalence predicate for a value
func Isnt[T comparable](val T) func(T) bool {
	return func(arg T) bool {
		return arg != val
	}
}

// IsFunc wraps an equivalence predicate for a value of non-comparable type
func IsFunc[T any](val T, f func(T, T) bool) func(T) bool {
	return func(arg T) bool {
		return f(arg, val)
	}
}
