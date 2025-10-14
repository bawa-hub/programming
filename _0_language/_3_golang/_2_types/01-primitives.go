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

// PrimitiveTypes demonstrates all basic Go primitive types
type PrimitiveTypes struct {
	// Integer types (signed and unsigned)
	Int8    int8    // -128 to 127
	Int16   int16   // -32,768 to 32,767
	Int32   int32   // -2,147,483,648 to 2,147,483,647
	Int64   int64   // -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
	Int     int     // Platform dependent (32 or 64 bits)
	
	Uint8   uint8   // 0 to 255 (also called byte)
	Uint16  uint16  // 0 to 65,535
	Uint32  uint32  // 0 to 4,294,967,295
	Uint64  uint64  // 0 to 18,446,744,073,709,551,615
	Uint    uint    // Platform dependent (32 or 64 bits)
	Uintptr uintptr // Unsigned integer type for storing pointer values
	
	// Floating point types
	Float32 float32 // IEEE-754 32-bit floating-point numbers
	Float64 float64 // IEEE-754 64-bit floating-point numbers
	
	// Complex number types
	Complex64  complex64  // Complex numbers with float32 real and imaginary parts
	Complex128 complex128 // Complex numbers with float64 real and imaginary parts
	
	// Boolean type
	Bool bool // true or false
	
	// String type
	String string // Immutable sequence of bytes
	
	// Character types
	Byte byte   // Alias for uint8
	Rune rune   // Alias for int32, represents a Unicode code point
}

func NewPrimitiveTypes() *PrimitiveTypes {
	return &PrimitiveTypes{
		Int8:      42,
		Int16:     1000,
		Int32:     100000,
		Int64:     1000000000,
		Int:       1000000000,
		Uint8:     200,
		Uint16:    50000,
		Uint32:    2000000000,
		Uint64:    9000000000000000000,
		Uint:      2000000000,
		Uintptr:   0x12345678,
		Float32:   3.14159,
		Float64:   3.141592653589793,
		Complex64: complex(1, 2),
		Complex128: complex(1.5, 2.5),
		Bool:      true,
		String:    "Hello, 世界! 🌍",
		Byte:      'A',
		Rune:      '🚀',
	}
}

func (pt *PrimitiveTypes) Create() {
	fmt.Println("🔧 Creating primitive type values...")
	
	// Demonstrate type conversion
	pt.Int8 = int8(pt.Int16) // Explicit conversion required
	pt.Float32 = float32(pt.Float64)
	pt.Complex64 = complex64(pt.Complex128)
	
	// String operations
	pt.String = "Golang Mastery: " + strconv.Itoa(int(pt.Int))
	
	fmt.Printf("✅ Created primitive values successfully\n")
}

// Read - Display all primitive values
func (pt *PrimitiveTypes) Read() {
	fmt.Println("\n📖 READING PRIMITIVE TYPES:")
	fmt.Println("==========================")
	
	// Integer types
	fmt.Printf("Integers:\n")
	fmt.Printf("  int8:   %d (range: %d to %d)\n", pt.Int8, math.MinInt8, math.MaxInt8)
	fmt.Printf("  int16:  %d (range: %d to %d)\n", pt.Int16, math.MinInt16, math.MaxInt16)
	fmt.Printf("  int32:  %d (range: %d to %d)\n", pt.Int32, math.MinInt32, math.MaxInt32)
	fmt.Printf("  int64:  %d (range: %d to %d)\n", pt.Int64, math.MinInt64, math.MaxInt64)
	fmt.Printf("  int:    %d\n", pt.Int)
	
	fmt.Printf("\nUnsigned Integers:\n")
	fmt.Printf("  uint8:   %d (range: 0 to %d)\n", pt.Uint8, math.MaxUint8)
	fmt.Printf("  uint16:  %d (range: 0 to %d)\n", pt.Uint16, math.MaxUint16)
	fmt.Printf("  uint32:  %d (range: 0 to %d)\n", pt.Uint32, math.MaxUint32)
	fmt.Printf("  uint64:  %d (range: 0 to %d)\n", pt.Uint64, uint64(math.MaxUint64))
	fmt.Printf("  uint:    %d\n", pt.Uint)
	fmt.Printf("  uintptr: 0x%x\n", pt.Uintptr)
	
	// Floating point types
	fmt.Printf("\nFloating Point:\n")
	fmt.Printf("  float32: %f (precision: 7 digits)\n", pt.Float32)
	fmt.Printf("  float64: %.15f (precision: 15 digits)\n", pt.Float64)
	
	// Complex numbers
	fmt.Printf("\nComplex Numbers:\n")
	fmt.Printf("  complex64:  %v\n", pt.Complex64)
	fmt.Printf("  complex128: %v\n", pt.Complex128)
	fmt.Printf("  real part:  %.2f\n", real(pt.Complex128))
	fmt.Printf("  imag part:  %.2f\n", imag(pt.Complex128))
	
	// Boolean
	fmt.Printf("\nBoolean:\n")
	fmt.Printf("  bool: %t\n", pt.Bool)
	
	// String and characters
	fmt.Printf("\nString and Characters:\n")
	fmt.Printf("  string: %s\n", pt.String)
	fmt.Printf("  length: %d bytes, %d runes\n", len(pt.String), utf8.RuneCountInString(pt.String))
	fmt.Printf("  byte:   %c (%d)\n", pt.Byte, pt.Byte)
	fmt.Printf("  rune:   %c (%d)\n", pt.Rune, pt.Rune)
}

