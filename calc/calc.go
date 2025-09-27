package calc

import (
	"errors"
)

// Add returns the sum of a and b.
func Add(a, b int) int {
	return a + b
}

// Sub returns a - b.
func Sub(a, b int) int {
	return a - b
}

// Mul returns a * b.
func Mul(a, b int) int {
	return a * b
}

// Div returns a / b or an error when dividing by zero.
func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

// Factorial returns n! for n >= 0. Returns error for negative n.
func Factorial(n int) (int64, error) {
	if n < 0 {
		return 0, errors.New("negative input")
	}
	var res int64 = 1
	for i := 2; i <= n; i++ {
		res *= int64(i)
	}
	return res, nil
}

// Reverse returns the reversed string (handles UTF-8 correctly).
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPrime reports whether n is a prime (n <= 1 -> false).
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
