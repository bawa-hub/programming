package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

// Custom types for demonstration
type Person struct {
	Name string
	Age  int
	City string
}

type Point struct {
	X, Y float64
}

type Data struct {
	ID    int
	Value float64
	Flag  bool
	Text  string
}

// Custom memory operations
func InspectStruct(v interface{}) {
	val := reflect.ValueOf(v)
	typ := val.Type()
	
	fmt.Printf("Type: %s\n", typ.Name())
	fmt.Printf("Size: %d bytes\n", unsafe.Sizeof(v))
	fmt.Printf("Alignment: %d bytes\n", unsafe.Alignof(v))
	
		if val.Kind() == reflect.Struct {
			for i := 0; i < val.NumField(); i++ {
				field := val.Field(i)
				fieldType := typ.Field(i)
				// Note: unsafe.Offsetof requires a struct field, not a value
				fmt.Printf("  %s: size=%d, align=%d\n", 
					fieldType.Name, unsafe.Sizeof(field.Interface()), 
					unsafe.Alignof(field.Interface()))
			}
		}
}

func AccessBytes(ptr unsafe.Pointer, size int) []byte {
	return unsafe.Slice((*byte)(ptr), size)
}

func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func BytesToString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

func main() {
	fmt.Println("🚀 Go unsafe Package Mastery Examples")
	fmt.Println("=====================================")

	// 1. Basic Pointer Operations
	fmt.Println("\n1. Basic Pointer Operations:")
	
	var x int = 42
	fmt.Printf("Original value: %d\n", x)
	
	// Get pointer to x
	ptr := unsafe.Pointer(&x)
	fmt.Printf("Pointer address: %p\n", ptr)
	
	// Convert back to int pointer
	intPtr := (*int)(ptr)
	fmt.Printf("Dereferenced value: %d\n", *intPtr)
	
	// Modify through unsafe pointer
	*intPtr = 100
	fmt.Printf("Modified value: %d\n", x)

	// 2. Size and Alignment
	fmt.Println("\n2. Size and Alignment:")
	
	types := []interface{}{
		bool(false),
		int8(0),
		int16(0),
		int32(0),
		int64(0),
		int(0),
		uint(0),
		uintptr(0),
		float32(0),
		float64(0),
		complex64(0),
		complex128(0),
		string(""),
		[]int{},
		map[string]int{},
		chan int(nil),
		func() {},
		unsafe.Pointer(nil),
	}
	
	for _, v := range types {
		fmt.Printf("%-15T: size=%2d, align=%2d\n", 
			v, unsafe.Sizeof(v), unsafe.Alignof(v))
	}

	// 3. Struct Field Offsets
	fmt.Println("\n3. Struct Field Offsets:")
	
	person := Person{Name: "John", Age: 30, City: "New York"}
	fmt.Printf("Person struct:\n")
	fmt.Printf("  Name: offset=%d, size=%d, align=%d\n", 
		unsafe.Offsetof(person.Name), 
		unsafe.Sizeof(person.Name), 
		unsafe.Alignof(person.Name))
	fmt.Printf("  Age: offset=%d, size=%d, align=%d\n", 
		unsafe.Offsetof(person.Age), 
		unsafe.Sizeof(person.Age), 
		unsafe.Alignof(person.Age))
	fmt.Printf("  City: offset=%d, size=%d, align=%d\n", 
		unsafe.Offsetof(person.City), 
		unsafe.Sizeof(person.City), 
		unsafe.Alignof(person.City))
	
	fmt.Printf("Total struct size: %d bytes\n", unsafe.Sizeof(person))

	// 4. Pointer Arithmetic
	fmt.Println("\n4. Pointer Arithmetic:")
	
	arr := []int{1, 2, 3, 4, 5}
	fmt.Printf("Array: %v\n", arr)
	
	// Get pointer to first element
	ptr = unsafe.Pointer(&arr[0])
	fmt.Printf("First element: %d\n", *(*int)(ptr))
	
	// Move to second element
	ptr = unsafe.Add(ptr, unsafe.Sizeof(int(0)))
	fmt.Printf("Second element: %d\n", *(*int)(ptr))
	
	// Move to third element
	ptr = unsafe.Add(ptr, unsafe.Sizeof(int(0)))
	fmt.Printf("Third element: %d\n", *(*int)(ptr))

	// 5. Slice Operations
	fmt.Println("\n5. Slice Operations:")
	
	// Create slice from pointer
	slice := unsafe.Slice((*int)(unsafe.Pointer(&arr[0])), len(arr))
	fmt.Printf("Slice from pointer: %v\n", slice)
	
	// Get pointer to slice data
	sliceData := unsafe.SliceData(slice)
	fmt.Printf("Slice data pointer: %p\n", sliceData)
	
	// Access slice elements through pointer
	for i := 0; i < len(slice); i++ {
		elementPtr := unsafe.Add(unsafe.Pointer(sliceData), int(i)*int(unsafe.Sizeof(int(0))))
		element := *(*int)(elementPtr)
		fmt.Printf("Element %d: %d\n", i, element)
	}

	// 6. String Operations
	fmt.Println("\n6. String Operations:")
	
	s := "Hello, World!"
	fmt.Printf("Original string: %s\n", s)
	
	// Convert string to bytes
	bytes := StringToBytes(s)
	fmt.Printf("String as bytes: %v\n", bytes)
	
	// Convert bytes back to string
	s2 := BytesToString(bytes)
	fmt.Printf("Bytes as string: %s\n", s2)
	
	// Get string data pointer
	strData := unsafe.StringData(s)
	fmt.Printf("String data pointer: %p\n", strData)
	
	// Access string characters through pointer
	for i := 0; i < len(s); i++ {
		charPtr := unsafe.Add(unsafe.Pointer(strData), i)
		char := *(*byte)(charPtr)
		fmt.Printf("Character %d: %c\n", i, char)
	}

	// 7. Memory Layout Inspection
	fmt.Println("\n7. Memory Layout Inspection:")
	
	InspectStruct(Person{Name: "Alice", Age: 25, City: "Boston"})
	InspectStruct(Point{X: 1.0, Y: 2.0})
	InspectStruct(Data{ID: 1, Value: 3.14, Flag: true, Text: "test"})

	// 8. Type Conversions
	fmt.Println("\n8. Type Conversions:")
	
	// Convert between different types
	var i int = 42
	var f float64 = 3.14
	
	// Convert int to float64 through unsafe
	intPtr2 := unsafe.Pointer(&i)
	floatPtr := (*float64)(intPtr2)
	fmt.Printf("Int %d as float64: %f\n", i, *floatPtr)
	
	// Convert float64 to int through unsafe
	floatPtr2 := unsafe.Pointer(&f)
	intPtr3 := (*int)(floatPtr2)
	fmt.Printf("Float64 %f as int: %d\n", f, *intPtr3)

	// 9. Memory Access Patterns
	fmt.Println("\n9. Memory Access Patterns:")
	
	// Access memory as different types
	data := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f} // "Hello"
	ptr = unsafe.Pointer(&data[0])
	
	// Access as bytes
	bytes2 := unsafe.Slice((*byte)(ptr), len(data))
	fmt.Printf("As bytes: %v\n", bytes2)
	
	// Access as string
	str := unsafe.String(unsafe.SliceData(data), len(data))
	fmt.Printf("As string: %s\n", str)
	
	// Access as int32 (first 4 bytes)
	int32Ptr := (*int32)(ptr)
	fmt.Printf("As int32: %d\n", *int32Ptr)

	// 10. Struct Field Access
	fmt.Println("\n10. Struct Field Access:")
	
	p := Person{Name: "Bob", Age: 35, City: "Chicago"}
	ptr = unsafe.Pointer(&p)
	
	// Access Name field
	namePtr := unsafe.Add(ptr, unsafe.Offsetof(p.Name))
	name := *(*string)(namePtr)
	fmt.Printf("Name: %s\n", name)
	
	// Access Age field
	agePtr := unsafe.Add(ptr, unsafe.Offsetof(p.Age))
	age := *(*int)(agePtr)
	fmt.Printf("Age: %d\n", age)
	
	// Access City field
	cityPtr := unsafe.Add(ptr, unsafe.Offsetof(p.City))
	city := *(*string)(cityPtr)
	fmt.Printf("City: %s\n", city)

	// 11. Array Access
	fmt.Println("\n11. Array Access:")
	
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	
	ptr = unsafe.Pointer(&matrix[0][0])
	
		// Access elements through pointer arithmetic
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				offset := i*3 + j
				elementPtr := unsafe.Add(ptr, int(offset)*int(unsafe.Sizeof(int(0))))
				element := *(*int)(elementPtr)
				fmt.Printf("matrix[%d][%d] = %d\n", i, j, element)
			}
		}

	// 12. Pointer Comparison
	fmt.Println("\n12. Pointer Comparison:")
	
	var a, b int = 10, 20
	ptrA := unsafe.Pointer(&a)
	ptrB := unsafe.Pointer(&b)
	
	fmt.Printf("Pointer A: %p\n", ptrA)
	fmt.Printf("Pointer B: %p\n", ptrB)
	fmt.Printf("Pointers equal: %t\n", ptrA == ptrB)
	
	// Compare with uintptr
	uintptrA := uintptr(ptrA)
	uintptrB := uintptr(ptrB)
	fmt.Printf("Uintptr A: %d\n", uintptrA)
	fmt.Printf("Uintptr B: %d\n", uintptrB)
	fmt.Printf("Difference: %d bytes\n", uintptrB-uintptrA)

	// 13. Memory Safety Demonstration
	fmt.Println("\n13. Memory Safety Demonstration:")
	
	// This demonstrates why unsafe operations are dangerous
	arr2 := []int{1, 2, 3}
	ptr = unsafe.Pointer(&arr2[0])
	
	// Access valid element
	fmt.Printf("Valid element: %d\n", *(*int)(ptr))
	
	// Access element beyond bounds (dangerous!)
	// This could cause a panic or undefined behavior
	ptr = unsafe.Add(ptr, 10*unsafe.Sizeof(int(0)))
	fmt.Printf("Beyond bounds (dangerous): %d\n", *(*int)(ptr))

	// 14. Custom Memory Operations
	fmt.Println("\n14. Custom Memory Operations:")
	
	// Create a custom memory block
	size := 100
	memory := make([]byte, size)
	ptr = unsafe.Pointer(&memory[0])
	
	// Fill memory with pattern
	for i := 0; i < size; i++ {
		bytePtr := unsafe.Add(ptr, i)
		*(*byte)(bytePtr) = byte(i % 256)
	}
	
	// Read back the pattern
	fmt.Printf("Memory pattern (first 20 bytes): ")
	for i := 0; i < 20; i++ {
		bytePtr := unsafe.Add(ptr, i)
		fmt.Printf("%02x ", *(*byte)(bytePtr))
	}
	fmt.Println()

	// 15. Performance Comparison
	fmt.Println("\n15. Performance Comparison:")
	
	// Test unsafe vs safe operations
	data2 := make([]int, 1000)
	for i := range data2 {
		data2[i] = i
	}
	
	// Safe access
	start := time.Now()
	sum := 0
	for _, v := range data2 {
		sum += v
	}
	safeTime := time.Since(start)
	
	// Unsafe access
	start = time.Now()
	sum2 := 0
	ptr = unsafe.Pointer(&data2[0])
	for i := 0; i < len(data2); i++ {
		elementPtr := unsafe.Add(ptr, int(i)*int(unsafe.Sizeof(int(0))))
		sum2 += *(*int)(elementPtr)
	}
	unsafeTime := time.Since(start)
	
	fmt.Printf("Safe access time: %v\n", safeTime)
	fmt.Printf("Unsafe access time: %v\n", unsafeTime)
	fmt.Printf("Sum (safe): %d\n", sum)
	fmt.Printf("Sum (unsafe): %d\n", sum2)

	// 16. String Header Inspection
	fmt.Println("\n16. String Header Inspection:")
	
	s3 := "Hello, World!"
	
	// Get string header information
	strData2 := unsafe.StringData(s3)
	length := len(s3)
	
	fmt.Printf("String: %s\n", s3)
	fmt.Printf("Data pointer: %p\n", strData2)
	fmt.Printf("Length: %d\n", length)
	
	// Access string header directly
	header := (*reflect.StringHeader)(unsafe.Pointer(&s3))
	fmt.Printf("Header data: %p\n", unsafe.Pointer(header.Data))
	fmt.Printf("Header len: %d\n", header.Len)

	// 17. Slice Header Inspection
	fmt.Println("\n17. Slice Header Inspection:")
	
	slice2 := []int{1, 2, 3, 4, 5}
	
	// Get slice header information
	sliceData2 := unsafe.SliceData(slice2)
	len2 := len(slice2)
	cap2 := cap(slice2)
	
	fmt.Printf("Slice: %v\n", slice2)
	fmt.Printf("Data pointer: %p\n", sliceData2)
	fmt.Printf("Length: %d\n", len2)
	fmt.Printf("Capacity: %d\n", cap2)
	
	// Access slice header directly
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice2))
	fmt.Printf("Header data: %p\n", unsafe.Pointer(sliceHeader.Data))
	fmt.Printf("Header len: %d\n", sliceHeader.Len)
	fmt.Printf("Header cap: %d\n", sliceHeader.Cap)

	// 18. Interface Inspection
	fmt.Println("\n18. Interface Inspection:")
	
	// Note: InterfaceHeader is not available in all Go versions
	fmt.Printf("Interface inspection not available in this Go version\n")

	// 19. Memory Alignment
	fmt.Println("\n19. Memory Alignment:")
	
	type AlignedStruct struct {
		a bool    // 1 byte
		b int32   // 4 bytes
		c bool    // 1 byte
		d int64   // 8 bytes
	}
	
	aligned := AlignedStruct{}
	fmt.Printf("AlignedStruct size: %d bytes\n", unsafe.Sizeof(aligned))
	fmt.Printf("Field a offset: %d\n", unsafe.Offsetof(aligned.a))
	fmt.Printf("Field b offset: %d\n", unsafe.Offsetof(aligned.b))
	fmt.Printf("Field c offset: %d\n", unsafe.Offsetof(aligned.c))
	fmt.Printf("Field d offset: %d\n", unsafe.Offsetof(aligned.d))

	// 20. Final Warning
	fmt.Println("\n20. Final Warning:")
	fmt.Println("⚠️  WARNING: The unsafe package bypasses Go's type safety!")
	fmt.Println("⚠️  Use only when absolutely necessary and with extreme caution!")
	fmt.Println("⚠️  Unsafe operations can cause crashes, memory corruption, and undefined behavior!")
	fmt.Println("⚠️  Always test thoroughly and document your unsafe code!")

	fmt.Println("\n🎉 unsafe Package Mastery Complete!")
}
