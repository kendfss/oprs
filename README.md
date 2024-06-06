oprs
---

Offers tools to support functional programming in go

## Exports
```go
package oprs // import "github.com/kendfss/oprs"


VARIABLES

var DefaultPrinter = os.Stdout

FUNCTIONS

func After[I, O any](fn func(I) O, op func()) func(I) O
func All[T any](preds ...func(T) bool) func(T) bool
    All AND-concatenates a sequence of boolean operators with of a shared type
    If no argument is given, the returned predicate will check if its argument's
    pointer is nil

func And[L, R any](one func(L) bool, two func(R) bool) func(L, R) bool
    And returns the && gate for non-boolean types

func Any[T any](preds ...func(T) bool) func(T) bool
    Any OR-concatenates a sequence of boolean operators with of a shared type If
    no argument is given, the returned predicate will check if its argument's
    pointer is nil

func AreAll[I any](f func(I) bool) func(...I) bool
func Assert[T any](arg any) *T
    Assert uses type assertion to convert a given interface{} into a literal of
    the given T

func Before[I, O any](fn func(I) O, op func()) func(I) O
func Bind[L, R, T any](f func(L, R) T, val R) func(L) T
    Bind turns a diadic function of upto two types into a monadic one It is like
    Method, except it targets the right operand

func BindEach[L, R, T any](fn func(L, R) T, vals ...R) []func(L) T
    BindEach creates a collection of unary operators from a binary operator and
    a collection of operands

func BindOp[T, U, V, L, R any](f func(T, func(L, R) U) V, op func(L, R) U) func(T) V
    BindOp turns a diadic function of some type and binary operator into a
    monadic one It is like MethodOp, except it targets the right operand

func BindVar[L, R, T any](f func(L, R) T, v func() R) func(L) T
    BindVar turns a diadic function of upto two types into a monadic one It is
    like MethodVar, except it targets the right operand

func Both[T any](one, two func(T) bool) func(T) bool
    Both offers the AND gate

func Capy[T any](dst *[]T, src []T) int
    Capy copies elements from a source slice into a destination slice with
    respect to its capacity. The source and destination may overlap. It is a
    no-op if the destination's length equals its capacity. Returns the number of
    elements copied, which will be:

        		0				if src is empty or dst is full
        cap(dst) - len(dst); 	unless
        cap(dst) - len(src)		is smaller but non-negative

func Cast[O, I rules.Real](arg I) O
    Cast allows you to convert between real number types

func Chars[char byte | rune](w string) (out []char)
    Chars returns the byes or runes associated with a string

func Clone[T any](arg *T) *T
    Clone duplicates an object in memory

func CurryL[L, R, T any](f func(L, R) T) func(L) func(R) T
func CurryLV[L, R, T any](f func(L, ...R) T) func(L) func(...R) T
func CurryR[L, R, T any](f func(L, R) T) func(R) func(L) T
func CurryRV[L, R, T any](f func(L, ...R) T) func(...R) func(L) T
func Denull[T comparable](ptr *T, val T)
    Denull sets the value of a pointer if it's current value is zero-like See
    ToBool

func DenullLener[Seq rules.Lener[K, V], K comparable, V any](ptr *Seq, val Seq)
    DenullLener sets the value of a pointer if it's current value is zero-like
    See ToBoolLener

func DropLeft[I, L, R any](fn func(I) (L, R)) func(I) R
    DropLeft strips a function of its left-most return value

func DropRight[I, L, R any](fn func(I) (L, R)) func(I) L
    DropRight strips a function of its right-most return value

func Either[T any](one, two func(T) bool) func(T) bool
    Either offers the OR gate

func Eq[T comparable](a, b T) bool
    Eq compares two values for equality

func Flip[L, R, T any](op func(L, R) T) func(R, L) T
    Flip changes the order of arguments for a binary operator

func Fprinter(w io.Writer) func(...any) (int, error)
    Fprinter returns a closure that prints to the given writer defaults to
    DefaultPrinter if writer is nil

func Fprinterf(w io.Writer, format string) func(...any) (int, error)
    Fprinterf returns a closure that prints a format string to the given writer
    defaults to DefaultPrinter if writer is nil

func Fprinterln(w io.Writer) func(...any) (int, error)
    Fprinterln returns a closure that prints a line to the given writer defaults
    to DefaultPrinter if writer is nil

func Ge[T rules.Ordered](a, b T) bool
    Ge calls ">="

func Gt[T rules.Ordered](a, b T) bool
    Gt calls ">"

func Imag[N rules.Real, C rules.Complex](arg C) N
    Imag extracts the imaginary part of a complex number

func Integrate[I, O any](f func(I) O) func([]I) []O
    Integrate transforms a function between two types into one between arrays of
    their instances

func Is[T comparable](val T) func(T) bool
    Is creates an equivalence predicate for a value

func IsEven[T rules.Integer](a T) bool
    IsEven returns true iff the argument is an even number

func IsFalse(b bool) bool
    IsFalse checks if a bool is false, for declarative testing

func IsFunc[T any](val T, f func(T, T) bool) func(T) bool
    IsFunc wraps an equivalence predicate for a value of non-comparable type

func IsNil[T Niler[*any, any]](arg T) bool
    IsNil checks if some value is nil

func IsOdd[T rules.Integer](a T) bool
    IsOdd returns true iff the argument is an odd number

func IsTrue(b bool) bool
    IsTrue checks if a bool is true, for declarative testing

func Isnt[T comparable](val T) func(T) bool
    Isnt creates an anti-equivalence predicate for a value

func Le[T rules.Ordered](a, b T) bool
    Le calls "<="

func Len[T any](arg []T) int
    Len returns the length of a slice

func Lsb[T rules.Int](arg T) bool
    Lsb checks if Least Significant Bit equals 1

func Lt[T rules.Ordered](a, b T) bool
    Lt calls "<"

func Method[L, R, T any](val L, f func(L, R) T) func(R) T
    Method returns a detatched method for some value calling the given binary
    operator with it as the left operand similar to Bind except the value is the
    left operand see Bind for more info

func MethodEach[L, R, T any](fn func(L, R) T, vals ...L) []func(R) T
    MethodEach creates a collection of unary operators from a binary operator
    and a collection of operands

func MethodOp[T, U, V, L, R any](op func(L, R) U, f func(func(L, R) U, T) V) func(T) V
    MethodOp returns a detatched method for some binary operator calling the
    given binary operator with it as the left operand similar to BindVal except
    the value is the left operand see BindOp for more info

func MethodVar[R, L, T any](v func() L, f func(L, R) T) func(R) T
    MethodVar returns a detatched method for some variable calling the given
    binary operator with it as the left operand similar to BindVal except the
    value is the left operand see BindVar for more info

func Move[T any](dst, src *[]T) int
    Move moves elements from a source slice into a destination slice with
    respect to its capacity. The source and destination may overlap. It is a
    no-op if the destination's length equals its capacity. Returns the number of
    elements copied, which will be:

        		0				if src is empty or dst is full
        cap(dst) - len(dst); 	unless
        cap(dst) - len(src)		is smaller but non-negative

func Msb[T rules.Int](arg T) bool
    Msb checks if Most Significant Bit equals 1

func Must[I, O any](f func(I) (O, error)) func(I) O
func Ne[T comparable](a, b T) bool
    Ne compares two values for inequality

func Neither[T any](one, two func(T) bool) func(T) bool
    Neither offers the !AND gate

func New[T any]() *T
    New returns a pointer to a new instance of the given type

func Not[T any](pred func(T) bool) func(T) bool
    Not returns the negation of a predicate

func NotBoth[T any](one, two func(T) bool) func(T) bool
    NotBoth offers the XOR gate

func NotNil[T Niler[*any, any]](arg T) bool
    NotNil checks if some value is nil

func Nothing(args ...any)
    Nothing is a place holder for imports so that you can run a script even
    though you've removed all references to an import

func One[T any](preds ...func(T) bool) func(T) bool
    Any XOR-concatenates a sequence of boolean operators with of a shared type
    If no argument is given, the returned predicate will check if its argument's
    pointer is nil

func Onead[I, O any](f func(...I) O) func(I) O
    Onead transforms a variadic function to a 1-adic

func Or[L, R any](one func(L) bool, two func(R) bool) func(L, R) bool
    Or returns the || gate for non-boolean types

func ParseBin[T rules.Real](s string) T
    ParseBin parses a base 2 integer-string into the given type under the hood,
    it's a panicky-wrapper on strconv.ParseInt

func ParseFloat[T rules.Real](s string) T
    ParseFloat parses a decimal-string into the given type under the hood,
    it's a panicky-wrapper on strconv.ParseInt

func ParseHex[T rules.Real](s string) T
    ParseHex parses a base 16 integer-string into the given type under the hood,
    it's a panicky-wrapper on strconv.ParseInt

func ParseInt[T rules.Real](s string) T
    ParseInt parses a base 10 integer-string into the given type under the hood,
    it's a panicky-wrapper on strconv.ParseInt

func ParseIntn[T rules.Real](n int, s string) T
    ParseIntn parses an integer-string of arbitrary base into the given type
    under the hood, it's a panicky-wrapper on strconv.ParseInt

func Pipe[L, R, T any](one func(L) R, two func(R) T) func(L) T
    Pipe concatenates two unary operators

func Pointer[T any](p T) *T
    Pointer returns a pointer to the given value

func Pred[T, U any](pred func(T) bool, whenTrue, whenFalse U) func(T) U
func Printer() func(...any) (int, error)
    Printer returns a closure that prints to the given writer

func Printerf(format string) func(...any) (int, error)
    Printerf returns a closure that prints a format string to the given writer

func Printerln() func(...any) (int, error)
    Printerln returns a closure that prints a line to the given writer

func Real[N rules.Real, C rules.Complex](arg C) N
    Real extracts the real part of a complex number

func Returner[T any](arg T) T
    Returner is a place holder for pipelines that depend on casters

func Sizeof[T any]() uintptr
    Sizeof returns the size, in bits, of an instance of the given type

func Slicead[I, O any](f func(...I) O) func([]I) O
    Slicead transforms a variadic function to one that accepts a slice

func Ternary[T any](pred bool, whenTrue, whenFalse T) T
    Ternary reduces the if-else statement to a one-liner

func ToBool[T comparable](arg T) bool
    ToBool returns true iff the argument is non-zero a string is zero iff it's
    length is zero

func ToBoolLener[Seq rules.Lener[K, V], K comparable, V any](arg Seq) bool
    ToBoolIter converts indexable types to booleans if it's length is zero,
    it is false

func ToByte[T rules.Integer](arg T) byte
    ToByte casts the argument to a byte value

func ToComplex128[T rules.Real](arg T) complex128
    ToComplex128 casts the argument to a complex128 value

func ToComplex64[T rules.Real](arg T) complex64
    ToComplex64 casts the argument to a complex64 value

func ToFloat32[T rules.Integer](arg T) float32
    ToFloat32 casts the argument to a float32 value

func ToFloat64[T rules.Integer](arg T) float64
    ToFloat64 casts the argument to a float64 value

func ToInt[T rules.Integer](arg T) int
    ToInt casts the argument to a int value

func ToInt16[T rules.Integer](arg T) int16
    ToInt16 casts the argument to a int16 value

func ToInt32[T rules.Integer](arg T) int32
    ToInt32 casts the argument to a int32 value

func ToInt64[T rules.Integer](arg T) int64
    ToInt64 casts the argument to a int64 value

func ToInt8[T rules.Integer](arg T) int8
    ToInt8 casts the argument to a int8 value

func ToRune[T rules.Integer](arg T) rune
    ToRune casts the argument to a rune value

func ToString[T any](arg T) string
    ToString casts the argument to a string value

func ToString2[T any](arg T) string
    ToString2 casts the argument to a string2 value with type information

func ToUint[T rules.Integer](arg T) uint
    ToUint casts the argument to a uint value

func ToUint16[T rules.Integer](arg T) uint16
    ToUint16 casts the argument to a uint16 value

func ToUint32[T rules.Integer](arg T) uint32
    ToUint32 casts the argument to a uint32 value

func ToUint64[T rules.Integer](arg T) uint64
    ToUint64 casts the argument to a uint64 value

func ToUint8[T rules.Integer](arg T) uint8
    ToUint8 casts the argument to a uint8 value

func ToUintptr[T rules.Integer](arg T) uintptr
    ToUintptr casts the argument to a uintptr value

func Typeis[want, have any](arg have) bool
    Typeis checks if a given value has wanted type

func Value[T any](p *T) T
    Value returns the value to the given pointer

func Variad[I, O any](f func([]I) O) func(...I) O
    Variad transforms a variadic function to one that accepts a slice

func WrapAny[T, L, R any](fn func(L, R)) func(L, R) *T
func Xor[L, R any](one func(L) bool, two func(R) bool) func(L, R) bool
    Or returns the && gate for non-boolean types


TYPES

type BinOp[L, R, T any] func(L, R) T

type BinVar[L, R any] func() (L, R)

type Caster[I, O any] func(I) O

type Niler[K comparable, V any] interface {
	~*V | ~[]V | ~map[K]V | ~chan V //| error // compiler: cannot use error in union (error contains methods)
}
    Niler is a constraint satisfied by nullable types

type Op[T any] func(T) T

type Option[T any] BinVar[T, error]

type TernOp[T any] func(bool, T, T) T

type Var[T any] func() T


```