// DemonstrateTypeConversions shows various type conversion techniques
func (pt *PrimitiveTypes) DemonstrateTypeConversions() {
	fmt.Println("\n🔄 TYPE CONVERSIONS DEMONSTRATION:")
	fmt.Println("==================================")
	
	// String to number conversions
	str := "12345"
	if num, err := strconv.Atoi(str); err == nil {
		fmt.Printf("String '%s' converted to int: %d\n", str, num)
	}
	
	// Number to string conversions
	fmt.Printf("Int to string: %s\n", strconv.Itoa(42))
	fmt.Printf("Float to string: %s\n", strconv.FormatFloat(3.14159, 'f', 2, 64))
	
	// Explicit type conversions
	var i int = 42
	var f float64 = float64(i)
	var s string = string(rune(i))
	fmt.Printf("int %d -> float64 %f -> string %s\n", i, f, s)
	
	// Rune and string conversions
	runes := []rune("Hello, 世界!")
	fmt.Printf("String to runes: %v\n", runes)
	fmt.Printf("Runes to string: %s\n", string(runes))
	
	// Byte and string conversions
	bytes := []byte("Hello")
	fmt.Printf("String to bytes: %v\n", bytes)
	fmt.Printf("Bytes to string: %s\n", string(bytes))
}

// DemonstrateConstants shows how to work with constants
func (pt *PrimitiveTypes) DemonstrateConstants() {
	fmt.Println("\n📌 CONSTANTS DEMONSTRATION:")
	fmt.Println("===========================")
	
	// Typed constants
	const (
		MaxUsers    = 1000
		Pi          = 3.14159
		AppName     = "Golang CRUD Mastery"
		IsActive    = true
	)
	
	// Untyped constants (can be used with any compatible type)
	const (
		untypedInt = 42
		untypedFloat = 3.14
		untypedString = "Hello"
	)
	
	// Using untyped constants with different types
	var intVar int = untypedInt
	var floatVar float64 = untypedFloat
	var stringVar string = untypedString
	
	fmt.Printf("Typed constants: %d, %.2f, %s, %t\n", MaxUsers, Pi, AppName, IsActive)
	fmt.Printf("Untyped constants: %d, %.2f, %s\n", intVar, floatVar, stringVar)
	
	// Iota for enumerated constants
	const (
		StatusPending = iota
		StatusActive
		StatusInactive
		StatusDeleted
	)
	
	fmt.Printf("Enum constants: Pending=%d, Active=%d, Inactive=%d, Deleted=%d\n",
		StatusPending, StatusActive, StatusInactive, StatusDeleted)
}

