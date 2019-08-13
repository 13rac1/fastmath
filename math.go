package fastmath

// Fast, efficient 8-bit math functions specifically
// designed for high-performance LED programming.

// Note: Only functions not provided by the Go runtime are implemented.

// QAdd8 adds one byte to another, saturating at 0xFF.
// Accepts:
// * i - first byte to add
// * j - second byte to add
// Returns the sum of i & j, capped at 0xFF
func QAdd8(i, j uint8) uint8 {
	t := uint16(i) + uint16(j)
	if t > 255 {
		t = 255
	}
	return uint8(t)
}

// QSub8 subtracts one byte from another, saturating at 0x00.
// Returns i - j with a floor of 0
func QSub8(i, j uint8) uint8 {
	t := int16(i) - int16(j)
	if t < 0 {
		t = 0
	}
	return uint8(t)
}

// QMul8 performs saturating 8x8 bit multiplication, with 8 bit result.
// Returns the product of i * j, capping at 0xFF
func QMul8(i, j uint8) uint8 {
	p := uint16(i) * uint16(j)
	if p > 255 {
		p = 255
	}
	return uint8(p)
}

// Abs8 finds the absolute value of a signed 8-bit int8
func Abs8(i int8) int8 {
	if i < 0 {
		i = -i
	}
	return i
}

// Sqrt16 finds the square root for 16-bit integers.
// This function is significantly slower, ~20X, on Intel/AMD CPUs. It should
// be much faster on a microcontroller though.
//
// Note: Sqrt is implemented in assembly on some systems.
// Others have assembly stubs that jump to func sqrt below.
// On systems where Sqrt is a single instruction, the compiler
// may turn a direct call into a direct use of that instruction instead.
// src: https://golang.org/src/math/sqrt.go
func Sqrt16(x uint16) uint8 {
	if x <= 1 {
		return uint8(x)
	}

	var low uint8 = 1 // lower bound
	var hi, mid uint8

	if x > 7904 {
		hi = 255
	} else {
		// initial estimate for upper bound
		hi = uint8((x >> 5) + 8)
	}

	for hi >= low {
		mid = uint8((uint16(low) + uint16(hi)) >> 1)
		if uint16(mid)*uint16(mid) > x {
			hi = mid - 1
		} else {
			if mid == 255 {
				return 255
			}
			low = mid + 1
		}
	}

	return low - 1
}
