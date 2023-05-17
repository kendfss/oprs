package oprs

import "github.com/kendfss/rules"

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
