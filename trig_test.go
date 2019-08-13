package fastmath_test

import (
	"math"
	"testing"

	"github.com/13rac1/fastmath"
)

// Complete range of output from the FastLED sin8_C() function.
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

// Range of output from the FastLED sin16_C() function.
// for (int x = 0; x < 65535; x+=256) {
//   printf("%d,", sin16(x));
//   if (x%2048 == 1792) {
//     printf("\n");
//   }
// }
var fastLEDSin16 = [256]int16{
	0, 784, 1568, 2352, 3136, 3920, 4704, 5488,
	6393, 7161, 7929, 8697, 9465, 10233, 11001, 11769,
	12539, 13243, 13947, 14651, 15355, 16059, 16763, 17467,
	18204, 18812, 19420, 20028, 20636, 21244, 21852, 22460,
	23170, 23666, 24162, 24658, 25154, 25650, 26146, 26642,
	27245, 27613, 27981, 28349, 28717, 29085, 29453, 29821,
	30273, 30497, 30721, 30945, 31169, 31393, 31617, 31841,
	32137, 32201, 32265, 32329, 32393, 32457, 32521, 32585,
	32645, 32581, 32517, 32453, 32389, 32325, 32261, 32197,
	32051, 31827, 31603, 31379, 31155, 30931, 30707, 30483,
	30166, 29798, 29430, 29062, 28694, 28326, 27958, 27590,
	27107, 26611, 26115, 25619, 25123, 24627, 24131, 23635,
	23030, 22422, 21814, 21206, 20598, 19990, 19382, 18774,
	18127, 17423, 16719, 16015, 15311, 14607, 13903, 13199,
	12489, 11721, 10953, 10185, 9417, 8649, 7881, 7113,
	6223, 5439, 4655, 3871, 3087, 2303, 1519, 735,
	0, -784, -1568, -2352, -3136, -3920, -4704, -5488,
	-6393, -7161, -7929, -8697, -9465, -10233, -11001, -11769,
	-12539, -13243, -13947, -14651, -15355, -16059, -16763, -17467,
	-18204, -18812, -19420, -20028, -20636, -21244, -21852, -22460,
	-23170, -23666, -24162, -24658, -25154, -25650, -26146, -26642,
	-27245, -27613, -27981, -28349, -28717, -29085, -29453, -29821,
	-30273, -30497, -30721, -30945, -31169, -31393, -31617, -31841,
	-32137, -32201, -32265, -32329, -32393, -32457, -32521, -32585,
	-32645, -32581, -32517, -32453, -32389, -32325, -32261, -32197,
	-32051, -31827, -31603, -31379, -31155, -30931, -30707, -30483,
	-30166, -29798, -29430, -29062, -28694, -28326, -27958, -27590,
	-27107, -26611, -26115, -25619, -25123, -24627, -24131, -23635,
	-23030, -22422, -21814, -21206, -20598, -19990, -19382, -18774,
	-18127, -17423, -16719, -16015, -15311, -14607, -13903, -13199,
	-12489, -11721, -10953, -10185, -9417, -8649, -7881, -7113,
	-6223, -5439, -4655, -3871, -3087, -2303, -1519, -735,
}

func TestSin8(t *testing.T) {
	for x := 0; x < 256; x++ {
		if fastmath.Sin8(uint8(x)) != fastLEDSin8[x] {
			t.Errorf("sin8(%d) expected: %d, found: %d", x, fastLEDSin8[x], fastmath.Sin8(uint8(x)))
		}
	}
}

func TestSin8Delta(t *testing.T) {
	var totalDelta float64 = 0
	var maxDelta float64 = 0
	for x := 0; x <= 256; x++ {
		sin8 := fastmath.Sin8(uint8(x))
		xInRadians := float64(x) * 2 * math.Pi / 256
		stdSin8 := (math.Sin(xInRadians) + 1) / 2 * 256
		if sin8 == 0 || math.Round(stdSin8) == 0 {
			// Avoid division with zero.
			continue
		}
		delta := float64(sin8) - stdSin8
		if delta > maxDelta {
			maxDelta = delta
		}
		totalDelta += math.Abs(delta)

		// t.Logf("sin8(%d): %d, sin float: %.02f, delta: %.02f", x, sin8, stdSin8, delta)
	}
	t.Logf("max delta: %.02f, average delta: %.02f", maxDelta, totalDelta/256)
	averageError := totalDelta / 256 / 256 * 100
	t.Logf("max error: %.02f%%, average error: %.02f%%", maxDelta/256*100, averageError)
	if averageError > 0.8 {
		// TODO: Can we lower this further?
		t.Fatal("average error is too high")
	}
}

func TestSin16(t *testing.T) {
	for x := 0; x < 65535; x += 256 {
		if fastmath.Sin16(uint16(x)) != fastLEDSin16[x/256] {
			t.Errorf("sin16(%d) expected: %d, found: %d", x, fastLEDSin16[x/256], fastmath.Sin16(uint16(x)))
		}
	}
}

func TestSin16Delta(t *testing.T) {
	var totalDelta float64 = 0
	var maxDelta float64 = 0
	for x := 0; x <= 65535; x += 256 {
		sin16 := fastmath.Sin16(uint16(x))
		xInRadians := float64(x) * 2 * math.Pi / 65535
		stdSin16 := math.Sin(xInRadians) / 2 * 65535
		if sin16 == 0 || math.Round(stdSin16) == 0 {
			// Avoid division with zero.
			continue
		}
		delta := float64(sin16) - stdSin16
		if math.Abs(delta) > maxDelta {
			maxDelta = delta
		}
		totalDelta += math.Abs(delta)

		// t.Logf("sin16(%d): %d, sin float: %.02f, delta: %.02f", x, sin16, stdSin16, delta)
	}
	t.Logf("max delta: %.02f, average delta: %.02f", maxDelta, totalDelta/256)
	averageError := totalDelta / 256 / 65536 * 100
	t.Logf("max error: %.02f%%, average error: %.02f%%", maxDelta/65536*100, averageError)
	if averageError > 0.2 {
		// TODO: Can we lower this further?
		t.Fatal("average error is too high")
	}
}

func stdLibSin8(theta uint8) uint8 {
	// Find the same 0-255 range as Sin8()
	xInRadians := float64(theta) / 255 * 2 * math.Pi
	return uint8(math.Round((math.Sin(xInRadians) + 1) / 2 * 255))
}

func BenchmarkStdLibSin8(b *testing.B) {
	var r uint8
	x := fastmath.PI8
	for n := 0; n < b.N; n++ {
		r = stdLibSin8(x)
	}
	result8 = r
}

func BenchmarkSin8(b *testing.B) {
	var r uint8
	x := fastmath.PI8
	for n := 0; n < b.N; n++ {
		r = fastmath.Sin8(x)
	}
	result8 = r
}

func stdLibSin16(theta uint16) int16 {
	// Find the same -32767 to 32767 range as Sin16()
	xInRadians := float64(theta) / 65535 * 2 * math.Pi
	return int16(math.Round((math.Sin(xInRadians)) / 2 * 65535))
}

func BenchmarkStdLibSin16(b *testing.B) {
	var r int16
	x := fastmath.PI16
	for n := 0; n < b.N; n++ {
		r = stdLibSin16(x)
	}
	result16 = r
}

func BenchmarkSin16(b *testing.B) {
	var r int16
	x := fastmath.PI16
	for n := 0; n < b.N; n++ {
		r = fastmath.Sin16(x)
	}
	result16 = r
}
