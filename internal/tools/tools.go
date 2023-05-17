package tools

import (
	"fmt"
	"math/rand"

	"github.com/kendfss/rules"
)

const (
	nTests = 10
	nItems = 10
	nMax   = 10
)

// Add returns the sum of two values
func Add[T rules.Ordered](a, b T) T {
	return a + b
}

// Eq compares two values for equality
func Eq[T comparable](a, b T) bool {
	return a == b
}

// Method returns a detatched method for some value
// calling the given binary operator with it as the left operand
// similar to Bind except the value is the left operand
// see Bind for more info
func Method[L, R, T any](val L, f func(L, R) T) func(R) T {
	return func(arg R) T {
		return f(val, arg)
	}
}

func RandSign[T rules.SignedNumber](val T) T {
	if rand.Intn(nMax)%2 == 0 {
		return val * -1
	}
	return val
}

func Randints(count int) []int {
	out := make([]int, count)
	for i := range out {
		out[i] = rand.Intn(nMax)
	}
	return out
}

// Send calls an impure function on each element of a slice
func Send[T any](f func(T), args []T) {
	for _, arg := range args {
		f(arg)
	}
}

// Index returns the Index of the first occurrence of v in s,
// or -1 if not present.
func Index[E comparable](val E, s []E) int {
	return IndexFunc(Eq[E], val, s)
}

// Contains reports whether v is present in s.
func Contains[E comparable](s []E, v E) bool {
	return Index(v, s) >= 0
}

// IndexFunc returns the first index i satisfying f(s[i]),
// or -1 if none do.
func IndexFunc[E any](eq func(E, E) bool, val E, s []E) int {
	return IndexPred(Method(val, eq), s)
}

// IndexPred returns the first index i satisfying f(s[i]),
// or -1 if none do.
func IndexPred[E any](eq func(E) bool, s []E) int {
	for i, v := range s {
		if eq(v) {
			return i
		}
	}
	return -1
}

// FilterPred returns a slice featuring all elements of the incident that satisfy the given predicate
func FilterPred[E any](f func(E) bool, args []E) (out []E) {
	for _, e := range args {
		if f(e) {
			out = append(out, e)
		}
	}
	return out
}

// Uptoch returns an iterator whose content depends on the number of arguments as follows
// 		# of args 	|| 	behaviour
//	 		 1	 	|| 	stop
//		 	 2	 	|| 	start, stop
//		 	 3	 	|| 	start, stop, step
//         else 	|| 	error
func Uptoch[T rules.Real](args ...T) (chan T, error) {
	switch len(args) {
	case 1:
		return Uptoch(0, args[0], 1)
	case 2:
		return Uptoch(args[0], args[1], 1)
	case 3:
		out := make(chan T)
		go func() {
			start, stop, delta := args[0], args[1], args[2]
			for stop-delta >= start {
				out <- start
				start += delta
			}
		}()
		return out, nil
	case 0:
		return nil, ErrUptoUnder(args)
	default:
		return nil, ErrUptoOver(args)
	}
}

// MustUptoch returns an iterator whose behaviour is equivalent to that of Range
func MustUptoch[T rules.Real](args ...T) chan T {
	out, err := Uptoch(args...)
	if err == nil {
		return out
	}
	panic(err)
}

// Upto returns an iterator whose content depends on the number of arguments as follows
// 		# of args 	|| 	behaviour
//	 		 1	 	|| 	stop
//		 	 2	 	|| 	start, stop
//		 	 3	 	|| 	start, stop, step
//         else 	|| 	error
func Upto[T rules.Real](args ...T) ([]T, error) {
	switch len(args) {
	case 1:
		return Upto(0, args[0], 1)
	case 2:
		return Upto(args[0], args[1], 1)
	case 3:
		out := make([]T, 0)
		start, stop, delta := args[0], args[1], args[2]
		for stop-delta >= start {
			out = append(out, start)
			start += delta
		}
		return out, nil
	case 0:
		return nil, ErrUptoUnder(args)
	default:
		return nil, ErrUptoOver(args)
	}
}

// MustUpto returns an iterator whose behaviour is equivalent to that of Range
func MustUpto[T rules.Real](args ...T) []T {
	out, err := Upto(args...)
	if err == nil {
		return out
	}
	panic(err)
}

func ErrUptoUnder[T any](args []T) error {
	return fmt.Errorf("oprs.upto: not enough args (%d). want 1, 2, or 3", len(args))
}

func ErrUptoOver[T any](args []T) error {
	return fmt.Errorf("oprs.upto: too many args (%d). want 1, 2, or 3", len(args))
}

func Consume[T any](ch chan T) (out []T) {
	for e := range ch {
		out = append(out, e)
	}
	return
}
