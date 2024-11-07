package oprs

import "github.com/kendfss/rules"

// Add returns the sum of the given argumens
func Add[T rules.Num](a, b T) T {
	return a + b
}

// Mul multiplies the first argument by the second argument
func Mul[T rules.Num](a, b T) T {
	return a * b
}

// Sub subtracts the first argument by the second argument
func Sub[T rules.Num](a, b T) T {
	return a - b
}

// Div divides the first argument by the second argument
func Div[T rules.Num](a, b T) T {
	return a / b
}

// Lt wraps calls to "<"
func Lt[T rules.Ordered](a, b T) bool {
	return a < b
}

// Le wraps calls to "<="
func Le[T rules.Ordered](a, b T) bool {
	return a <= b
}

// Gt wraps calls to ">"
func Gt[T rules.Ordered](a, b T) bool {
	return a > b
}

// Ge wraps calls to ">="
func Ge[T rules.Ordered](a, b T) bool {
	return a >= b
}

// Eq wraps compares two values for equality
func Eq[T comparable](a, b T) bool {
	return a == b
}

// Ne wraps compares two values for inequality
func Ne[T comparable](a, b T) bool {
	return a != b
}

// Lor wraps calls to "||"
func Lor(a, b bool) bool {
	return a || b
}

// Land wraps calls to "&&"
func Land(a, b bool) bool {
	return a && b
}

// Bor wraps calls to "|"
func Bor[T rules.Int](a, b T) T {
	return a | b
}

// Band wraps calls to "&"
func Band[T rules.Int](a, b T) T {
	return a & b
}

// Bxor wraps calls to "^"
func Bxor[T rules.Int](a, b T) T {
	return a ^ b
}

// Rshift wraps calls to ">>"
func Rshift[T rules.Int](a, b T) T {
	return a >> b
}

// Lshift wraps calls to "<<"
func Lshift[T rules.Int](a, b T) T {
	return a << b
}

// Bclear wraps calls to "&^" aka bitwise clear
func Bclear[T rules.Int](a, b T) T {
	return a &^ b
}

// Mod wraps calls "%", the modulo operator
func Mod[T rules.Int](a, b T) T {
	return a % b
}

// Negative wraps "-" the negation operator
func Negative[T rules.Int](a T) T {
	return -a
}

// Complement wraps "^" the bitwise complement operator
func Complement[T rules.Int](a T) T {
	return ^a
}

// Lnot wraps the logical not "!" prefix operator
func Lnot(b bool) bool {
	return !b
}
