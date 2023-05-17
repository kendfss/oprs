package oprs

import (
	"unsafe"

	"github.com/kendfss/rules"
)

// Denull sets the value of a pointer if it's current value is zero-like
// See ToBool
func Denull[T comparable](ptr *T, val T) {
	if !ToBool(*ptr) {
		*ptr = val
	}
}

// DenullLener sets the value of a pointer if it's current value is zero-like
// See ToBoolLener
func DenullLener[Seq rules.Lener[K, V], K comparable, V any](ptr *Seq, val Seq) {
	if !ToBoolLener[Seq, K, V](*ptr) {
		*ptr = val
	}
}

// Clone duplicates an object in memory
func Clone[T any](arg *T) *T {
	out := new(T)
	*out = *(*T)(unsafe.Pointer(&arg))
	return out
}

// Value returns the value to the given pointer
func Value[T any](p *T) T {
	return *p
}

// Pointer returns a pointer to the given value
func Pointer[T any](p T) *T {
	return &p
}

// New returns a pointer to a new instance of the given type
func New[T any]() *T {
	return new(T)
}

// Sizeof returns the size, in bits, of an instance of the given type
func Sizeof[T any]() uintptr {
	return unsafe.Sizeof(*new(T)) * 8
}

// IsNil checks if some value is nil
func IsNil[T Niler[*any, any]](arg T) bool {
	return arg == nil
}

// NotNil checks if some value is nil
func NotNil[T Niler[*any, any]](arg T) bool {
	return arg != nil
}

// Msb checks if Most Significant Bit equals 1
func Msb[T rules.Int](arg T) bool {
	return (1 << Sizeof[T]() & arg) == 1
}

// Lsb checks if Least Significant Bit equals 1
func Lsb[T rules.Int](arg T) bool {
	return (1 & arg) == 1
}

// func WedgeFor[T any]() []byte {

// }

// Niler is a constraint satisfied by nullable types
type Niler[K comparable, V any] interface {
	~*V | ~[]V | ~map[K]V | ~chan V //| error // compiler: cannot use error in union (error contains methods)
}

// Capy copies elements from a source slice into a destination slice
// with respect to its capacity.
// The source and destination may overlap.
// It is a no-op if the destination's length equals its capacity.
// Returns the number of elements copied, which will be:
//				0				if src is empty or dst is full
// 		cap(dst) - len(dst); 	unless
// 		cap(dst) - len(src)		is smaller but non-negative
func Capy[T any](dst *[]T, src []T) int {
	n := 0
	for cap(*dst) > len(*dst) && n < len(src) {
		*dst = append(*dst, src[n])
		n++
	}
	return n
}

// Move moves elements from a source slice into a destination slice
// with respect to its capacity.
// The source and destination may overlap.
// It is a no-op if the destination's length equals its capacity.
// Returns the number of elements copied, which will be:
//				0				if src is empty or dst is full
// 		cap(dst) - len(dst); 	unless
// 		cap(dst) - len(src)		is smaller but non-negative
func Move[T any](dst, src *[]T) int {
	n := Capy(dst, *src)
	*src = (*src)[n:]
	return n
}
