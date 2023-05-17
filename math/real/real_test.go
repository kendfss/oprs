package real

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kendfss/rules"
)

const (
	nTests = 10
	nItems = 10
	nMax   = 10
)

func TestIncrementer(t *testing.T) {
	nMax := 100
	for i := 0; i < nTests; i++ {
		seed, delta := rand.Intn(nMax), rand.Intn(nMax)
		inc := Incrementer(seed, delta)
		prev := inc()
		assert.Equal(t, seed, prev, "seed != now\ndelta: %d", delta)
		for j := 0; j < nItems; j++ {
			now := inc()
			assert.Equal(t, now, prev+delta)
			prev = now
		}
	}
}

func negTester[N rules.Negable](t *testing.T, val N) bool {
	return assert.Equal(t, -val, Neg(val))
}

func TestNeg(t *testing.T) {
	negTester(t, int(1))
	negTester(t, int8(2))
	negTester(t, int16(3))
	negTester(t, int32(4))
	negTester(t, int64(5))
	negTester(t, float32(6))
	negTester(t, float64(7))
	negTester(t, complex64(8))
	negTester(t, complex128(9))
	negTester(t, complex(10, 11))
}
