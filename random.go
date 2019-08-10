package fastmath

// Random Fast random number generators
// Fast 8- and 16- bit unsigned random numbers.
// Significantly faster than StdLib random(), but
// also somewhat less random. You can add entropy.

// X(n+1) = (2053 * X(n)) + 13849)
const rand2053 uint16 = 2053
const rand13849 uint16 = 13849

// random number seed
var rand16seed uint16 = 4

// Random8 generates an 8-bit random number.
func Random8() uint8 {
	rand16seed = (rand16seed * rand2053) + rand13849
	// return the sum of the high and low bytes, for better
	// mixing and non-sequential correlation.
	return uint8(uint8(rand16seed&0xFF) + uint8(rand16seed>>8))
}

// Random16 generates a 16 bit random number.
func Random16() uint16 {
	rand16seed = (rand16seed * rand2053) + rand13849
	return rand16seed
}

// Random8Limit generates an 8-bit random number between 0 and lim.
// Accept lim the upper bound for the result
func Random8Limit(lim uint8) uint8 {
	if lim == 0 {
		return 0
	}
	r := Random8()
	r = uint8((uint16(r) * uint16(lim)) >> 8)
	return r
}

// Random8Range generates an 8-bit random number in the given range.
// Accepts min the lower bound for the random number
// Accepts lim the upper bound for the random number
func Random8Range(min, lim uint8) uint8 {
	if min > lim {
		return lim
	}
	delta := lim - min
	r := Random8Limit(delta) + min
	return r
}

// Random16Limit generates an 16-bit random number between 0 and lim.
// Accepts lim the upper bound for the result
func Random16Limit(lim uint16) uint16 {
	if lim == 0 {
		return 0
	}
	r := Random16()
	p := uint32(lim) * uint32(r)
	r = uint16(p >> 16)
	return uint16(r)
}

// Random16Range generates an 16-bit random number in the given range.
// Accept min the lower bound for the random number
// Return lim the upper bound for the random number
func Random16Range(min, lim uint16) uint16 {
	delta := lim - min
	r := Random16Limit(delta) + min
	return r
}

// Random16SetSeed sets the 16-bit seed used for the random number generator.
func Random16SetSeed(seed uint16) {
	rand16seed = seed
}

// Random16GetSeed gets the current seed value for the random number generator.
func Random16GetSeed() uint16 {
	return rand16seed
}

// Random16AddEntropy adds entropy into the random number generator.
func Random16AddEntropy(entropy uint16) {
	rand16seed += entropy
}
