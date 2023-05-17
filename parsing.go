package oprs

import (
	"strconv"
	"unsafe"

	"github.com/kendfss/rules"
)

// ParseIntn parses an integer-string of arbitrary base into the given type
// under the hood, it's a panicky-wrapper on strconv.ParseInt
func ParseIntn[T rules.Real](n int, s string) T {
	i, err := strconv.ParseInt(s, n, int(unsafe.Sizeof(*new(T))*8))
	if err != nil {
		panic(err)
	}
	return T(i)
}

// ParseInt parses a base 10 integer-string into the given type
// under the hood, it's a panicky-wrapper on strconv.ParseInt
func ParseInt[T rules.Real](s string) T {
	return ParseIntn[T](10, s)
}

// ParseBin parses a base 2 integer-string into the given type
// under the hood, it's a panicky-wrapper on strconv.ParseInt
func ParseBin[T rules.Real](s string) T {
	return ParseIntn[T](2, s)
}

// ParseHex parses a base 16 integer-string into the given type
// under the hood, it's a panicky-wrapper on strconv.ParseInt
func ParseHex[T rules.Real](s string) T {
	return ParseIntn[T](16, s)
}

// ParseFloat parses a decimal-string into the given type
// under the hood, it's a panicky-wrapper on strconv.ParseInt
func ParseFloat[T rules.Real](s string) T {
	i, err := strconv.ParseFloat(s, int(unsafe.Sizeof(*new(T))*8))
	if err != nil {
		panic(err)
	}
	return T(i)
}
