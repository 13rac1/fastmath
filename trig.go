package fastmath

// PI8 is the value of π in 8 bit math.
const PI8 uint8 = 127

var bM16Interleave = []uint8{0, 49, 49, 41, 90, 27, 117, 10}

// Sin8 is a fast 8-bit approximation of sin(x). This approximation never varies
// more than 2% from the floating point value.
//
// This is a Golang translation of the FastLED lib8tion sin8_C() function.
// https://raw.githubusercontent.com/FastLED/FastLED/dcbf3993/lib8tion/trig8.h
//
// @param theta input angle from 0-255.
// @returns sin of theta, value between 0 and 255
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
