package math

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
