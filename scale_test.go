package fastmath_test

import (
	"fmt"
	"testing"

	"github.com/13rac1/fastmath"
)

func TestScale8(t *testing.T) {
	testCases := []struct {
		in       uint8
		scale    uint8
		expected uint8
	}{{
		in:       0,
		scale:    0,
		expected: 0,
	}, {
		in:       0,
		scale:    255,
		expected: 0,
	}, {
		in:       255,
		scale:    0,
		expected: 0,
	}, {
		in:       255,
		scale:    255,
		expected: 255,
	}, {
		in:       255,
		scale:    128,
		expected: 128,
	}, {
		in:       128,
		scale:    255,
		expected: 128,
	}, {
		in:       128,
		scale:    128,
		expected: 64,
	}, {
		in:       64,
		scale:    128,
		expected: 32,
	}, {
		in:       1,
		scale:    1,
		expected: 0, // Primary difference vs Scale8Video
	}}

	for _, test := range testCases {
		name := fmt.Sprintf("Scale8(%d,%d)", test.in, test.scale)
		t.Run(name, func(t *testing.T) {
			r := fastmath.Scale8(test.in, test.scale)
			if test.expected != r {
				t.Fatalf("expected: %d, found: %d", test.expected, r)
			}
		})

		name = fmt.Sprintf("NScale8x3(%d,%d,%d,%d)", test.in, test.in, test.in, test.scale)
		t.Run(name, func(t *testing.T) {
			r := test.in
			g := test.in
			b := test.in
			fastmath.NScale8x3(&r, &g, &b, test.scale)
			if test.expected != r || test.expected != g || test.expected != b {
				t.Fatalf("expected: %d, found: (%d,%d,%d)", test.expected, r, g, b)
			}
		})
	}
}

func TestScale8Video(t *testing.T) {
	testCases := []struct {
		in       uint8
		scale    uint8
		expected uint8
	}{{
		in:       0,
		scale:    0,
		expected: 0,
	}, {
		in:       0,
		scale:    255,
		expected: 0,
	}, {
		in:       255,
		scale:    0,
		expected: 0,
	}, {
		in:       255,
		scale:    255,
		expected: 255,
	}, {
		in:       255,
		scale:    128,
		expected: 128,
	}, {
		in:       128,
		scale:    255,
		expected: 128,
	}, {
		in:       128,
		scale:    128,
		expected: 65,
	}, {
		in:       64,
		scale:    128,
		expected: 33,
	}, {
		in:       1,
		scale:    1,
		expected: 1, // Primary difference vs Scale8
	}}

	for _, test := range testCases {
		name := fmt.Sprintf("Scale8Video(%d,%d)", test.in, test.scale)
		t.Run(name, func(t *testing.T) {
			r := fastmath.Scale8Video(test.in, test.scale)
			if test.expected != r {
				t.Fatalf("expected: %d, found: %d", test.expected, r)
			}
		})

		name = fmt.Sprintf("NScale8x3Video(%d,%d,%d,%d)", test.in, test.in, test.in, test.scale)
		t.Run(name, func(t *testing.T) {
			r := test.in
			g := test.in
			b := test.in
			fastmath.NScale8x3Video(&r, &g, &b, test.scale)
			if test.expected != r || test.expected != g || test.expected != b {
				t.Fatalf("expected: %d, found: (%d,%d,%d)", test.expected, r, g, b)
			}
		})
	}
}

func TestScale16By8(t *testing.T) {
	testCases := []struct {
		in       uint16
		scale    uint8
		expected uint16
	}{{
		in:       0,
		scale:    0,
		expected: 0,
	}, {
		in:       0,
		scale:    255,
		expected: 0,
	}, {
		in:       255,
		scale:    0,
		expected: 0,
	}, {
		in:       255,
		scale:    255,
		expected: 255,
	}, {
		in:       65535,
		scale:    255,
		expected: 65535,
	}, {
		in:       65535,
		scale:    127,
		expected: 32767,
	}}

	for _, test := range testCases {
		name := fmt.Sprintf("Scale16By8(%d,%d)", test.in, test.scale)
		t.Run(name, func(t *testing.T) {
			r := fastmath.Scale16By8(test.in, test.scale)
			if test.expected != r {
				t.Fatalf("expected: %d, found: %d", test.expected, r)
			}
		})
	}
}

func TestScale16(t *testing.T) {
	testCases := []struct {
		in       uint16
		scale    uint16
		expected uint16
	}{{
		in:       0,
		scale:    0,
		expected: 0,
	}, {
		in:       0,
		scale:    65535,
		expected: 0,
	}, {
		in:       65535,
		scale:    0,
		expected: 0,
	}, {
		in:       255,
		scale:    65535,
		expected: 255,
	}, {
		in:       65535,
		scale:    255,
		expected: 255,
	}, {
		in:       65535,
		scale:    32767,
		expected: 32767,
	}, {
		in:       256,
		scale:    32767,
		expected: 128,
	}}

	for _, test := range testCases {
		name := fmt.Sprintf("Scale16By8(%d,%d)", test.in, test.scale)
		t.Run(name, func(t *testing.T) {
			r := fastmath.Scale16(test.in, test.scale)
			if test.expected != r {
				t.Fatalf("expected: %d, found: %d", test.expected, r)
			}
		})
	}
}
