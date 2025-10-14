package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
)

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
