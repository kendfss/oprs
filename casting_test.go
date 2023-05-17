package oprs

import (
	"testing"

	"github.com/kendfss/oprs/internal/tools"
	"github.com/stretchr/testify/assert"
)

const (
	nTests = 10
	nItems = 10
	nMax   = 10
)

// Appender returns a detatched-method that appends arguments to the given slice
func Appender[T any](slice *[]T) func(T) {
	return func(arg T) {
		*slice = append(*slice, arg)
	}
}

func TestMethod(t *testing.T) {
	for i := range tools.Randints(nTests) {
		want := tools.Randints(nItems)
		have := []int{}
		tools.Send(Appender(&have), want)
		assert.Equal(t, want, have, "#%d", i)
	}
}

func TestBeforeAfter(t *testing.T) {
	val := 0
	fn := CurryL(tools.Add[int])(val)
	op1 := func() { val++; println("op1:", val) }
	op2 := func() { val--; println("op2:", val) }
	wrapped := After(Before(fn, op1), op2)
	for i, want := range tools.Randints(nTests) {
		have := wrapped(want)
		assert.Equal(t, 0, val, "#%d\n\tval %d != 0", i, val)
		assert.Equal(t, want, have, "#%d\n\twant %d, have %d", i, want, have)
	}
}
