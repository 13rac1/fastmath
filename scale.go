package fastmath

// Fast, efficient 8-bit scaling functions specifically
// designed for high-performance LED programming.

// Scale8 scales one byte by a second one, which is treated as the numerator of
// a fraction whose denominator is 256. In other words, it computes i * (scale /
// 256)
func Scale8(i uint8, scale uint8) uint8 {
	return uint8((uint16(i) * (1 + uint16(scale))) >> 8)
}

// Scale8Video is the "video" version of Scale8. Guarantees the output will be
// only be zero if one or both of the inputs are zero.  If both inputs are
// non-zero, the output is guaranteed to be non-zero. This makes for better
// 'video'/LED dimming, at the cost of several additional cycles.
func Scale8Video(i, scale uint8) uint8 {
	var scaleFixed uint16
	if i != 0 && scale != 0 {
		scaleFixed = 1
	}
	return uint8((uint16(i)*uint16(scale))>>8 + scaleFixed)
}

// NScale8x3 scales three one byte values by a fourth one, which is treated as
// the numerator of a fraction whose demominator is 256. In other words, it
// computes r,g,b * (scale / 256).
func NScale8x3(r, g, b *uint8, scale uint8) {
	scaleFixed := uint16(scale) + 1
	*r = uint8((uint16(*r) * scaleFixed) >> 8)
	*g = uint8((uint16(*g) * scaleFixed) >> 8)
	*b = uint8((uint16(*b) * scaleFixed) >> 8)
}

// NScale8x3Video scale three one byte values by a fourth one, which is treated
// as the numerator of a fraction whose demominator is 256. In other words, it
// computes r,g,b * (scale / 256), ensuring that non-zero values passed in
// remain non zero, no matter how low the scale argument.
func NScale8x3Video(r, g, b *uint8, scale uint8) {
	var nonZeroScale uint16
	if scale != 0 {
		nonZeroScale = 1
	}
	if *r != 0 {
		*r = uint8((uint16(*r)*uint16(scale))>>8 + nonZeroScale)
	}
	if *g != 0 {
		*g = uint8((uint16(*g)*uint16(scale))>>8 + nonZeroScale)
	}
	if *b != 0 {
		*b = uint8((uint16(*b)*uint16(scale))>>8 + nonZeroScale)
	}
}

// Scale16By8 scales a 16-bit unsigned value by an 8-bit value, considered as
// numerator of a fraction whose denominator is 256. In other words, it computes
// i * (scale / 256).
func Scale16By8(i uint16, scale uint8) uint16 {
	return uint16((uint32(i) * (1 + uint32(scale))) >> 8)
}

// Scale16 scales a 16-bit unsigned value by a 16-bit value, considered as
// numerator of a fraction whose denominator is 65536. In other words, it
// computes i * (scale / 65536).
func Scale16(i, scale uint16) uint16 {
	return uint16((uint32(i) * (1 + uint32(scale))) / 65536)
}
