package pointers

import (
	"fmt"
	"unsafe"
)

// Person represents a person for pointer examples
type Person struct {
	Name string
	Age  int
}

// Node represents a linked list node
type Node struct {
	Value int
	Next  *Node
}

// DemonstrateBasicPointers shows basic pointer operations
func DemonstrateBasicPointers() {
	fmt.Println("=== BASIC POINTERS ===")
	
	// 1. Pointer Declaration and Initialization
	fmt.Println("\n--- Pointer Declaration ---")
	
	var p *int                    // nil pointer
	var x int = 42
	var ptr *int = &x             // pointer to x
	
	fmt.Printf("p (nil pointer): %v\n", p)
	fmt.Printf("x: %d\n", x)
	fmt.Printf("ptr (pointer to x): %v\n", ptr)
	fmt.Printf("ptr points to: %d\n", *ptr)
	
	// 2. Address and Dereference Operators
	fmt.Println("\n--- Address and Dereference ---")
	
	y := 100
	ptrY := &y
	
	fmt.Printf("y: %d\n", y)
	fmt.Printf("Address of y: %p\n", &y)
	fmt.Printf("ptrY: %p\n", ptrY)
	fmt.Printf("Value at ptrY: %d\n", *ptrY)
	
	// 3. Modifying Values Through Pointers
	fmt.Println("\n--- Modifying Through Pointers ---")
	
	z := 50
	ptrZ := &z
	
	fmt.Printf("Before: z = %d\n", z)
	*ptrZ = 75
	fmt.Printf("After *ptrZ = 75: z = %d\n", z)
	
	// 4. Pointer Arithmetic (Limited in Go)
	fmt.Println("\n--- Pointer Arithmetic ---")
	
	arr := [5]int{1, 2, 3, 4, 5}
	ptrArr := &arr[0]
	
	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("Pointer to first element: %p\n", ptrArr)
	fmt.Printf("Value at first element: %d\n", *ptrArr)
	
	// Go doesn't allow pointer arithmetic like C, but we can use unsafe
	// This is generally not recommended in production code
	ptrArr2 := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptrArr)) + unsafe.Sizeof(int(0))))
	fmt.Printf("Pointer to second element (unsafe): %p\n", ptrArr2)
	fmt.Printf("Value at second element: %d\n", *ptrArr2)
}

// DemonstratePointerToStruct shows pointers to structs
func DemonstratePointerToStruct() {
	fmt.Println("\n=== POINTERS TO STRUCTS ===")
	
	// 1. Creating pointers to structs
	fmt.Println("\n--- Creating Pointers to Structs ---")
	
	// Method 1: Create struct, then get pointer
	person1 := Person{Name: "Alice", Age: 30}
	ptrPerson1 := &person1
	
	fmt.Printf("person1: %+v\n", person1)
	fmt.Printf("ptrPerson1: %p\n", ptrPerson1)
	fmt.Printf("Value at ptrPerson1: %+v\n", *ptrPerson1)
	
	// Method 2: Create pointer directly
	ptrPerson2 := &Person{Name: "Bob", Age: 25}
	fmt.Printf("ptrPerson2: %p\n", ptrPerson2)
	fmt.Printf("Value at ptrPerson2: %+v\n", *ptrPerson2)
	
	// Method 3: Using new()
	ptrPerson3 := new(Person)
	ptrPerson3.Name = "Charlie"
	ptrPerson3.Age = 35
	fmt.Printf("ptrPerson3: %p\n", ptrPerson3)
	fmt.Printf("Value at ptrPerson3: %+v\n", *ptrPerson3)
	
	// 2. Accessing struct fields through pointers
	fmt.Println("\n--- Accessing Fields Through Pointers ---")
	
	person := &Person{Name: "David", Age: 28}
	
	// Both syntaxes are equivalent
	fmt.Printf("Name: %s\n", (*person).Name)  // Explicit dereference
	fmt.Printf("Name: %s\n", person.Name)     // Go automatically dereferences
	
	fmt.Printf("Age: %d\n", person.Age)
	
	// 3. Modifying struct fields through pointers
	fmt.Println("\n--- Modifying Fields Through Pointers ---")
	
	fmt.Printf("Before: %+v\n", *person)
	person.Age = 29
	person.Name = "David Smith"
	fmt.Printf("After: %+v\n", *person)
}

