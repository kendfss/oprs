package oprs

import (
	"fmt"
	"reflect"
	"unicode/utf8"

	"github.com/kendfss/rules"
)

// ToBool returns true iff the argument is non-zero
// a string is zero iff it's length is zero
func ToBool[T comparable](arg T) bool {
	return arg != *new(T)
}

// ToBoolIter converts indexable types to booleans
// if it's length is zero, it is false
func ToBoolLener[Seq rules.Lener[K, V], K comparable, V any](arg Seq) bool {
	return len(arg) != 0
}

// ToUint8 casts the argument to a uint8 value
func ToUint8[T rules.Integer](arg T) uint8 {
	return uint8(arg)
}

// ToUint16 casts the argument to a uint16 value
func ToUint16[T rules.Integer](arg T) uint16 {
	return uint16(arg)
}

// ToUint32 casts the argument to a uint32 value
func ToUint32[T rules.Integer](arg T) uint32 {
	return uint32(arg)
}

// ToUint64 casts the argument to a uint64 value
func ToUint64[T rules.Integer](arg T) uint64 {
	return uint64(arg)
}

// ToInt8 casts the argument to a int8 value
func ToInt8[T rules.Integer](arg T) int8 {
	return int8(arg)
}

// ToInt16 casts the argument to a int16 value
func ToInt16[T rules.Integer](arg T) int16 {
	return int16(arg)
}

// ToInt32 casts the argument to a int32 value
func ToInt32[T rules.Integer](arg T) int32 {
	return int32(arg)
}

// ToInt64 casts the argument to a int64 value
func ToInt64[T rules.Integer](arg T) int64 {
	return int64(arg)
}

// ToFloat32 casts the argument to a float32 value
func ToFloat32[T rules.Integer](arg T) float32 {
	return float32(arg)
}

// ToFloat64 casts the argument to a float64 value
func ToFloat64[T rules.Integer](arg T) float64 {
	return float64(arg)
}

// ToComplex64 casts the argument to a complex64 value
func ToComplex64[T rules.Real](arg T) complex64 {
	return complex(float32(arg), 0)
}

// ToComplex128 casts the argument to a complex128 value
func ToComplex128[T rules.Real](arg T) complex128 {
	return complex(float64(arg), 0)
}

// ToString casts the argument to a string value
func ToString[T any](arg T) string {
	return fmt.Sprintf("%v", arg)
}

// ToString2 casts the argument to a string2 value with type information
func ToString2[T any](arg T) string {
	return fmt.Sprintf("%#v", arg)
}

// ToInt casts the argument to a int value
func ToInt[T rules.Integer](arg T) int {
	return int(arg)
}

// ToUint casts the argument to a uint value
func ToUint[T rules.Integer](arg T) uint {
	return uint(arg)
}

// ToUintptr casts the argument to a uintptr value
func ToUintptr[T rules.Integer](arg T) uintptr {
	return uintptr(arg)
}

// ToByte casts the argument to a byte value
func ToByte[T rules.Integer](arg T) byte {
	return byte(arg)
}

// ToRune casts the argument to a rune value
func ToRune[T rules.Integer](arg T) rune {
	return rune(arg)
}

// BindOp turns a diadic function of some type and binary operator into a monadic one
// It is like MethodOp, except it targets the right operand
func BindOp[T, U, V, L, R any](f func(T, func(L, R) U) V, op func(L, R) U) func(T) V {
	return func(arg T) V {
		return f(arg, op)
	}
}

// BindVar turns a diadic function of upto two types into a monadic one
// It is like MethodVar, except it targets the right operand
func BindVar[L, R, T any](f func(L, R) T, v func() R) func(L) T {
	return func(arg L) T {
		return f(arg, v())
	}
}

// Bind turns a diadic function of upto two types into a monadic one
// It is like Method, except it targets the right operand
func Bind[L, R, T any](f func(L, R) T, val R) func(L) T {
	return func(arg L) T {
		return f(arg, val)
	}
}

// MethodOp returns a detatched method for some binary operator
// calling the given binary operator with it as the left operand
// similar to BindVal except the value is the left operand
// see BindOp for more info
func MethodOp[T, U, V, L, R any](op func(L, R) U, f func(func(L, R) U, T) V) func(T) V {
	return func(arg T) V {
		return f(op, arg)
	}
}

