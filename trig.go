package fastmath

// Trig Fast trig functions
//
// Fast 8 and 16-bit approximations of sin(x) and cos(x).
// Don't use these approximations for calculating the
// Trajectory of a rocket to Mars, but they're great
// for art projects and LED displays.

// PI8 is the value of π in 8 bit math.
const PI8 uint8 = 127

// PI16 is the value of π in 16 bit math.
const PI16 uint16 = 32768

var bM16Interleave = []uint8{0, 49, 49, 41, 90, 27, 117, 10}

// Sin8 is a fast 8-bit approximation of sin(x). This approximation never varies
// more than 2% from the floating point value.
//
// This is a Golang translation of the FastLED lib8tion sin8_C() function.
// https://raw.githubusercontent.com/FastLED/FastLED/dcbf3993/lib8tion/trig8.h
//
// Accepts theta input angle from 0-255.
// Returns sin of theta, value between 0 and 255
func Sin8(theta uint8) uint8 {
	offset := theta
	if theta&0x40 != 0 {
		offset = 255 - offset
	}
	offset &= 0x3F // 0..63

	secoffset := offset & 0x0F // 0..15
	if theta&0x40 != 0 {
		secoffset++
	}

	section := offset >> 4 // 0..3
	s2 := section * 2

	var p uint8
	p += s2
	b := bM16Interleave[p]
	p++
	m16 := bM16Interleave[p]

	// Must be cast int16 so multiplication can result be greater than 255.
	mx := (int16(m16) * int16(secoffset)) >> 4

	// Must be int16 to be negative and allow adding 128.
	var y int16
	y = mx + int16(b)
	if theta&0x80 != 0 {
		y = -y
	}

	y += 128
	return uint8(y)
}

// Cos8 is a fast 8-bit approximation of cos(x). This approximation never varies
// more than 2% from the floating point value.
//
// Accepts theta input angle from 0-255
// Returns sin of theta, value between 0 and 255
func Cos8(theta uint8) uint8 {
	return Sin8(theta + 64)
}

var base = []uint16{0, 6393, 12539, 18204, 23170, 27245, 30273, 32137}
var slope = []uint8{49, 48, 44, 38, 31, 23, 14, 4}

// Sin16 is a fast 16-bit approximation of sin(x). This approximation never
// varies more than 0.69% from the floating point value.
//
// This is a Golang translation of the FastLED lib8tion sin16_C() function.
// https://raw.githubusercontent.com/FastLED/FastLED/dcbf3993/lib8tion/trig8.h
//
// Accepts theta input angle from 0-65535.
// Returns sin of theta, value between -32767 to 32767.
func Sin16(theta uint16) int16 {
	offset := (theta & 0x3FFF) >> 3 // 0..2047
	if theta&0x4000 != 0 {
		offset = 2047 - offset
	}

	section := offset / 256 // 0..7
	b := base[section]
	m := slope[section]

	secoffset8 := uint8(offset) / 2

	mx := int16(m) * int16(secoffset8)
	y := int16(mx) + int16(b)

	if theta&0x8000 != 0 {
		y = -y
	}

	return y
}

// Cos16 is a fast 16-bit approximation of cos(x). This approximation never
// varies more than 0.69% from the floating point value.
//
// Accepts theta input angle from 0-65535.
// Returns sin of theta, value between -32767 to 32767.
func Cos16(theta uint16) int16 {
	return Sin16(theta + 16384)
}
