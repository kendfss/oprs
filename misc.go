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

// Lt calls "<"
func Lt[T rules.Ordered](a, b T) bool {
	return a < b
}

// Le calls "<="
func Le[T rules.Ordered](a, b T) bool {
	return a <= b
}

// Gt calls ">"
func Gt[T rules.Ordered](a, b T) bool {
	return a > b
}

// Ge calls ">="
func Ge[T rules.Ordered](a, b T) bool {
	return a >= b
}

// Eq compares two values for equality
func Eq[T comparable](a, b T) bool {
	return a == b
}

// Ne compares two values for inequality
func Ne[T comparable](a, b T) bool {
	return a != b
}