// DemonstratePrimitiveTypes shows all primitive data types in Go
func DemonstratePrimitiveTypes() {
	fmt.Println("=== PRIMITIVE DATA TYPES ===")
	
	// 1. INTEGER TYPES
	fmt.Println("\n--- Integer Types ---")
	
	// Signed integers
	var int8Var int8 = 127                    // -128 to 127
	var int16Var int16 = 32767                // -32768 to 32767
	var int32Var int32 = 2147483647           // -2^31 to 2^31-1
	var int64Var int64 = 9223372036854775807  // -2^63 to 2^63-1
	var intVar int = 42                       // Platform dependent (32 or 64 bit)
	
	fmt.Printf("int8: %d (size: %d bytes)\n", int8Var, unsafe.Sizeof(int8Var))
	fmt.Printf("int16: %d (size: %d bytes)\n", int16Var, unsafe.Sizeof(int16Var))
	fmt.Printf("int32: %d (size: %d bytes)\n", int32Var, unsafe.Sizeof(int32Var))
	fmt.Printf("int64: %d (size: %d bytes)\n", int64Var, unsafe.Sizeof(int64Var))
	fmt.Printf("int: %d (size: %d bytes)\n", intVar, unsafe.Sizeof(intVar))
	
	// Unsigned integers
	var uint8Var uint8 = 255                  // 0 to 255
	var uint16Var uint16 = 65535              // 0 to 65535
	var uint32Var uint32 = 4294967295         // 0 to 2^32-1
	var uint64Var uint64 = 18446744073709551615 // 0 to 2^64-1
	var uintVar uint = 42                     // Platform dependent
	
	fmt.Printf("uint8: %d (size: %d bytes)\n", uint8Var, unsafe.Sizeof(uint8Var))
	fmt.Printf("uint16: %d (size: %d bytes)\n", uint16Var, unsafe.Sizeof(uint16Var))
	fmt.Printf("uint32: %d (size: %d bytes)\n", uint32Var, unsafe.Sizeof(uint32Var))
	fmt.Printf("uint64: %d (size: %d bytes)\n", uint64Var, unsafe.Sizeof(uint64Var))
	fmt.Printf("uint: %d (size: %d bytes)\n", uintVar, unsafe.Sizeof(uintVar))
	
	// 2. FLOATING POINT TYPES
	fmt.Println("\n--- Floating Point Types ---")
	
	var float32Var float32 = 3.14159
	var float64Var float64 = 3.141592653589793
	
	fmt.Printf("float32: %f (size: %d bytes)\n", float32Var, unsafe.Sizeof(float32Var))
	fmt.Printf("float64: %f (size: %d bytes)\n", float64Var, unsafe.Sizeof(float64Var))
	
	// Special floating point values
	var inf float64 = math.Inf(1)   // Positive infinity
	var negInf float64 = math.Inf(-1) // Negative infinity
	var nan float64 = math.NaN()    // Not a Number
	
	fmt.Printf("Positive Infinity: %f\n", inf)
	fmt.Printf("Negative Infinity: %f\n", negInf)
	fmt.Printf("NaN: %f\n", nan)
	fmt.Printf("Is NaN: %t\n", math.IsNaN(nan))
	
	// 3. COMPLEX TYPES
	fmt.Println("\n--- Complex Types ---")
	
	var complex64Var complex64 = 3 + 4i
	var complex128Var complex128 = 3.14159 + 2.71828i
	
	fmt.Printf("complex64: %v (size: %d bytes)\n", complex64Var, unsafe.Sizeof(complex64Var))
	fmt.Printf("complex128: %v (size: %d bytes)\n", complex128Var, unsafe.Sizeof(complex128Var))
	
	// Complex number operations
	realPart := real(complex64Var)
	imagPart := imag(complex64Var)
	magnitude := math.Sqrt(float64(realPart*realPart + imagPart*imagPart))
	
	fmt.Printf("Real part: %f\n", realPart)
	fmt.Printf("Imaginary part: %f\n", imagPart)
	fmt.Printf("Magnitude: %f\n", magnitude)
	
	// 4. BOOLEAN TYPE
	fmt.Println("\n--- Boolean Type ---")
	
	var boolVar bool = true
	var falseVar bool = false
	
	fmt.Printf("bool: %t (size: %d bytes)\n", boolVar, unsafe.Sizeof(boolVar))
	fmt.Printf("false: %t\n", falseVar)
	
	// Boolean operations
	fmt.Printf("true && false: %t\n", boolVar && falseVar)
	fmt.Printf("true || false: %t\n", boolVar || falseVar)
	fmt.Printf("!true: %t\n", !boolVar)
	
	// 5. STRING TYPE
	fmt.Println("\n--- String Type ---")
	
	var stringVar string = "Hello, Go!"
	var emptyString string = ""
	
	fmt.Printf("string: %s (size: %d bytes)\n", stringVar, unsafe.Sizeof(stringVar))
	fmt.Printf("empty string: '%s' (size: %d bytes)\n", emptyString, unsafe.Sizeof(emptyString))
	
	// String length (in bytes, not characters)
	fmt.Printf("String length: %d bytes\n", len(stringVar))
	
	// 6. RUNE TYPE (Unicode code point)
	fmt.Println("\n--- Rune Type ---")
	
	var runeVar rune = 'A'
	var unicodeRune rune = '🚀'
	
	fmt.Printf("rune: %c (Unicode: %U, size: %d bytes)\n", runeVar, runeVar, unsafe.Sizeof(runeVar))
	fmt.Printf("Unicode rune: %c (Unicode: %U)\n", unicodeRune, unicodeRune)
	
	// String to rune slice
	runeSlice := []rune(stringVar)
	fmt.Printf("String as rune slice: %v\n", runeSlice)
}

