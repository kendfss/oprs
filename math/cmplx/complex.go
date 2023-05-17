package cmplx

import (
	"math/cmplx" // Package cmplx provides basic constants and mathematical functions for complex

	"github.com/kendfss/rules"
)

// numbers. Special case handling conforms to the C99 standard Annex G IEC
// 60559-compatible complex arithmetic.

func Real[R rules.Real, C rules.Complex](z C) R {
	return R(real(complex128(z)))
}

func Imag[R rules.Real, C rules.Complex](z C) R {
	return R(imag(complex128(z)))
}

// Abs returns the absolute value (also called the modulus) of x.
func Abs[R rules.Real, C rules.Complex](x C) R {
	return R(cmplx.Abs(complex128(x)))
}

// Acos returns the inverse cosine of x.
func Acos[C rules.Complex](x C) C {
	return C(cmplx.Acos(complex128(x)))
}

// Acosh returns the inverse hyperbolic cosine of x.
func Acosh[C rules.Complex](x C) C {
	return C(cmplx.Acosh(complex128(x)))
}

// Asin returns the inverse sine of x.
func Asin[C rules.Complex](x C) C {
	return C(cmplx.Asin(complex128(x)))
}

// Asinh returns the inverse hyperbolic sine of x.
func Asinh[C rules.Complex](x C) C {
	return C(cmplx.Asinh(complex128(x)))
}

// Atan returns the inverse tangent of x.
func Atan[C rules.Complex](x C) C {
	return C(cmplx.Atan(complex128(x)))
}

// Atanh returns the inverse hyperbolic tangent of x.
func Atanh[C rules.Complex](x C) C {
	return C(cmplx.Atanh(complex128(x)))
}

// Conj returns the complex conjugate of x.
func Conj[C rules.Complex](x C) C {
	return C(cmplx.Conj(complex128(x)))
}

// Cos returns the cosine of x.
func Cos[C rules.Complex](x C) C {
	return C(cmplx.Cos(complex128(x)))
}

// Cosh returns the hyperbolic cosine of x.
func Cosh[C rules.Complex](x C) C {
	return C(cmplx.Cosh(complex128(x)))
}

// Cot returns the cotangent of x.
func Cot[C rules.Complex](x C) C {
	return C(cmplx.Cot(complex128(x)))
}

// Exp returns e**x, the base-e exponential of x.
func Exp[C rules.Complex](x C) C {
	return C(cmplx.Exp(complex128(x)))
}

// Inf returns a complex infinity, complex(+Inf, +Inf).
func Inf[C rules.Complex]() C {
	return C(cmplx.Inf())
}

// IsInf reports whether either real(x) or imag(x) is an infinity.
func IsInf[C rules.Complex](x C) bool {
	return cmplx.IsInf(complex128(x))
}

// IsNaN reports whether either real(x) or imag(x) is NaN
// and neither is an infinity.
func IsNaN[C rules.Complex](x C) bool {
	return cmplx.IsNaN(complex128(x))
}

// Log returns the natural logarithm of x.
func Log[C rules.Complex](x C) C {
	return C(cmplx.Log(complex128(x)))
}

// Log10 returns the decimal logarithm of x.
func Log10[C rules.Complex](x C) C {
	return C(cmplx.Log10(complex128(x)))
}

// NaN returns a complex “not-a-number” value.
func NaN[C rules.Complex]() C {
	return C(cmplx.NaN())
}

// Phase returns the phase (also called the argument) of x.
// The returned value is in the range [-Pi, Pi].
func Phase[R rules.Real, C rules.Complex](x C) R {
	return R(cmplx.Phase(complex128(x)))
}

// Polar returns the absolute value r and phase θ of x,
// such that x = r * e**θi.
// The phase is in the range [-Pi, Pi].
func Polar[R rules.Real, C rules.Complex](x C) (r, θ R) {
	r2, θ2 := cmplx.Polar(complex128(x))
	return R(r2), R(θ2)
}

// Pow returns x**y, the base-x exponential of y.
// For generalized compatibility with math.Pow:
//
//	Pow(0, ±0) returns 1+0i
//	Pow(0, c) for real(c)<0 returns Inf+0i if imag(c) is zero, otherwise Inf+Inf i.
func Pow[C rules.Complex](x, y C) C {
	return C(cmplx.Pow(complex128(x), complex128(y)))
}

func Rect[C rules.Complex, R rules.Real](r, θ R) C {
	return C(cmplx.Rect(float64(r), float64(θ)))
}

// Sin returns the sine of x.
func Sin[C rules.Complex](x C) C {
	return C(cmplx.Sin(complex128(x)))
}

// Sinh returns the hyperbolic sine of x.
func Sinh[C rules.Complex](x C) C {
	return C(cmplx.Sinh(complex128(x)))
}

// Sqrt returns the square root of x.
// The result r is chosen so that real(r) ≥ 0 and imag(r) has the same sign as imag(x).
func Sqrt[C rules.Complex](x C) C {
	return C(cmplx.Sqrt(complex128(x)))
}

// Tan returns the tangent of x.
func Tan[C rules.Complex](x C) C {
	return C(cmplx.Tan(complex128(x)))
}

// Tanh returns the hyperbolic tangent of x.
func Tanh[C rules.Complex](x C) C {
	return C(cmplx.Tanh(complex128(x)))
}