// DemonstratePointerMethods shows pointer receivers
func DemonstratePointerMethods() {
	fmt.Println("\n=== POINTER RECEIVERS ===")
	
	// 1. Value vs Pointer Receivers
	fmt.Println("\n--- Value vs Pointer Receivers ---")
	
	person := Person{Name: "Eve", Age: 32}
	fmt.Printf("Original: %+v\n", person)
	
	// Value receiver - doesn't modify original
	person.SetAgeValue(33)
	fmt.Printf("After SetAgeValue(33): %+v\n", person)
	
	// Pointer receiver - modifies original
	person.SetAgePointer(34)
	fmt.Printf("After SetAgePointer(34): %+v\n", person)
	
	// 2. Method calls on pointers
	fmt.Println("\n--- Method Calls on Pointers ---")
	
	ptrPerson := &Person{Name: "Frank", Age: 40}
	fmt.Printf("Original: %+v\n", *ptrPerson)
	
	// Can call both value and pointer methods on pointer
	ptrPerson.SetAgeValue(41)  // Go automatically dereferences
	ptrPerson.SetAgePointer(42)
	fmt.Printf("After both methods: %+v\n", *ptrPerson)
	
	// 3. Method calls on values
	fmt.Println("\n--- Method Calls on Values ---")
	
	valuePerson := Person{Name: "Grace", Age: 45}
	fmt.Printf("Original: %+v\n", valuePerson)
	
	// Can call both value and pointer methods on value
	valuePerson.SetAgeValue(46)  // Works directly
	valuePerson.SetAgePointer(47) // Go automatically takes address
	fmt.Printf("After both methods: %+v\n", valuePerson)
}

// DemonstratePointerToSlice shows pointers to slices
func DemonstratePointerToSlice() {
	fmt.Println("\n=== POINTERS TO SLICES ===")
	
	// 1. Pointer to slice
	fmt.Println("\n--- Pointer to Slice ---")
	
	slice := []int{1, 2, 3, 4, 5}
	ptrSlice := &slice
	
	fmt.Printf("slice: %v\n", slice)
	fmt.Printf("ptrSlice: %p\n", ptrSlice)
	fmt.Printf("Value at ptrSlice: %v\n", *ptrSlice)
	
	// 2. Modifying slice through pointer
	fmt.Println("\n--- Modifying Slice Through Pointer ---")
	
	fmt.Printf("Before: %v\n", *ptrSlice)
	*ptrSlice = append(*ptrSlice, 6, 7, 8)
	fmt.Printf("After append: %v\n", *ptrSlice)
	
	// 3. Slice of pointers
	fmt.Println("\n--- Slice of Pointers ---")
	
	people := []*Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}
	
	fmt.Println("Slice of pointers:")
	for i, person := range people {
		fmt.Printf("  [%d]: %+v (address: %p)\n", i, *person, person)
	}
	
	// 4. Modifying elements in slice of pointers
	fmt.Println("\n--- Modifying Elements ---")
	
	fmt.Printf("Before: %+v\n", *people[0])
	people[0].Age = 31
	fmt.Printf("After: %+v\n", *people[0])
}

// DemonstratePointerToArray shows pointers to arrays
func DemonstratePointerToArray() {
	fmt.Println("\n=== POINTERS TO ARRAYS ===")
	
	// 1. Pointer to array
	fmt.Println("\n--- Pointer to Array ---")
	
	arr := [5]int{1, 2, 3, 4, 5}
	ptrArr := &arr
	
	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("ptrArr: %p\n", ptrArr)
	fmt.Printf("Value at ptrArr: %v\n", *ptrArr)
	
	// 2. Accessing array elements through pointer
	fmt.Println("\n--- Accessing Elements ---")
	
	fmt.Printf("First element: %d\n", (*ptrArr)[0])
	fmt.Printf("Second element: %d\n", (*ptrArr)[1])
	
	// 3. Modifying array elements through pointer
	fmt.Println("\n--- Modifying Elements ---")
	
	fmt.Printf("Before: %v\n", *ptrArr)
	(*ptrArr)[0] = 10
	(*ptrArr)[1] = 20
	fmt.Printf("After: %v\n", *ptrArr)
}

// DemonstratePointerChaining shows pointer chaining
func DemonstratePointerChaining() {
	fmt.Println("\n=== POINTER CHAINING ===")
	
	// 1. Pointer to pointer
	fmt.Println("\n--- Pointer to Pointer ---")
	
	x := 42
	ptrX := &x
	ptrPtrX := &ptrX
	
	fmt.Printf("x: %d\n", x)
	fmt.Printf("ptrX: %p\n", ptrX)
	fmt.Printf("ptrPtrX: %p\n", ptrPtrX)
	fmt.Printf("*ptrX: %d\n", *ptrX)
	fmt.Printf("**ptrPtrX: %d\n", **ptrPtrX)
	
	// 2. Modifying through double pointer
	fmt.Println("\n--- Modifying Through Double Pointer ---")
	
	fmt.Printf("Before: x = %d\n", x)
	**ptrPtrX = 100
	fmt.Printf("After **ptrPtrX = 100: x = %d\n", x)
	
	// 3. Linked list example
	fmt.Println("\n--- Linked List Example ---")
	
	// Create a simple linked list
	head := &Node{Value: 1}
	head.Next = &Node{Value: 2}
	head.Next.Next = &Node{Value: 3}
	
	fmt.Println("Linked list:")
	current := head
	for current != nil {
		fmt.Printf("  Node: %d (address: %p)\n", current.Value, current)
		current = current.Next
	}
}

