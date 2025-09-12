package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

// Custom conversion functions for demonstration
func SafeAtoi(s string) (int, error) {
	if s == "" {
		return 0, fmt.Errorf("empty string")
	}
	return strconv.Atoi(s)
}

func IsNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func IsFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func FormatNumber(n int64) string {
	return strconv.FormatInt(n, 10)
}

func FormatFloat(f float64, precision int) string {
	return strconv.FormatFloat(f, 'f', precision, 64)
}

func ParseHex(s string) (int64, error) {
	return strconv.ParseInt(s, 16, 64)
}

func ParseBinary(s string) (int64, error) {
	return strconv.ParseInt(s, 2, 64)
}

func ParseOctal(s string) (int64, error) {
	return strconv.ParseInt(s, 8, 64)
}

func IsValidBase(s string, base int) bool {
	_, err := strconv.ParseInt(s, base, 64)
	return err == nil
}

func ConvertBase(s string, fromBase, toBase int) (string, error) {
	// Parse from source base
	val, err := strconv.ParseInt(s, fromBase, 64)
	if err != nil {
		return "", err
	}
	
	// Format to target base
	return strconv.FormatInt(val, toBase), nil
}

func FormatComplex(c complex128) string {
	return strconv.FormatComplex(c, 'f', 2, 128)
}

func ParseComplex(s string) (complex128, error) {
	return strconv.ParseComplex(s, 128)
}

func IsPrintable(s string) bool {
	for _, r := range s {
		if !strconv.IsPrint(r) {
			return false
		}
	}
	return true
}

func IsGraphic(s string) bool {
	for _, r := range s {
		if !strconv.IsGraphic(r) {
			return false
		}
	}
	return true
}

func CanBackquote(s string) bool {
	return strconv.CanBackquote(s)
}

func QuoteString(s string) string {
	return strconv.Quote(s)
}

func UnquoteString(s string) (string, error) {
	return strconv.Unquote(s)
}

func QuoteRune(r rune) string {
	return strconv.QuoteRune(r)
}

func UnquoteRune(s string) (rune, error) {
	r, _, _, err := strconv.UnquoteChar(s, 0)
	return r, err
}

func AppendIntToSlice(slice []byte, n int64) []byte {
	return strconv.AppendInt(slice, n, 10)
}

func AppendFloatToSlice(slice []byte, f float64) []byte {
	return strconv.AppendFloat(slice, f, 'f', 2, 64)
}

func AppendBoolToSlice(slice []byte, b bool) []byte {
	return strconv.AppendBool(slice, b)
}

func AppendQuoteToSlice(slice []byte, s string) []byte {
	return strconv.AppendQuote(slice, s)
}

