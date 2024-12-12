package real

import (
	"math"

	"github.com/kendfss/oprs/internal/tools"
	"github.com/kendfss/rules"
)

// WrapInt interpolates val as though it were referring to an element of
// an array indexed by the range [min, max)
func WrapInt[T rules.Int](val, max, min T) T {
	if val >= min && val < max {
		return val
	}
	if val >= max {
		return WrapInt(val%max, max, min)
	}
	for val <= min {
		val += max
	}
	return WrapInt(val, max, min)
}

// Add returns the sum of two values
func Add[T rules.Ordered](a, b T) T {
	return a + b
}

// Mul returns the product of two values
func Mul[T rules.Number](a, b T) T {
	return a * b
}

// Sub returns the difference between two values
func Sub[T rules.Number](a, b T) T {
	return a - b
}

// Div returns the quotient of two values
func Div[T rules.Number](a, b T) T {
	return a / b
}

// Divs returns the quotient of two values
func Divs[T rules.Int](a, b T) bool {
	return a%b == 0
}

// Succ returns the successor of a number
func Succ[T rules.Number](a T) T {
	return a + 1
}

// Prev returns the predecessor of a number
func Prev[T rules.Number](a T) T {
	return a - 1
}

// Remainders counts the number of operations needed to divide a numerator by a denominator via subtraction
func Remainders[T rules.Integer](num, den T) T {
	count := num / den
	if num%den != 0 {
		count++
	}
	return count
}

func Mod[I rules.Int](a, b I) I {
	return a % b
}

// Floating point modulus
// less accurate for large numbers
func Fmod[R rules.Real](value, modulus R) R {
	return value - R(int64(value/modulus))*modulus
}

// the sequence of terms generated repeatedly subtracting
// the divisor from the value and then subtracting any non-zero remainder
func Subtractions[T rules.Real](value, divisor T) (out []T) {
	if divisor == 0 {
		return []T{value}
	}
	for value > 0 {
		if value >= divisor {
			out = append(out, divisor)
			value -= divisor
		} else {
			out = append(out, value)
			value = 0
		}
	}
	return out
}

// Neg returns the arithmetic negation of
func Neg[T rules.Negable](a T) T {
	n := *new(T)
	n -= 1
	return a * n
}

// Returns a binary shift operator for the referenced value
func Shifter[T rules.Integer](arg *T, left bool) func(distance T) T {
	if left {
		return func(dist T) T {
			return *arg << dist
		}
	}
	return func(dist T) T {
		return *arg >> dist
	}
}

// Incrementer encloses a function that returns a number whose value
// is incremented, by delta, between successive calls
func Incrementer[T rules.Number](seed, delta T) func() T {
	seed -= delta
	return func() T {
		seed += delta
		return seed
	}
}

// Realf allows for converting within the real numbers
func Realf[I, O rules.Real](val I) O {
	return O(val)
}

// MapVal implements "map" from Java's Processing framework.
// It returns the following:
//
//	min2 + (max2-min2) * ((n - min1) / (max1 - min1))
func MapVal[N rules.Number](n, min1, max1, min2, max2 N) N {
	return min2 + (max2-min2)*((n-min1)/(max1-min1))
}

// ValMapper is a wrapper on MapVal that creates a reusable mapping
func ValMapper[N rules.Real](min1, max1, min2, max2 N) func(N) N {
	return func(n N) N {
		return MapVal(n, min1, max1, min2, max2)
	}
}

// Abs computes the absolute value of a real number
func Abs[R rules.Real](val R) R {
	if val >= 0 {
		return val
	}
	return -val
	// return R(math.Abs(float64(val)))
}

// Diff computes the absolute difference between a pair of real numbers
func Diff[R rules.Real](a, b R) R {
	if a < b {
		return Abs(b - a)
	}
	return Abs(a - b)
}

