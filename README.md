# FastMath for Go

> "I can make things very fast if they don’t have to be correct." — Russ Cox

8 and 16 bit math functions for when speed matters more than precision. Potential use cases include LED displays, 2D/3D graphics, and games.

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

## TODO

* Validate sin/cos function approximation errors.
* Add assembly implementations from upstream, benchmark difference.

## License

Licensed MIT

© 2019 Brad Erickson

Based on FastLED MIT-licensed code

© 2013 FastLED