// DemonstratePointerSafety shows pointer safety considerations
func DemonstratePointerSafety() {
	fmt.Println("\n=== POINTER SAFETY ===")
	
	// 1. Nil pointer checks
	fmt.Println("\n--- Nil Pointer Checks ---")
	
	var ptr *int
	fmt.Printf("ptr == nil: %t\n", ptr == nil)
	
	// This would cause a panic if we tried to dereference
	// fmt.Printf("*ptr: %d\n", *ptr) // PANIC!
	
	// Safe dereference
	if ptr != nil {
		fmt.Printf("*ptr: %d\n", *ptr)
	} else {
		fmt.Println("ptr is nil, cannot dereference")
	}
	
	// 2. Pointer comparison
	fmt.Println("\n--- Pointer Comparison ---")
	
	x := 42
	y := 42
	ptrX := &x
	ptrY := &y
	ptrX2 := &x
	
	fmt.Printf("ptrX == ptrY: %t\n", ptrX == ptrY)     // Different addresses
	fmt.Printf("ptrX == ptrX2: %t\n", ptrX == ptrX2)   // Same address
	fmt.Printf("*ptrX == *ptrY: %t\n", *ptrX == *ptrY) // Same values
	
	// 3. Pointer to local variable (be careful!)
	fmt.Println("\n--- Pointer to Local Variable ---")
	
	ptr = getPointerToLocal()
	if ptr != nil {
		fmt.Printf("Value through pointer: %d\n", *ptr)
		fmt.Println("WARNING: This is dangerous! The local variable may no longer exist.")
	}
}

// DemonstratePointerMemory shows memory characteristics
func DemonstratePointerMemory() {
	fmt.Println("\n=== POINTER MEMORY ===")
	
	// 1. Pointer size
	fmt.Println("\n--- Pointer Size ---")
	
	var ptrInt *int
	var ptrString *string
	var ptrStruct *Person
	
	fmt.Printf("Size of *int: %d bytes\n", unsafe.Sizeof(ptrInt))
	fmt.Printf("Size of *string: %d bytes\n", unsafe.Sizeof(ptrString))
	fmt.Printf("Size of *Person: %d bytes\n", unsafe.Sizeof(ptrStruct))
	
	// 2. Memory addresses
	fmt.Println("\n--- Memory Addresses ---")
	
	x := 42
	y := 100
	ptrX := &x
	ptrY := &y
	
	fmt.Printf("x address: %p\n", &x)
	fmt.Printf("y address: %p\n", &y)
	fmt.Printf("ptrX value: %p\n", ptrX)
	fmt.Printf("ptrY value: %p\n", ptrY)
	
	// 3. Pointer arithmetic (unsafe, not recommended)
	fmt.Println("\n--- Pointer Arithmetic (Unsafe) ---")
	
	arr := [3]int{10, 20, 30}
	ptr := &arr[0]
	
	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("First element address: %p\n", ptr)
	
	// Calculate address of second element
	ptr2 := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(int(0))))
	fmt.Printf("Second element address: %p\n", ptr2)
	fmt.Printf("Second element value: %d\n", *ptr2)
}

// Helper functions

// getPointerToLocal returns a pointer to a local variable (dangerous!)
func getPointerToLocal() *int {
	localVar := 42
	return &localVar // This is dangerous! localVar will be destroyed when function returns
}

// Person methods

// SetAgeValue is a value receiver method
func (p Person) SetAgeValue(age int) {
	p.Age = age // This doesn't modify the original
}

// SetAgePointer is a pointer receiver method
func (p *Person) SetAgePointer(age int) {
	p.Age = age // This modifies the original
}

// GetInfo returns person information
func (p *Person) GetInfo() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

// UpdateAge updates the person's age
func (p *Person) UpdateAge(newAge int) {
	if newAge >= 0 && newAge <= 150 {
		p.Age = newAge
	}
}

// RunAllPointerExamples runs all pointer examples
func RunAllPointerExamples() {
	DemonstrateBasicPointers()
	DemonstratePointerToStruct()
	DemonstratePointerMethods()
	DemonstratePointerToSlice()
	DemonstratePointerToArray()
	DemonstratePointerChaining()
	DemonstratePointerSafety()
	DemonstratePointerMemory()
}
