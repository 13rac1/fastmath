package fastmath_test

import (
	"math/rand"
	"testing"

	"github.com/13rac1/fastmath"
)

func TestRandom8(t *testing.T) {
	t.Run("SetSeed", func(t *testing.T) {
		fastmath.Random16SetSeed(0)

		a := fastmath.Random8()
		fastmath.Random16SetSeed(0)
		b := fastmath.Random8()
		if a != b {
			t.Fatal("same seed random did not match")
		}
	})
	t.Run("Limit", func(t *testing.T) {
		for lim := 1; lim < 256; lim++ {
			for x := 0; x < 1000; x++ {
				if fastmath.Random8Limit(uint8(lim)) > uint8(lim) {
					t.Fatalf("over limit: %d", lim)
				}
			}
		}
	})
	t.Run("Range", func(t *testing.T) {
		for min := 1; min < 256; min++ {
			for lim := min; lim < 256; lim++ {
				for x := 0; x < 1000; x++ {
					v := fastmath.Random8Range(uint8(min), uint8(lim))
					if v < uint8(min) {
						t.Fatalf("under min: %d", min)
					}
					if v > uint8(lim) {
						t.Fatalf("over limit: %d", lim)
					}
				}
			}
		}

		if fastmath.Random8Range(10, 0) == 10 {
			t.Fatalf("expected limit, found min")
		}
	})
}

func TestRandom16(t *testing.T) {
	t.Run("Seed", func(t *testing.T) {
		fastmath.Random16SetSeed(0)
		a := fastmath.Random16()
		fastmath.Random16SetSeed(0)
		b := fastmath.Random16()
		if a != b {
			t.Fatal("same seed random did not match")
		}
	})
	t.Run("Limit", func(t *testing.T) {
		for lim := 1; lim < 65535; lim++ {
			for x := 0; x < 1000; x++ {
				if fastmath.Random16Limit(uint16(lim)) > uint16(lim) {
					t.Fatalf("over limit: %d", lim)
				}
			}
		}
	})
	t.Run("Range", func(t *testing.T) {
		for min := 1; min < 65535; min += 256 {
			for lim := min; lim < 65535; lim += 256 {
				for x := 0; x < 1000; x++ {
					v := fastmath.Random16Range(uint16(min), uint16(lim))
					if v < uint16(min) {
						t.Fatalf("under min: %d", min)
					}
					if v > uint16(lim) {
						t.Fatalf("over limit: %d", lim)
					}
				}
			}
		}

		if fastmath.Random16Range(10, 0) == 10 {
			t.Fatalf("expected limit, found min")
		}
	})

}

func BenchmarkStdLibRandom8(b *testing.B) {
	var r uint8
	for n := 0; n < b.N; n++ {
		r = uint8(rand.Intn(255))
	}
	result8 = r
}

func BenchmarkRandom8(b *testing.B) {
	var r uint8
	for n := 0; n < b.N; n++ {
		r = fastmath.Random8Limit(255)
	}
	result8 = r
}
