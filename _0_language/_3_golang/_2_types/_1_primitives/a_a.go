package main

import (
	"fmt"
	"unicode/utf8"
)

var pl = fmt.Println

func main() {

	// 1) zero values
	// Every Go variable has a static type and a zero value when not initialized (e.g., int → 0, string → "", bool → false, pointer/map/slice/channel/function → nil).
	// Typed vs untyped constants: const x = 1 is untyped until used; untyped constants can be more flexible (higher precision) than typed ones.

	var a int    // a == 0
	var s string // s == ""
	const c = 1  // untyped constant
	pl(a, s, c)

	// 2) Numeric types (integers, floats, complex)
	// Integer family
	// Signed: int8, int16, int32, int64, int (size depends on architecture; 64-bit on typical x86-64).
	// Unsigned: uint8 (alias byte), uint16, uint32, uint64, uint.
	// uintptr used to hold pointer addresses for low-level code.

	// Float
	// float32, float64 — use float64 by default for precision.

	// Complex
	// complex64 (two float32 parts), complex128 (two float64 parts).

	// // 3) Booleans
	// Type bool with true/false. No numeric-boolean intermixing (unlike C).

	// 4) Strings, bytes, and runes (UTF-8)
	// A Go string is an immutable sequence of bytes (UTF-8 by convention).
	// len(s) returns bytes, not runes (Unicode code points).
	// rune = alias for int32 and represents a Unicode code point.
	// Convert between bytes and runes:
	// []byte(s) for raw bytes
	// []rune(s) to operate on runes
	// For building strings efficiently, use strings.Builder (or bytes.Buffer).

	st := "日本"                     // bytes: 6, runes: 2
	pl(len(st))                    // prints 6
	pl(utf8.RuneCountInString(st)) // prints 2

	bt := []byte("hello")
	ru := []rune("héllo") // r[1] is 'é' as rune
	pl(bt, ru)

	// Common pitfalls:
	// Indexing a string gives bytes: s[0] is a byte.
	// To iterate runes correctly:

	str := "vikram"
	for i, r := range str {
		fmt.Printf("%d: %c\n", i, r) // i is byte index, r is rune
	}

	// Common pitfalls & best practices (interview-friendly)
	// Implicit conversions: none — always convert types explicitly.
	// Nil maps: writing causes panic — initialize with make.
	// Slices share memory: be careful when slicing or passing to goroutines.
	// Iteration order for map is randomized: don’t assume order.
	// Strings are immutable: convert to []byte or []rune to modify.
	// Use strings.Builder for heavy concatenation.
	// Prefer int for counts/indices unless you need a specific width; use int64 for DB IDs or when a specific width is required.
	// Avoid unsafe unless necessary.
	// For concurrency, remember that maps and slices are not safe for concurrent writes; synchronize.

	// Handy cheat-sheet (typical sizes on 64-bit)
	// bool — 1 byte
	// int, uint — 8 bytes (on 64-bit arch)
	// int8, uint8 (byte) — 1 byte
	// int16/uint16 — 2 bytes
	// int32/uint32 — 4 bytes
	// int64/uint64 — 8 bytes
	// float32 — 4 bytes; float64 — 8 bytes
	// complex64 — 8 bytes; complex128 — 16 bytes
	// string — 16 bytes (two-word header: pointer + length)
	// slice — 24 bytes (pointer + len + cap)
	// map/channel/func — runtime-dependent header pointers
	// (Use unsafe.Sizeof() in Go to measure on your machine.)

}
