package primitives

import (
	"fmt"
	"math"
	"unsafe"
)

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
