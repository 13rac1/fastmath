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

const (
	mask  = 0x7FF
	shift = 64 - 11 - 1
	bias  = 1023
)

// Copied from https://golang.org/src/math/sqrt.go to avoid compiler
// optimizations into a single assembly instruction on many architectures.
// Copyright 2009 The Go Authors. All rights reserved.
func sqrt(x float64) float64 {
	// special cases
	switch {
	case x == 0 || math.IsNaN(x) || math.IsInf(x, 1):
		return x
	case x < 0:
		return math.NaN()
	}
	ix := math.Float64bits(x)
	// normalize x
	exp := int((ix >> shift) & mask)
	if exp == 0 { // subnormal x
		for ix&(1<<shift) == 0 {
			ix <<= 1
			exp--
		}
		exp++
	}
	exp -= bias // unbias exponent
	ix &^= mask << shift
	ix |= 1 << shift
	if exp&1 == 1 { // odd exp, double x to make it even
		ix <<= 1
	}
	exp >>= 1 // exp = exp/2, exponent of square root
	// generate sqrt(x) bit by bit
	ix <<= 1
	var q, s uint64               // q = sqrt(x)
	r := uint64(1 << (shift + 1)) // r = moving bit from MSB to LSB
	for r != 0 {
		t := s + r
		if t <= ix {
			s = t + r
			ix -= t
			q += r
		}
		ix <<= 1
		r >>= 1
	}
	// final rounding
	if ix != 0 { // remainder, result not exact
		q += q & 1 // round according to extra bit
	}
	ix = q>>1 + uint64(exp-1+bias)<<shift // significand + biased exponent
	return math.Float64frombits(ix)
}

func BenchmarkStdLibFallbackSqrt(b *testing.B) {
	var r uint8
	x := fastmath.PI16
	for n := 0; n < b.N; n++ {
		r = uint8(sqrt(float64(x)))
	}
	result8 = r
}

func BenchmarkStdLibDefaultSqrt(b *testing.B) {
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