// DemonstrateTypeConversions shows explicit and implicit type conversions
func DemonstrateTypeConversions() {
	fmt.Println("\n=== TYPE CONVERSIONS ===")
	
	// Explicit conversions
	var intVar int = 42
	var floatVar float64 = float64(intVar)
	var stringVar string = string(intVar) // This converts to Unicode character!
	
	fmt.Printf("int to float64: %d -> %f\n", intVar, floatVar)
	fmt.Printf("int to string: %d -> '%s' (Unicode char)\n", intVar, stringVar)
	
	// Proper int to string conversion
	stringVar = fmt.Sprintf("%d", intVar)
	fmt.Printf("int to string (proper): %d -> '%s'\n", intVar, stringVar)
	
	// String to int conversion
	var str string = "123"
	var convertedInt int
	_, err := fmt.Sscanf(str, "%d", &convertedInt)
	if err == nil {
		fmt.Printf("string to int: '%s' -> %d\n", str, convertedInt)
	}
	
	// Float conversions
	var float32Var float32 = 3.14
	var float64Var float64 = float64(float32Var)
	var intFromFloat int = int(float64Var) // Truncates decimal part
	
	fmt.Printf("float32 to float64: %f -> %f\n", float32Var, float64Var)
	fmt.Printf("float64 to int: %f -> %d (truncated)\n", float64Var, intFromFloat)
}

// DemonstrateConstants shows const declarations and iota
func DemonstrateConstants() {
	fmt.Println("\n=== CONSTANTS ===")
	
	// Basic constants
	const pi = 3.14159
	const e = 2.71828
	const greeting = "Hello, World!"
	
	fmt.Printf("Pi: %f\n", pi)
	fmt.Printf("E: %f\n", e)
	fmt.Printf("Greeting: %s\n", greeting)
	
	// Typed constants
	const typedPi float64 = 3.14159
	const typedInt int = 42
	
	fmt.Printf("Typed Pi: %f (type: %T)\n", typedPi, typedPi)
	fmt.Printf("Typed Int: %d (type: %T)\n", typedInt, typedInt)
	
	// Iota - Go's enum-like constant generator
	const (
		Sunday = iota    // 0
		Monday           // 1
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)
	
	fmt.Println("\nDays of the week (using iota):")
	fmt.Printf("Sunday: %d\n", Sunday)
	fmt.Printf("Monday: %d\n", Monday)
	fmt.Printf("Tuesday: %d\n", Tuesday)
	fmt.Printf("Wednesday: %d\n", Wednesday)
	fmt.Printf("Thursday: %d\n", Thursday)
	fmt.Printf("Friday: %d\n", Friday)
	fmt.Printf("Saturday: %d\n", Saturday)
	
	// Iota with expressions
	const (
		_ = iota             // Skip first value
		KB = 1 << (10 * iota) // 1024
		MB                    // 2048
		GB                    // 4096
		TB                    // 8192
	)
	
	fmt.Println("\nFile sizes (using iota with expressions):")
	fmt.Printf("KB: %d\n", KB)
	fmt.Printf("MB: %d\n", MB)
	fmt.Printf("GB: %d\n", GB)
	fmt.Printf("TB: %d\n", TB)
}

// DemonstrateZeroValues shows zero values for all types
func DemonstrateZeroValues() {
	fmt.Println("\n=== ZERO VALUES ===")
	
	var intZero int
	var floatZero float64
	var boolZero bool
	var stringZero string
	var runeZero rune
	var complexZero complex128
	
	fmt.Printf("int zero value: %d\n", intZero)
	fmt.Printf("float64 zero value: %f\n", floatZero)
	fmt.Printf("bool zero value: %t\n", boolZero)
	fmt.Printf("string zero value: '%s' (empty string)\n", stringZero)
	fmt.Printf("rune zero value: %c (Unicode: %U)\n", runeZero, runeZero)
	fmt.Printf("complex128 zero value: %v\n", complexZero)
}

// RunAllPrimitiveExamples runs all primitive type examples
func RunAllPrimitiveExamples() {
	DemonstratePrimitiveTypes()
	DemonstrateTypeConversions()
	DemonstrateConstants()
	DemonstrateZeroValues()
}