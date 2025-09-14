package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestFmtPrintFunctions(t *testing.T) {
	// Test Print
	var buf bytes.Buffer
	fmt.Fprint(&buf, "Hello", " ", "World")
	expected := "Hello World"
	if buf.String() != expected {
		t.Errorf("Print failed. Expected: %s, Got: %s", expected, buf.String())
	}

	// Test Println
	buf.Reset()
	fmt.Fprintln(&buf, "Hello", "World")
	expected = "Hello World\n"
	if buf.String() != expected {
		t.Errorf("Println failed. Expected: %s, Got: %s", expected, buf.String())
	}

	// Test Printf
	buf.Reset()
	fmt.Fprintf(&buf, "Name: %s, Age: %d", "Alice", 30)
	expected = "Name: Alice, Age: 30"
	if buf.String() != expected {
		t.Errorf("Printf failed. Expected: %s, Got: %s", expected, buf.String())
	}
}

func TestFmtSprintFunctions(t *testing.T) {
	// Test Sprint
	result := fmt.Sprint("Hello", " ", "World")
	expected := "Hello World"
	if result != expected {
		t.Errorf("Sprint failed. Expected: %s, Got: %s", expected, result)
	}

	// Test Sprintln
	result = fmt.Sprintln("Hello", "World")
	expected = "Hello World\n"
	if result != expected {
		t.Errorf("Sprintln failed. Expected: %s, Got: %s", expected, result)
	}

	// Test Sprintf
	result = fmt.Sprintf("Name: %s, Age: %d", "Alice", 30)
	expected = "Name: Alice, Age: 30"
	if result != expected {
		t.Errorf("Sprintf failed. Expected: %s, Got: %s", expected, result)
	}
}

func TestFmtFormatVerbs(t *testing.T) {
	tests := []struct {
		format string
		value  interface{}
		expected string
	}{
		{"%v", 42, "42"},
		{"%d", 42, "42"},
		{"%b", 42, "101010"},
		{"%o", 42, "52"},
		{"%x", 42, "2a"},
		{"%X", 42, "2A"},
		{"%f", 3.14, "3.140000"},
		{"%.2f", 3.14, "3.14"},
		{"%s", "hello", "hello"},
		{"%q", "hello", "\"hello\""},
		{"%t", true, "true"},
		{"%T", 42, "int"},
	}

	for _, test := range tests {
		result := fmt.Sprintf(test.format, test.value)
		if result != test.expected {
			t.Errorf("Format %s failed. Expected: %s, Got: %s", 
				test.format, test.expected, result)
		}
	}
}

func TestFmtStructFormatting(t *testing.T) {
	person := Person{Name: "Bob", Age: 25, City: "New York"}
	
	// Test %v
	result := fmt.Sprintf("%v", person)
	expected := "Person{Name: Bob, Age: 25, City: New York}"
	if result != expected {
		t.Errorf("Struct %v failed. Expected: %s, Got: %s", expected, result)
	}

	// Test %+v
	result = fmt.Sprintf("%+v", person)
	expected = "{Name:Bob Age:25 City:New York}"
	if result != expected {
		t.Errorf("Struct %+v failed. Expected: %s, Got: %s", expected, result)
	}

	// Test %#v
	result = fmt.Sprintf("%#v", person)
	expected = "main.Person{Name:\"Bob\", Age:25, City:\"New York\"}"
	if result != expected {
		t.Errorf("Struct %#v failed. Expected: %s, Got: %s", expected, result)
	}
}

func TestFmtErrorFormatting(t *testing.T) {
	err := CustomError{Code: 404, Message: "Not Found"}
	
	result := fmt.Sprintf("%v", err)
	expected := "Error 404: Not Found"
	if result != expected {
		t.Errorf("Error formatting failed. Expected: %s, Got: %s", expected, result)
	}
}

func TestFmtWidthAndPrecision(t *testing.T) {
	tests := []struct {
		format string
		value  interface{}
		expected string
	}{
		{"%5d", 42, "   42"},
		{"%05d", 42, "00042"},
		{"%-5d", 42, "42   "},
		{"%5.2f", 3.14159, " 3.14"},
		{"%05.2f", 3.14159, "03.14"},
		{"%-5.2f", 3.14159, "3.14 "},
	}

	for _, test := range tests {
		result := fmt.Sprintf(test.format, test.value)
		if result != test.expected {
			t.Errorf("Width/Precision %s failed. Expected: %s, Got: %s", 
				test.format, test.expected, result)
		}
	}
}