func main() {
	fmt.Println("🚀 Go strconv Package Mastery Examples")
	fmt.Println("=====================================")

	// 1. Basic String to Number Conversion
	fmt.Println("\n1. Basic String to Number Conversion:")
	
	// String to int
	s := "123"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("Error converting '%s': %v\n", s, err)
	} else {
		fmt.Printf("'%s' -> %d\n", s, i)
	}
	
	// String to int64
	s = "9223372036854775807"
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Printf("Error converting '%s': %v\n", s, err)
	} else {
		fmt.Printf("'%s' -> %d\n", s, i64)
	}
	
	// String to uint64
	s = "18446744073709551615"
	u64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		fmt.Printf("Error converting '%s': %v\n", s, err)
	} else {
		fmt.Printf("'%s' -> %d\n", s, u64)
	}

	// 2. String to Float Conversion
	fmt.Println("\n2. String to Float Conversion:")
	
	// String to float64
	s = "3.14159"
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("Error converting '%s': %v\n", s, err)
	} else {
		fmt.Printf("'%s' -> %f\n", s, f)
	}
	
	// String to float32
	s = "2.71828"
	f32, err := strconv.ParseFloat(s, 32)
	if err != nil {
		fmt.Printf("Error converting '%s': %v\n", s, err)
	} else {
		fmt.Printf("'%s' -> %f\n", s, f32)
	}

	// 3. String to Bool Conversion
	fmt.Println("\n3. String to Bool Conversion:")
	
	boolStrings := []string{"true", "false", "1", "0", "TRUE", "FALSE", "yes", "no"}
	for _, s := range boolStrings {
		b, err := strconv.ParseBool(s)
		if err != nil {
			fmt.Printf("'%s' -> error: %v\n", s, err)
		} else {
			fmt.Printf("'%s' -> %t\n", s, b)
		}
	}

	// 4. String to Complex Conversion
	fmt.Println("\n4. String to Complex Conversion:")
	
	complexStrings := []string{"1+2i", "3.14+2.71i", "1", "2i", "0+1i"}
	for _, s := range complexStrings {
		c, err := strconv.ParseComplex(s, 128)
		if err != nil {
			fmt.Printf("'%s' -> error: %v\n", s, err)
		} else {
			fmt.Printf("'%s' -> %v\n", s, c)
		}
	}

	// 5. Number to String Conversion
	fmt.Println("\n5. Number to String Conversion:")
	
	// Int to string
	i = 123
	s = strconv.Itoa(i)
	fmt.Printf("%d -> '%s'\n", i, s)
	
	// Int64 to string
	i64 = 9223372036854775807
	s = strconv.FormatInt(i64, 10)
	fmt.Printf("%d -> '%s'\n", i64, s)
	
	// Uint64 to string
	u64 = 18446744073709551615
	s = strconv.FormatUint(u64, 10)
	fmt.Printf("%d -> '%s'\n", u64, s)

	// 6. Float to String Conversion
	fmt.Println("\n6. Float to String Conversion:")
	
	// Float64 to string
	f = 3.14159
	s = strconv.FormatFloat(f, 'f', 2, 64)
	fmt.Printf("%f -> '%s' (2 decimal places)\n", f, s)
	
	s = strconv.FormatFloat(f, 'e', 2, 64)
	fmt.Printf("%f -> '%s' (scientific notation)\n", f, s)
	
	s = strconv.FormatFloat(f, 'g', 2, 64)
	fmt.Printf("%f -> '%s' (compact notation)\n", f, s)

	// 7. Bool to String Conversion
	fmt.Println("\n7. Bool to String Conversion:")
	
	b := true
	s = strconv.FormatBool(b)
	fmt.Printf("%t -> '%s'\n", b, s)
	
	b = false
	s = strconv.FormatBool(b)
	fmt.Printf("%t -> '%s'\n", b, s)

	// 8. Complex to String Conversion
	fmt.Println("\n8. Complex to String Conversion:")
	
	c := complex(1, 2)
	s = strconv.FormatComplex(c, 'f', 2, 128)
	fmt.Printf("%v -> '%s'\n", c, s)
	
	c = complex(3.14, 2.71)
	s = strconv.FormatComplex(c, 'f', 2, 128)
	fmt.Printf("%v -> '%s'\n", c, s)

	// 9. String Quoting
	fmt.Println("\n9. String Quoting:")
	
	strings := []string{"hello", "world", "hello world", "hello\nworld", "hello\tworld"}
	for _, s := range strings {
		quoted := strconv.Quote(s)
		fmt.Printf("'%s' -> %s\n", s, quoted)
	}

	// 10. String Unquoting
	fmt.Println("\n10. String Unquoting:")
	
	quotedStrings := []string{`"hello"`, `"hello world"`, `"hello\nworld"`, `"hello\tworld"`}
	for _, s := range quotedStrings {
		unquoted, err := strconv.Unquote(s)
		if err != nil {
			fmt.Printf("%s -> error: %v\n", s, err)
		} else {
			fmt.Printf("%s -> '%s'\n", s, unquoted)
		}
	}

	// 11. Rune Quoting
	fmt.Println("\n11. Rune Quoting:")
	
	runes := []rune{'a', 'A', '1', ' ', '\n', '\t', '中', '🌍'}
	for _, r := range runes {
		quoted := strconv.QuoteRune(r)
		fmt.Printf("'%c' -> %s\n", r, quoted)
	}

	// 12. Different Number Bases
	fmt.Println("\n12. Different Number Bases:")
	
	// Binary
	s = "1010"
	i64, err = strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Printf("Error parsing binary '%s': %v\n", s, err)
	} else {
		fmt.Printf("Binary '%s' -> %d\n", s, i64)
	}
	
	// Octal
	s = "777"
	i64, err = strconv.ParseInt(s, 8, 64)
	if err != nil {
		fmt.Printf("Error parsing octal '%s': %v\n", s, err)
	} else {
		fmt.Printf("Octal '%s' -> %d\n", s, i64)
	}
	
	// Hexadecimal
	s = "FF"
	i64, err = strconv.ParseInt(s, 16, 64)
	if err != nil {
		fmt.Printf("Error parsing hex '%s': %v\n", s, err)
	} else {
		fmt.Printf("Hex '%s' -> %d\n", s, i64)
	}

	// 13. Number to Different Bases
	fmt.Println("\n13. Number to Different Bases:")
	
	n := 255
	fmt.Printf("%d in binary: %s\n", n, strconv.FormatInt(int64(n), 2))
	fmt.Printf("%d in octal: %s\n", n, strconv.FormatInt(int64(n), 8))
	fmt.Printf("%d in decimal: %s\n", n, strconv.FormatInt(int64(n), 10))
	fmt.Printf("%d in hex: %s\n", n, strconv.FormatInt(int64(n), 16))

	// 14. String Appending
	fmt.Println("\n14. String Appending:")
	
	slice := []byte("Numbers: ")
	slice = strconv.AppendInt(slice, 123, 10)
	slice = append(slice, []byte(", ")...)
	slice = strconv.AppendFloat(slice, 3.14, 'f', 2, 64)
	slice = append(slice, []byte(", ")...)
	slice = strconv.AppendBool(slice, true)
	
	fmt.Printf("Appended: %s\n", string(slice))

	// 15. String Validation
	fmt.Println("\n15. String Validation:")
	
	testStrings := []string{"hello", "hello world", "hello\nworld", "hello\tworld", "hello世界"}
	for _, s := range testStrings {
		fmt.Printf("'%s': printable=%t, graphic=%t, backquote=%t\n", 
			s, IsPrintable(s), IsGraphic(s), CanBackquote(s))
	}

	// 16. Error Handling
	fmt.Println("\n16. Error Handling:")
	
	invalidStrings := []string{"", "abc", "12.34.56", "not-a-number"}
	for _, s := range invalidStrings {
		_, err := strconv.Atoi(s)
		if err != nil {
			switch e := err.(type) {
			case *strconv.NumError:
				fmt.Printf("'%s': %s (func=%s, num=%s)\n", s, e.Err, e.Func, e.Num)
			default:
				fmt.Printf("'%s': %v\n", s, err)
			}
		}
	}

	// 17. Custom Conversion Functions
	fmt.Println("\n17. Custom Conversion Functions:")
	
	// Safe Atoi
	numbers := []string{"123", "", "abc", "12.34"}
	for _, s := range numbers {
		n, err := SafeAtoi(s)
		if err != nil {
			fmt.Printf("SafeAtoi('%s'): error: %v\n", s, err)
		} else {
			fmt.Printf("SafeAtoi('%s'): %d\n", s, n)
		}
	}
	
	// Is Numeric
	strings = []string{"123", "12.34", "abc", "12.34.56"}
	for _, s := range strings {
		fmt.Printf("IsNumeric('%s'): %t\n", s, IsNumeric(s))
		fmt.Printf("IsFloat('%s'): %t\n", s, IsFloat(s))
	}

	// 18. Base Conversion
	fmt.Println("\n18. Base Conversion:")
	
	// Convert between bases
	number := "255"
	bases := []struct{ from, to int }{
		{10, 2},  // decimal to binary
		{10, 8},  // decimal to octal
		{10, 16}, // decimal to hex
		{2, 10},  // binary to decimal
		{8, 10},  // octal to decimal
		{16, 10}, // hex to decimal
	}
	
	for _, base := range bases {
		result, err := ConvertBase(number, base.from, base.to)
		if err != nil {
			fmt.Printf("Convert %s from base %d to %d: error: %v\n", number, base.from, base.to, err)
		} else {
			fmt.Printf("Convert %s from base %d to %d: %s\n", number, base.from, base.to, result)
		}
	}

	// 19. Float Formatting Options
	fmt.Println("\n19. Float Formatting Options:")
	
	f = 3.141592653589793
	formats := []struct{ format byte; precision int; name string }{
		{'f', 2, "fixed"},
		{'f', 6, "fixed (6 decimal)"},
		{'e', 2, "scientific"},
		{'e', 6, "scientific (6 decimal)"},
		{'g', 2, "compact"},
		{'g', 6, "compact (6 decimal)"},
	}
	
	for _, format := range formats {
		s := strconv.FormatFloat(f, format.format, format.precision, 64)
		fmt.Printf("%s: %s\n", format.name, s)
	}

	// 20. Performance Test
	fmt.Println("\n20. Performance Test:")
	
	// Test string conversion performance
	numbersPerf := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	// Method 1: Using strconv.Itoa
	start := time.Now()
	var result1 []string
	for _, n := range numbersPerf {
		result1 = append(result1, strconv.Itoa(n))
	}
	time1 := time.Since(start)
	
	// Method 2: Using strconv.FormatInt
	start = time.Now()
	var result2 []string
	for _, n := range numbersPerf {
		result2 = append(result2, strconv.FormatInt(int64(n), 10))
	}
	time2 := time.Since(start)
	
	fmt.Printf("Itoa method: %v (result: %v)\n", time1, result1)
	fmt.Printf("FormatInt method: %v (result: %v)\n", time2, result2)
	
	if time2 < time1 {
		fmt.Printf("FormatInt is %.2fx faster than Itoa\n", float64(time1)/float64(time2))
	} else {
		fmt.Printf("Itoa is %.2fx faster than FormatInt\n", float64(time2)/float64(time1))
	}

	// 21. Edge Cases
	fmt.Println("\n21. Edge Cases:")
	
	// Maximum values
	maxInt64 := int64(9223372036854775807)
	maxUint64 := uint64(18446744073709551615)
	maxFloat64 := math.MaxFloat64
	
	fmt.Printf("Max int64: %d\n", maxInt64)
	fmt.Printf("Max uint64: %d\n", maxUint64)
	fmt.Printf("Max float64: %e\n", maxFloat64)
	
	// Minimum values
	minInt64 := int64(-9223372036854775808)
	minFloat64 := math.SmallestNonzeroFloat64
	
	fmt.Printf("Min int64: %d\n", minInt64)
	fmt.Printf("Min float64: %e\n", minFloat64)

	// 22. String to Rune Conversion
	fmt.Println("\n22. String to Rune Conversion:")
	
	s = "Hello 世界 🌍"
	for i, r := range s {
		fmt.Printf("Position %d: '%c' (rune: %d)\n", i, r, r)
	}

	// 23. Rune to String Conversion
	fmt.Println("\n23. Rune to String Conversion:")
	
	runesConv := []rune{'H', 'e', 'l', 'l', 'o', ' ', '世', '界', ' ', '🌍'}
	sRunes := string(runesConv)
	fmt.Printf("Runes to string: '%s'\n", sRunes)

	// 24. String Length vs Rune Count
	fmt.Println("\n24. String Length vs Rune Count:")
	
	stringsLen := []string{"hello", "世界", "🌍", "hello世界🌍"}
	for _, s := range stringsLen {
		fmt.Printf("'%s': bytes=%d, runes=%d\n", s, len(s), len([]rune(s)))
	}

	// 25. Custom Number Formatting
	fmt.Println("\n25. Custom Number Formatting:")
	
	numbersFmt := []int64{123, 1234, 12345, 123456, 1234567}
	for _, n := range numbersFmt {
		fmt.Printf("Number: %d, Formatted: %s\n", n, FormatNumber(n))
	}
	
	floats := []float64{3.14159, 2.71828, 1.41421, 1.73205}
	for _, f := range floats {
		fmt.Printf("Float: %f, Formatted (2): %s\n", f, FormatFloat(f, 2))
		fmt.Printf("Float: %f, Formatted (4): %s\n", f, FormatFloat(f, 4))
	}

	fmt.Println("\n🎉 strconv Package Mastery Complete!")
}
