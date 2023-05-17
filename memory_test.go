package oprs

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClone(t *testing.T) {
	x := big.NewInt(100)
	y := Clone(x)
	assert.NotEqual(t, x, y, "rereferencing does not work for big.Int %#v", x == y)
	assert.NotEqual(t, x.Cmp(y), 0, "rereferencing does not work for big.Int")
}

func TestByteSliceOperators(t *testing.T) {
	type test struct {
		dst, src []byte
		want     int
	}

	run := func(tests []test, name string) {
		ref := map[string]interface{}{
			"Capy": Capy[byte],
			"Move": Move[byte],
		}[name]

		for i, test := range tests {
			src0, dst0 := test.src, test.dst

			var have int
			switch ref.(type) {
			case func(*[]byte, []byte) int:
				fn, _ := ref.(func(*[]byte, []byte) int)
				have = fn(&test.dst, test.src)
			case func(*[]byte, *[]byte) int:
				fn, _ := ref.(func(*[]byte, *[]byte) int)
				have = fn(&test.dst, &test.src)
			}

			if have != test.want || (name == "Move" && len(test.src) != len(src0)-have) || (len(dst0) == 0 && cap(dst0) >= len(src0) && !bytes.Equal(src0, test.dst)) {
				t.Fatalf("\ntest #%d:\n\thave:\t%d\n\twant:\t%d\n\tsrc0:\t%d\t\n\tdst0:\t%d\t(cap=%d, len=%d)\n\tsrc:\t%d\t\n\tdst:\t%d\t(cap=%d, len=%d)\n", i, have, test.want, src0, dst0, cap(dst0), len(dst0), test.src, test.dst, cap(test.dst), len(test.dst))
				// } else {
				// 	fmt.Printf("\ntest #%d:\n\thave:\t%d\n\twant:\t%d\n\tsrc0:\t%d\t\n\tdst0:\t%d\t(cap=%d, len=%d)\n\tsrc:\t%d\t\n\tdst:\t%d\t(cap=%d, len=%d)\n", i, have, test.want, src0, dst0, cap(dst0), len(dst0), test.src, test.dst, cap(test.dst), len(test.dst))
			}
		}
	}

	t.Run("Capy", func(t *testing.T) {
		tests := []test{
			{src: []byte{5, 3, 2}, dst: make([]byte, 0, 0), want: 0},
			{src: []byte{5, 3, 2}, dst: make([]byte, 0, 1), want: 1},
			{src: []byte{5, 3, 2}, dst: make([]byte, 0, 2), want: 2},
			{src: []byte{5, 3, 2}, dst: make([]byte, 0, 3), want: 3},
			{src: []byte{5, 3, 2}, dst: make([]byte, 0, 4), want: 3},
		}
		run(tests, "Capy")
	})

	t.Run("Move", func(t *testing.T) {
		tests := []test{
			{src: []byte{5, 3, 2}, dst: make([]byte, 0, 0), want: 0},
			{src: []byte{5, 3, 2}, dst: make([]byte, 0, 1), want: 1},
			{src: []byte{5, 3, 2}, dst: make([]byte, 0, 2), want: 2},
			{src: []byte{5, 3, 2}, dst: make([]byte, 0, 3), want: 3},
			{src: []byte{5, 3, 2}, dst: make([]byte, 0, 4), want: 3},

			{src: []byte{5, 3, 2}, dst: append(make([]byte, 0, 4), 0), want: 3},
			{src: []byte{5, 3, 2}, dst: append(make([]byte, 0, 3), 0), want: 2},
			{src: []byte{5, 3, 2}, dst: append(make([]byte, 0, 2), 0), want: 1},
			{src: []byte{5, 3, 2}, dst: append(make([]byte, 0, 1), 0), want: 0},
			// {src: []byte{5, 3, 2}, dst: append(make([]byte, 0, 0),), want: 0},
		}
		run(tests, "Move")
	})
}