func TestFmtSliceAndMapFormatting(t *testing.T) {
	// Test slice formatting
	numbers := []int{1, 2, 3, 4, 5}
	result := fmt.Sprintf("%v", numbers)
	expected := "[1 2 3 4 5]"
	if result != expected {
		t.Errorf("Slice formatting failed. Expected: %s, Got: %s", expected, result)
	}

	// Test map formatting
	colors := map[string]string{"red": "#FF0000", "green": "#00FF00"}
	result = fmt.Sprintf("%v", colors)
	// Map order is not guaranteed, so we check if it contains the expected elements
	if !strings.Contains(result, "red:#FF0000") || !strings.Contains(result, "green:#00FF00") {
		t.Errorf("Map formatting failed. Got: %s", result)
	}
}

func TestFmtScanFunctions(t *testing.T) {
	// Test Sscan
	var name string
	var age int
	n, err := fmt.Sscan("Alice 30", &name, &age)
	if err != nil {
		t.Errorf("Sscan failed: %v", err)
	}
	if n != 2 {
		t.Errorf("Sscan returned wrong count. Expected: 2, Got: %d", n)
	}
	if name != "Alice" || age != 30 {
		t.Errorf("Sscan values wrong. Expected: Alice 30, Got: %s %d", name, age)
	}

	// Test Sscanf
	var name2 string
	var age2 int
	n, err = fmt.Sscanf("Name: Bob, Age: 25", "Name: %s, Age: %d", &name2, &age2)
	if err != nil {
		t.Errorf("Sscanf failed: %v", err)
	}
	if n != 2 {
		t.Errorf("Sscanf returned wrong count. Expected: 2, Got: %d", n)
	}
	if name2 != "Bob" || age2 != 25 {
		t.Errorf("Sscanf values wrong. Expected: Bob 25, Got: %s %d", name2, age2)
	}
}

func TestFmtStringMethod(t *testing.T) {
	person := Person{Name: "Charlie", Age: 35, City: "London"}
	
	// Test String() method
	result := person.String()
	expected := "Person{Name: Charlie, Age: 35, City: London}"
	if result != expected {
		t.Errorf("String() method failed. Expected: %s, Got: %s", expected, result)
	}
}

func TestFmtErrorMethod(t *testing.T) {
	err := CustomError{Code: 500, Message: "Internal Server Error"}
	
	// Test Error() method
	result := err.Error()
	expected := "Error 500: Internal Server Error"
	if result != expected {
		t.Errorf("Error() method failed. Expected: %s, Got: %s", expected, result)
	}
}

func TestFmtPointerFormatting(t *testing.T) {
	number := 42
	ptr := &number
	
	result := fmt.Sprintf("%p", ptr)
	// Pointer address is not predictable, so we just check it's not empty
	if result == "" {
		t.Errorf("Pointer formatting failed. Got empty string")
	}
	
	// Test pointer value
	result = fmt.Sprintf("%v", ptr)
	expected := "0x" // Should start with 0x
	if !strings.HasPrefix(result, "0x") {
		t.Errorf("Pointer value formatting failed. Expected to start with 0x, Got: %s", result)
	}
}

func TestFmtTypeFormatting(t *testing.T) {
	tests := []struct {
		value    interface{}
		expected string
	}{
		{42, "int"},
		{"hello", "string"},
		{true, "bool"},
		{3.14, "float64"},
		{[]int{1, 2, 3}, "[]int"},
		{map[string]int{"a": 1}, "map[string]int"},
	}

	for _, test := range tests {
		result := fmt.Sprintf("%T", test.value)
		if result != test.expected {
			t.Errorf("Type formatting failed. Expected: %s, Got: %s", test.expected, result)
		}
	}
}

// Benchmark tests
func BenchmarkFmtSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("Name: %s, Age: %d", "Alice", 30)
	}
}

func BenchmarkFmtSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint("Name: ", "Alice", ", Age: ", 30)
	}
}

func BenchmarkFmtSprintln(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintln("Name:", "Alice", "Age:", 30)
	}
}
