# FastMath for Go

8 and 16 bit math functions for when speed matters more than precision.

> "I can make things very fast if they don’t have to be correct." — Russ Cox

* Designed for use with TinyGo or Web Assembly.
* Based on FastLED's lib8tion

## Functions

* `Sin8()` / `Sin16()`
* `Cos8()` / `Cos16()`
* `Random8()` / `Random16()`
* `Random8Limit()` / `Random16Limit()`
* `Random8Range()` / `Random16Range()`
* `Random16SetSeed()` / `Random16GetSeed()` / `Random16AddEntropy()`
* `Scale8()` / `Scale8Video()`
* `NScale8x3()` / `NScale8x3Video()`
* `Scale16()` / `Scale16By8()`

## TODO

* Add ARM Assembly implementations, benchmark difference.
