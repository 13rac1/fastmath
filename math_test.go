package fastmath_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/13rac1/fastmath"
)

func TestQAdd8(t *testing.T) {
	if fastmath.QAdd8(255, 128) != 255 {
		t.Fatal("QAdd8() is not saturating")
	}
}

func TestQSub8(t *testing.T) {
	if fastmath.QSub8(128, 255) != 0 {
		t.Fatal("QSub8() is not saturating")
	}
}

func TestQMul8(t *testing.T) {
	if fastmath.QMul8(255, 255) != 255 {
		t.Fatal("QMul8() is not saturating")
	}
}

func TestAbs8(t *testing.T) {
	if fastmath.Abs8(-100) != 100 {
		t.Fatal("Abs8() absolute value not found")
	}
}

func TestSqrt16(t *testing.T) {
	testCases := []struct {
		in       uint16
		expected uint8
	}{{
		in:       0,
		expected: 0,
	}, {
		in:       2,
		expected: 1,
	}, {
		in:       255,
		expected: 15,
	}, {
		in:       128,
		expected: 11,
	}, {
		in:       2500,
		expected: 50,
	}, {
		in:       5000,
		expected: 70,
	}, {
		in:       7500,
		expected: 86,
	}, {
		in:       10000,
		expected: 100,
	}}
	for _, test := range testCases {
		name := fmt.Sprintf("Sqrt16(%d)", test.in)
		t.Run(name, func(t *testing.T) {
			r := fastmath.Sqrt16(test.in)
			if test.expected != r {
				t.Fatalf("expected: %d, found: %d", test.expected, r)
			}
		})
	}
}

func BenchmarkStdLibSqrt(b *testing.B) {
	var r uint8
	x := fastmath.PI16
	for n := 0; n < b.N; n++ {
		r = uint8(math.Sqrt(float64(x)))
	}
	result8 = r
}

func BenchmarkSqrt16(b *testing.B) {
	var r uint8
	x := fastmath.PI16
	for n := 0; n < b.N; n++ {
		r = fastmath.Sqrt16(x)
	}
	result8 = r
}