// GCD returns the Greatest Common Divisor as per the Euclidean algorithm
func GCD[I rules.Int](a, b I) I {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

// Eratosthenes' prime sieve
func Eratosthenes[R rules.Real](r R) (out []R) {
	return tools.Consume(Eratosthenesch(r))
}

// Eratosthenesch executes Eratosthenes' prime sieve (single use, non-blocking)
func Eratosthenesch[R rules.Real](r R) chan R {
	out := make(chan R)
	go func() {
		rack := tools.MustUpto(2, r, 1)
		marked := []R{}
		for _, e := range rack {
			if !tools.Contains(marked, e) {
				marked = append(marked, tools.FilterPred(func(e2 R) bool {
					return Fmod(e2, e) == 0 && e2 > e
				}, rack)...)
				out <- e
			}
		}
	}()
	return out
}

// IsPrime checks if a real number is prime
func IsPrime[R rules.Real](r R) bool {
	for e := range Eratosthenesch(r + 1) {
		if e == r {
			return true
		}
	}
	return false
}

// GPF returns the largest prime factor
func GPF[I rules.Int](n I) I {
	factors := tools.FilterPred(tools.Method(n, Divs[I]), Eratosthenes(n))
	return factors[len(factors)-1]
}

func Sin[R rules.Real](r R) R {
	return R(math.Sin(float64(r)))
}

func Cos[R rules.Real](r R) R {
	return R(math.Cos(float64(r)))
}

func Tan[R rules.Real](r R) R {
	return R(math.Tan(float64(r)))
}

func Asin[R rules.Real](r R) R {
	return R(math.Asin(float64(r)))
}

func Acos[R rules.Real](r R) R {
	return R(math.Acos(float64(r)))
}

func Atan[R rules.Real](r R) R {
	return R(math.Atan(float64(r)))
}

func Sinh[R rules.Real](r R) R {
	return R(math.Sinh(float64(r)))
}

func Cosh[R rules.Real](r R) R {
	return R(math.Cosh(float64(r)))
}

func Tanh[R rules.Real](r R) R {
	return R(math.Tanh(float64(r)))
}

func Asinh[R rules.Real](r R) R {
	return R(math.Asinh(float64(r)))
}

func Acosh[R rules.Real](r R) R {
	return R(math.Acosh(float64(r)))
}

func Atanh[R rules.Real](r R) R {
	return R(math.Atanh(float64(r)))
}

func Cbrt[R rules.Real](r R) R {
	return R(math.Cbrt(float64(r)))
}

func Ceil[R rules.Real](r R) R {
	return R(math.Ceil(float64(r)))
}

func Copysign[R rules.Real](r, sign R) R {
	return R(math.Copysign(float64(r), float64(sign)))
}

func Dim[R rules.Real](x, y R) R {
	return R(math.Dim(float64(x), float64(y)))
}

/*
	--------------------------------------------------
*/

func Erf[R rules.Real](r R) R {
	return R(math.Erf(float64(r)))
}

func Erfc[R rules.Real](r R) R {
	return R(math.Erfc(float64(r)))
}

func Erfcinv[R rules.Real](r R) R {
	return R(math.Erfcinv(float64(r)))
}

func Erfinv[R rules.Real](r R) R {
	return R(math.Erfinv(float64(r)))
}

func Exp[R rules.Real](r R) R {
	return R(math.Exp(float64(r)))
}

func Exp2[R rules.Real](r R) R {
	return R(math.Exp2(float64(r)))
}

func Expm1[R rules.Real](r R) R {
	return R(math.Expm1(float64(r)))
}

func FMA[R rules.Real](x, y, z R) R {
	return R(math.FMA(float64(x), float64(y), float64(z)))
}

/*
	-------------------------------------------------
*/

func Float32bits[R rules.Real](r R) R {
	return R(math.Float32bits(float32(r)))
}

func Float32frombits[R rules.Real](r R) R {
	return R(math.Float32frombits(uint32(r)))
}

func Float64bits[R rules.Real](r R) R {
	return R(math.Float64bits(float64(r)))
}

func Float64frombits[R rules.Real](r R) R {
	return R(math.Float64frombits(uint64(r)))
}

/*
	-------------------------------------------------
*/

func Floor[R rules.Real](r R) R {
	return R(math.Floor(float64(r)))
}

func Frexp[R rules.Real](r R) (frac float64, exp int) {
	return math.Frexp(float64(r))
}

func Gamma[R rules.Real](r R) R {
	return R(math.Gamma(float64(r)))
}

func Hypot[R rules.Real](p, q R) R {
	return R(math.Hypot(float64(p), float64(q)))
}

func Ilogb[R rules.Real](r R) int {
	return math.Ilogb(float64(r))
}

func Inf[R rules.Real](sign R) R {
	return R(math.Inf(int(sign)))
}

func IsInf[R rules.Real](r R, sign int) bool {
	return math.IsInf(float64(r), sign)
}

func IsNaN[R rules.Real](r R) bool {
	return math.IsNaN(float64(r))
}

func J0[R rules.Real](r R) R {
	return R(math.J0(float64(r)))
}

func J1[R rules.Real](r R) R {
	return R(math.J1(float64(r)))
}

func Jn[R rules.Real](n int, r R) R {
	return R(math.Jn(n, float64(r)))
}

func Ldexp[R rules.Real](frac R, exp int) R {
	return R(math.Ldexp(float64(frac), exp))
}

func Lgamma[R rules.Real](r R) (lgamma R, sign int) {
	l, s := math.Lgamma(float64(r))
	return R(l), s
}

func Log[R rules.Real](r R) R {
	return R(math.Log(float64(r)))
}

func Log10[R rules.Real](r R) R {
	return R(math.Log10(float64(r)))
}

func Log1p[R rules.Real](r R) R {
	return R(math.Log1p(float64(r)))
}

func Log2[R rules.Real](r R) R {
	return R(math.Log2(float64(r)))
}

func Logb[R rules.Real](r R) R {
	return R(math.Logb(float64(r)))
}

/*
	-------------------------------------------------
*/

func Max[R rules.Real](x, y R) R {
	return R(math.Max(float64(x), float64(y)))
}

func Min[R rules.Real](x, y R) R {
	return R(math.Min(float64(x), float64(y)))
}

func Modf[R rules.Real](x R) (int R, frac float64) {
	i, f := math.Modf(float64(x))
	return R(i), f
}

func NaN[R rules.Real]() R {
	return R(math.NaN())
}

/*
	-------------------------------------------------
*/

func Nextafter[R rules.Real](x, y R) R {
	return R(math.Nextafter(float64(x), float64(y)))
}

func Nextafter32[R rules.Real](x, y R) R {
	return R(math.Nextafter32(float32(x), float32(y)))
}

func Pow[R rules.Real](base, height R) R {
	return R(math.Pow(float64(base), float64(height)))
}

func Pow10[R rules.Real](r R) R {
	return R(math.Pow10(int(r)))
}

func Remainder[R rules.Real](x, y R) R {
	return R(math.Remainder(float64(x), float64(y)))
}

func Round[R rules.Real](r R) R {
	return R(math.Round(float64(r)))
}

func RoundToEven[R rules.Real](r R) R {
	return R(math.RoundToEven(float64(r)))
}

func Signbit[R rules.Real](r R) bool {
	return math.Signbit(float64(r))
}

func Trunc[R rules.Real](r R) R {
	return R(math.Trunc(float64(r)))
}

func Y0[R rules.Real](r R) R {
	return R(math.Y0(float64(r)))
}

func Y1[R rules.Real](r R) R {
	return R(math.Y1(float64(r)))
}

func Yn[R rules.Real](n int, r R) R {
	return R(math.Yn(n, float64(r)))
}
