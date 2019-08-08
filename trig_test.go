package fastmath

import (
	"math"
	"testing"
)

// Complete range of output from the FastLED sin8_C() function.
// Reproduce:
// int main() {
// for (int x = 0; x < 256; x++) {
//   printf("%d,", sin8(x));
//   if (x%8 == 7) {
//     printf("\n");
//   }
// }
var fastLEDSin8 = [256]uint8{
	128, 131, 134, 137, 140, 143, 146, 149,
	152, 155, 158, 161, 164, 167, 170, 173,
	177, 179, 182, 184, 187, 189, 192, 194,
	197, 200, 202, 205, 207, 210, 212, 215,
	218, 219, 221, 223, 224, 226, 228, 229,
	231, 233, 234, 236, 238, 239, 241, 243,
	245, 245, 246, 246, 247, 248, 248, 249,
	250, 250, 251, 251, 252, 253, 253, 254,
	255, 254, 253, 253, 252, 251, 251, 250,
	250, 249, 248, 248, 247, 246, 246, 245,
	245, 243, 241, 239, 238, 236, 234, 233,
	231, 229, 228, 226, 224, 223, 221, 219,
	218, 215, 212, 210, 207, 205, 202, 200,
	197, 194, 192, 189, 187, 184, 182, 179,
	177, 173, 170, 167, 164, 161, 158, 155,
	152, 149, 146, 143, 140, 137, 134, 131,
	128, 125, 122, 119, 116, 113, 110, 107,
	104, 101, 98, 95, 92, 89, 86, 83,
	79, 77, 74, 72, 69, 67, 64, 62,
	59, 56, 54, 51, 49, 46, 44, 41,
	38, 37, 35, 33, 32, 30, 28, 27,
	25, 23, 22, 20, 18, 17, 15, 13,
	11, 11, 10, 10, 9, 8, 8, 7,
	6, 6, 5, 5, 4, 3, 3, 2,
	1, 2, 3, 3, 4, 5, 5, 6,
	6, 7, 8, 8, 9, 10, 10, 11,
	11, 13, 15, 17, 18, 20, 22, 23,
	25, 27, 28, 30, 32, 33, 35, 37,
	38, 41, 44, 46, 49, 51, 54, 56,
	59, 62, 64, 67, 69, 72, 74, 77,
	79, 83, 86, 89, 92, 95, 98, 101,
	104, 107, 110, 113, 116, 119, 122, 125,
}

// A global result variable to trick the compiler during benchmarks.
var result uint8

func TestSin8(t *testing.T) {
	for x := 0; x < 256; x++ {
		if Sin8(uint8(x)) != fastLEDSin8[x] {
			t.Errorf("sin(%d) expected: %d, found: %d", x, Sin8(uint8(x)), fastLEDSin8[x])
		}
	}
}

func BenchmarkStdLibSin(b *testing.B) {
	var r uint8
	x := PI8
	for n := 0; n < b.N; n++ {
		// Find the same 0-255 range as Sin8()
		xInRadians := float64(x) / 255 * 2 * math.Pi
		sinX := uint8(math.Round((math.Sin(xInRadians) + 1) / 2 * 255))
		r = sinX
	}
	result = r
}

func BenchmarkSin8(b *testing.B) {
	var r uint8
	x := PI8
	for n := 0; n < b.N; n++ {
		r = Sin8(x)
	}
	result = r
}
