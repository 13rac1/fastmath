# FastMath for Go

> "I can make things very fast if they don’t have to be correct." — Russ Cox

8 and 16 bit math functions for when speed matters more than precision.
Potential use cases include LED displays, 2D/3D graphics, and games.

* Designed for use with [TinyGo][tinygo] and/or [WebAssembly][go-wasm].
* Based on the [FastLED][fastled] [lib8tion][lib8ation-src] library.

[tinygo]:https://tinygo.org/
[go-wasm]:https://github.com/golang/go/wiki/WebAssembly
[fastled]:http://fastled.io/
[lib8ation-src]:https://raw.githubusercontent.com/FastLED/FastLED/dcbf3993/lib8tion/trig8.h

## Functions

* `Sin8()` / `Sin16()`
* `Cos8()` / `Cos16()`
* `Random8()` / `Random16()`
* `Random8Limit()` / `Random16Limit()`
* `Random8Range()` / `Random16Range()`
* `Random16SetSeed()` / `Random16GetSeed()` / `Random16AddEntropy()`
* `Scale8()` / `Scale8Video()`
* `NScale8x3()` / `NScale8x3Video()` - Scale three 8 bit integers at the same time.
* `Scale16()` / `Scale16By8()`
* `QAdd8()` / `QSub8()` / `QMul8()` - Saturating non-overflowing math functions.
* `Abs8()`
* `Sqrt16()`

Note: Functionality already well handled by the Go runtime has not be re-implemented.

## Approximation Error

Computer-based math functions have an error delta verses the pure mathematical
results. The Golang Standard Library's math functions are precise up to 64 bit
floats. The math functions provided by this library sacrifice additional
precision for speed by working with small integers.

* `Sin8()` - Max Error: 1.63%, Average Error: 0.78%
* `Sin16()` - Max Error: 0.34%, Average Error: 0.19%

## Benchmarks

Run on a Intel(R) Core(TM) i7-7600U CPU @ 2.80GHz.

```bash
BenchmarkStdLibFallbackSqrt-4   20000000                 50.7 ns/op
BenchmarkStdLibDefaultSqrt-4    2000000000               0.30 ns/op
BenchmarkSqrt16-4               200000000                9.12 ns/op
BenchmarkStdLibRandom8-4        50000000                 25.6 ns/op
BenchmarkRandom8-4              1000000000               2.12 ns/op
BenchmarkStdLibSin8-4           50000000                 20.8 ns/op
BenchmarkSin8-4                 300000000                4.07 ns/op
BenchmarkStdLibSin16-4          50000000                 20.0 ns/op
BenchmarkSin16-4                2000000000               0.94 ns/op
```

`Random8()`, `Sin8()` and `Sin16()` are significantly faster than using the
equivalent Go Standard Library's Math package functions.

`Sqrt16()` is compared against both the default compiled `math.Sqrt()` and a
copy of the fallback Standard Library `sqrt()` function. The default version is
optimized by the Go compiler into a single instruction on the AMD64
architecture, so the fallback version is used for a fair comparison.

## TODO

* Should `Sin8()` be a lookup table? Why is it 4ns/op vs `Sin16()`@0.9ns/op?
* Add ARM assembly implementations from upstream, benchmark difference.

## License

Licensed MIT

© 2019 Brad Erickson

Based on FastLED MIT-licensed code:

© 2013 FastLED

Parts of test-only BSD code:

© 2009 The Go Authors