// MethodVar returns a detatched method for some variable
// calling the given binary operator with it as the left operand
// similar to BindVal except the value is the left operand
// see BindVar for more info
func MethodVar[R, L, T any](v func() L, f func(L, R) T) func(R) T {
	return func(arg R) T {
		return f(v(), arg)
	}
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

// Cast allows you to convert between real number types
func Cast[O, I rules.Real](arg I) O {
	return O(arg)
}

// Real extracts the real part of a complex number
func Real[N rules.Real, C rules.Complex](arg C) N {
	return N(real(complex128(arg)))
}

// Imag extracts the imaginary part of a complex number
func Imag[N rules.Real, C rules.Complex](arg C) N {
	return N(imag(complex128(arg)))
}

// MethodEach creates a collection of unary operators from
// a binary operator and a collection of operands
func MethodEach[L, R, T any](fn func(L, R) T, vals ...L) []func(R) T {
	out := make([]func(R) T, len(vals))
	for i := range out {
		out[i] = Method(vals[i], fn)
	}
	return out
}

// BindEach creates a collection of unary operators from
// a binary operator and a collection of operands
func BindEach[L, R, T any](fn func(L, R) T, vals ...R) []func(L) T {
	out := make([]func(L) T, len(vals))
	for i := range out {
		out[i] = Bind(fn, vals[i])
	}
	return out
}

// Len returns the length of a slice
func Len[T any](arg []T) int {
	return len(arg)
}

// Pipe concatenates two unary operators
func Pipe[L, R, T any](one func(L) R, two func(R) T) func(L) T {
	return func(arg L) T {
		return two(one(arg))
	}
}

func CurryL[L, R, T any](f func(L, R) T) func(L) func(R) T {
	return func(left L) func(R) T {
		return func(right R) T {
			return f(left, right)
		}
	}
}

func CurryLV[L, R, T any](f func(L, ...R) T) func(L) func(...R) T {
	return func(l L) func(...R) T {
		return func(args ...R) T {
			return f(l, args...)
		}
	}
}

func CurryR[L, R, T any](f func(L, R) T) func(R) func(L) T {
	return func(right R) func(L) T {
		return func(left L) T {
			return f(left, right)
		}
	}
}

func CurryRV[L, R, T any](f func(L, ...R) T) func(...R) func(L) T {
	return func(args ...R) func(L) T {
		return func(l L) T {
			return f(l, args...)
		}
	}
}

// Onead transforms a variadic function to a 1-adic
func Onead[I, O any](f func(...I) O) func(I) O {
	return func(arg I) O {
		return f(arg)
	}
}

// func Variad[One, Two, Three any](operator func(One) Two, aggregator func(Two, Two) Three) func(...One) Three {
// 	rule := Integrate(operator)
// 	return func(args ...One) (out Three) {
// 		if len(args) > 0 {
// 			terms := rule(args)
// 			for i, term := range terms[1:] {
// 				out = aggregator(terms[i], term)
// 			}
// 		}
// 		return out
// 	}
// }

// Slicead transforms a variadic function to one that accepts a slice
func Slicead[I, O any](f func(...I) O) func([]I) O {
	return func(arg []I) O {
		return f(arg...)
	}
}

// Variad transforms a variadic function to one that accepts a slice
func Variad[I, O any](f func([]I) O) func(...I) O {
	return func(arg ...I) O {
		return f(arg)
	}
}

// Nothing is a place holder for imports
// so that you can run a script even though you've removed all references to an import
func Nothing(args ...any) {}

// Returner is a place holder for pipelines that depend on casters
func Returner[T any](arg T) T {
	return arg
}

// Integrate transforms a function between two types into one between arrays of their instances
func Integrate[I, O any](f func(I) O) func([]I) []O {
	return func(arg []I) []O {
		out := make([]O, len(arg))
		for i, e := range arg {
			out[i] = f(e)
		}
		return out
	}
}

func WrapAny[T, L, R any](fn func(L, R)) func(L, R) *T {
	return func(l L, r R) *T {
		fn(l, r)
		return new(T)
	}
}

// Flip changes the order of arguments for a binary operator
func Flip[L, R, T any](op func(L, R) T) func(R, L) T {
	return func(r R, l L) T {
		return op(l, r)
	}
}

func Before[I, O any](fn func(I) O, op func()) func(I) O {
	return func(i I) O {
		op()
		return fn(i)
	}
}

func After[I, O any](fn func(I) O, op func()) func(I) O {
	return func(i I) O {
		out := fn(i)
		op()
		return out
	}
}

// func RunMethod[Int, Bool, Slice any](arg Int, una func(Int) Slice, bin func(Int, Int) Bool, last func(func(Int) Bool, Slice) Slice) Slice {
// // compute prime factors
// 	return last(
// 		Method(arg, bin),
// 		una(arg),
// 	)
// }

func AreAll[I any](f func(I) bool) func(...I) bool {
	return func(args ...I) (out bool) {
		if len(args) > 0 {
			out = f(args[0])
			for _, arg := range args[1:] {
				out = f(arg) && out
			}
		}
		return out
	}
}

func Must[I, O any](f func(I) (O, error)) func(I) O {
	return func(arg I) O {
		o, err := f(arg)
		if err != nil {
			panic(err)
		}
		return o
	}
}

// Assert uses type assertion to convert a given
// interface{} into a literal of the given T
func Assert[T any](arg any) *T {
	val, ok := arg.(T)
	if ok {
		return &val
	}
	return nil
}

// Typeis checks if a given value has wanted type
func Typeis[want, have any](arg have) bool {
	return reflect.TypeOf(*new(want)) == reflect.TypeOf(arg)
}

// Chars returns the byes or runes associated with a string
func Chars[char byte | rune](w string) (out []char) {
	bytes := []byte(w)
	if Typeis[rune](*new(char)) {
		for len(bytes) > 0 {
			r, size := utf8.DecodeRune(bytes)
			out = append(out, char(r))
			bytes = bytes[size:]
		}
	} else {
		out = make([]char, len(w))
		for i, b := range bytes {
			out[i] = char(b)
		}
	}
	return
}
