package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

// 🎯 POINTER TYPES MASTERY
// This file demonstrates comprehensive pointer usage and memory management

// PointerManager handles various pointer operations and patterns
type PointerManager struct {
	// Basic pointer types
	IntPtr        *int
	FloatPtr      *float64
	StringPtr     *string
	BoolPtr       *bool
	
	// Pointer to structs
	PersonPtr     *Person
	CompanyPtr    *Company
	
	// Pointer to slices
	IntSlicePtr   *[]int
	StringSlicePtr *[]string
	
	// Pointer to maps
	MapPtr        *map[string]int
	
	// Pointer to functions
	FuncPtr       *func(int) int
	
	// Pointer to interfaces
	InterfacePtr  *interface{}
	
	// Pointer to pointers (double pointers)
	IntPtrPtr     **int
	StringPtrPtr  **string
	
	// Pointer arrays and slices
	IntPtrArray   [5]*int
	IntPtrSlice   []*int
	
	// Pointer to arrays
	ArrayPtr      *[5]int
	
	// Pointer to channels
	ChanPtr       *chan int
	
	// Pointer to unsafe operations
	UnsafePtr     unsafe.Pointer
}

// NewPointerManager creates a new pointer manager
func NewPointerManager() *PointerManager {
	return &PointerManager{
		IntPtrArray:  [5]*int{},
		IntPtrSlice:  make([]*int, 0),
	}
}

// CRUD Operations for Pointers

// Create - Initialize and create pointer instances
func (pm *PointerManager) Create() {
	fmt.Println("🔧 Creating pointer instances...")
	
	// Create basic pointers
	pm.createBasicPointers()
	
	// Create struct pointers
	pm.createStructPointers()
	
	// Create slice pointers
	pm.createSlicePointers()
	
	// Create map pointers
	pm.createMapPointers()
	
	// Create function pointers
	pm.createFunctionPointers()
	
	// Create interface pointers
	pm.createInterfacePointers()
	
	// Create double pointers
	pm.createDoublePointers()
	
	// Create pointer arrays and slices
	pm.createPointerArrays()
	
	// Create array pointers
	pm.createArrayPointers()
	
	// Create channel pointers
	pm.createChannelPointers()
	
	// Create unsafe pointers
	pm.createUnsafePointers()
	
	fmt.Println("✅ Pointer instances created successfully")
}

// createBasicPointers creates pointers to basic types
func (pm *PointerManager) createBasicPointers() {
	// Create pointers to basic types
	intVal := 42
	pm.IntPtr = &intVal
	
	floatVal := 3.14159
	pm.FloatPtr = &floatVal
	
	stringVal := "Hello, Pointers!"
	pm.StringPtr = &stringVal
	
	boolVal := true
	pm.BoolPtr = &boolVal
	
	fmt.Println("Created basic type pointers")
}

// createStructPointers creates pointers to structs
func (pm *PointerManager) createStructPointers() {
	// Create pointer to Person struct
	person := Person{
		ID:    1,
		Name:  "Alice Johnson",
		Email: "alice@example.com",
		Age:   30,
	}
	pm.PersonPtr = &person
	
	// Create pointer to Company struct
	company := Company{
		ID:   1,
		Name: "TechCorp Inc.",
		CEO:  person,
	}
	pm.CompanyPtr = &company
	
	fmt.Println("Created struct pointers")
}

// createSlicePointers creates pointers to slices
func (pm *PointerManager) createSlicePointers() {
	// Create pointer to int slice
	intSlice := []int{1, 2, 3, 4, 5}
	pm.IntSlicePtr = &intSlice
	
	// Create pointer to string slice
	stringSlice := []string{"apple", "banana", "cherry"}
	pm.StringSlicePtr = &stringSlice
	
	fmt.Println("Created slice pointers")
}

// createMapPointers creates pointers to maps
func (pm *PointerManager) createMapPointers() {
	// Create pointer to map
	mapVal := map[string]int{
		"apple":  5,
		"banana": 3,
		"cherry": 8,
	}
	pm.MapPtr = &mapVal
	
	fmt.Println("Created map pointers")
}

// createFunctionPointers creates pointers to functions
func (pm *PointerManager) createFunctionPointers() {
	// Create a function
	square := func(x int) int {
		return x * x
	}
	pm.FuncPtr = &square
	
	fmt.Println("Created function pointers")
}

// createInterfacePointers creates pointers to interfaces
func (pm *PointerManager) createInterfacePointers() {
	// Create pointer to interface
	var interfaceVal interface{} = "Hello, Interface!"
	pm.InterfacePtr = &interfaceVal
	
	fmt.Println("Created interface pointers")
}

// createDoublePointers creates double pointers
func (pm *PointerManager) createDoublePointers() {
	// Create double pointer to int
	intVal := 42
	intPtr := &intVal
	pm.IntPtrPtr = &intPtr
	
	// Create double pointer to string
	stringVal := "Double Pointer"
	stringPtr := &stringVal
	pm.StringPtrPtr = &stringPtr
	
	fmt.Println("Created double pointers")
}

// createPointerArrays creates pointer arrays and slices
func (pm *PointerManager) createPointerArrays() {
	// Create array of pointers to int
	for i := 0; i < 5; i++ {
		val := i * 10
		pm.IntPtrArray[i] = &val
	}
	
	// Create slice of pointers to int
	for i := 0; i < 3; i++ {
		val := i * 100
		pm.IntPtrSlice = append(pm.IntPtrSlice, &val)
	}
	
	fmt.Println("Created pointer arrays and slices")
}

// createArrayPointers creates pointers to arrays
func (pm *PointerManager) createArrayPointers() {
	// Create pointer to array
	array := [5]int{10, 20, 30, 40, 50}
	pm.ArrayPtr = &array
	
	fmt.Println("Created array pointers")
}

// createChannelPointers creates pointers to channels
func (pm *PointerManager) createChannelPointers() {
	// Create pointer to channel
	ch := make(chan int, 5)
	pm.ChanPtr = &ch
	
	fmt.Println("Created channel pointers")
}

// createUnsafePointers creates unsafe pointers
func (pm *PointerManager) createUnsafePointers() {
	// Create unsafe pointer
	intVal := 42
	pm.UnsafePtr = unsafe.Pointer(&intVal)
	
	fmt.Println("Created unsafe pointers")
}

// Read - Display pointer information
func (pm *PointerManager) Read() {
	fmt.Println("\n📖 READING POINTER INFORMATION:")
	fmt.Println("===============================")
	
	// Read basic pointers
	pm.readBasicPointers()
	
	// Read struct pointers
	pm.readStructPointers()
	
	// Read slice pointers
	pm.readSlicePointers()
	
	// Read map pointers
	pm.readMapPointers()
	
	// Read function pointers
	pm.readFunctionPointers()
	
	// Read interface pointers
	pm.readInterfacePointers()
	
	// Read double pointers
	pm.readDoublePointers()
	
	// Read pointer arrays
	pm.readPointerArrays()
	
	// Read array pointers
	pm.readArrayPointers()
	
	// Read channel pointers
	pm.readChannelPointers()
	
	// Read unsafe pointers
	pm.readUnsafePointers()
}

// readBasicPointers displays basic pointer information
func (pm *PointerManager) readBasicPointers() {
	fmt.Println("Basic Pointers:")
	
	if pm.IntPtr != nil {
		fmt.Printf("  IntPtr: %p -> %d\n", pm.IntPtr, *pm.IntPtr)
	} else {
		fmt.Println("  IntPtr: nil")
	}
	
	if pm.FloatPtr != nil {
		fmt.Printf("  FloatPtr: %p -> %.2f\n", pm.FloatPtr, *pm.FloatPtr)
	} else {
		fmt.Println("  FloatPtr: nil")
	}
	
	if pm.StringPtr != nil {
		fmt.Printf("  StringPtr: %p -> %s\n", pm.StringPtr, *pm.StringPtr)
	} else {
		fmt.Println("  StringPtr: nil")
	}
	
	if pm.BoolPtr != nil {
		fmt.Printf("  BoolPtr: %p -> %t\n", pm.BoolPtr, *pm.BoolPtr)
	} else {
		fmt.Println("  BoolPtr: nil")
	}
}

// readStructPointers displays struct pointer information
func (pm *PointerManager) readStructPointers() {
	fmt.Println("\nStruct Pointers:")
	
	if pm.PersonPtr != nil {
		fmt.Printf("  PersonPtr: %p -> %+v\n", pm.PersonPtr, *pm.PersonPtr)
	} else {
		fmt.Println("  PersonPtr: nil")
	}
	
	if pm.CompanyPtr != nil {
		fmt.Printf("  CompanyPtr: %p -> %+v\n", pm.CompanyPtr, *pm.CompanyPtr)
	} else {
		fmt.Println("  CompanyPtr: nil")
	}
}

// readSlicePointers displays slice pointer information
func (pm *PointerManager) readSlicePointers() {
	fmt.Println("\nSlice Pointers:")
	
	if pm.IntSlicePtr != nil {
		fmt.Printf("  IntSlicePtr: %p -> %v\n", pm.IntSlicePtr, *pm.IntSlicePtr)
	} else {
		fmt.Println("  IntSlicePtr: nil")
	}
	
	if pm.StringSlicePtr != nil {
		fmt.Printf("  StringSlicePtr: %p -> %v\n", pm.StringSlicePtr, *pm.StringSlicePtr)
	} else {
		fmt.Println("  StringSlicePtr: nil")
	}
}

// readMapPointers displays map pointer information
func (pm *PointerManager) readMapPointers() {
	fmt.Println("\nMap Pointers:")
	
	if pm.MapPtr != nil {
		fmt.Printf("  MapPtr: %p -> %v\n", pm.MapPtr, *pm.MapPtr)
	} else {
		fmt.Println("  MapPtr: nil")
	}
}

// readFunctionPointers displays function pointer information
func (pm *PointerManager) readFunctionPointers() {
	fmt.Println("\nFunction Pointers:")
	
	if pm.FuncPtr != nil {
		fmt.Printf("  FuncPtr: %p\n", pm.FuncPtr)
		// Call the function through the pointer
		result := (*pm.FuncPtr)(5)
		fmt.Printf("  (*FuncPtr)(5) = %d\n", result)
	} else {
		fmt.Println("  FuncPtr: nil")
	}
}

// readInterfacePointers displays interface pointer information
func (pm *PointerManager) readInterfacePointers() {
	fmt.Println("\nInterface Pointers:")
	
	if pm.InterfacePtr != nil {
		fmt.Printf("  InterfacePtr: %p -> %v (%T)\n", pm.InterfacePtr, *pm.InterfacePtr, *pm.InterfacePtr)
	} else {
		fmt.Println("  InterfacePtr: nil")
	}
}

// readDoublePointers displays double pointer information
func (pm *PointerManager) readDoublePointers() {
	fmt.Println("\nDouble Pointers:")
	
	if pm.IntPtrPtr != nil {
		fmt.Printf("  IntPtrPtr: %p -> %p -> %d\n", pm.IntPtrPtr, *pm.IntPtrPtr, **pm.IntPtrPtr)
	} else {
		fmt.Println("  IntPtrPtr: nil")
	}
	
	if pm.StringPtrPtr != nil {
		fmt.Printf("  StringPtrPtr: %p -> %p -> %s\n", pm.StringPtrPtr, *pm.StringPtrPtr, **pm.StringPtrPtr)
	} else {
		fmt.Println("  StringPtrPtr: nil")
	}
}

// readPointerArrays displays pointer array information
func (pm *PointerManager) readPointerArrays() {
	fmt.Println("\nPointer Arrays:")
	
	fmt.Printf("  IntPtrArray: %p\n", &pm.IntPtrArray)
	for i, ptr := range pm.IntPtrArray {
		if ptr != nil {
			fmt.Printf("    [%d]: %p -> %d\n", i, ptr, *ptr)
		} else {
			fmt.Printf("    [%d]: nil\n", i)
		}
	}
	
	fmt.Printf("  IntPtrSlice: %p\n", &pm.IntPtrSlice)
	for i, ptr := range pm.IntPtrSlice {
		if ptr != nil {
			fmt.Printf("    [%d]: %p -> %d\n", i, ptr, *ptr)
		} else {
			fmt.Printf("    [%d]: nil\n", i)
		}
	}
}

// readArrayPointers displays array pointer information
func (pm *PointerManager) readArrayPointers() {
	fmt.Println("\nArray Pointers:")
	
	if pm.ArrayPtr != nil {
		fmt.Printf("  ArrayPtr: %p -> %v\n", pm.ArrayPtr, *pm.ArrayPtr)
	} else {
		fmt.Println("  ArrayPtr: nil")
	}
}

// readChannelPointers displays channel pointer information
func (pm *PointerManager) readChannelPointers() {
	fmt.Println("\nChannel Pointers:")
	
	if pm.ChanPtr != nil {
		fmt.Printf("  ChanPtr: %p -> %p\n", pm.ChanPtr, *pm.ChanPtr)
	} else {
		fmt.Println("  ChanPtr: nil")
	}
}

// readUnsafePointers displays unsafe pointer information
func (pm *PointerManager) readUnsafePointers() {
	fmt.Println("\nUnsafe Pointers:")
	
	if pm.UnsafePtr != nil {
		fmt.Printf("  UnsafePtr: %p\n", pm.UnsafePtr)
		// Convert back to int pointer and read value
		intPtr := (*int)(pm.UnsafePtr)
		fmt.Printf("  Value: %d\n", *intPtr)
	} else {
		fmt.Println("  UnsafePtr: nil")
	}
}

// Update - Modify pointer values
func (pm *PointerManager) Update() {
	fmt.Println("\n🔄 UPDATING POINTER VALUES:")
	fmt.Println("===========================")
	
	// Update basic pointers
	pm.updateBasicPointers()
	
	// Update struct pointers
	pm.updateStructPointers()
	
	// Update slice pointers
	pm.updateSlicePointers()
	
	// Update map pointers
	pm.updateMapPointers()
	
	// Update function pointers
	pm.updateFunctionPointers()
	
	// Update interface pointers
	pm.updateInterfacePointers()
	
	// Update double pointers
	pm.updateDoublePointers()
	
	// Update pointer arrays
	pm.updatePointerArrays()
	
	// Update array pointers
	pm.updateArrayPointers()
	
	// Update channel pointers
	pm.updateChannelPointers()
	
	// Update unsafe pointers
	pm.updateUnsafePointers()
	
	fmt.Println("✅ Pointer values updated successfully")
}

// updateBasicPointers updates basic pointer values
func (pm *PointerManager) updateBasicPointers() {
	if pm.IntPtr != nil {
		*pm.IntPtr = 100
		fmt.Printf("Updated IntPtr to: %d\n", *pm.IntPtr)
	}
	
	if pm.FloatPtr != nil {
		*pm.FloatPtr = 2.71828
		fmt.Printf("Updated FloatPtr to: %.5f\n", *pm.FloatPtr)
	}
	
	if pm.StringPtr != nil {
		*pm.StringPtr = "Updated String!"
		fmt.Printf("Updated StringPtr to: %s\n", *pm.StringPtr)
	}
	
	if pm.BoolPtr != nil {
		*pm.BoolPtr = false
		fmt.Printf("Updated BoolPtr to: %t\n", *pm.BoolPtr)
	}
}

// updateStructPointers updates struct pointer values
func (pm *PointerManager) updateStructPointers() {
	if pm.PersonPtr != nil {
		pm.PersonPtr.Age = 31
		pm.PersonPtr.Name = "Alice Johnson-Updated"
		fmt.Printf("Updated PersonPtr: %+v\n", *pm.PersonPtr)
	}
	
	if pm.CompanyPtr != nil {
		pm.CompanyPtr.Name = "TechCorp Inc. - Updated"
		fmt.Printf("Updated CompanyPtr: %+v\n", *pm.CompanyPtr)
	}
}

// updateSlicePointers updates slice pointer values
func (pm *PointerManager) updateSlicePointers() {
	if pm.IntSlicePtr != nil {
		*pm.IntSlicePtr = append(*pm.IntSlicePtr, 6, 7, 8)
		fmt.Printf("Updated IntSlicePtr: %v\n", *pm.IntSlicePtr)
	}
	
	if pm.StringSlicePtr != nil {
		*pm.StringSlicePtr = append(*pm.StringSlicePtr, "date", "elderberry")
		fmt.Printf("Updated StringSlicePtr: %v\n", *pm.StringSlicePtr)
	}
}

// updateMapPointers updates map pointer values
func (pm *PointerManager) updateMapPointers() {
	if pm.MapPtr != nil {
		(*pm.MapPtr)["grape"] = 4
		(*pm.MapPtr)["kiwi"] = 2
		fmt.Printf("Updated MapPtr: %v\n", *pm.MapPtr)
	}
}

// updateFunctionPointers updates function pointer values
func (pm *PointerManager) updateFunctionPointers() {
	if pm.FuncPtr != nil {
		// Create new function
		cube := func(x int) int {
			return x * x * x
		}
		pm.FuncPtr = &cube
		
		// Test the new function
		result := (*pm.FuncPtr)(3)
		fmt.Printf("Updated FuncPtr, (*FuncPtr)(3) = %d\n", result)
	}
}

// updateInterfacePointers updates interface pointer values
func (pm *PointerManager) updateInterfacePointers() {
	if pm.InterfacePtr != nil {
		*pm.InterfacePtr = 42
		fmt.Printf("Updated InterfacePtr to: %v (%T)\n", *pm.InterfacePtr, *pm.InterfacePtr)
	}
}

// updateDoublePointers updates double pointer values
func (pm *PointerManager) updateDoublePointers() {
	if pm.IntPtrPtr != nil {
		**pm.IntPtrPtr = 200
		fmt.Printf("Updated IntPtrPtr to: %d\n", **pm.IntPtrPtr)
	}
	
	if pm.StringPtrPtr != nil {
		**pm.StringPtrPtr = "Updated Double Pointer"
		fmt.Printf("Updated StringPtrPtr to: %s\n", **pm.StringPtrPtr)
	}
}

// updatePointerArrays updates pointer array values
func (pm *PointerManager) updatePointerArrays() {
	// Update array of pointers
	for i := 0; i < len(pm.IntPtrArray); i++ {
		if pm.IntPtrArray[i] != nil {
			*pm.IntPtrArray[i] = i * 100
		}
	}
	fmt.Printf("Updated IntPtrArray: ")
	for i, ptr := range pm.IntPtrArray {
		if ptr != nil {
			fmt.Printf("[%d]=%d ", i, *ptr)
		}
	}
	fmt.Println()
	
	// Update slice of pointers
	for i := 0; i < len(pm.IntPtrSlice); i++ {
		if pm.IntPtrSlice[i] != nil {
			*pm.IntPtrSlice[i] = i * 1000
		}
	}
	fmt.Printf("Updated IntPtrSlice: ")
	for i, ptr := range pm.IntPtrSlice {
		if ptr != nil {
			fmt.Printf("[%d]=%d ", i, *ptr)
		}
	}
	fmt.Println()
}

// updateArrayPointers updates array pointer values
func (pm *PointerManager) updateArrayPointers() {
	if pm.ArrayPtr != nil {
		for i := 0; i < len(*pm.ArrayPtr); i++ {
			(*pm.ArrayPtr)[i] = i * 10
		}
		fmt.Printf("Updated ArrayPtr: %v\n", *pm.ArrayPtr)
	}
}

// updateChannelPointers updates channel pointer values
func (pm *PointerManager) updateChannelPointers() {
	if pm.ChanPtr != nil {
		// Send some values to the channel
		go func() {
			for i := 0; i < 3; i++ {
				*pm.ChanPtr <- i * 10
			}
			close(*pm.ChanPtr)
		}()
		
		// Read values from the channel
		fmt.Printf("Channel values: ")
		for val := range *pm.ChanPtr {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

// updateUnsafePointers updates unsafe pointer values
func (pm *PointerManager) updateUnsafePointers() {
	if pm.UnsafePtr != nil {
		// Convert to int pointer and update value
		intPtr := (*int)(pm.UnsafePtr)
		*intPtr = 999
		fmt.Printf("Updated UnsafePtr value to: %d\n", *intPtr)
	}
}

// Delete - Remove pointer references
func (pm *PointerManager) Delete() {
	fmt.Println("\n🗑️  DELETING POINTER REFERENCES:")
	fmt.Println("=================================")
	
	// Delete basic pointers
	pm.IntPtr = nil
	pm.FloatPtr = nil
	pm.StringPtr = nil
	pm.BoolPtr = nil
	fmt.Println("Deleted basic pointers")
	
	// Delete struct pointers
	pm.PersonPtr = nil
	pm.CompanyPtr = nil
	fmt.Println("Deleted struct pointers")
	
	// Delete slice pointers
	pm.IntSlicePtr = nil
	pm.StringSlicePtr = nil
	fmt.Println("Deleted slice pointers")
	
	// Delete map pointers
	pm.MapPtr = nil
	fmt.Println("Deleted map pointers")
	
	// Delete function pointers
	pm.FuncPtr = nil
	fmt.Println("Deleted function pointers")
	
	// Delete interface pointers
	pm.InterfacePtr = nil
	fmt.Println("Deleted interface pointers")
	
	// Delete double pointers
	pm.IntPtrPtr = nil
	pm.StringPtrPtr = nil
	fmt.Println("Deleted double pointers")
	
	// Delete pointer arrays
	for i := range pm.IntPtrArray {
		pm.IntPtrArray[i] = nil
	}
	pm.IntPtrSlice = nil
	fmt.Println("Deleted pointer arrays")
	
	// Delete array pointers
	pm.ArrayPtr = nil
	fmt.Println("Deleted array pointers")
	
	// Delete channel pointers
	pm.ChanPtr = nil
	fmt.Println("Deleted channel pointers")
	
	// Delete unsafe pointers
	pm.UnsafePtr = nil
	fmt.Println("Deleted unsafe pointers")
	
	fmt.Println("✅ Pointer references deleted successfully")
}

// Advanced Pointer Operations

// DemonstratePointerArithmetic shows pointer arithmetic (limited in Go)
func (pm *PointerManager) DemonstratePointerArithmetic() {
	fmt.Println("\n🔢 POINTER ARITHMETIC DEMONSTRATION:")
	fmt.Println("===================================")
	
	// Go doesn't have pointer arithmetic like C, but we can demonstrate
	// what's possible with unsafe operations
	
	// Create an array
	arr := [5]int{10, 20, 30, 40, 50}
	
	// Get pointer to first element
	firstPtr := &arr[0]
	fmt.Printf("First element: %d at %p\n", *firstPtr, firstPtr)
	
	// Get pointer to second element
	secondPtr := &arr[1]
	fmt.Printf("Second element: %d at %p\n", *secondPtr, secondPtr)
	
	// Calculate difference in addresses
	diff := uintptr(unsafe.Pointer(secondPtr)) - uintptr(unsafe.Pointer(firstPtr))
	fmt.Printf("Address difference: %d bytes\n", diff)
	
	// Demonstrate with unsafe pointer arithmetic
	unsafePtr := unsafe.Pointer(&arr[0])
	
	// Move to next element (unsafe)
	nextPtr := unsafe.Pointer(uintptr(unsafePtr) + unsafe.Sizeof(arr[0]))
	nextInt := (*int)(nextPtr)
	fmt.Printf("Next element (unsafe): %d at %p\n", *nextInt, nextInt)
	
	// Move to third element
	thirdPtr := unsafe.Pointer(uintptr(unsafePtr) + 2*unsafe.Sizeof(arr[0]))
	thirdInt := (*int)(thirdPtr)
	fmt.Printf("Third element (unsafe): %d at %p\n", *thirdInt, thirdInt)
}

// DemonstratePointerComparison shows pointer comparison
func (pm *PointerManager) DemonstratePointerComparison() {
	fmt.Println("\n⚖️  POINTER COMPARISON DEMONSTRATION:")
	fmt.Println("====================================")
	
	// Create two variables
	a := 42
	b := 42
	c := a
	
	// Create pointers
	ptrA := &a
	ptrB := &b
	ptrC := &c
	
	// Compare values
	fmt.Printf("a == b: %t (values are equal)\n", a == b)
	fmt.Printf("a == c: %t (values are equal)\n", a == c)
	
	// Compare pointers
	fmt.Printf("ptrA == ptrB: %t (pointers are different)\n", ptrA == ptrB)
	fmt.Printf("ptrA == ptrC: %t (pointers are different)\n", ptrA == ptrC)
	fmt.Printf("ptrA == &a: %t (pointers are same)\n", ptrA == &a)
	
	// Compare nil pointers
	var nilPtr *int
	fmt.Printf("nilPtr == nil: %t\n", nilPtr == nil)
	fmt.Printf("ptrA == nil: %t\n", ptrA == nil)
	
	// Compare addresses
	fmt.Printf("ptrA address: %p\n", ptrA)
	fmt.Printf("ptrB address: %p\n", ptrB)
	fmt.Printf("ptrC address: %p\n", ptrC)
}

// DemonstratePointerDereferencing shows pointer dereferencing
func (pm *PointerManager) DemonstratePointerDereferencing() {
	fmt.Println("\n🎯 POINTER DEREFERENCING DEMONSTRATION:")
	fmt.Println("======================================")
	
	// Create a variable and pointer
	value := 42
	ptr := &value
	
	fmt.Printf("Value: %d\n", value)
	fmt.Printf("Pointer: %p\n", ptr)
	fmt.Printf("Dereferenced: %d\n", *ptr)
	
	// Modify through pointer
	*ptr = 100
	fmt.Printf("After modification through pointer:\n")
	fmt.Printf("Value: %d\n", value)
	fmt.Printf("Dereferenced: %d\n", *ptr)
	
	// Demonstrate with struct
	person := Person{ID: 1, Name: "Alice", Age: 30}
	personPtr := &person
	
	fmt.Printf("\nPerson: %+v\n", person)
	fmt.Printf("Person pointer: %p\n", personPtr)
	fmt.Printf("Dereferenced person: %+v\n", *personPtr)
	
	// Modify struct through pointer
	personPtr.Age = 31
	personPtr.Name = "Alice Updated"
	
	fmt.Printf("After modification through pointer:\n")
	fmt.Printf("Person: %+v\n", person)
	fmt.Printf("Dereferenced person: %+v\n", *personPtr)
}

// DemonstratePointerPassing shows pointer passing in functions
func (pm *PointerManager) DemonstratePointerPassing() {
	fmt.Println("\n📤 POINTER PASSING DEMONSTRATION:")
	fmt.Println("=================================")
	
	// Pass by value
	value := 42
	fmt.Printf("Original value: %d\n", value)
	pm.passByValue(value)
	fmt.Printf("After pass by value: %d\n", value)
	
	// Pass by pointer
	fmt.Printf("Original value: %d\n", value)
	pm.passByPointer(&value)
	fmt.Printf("After pass by pointer: %d\n", value)
	
	// Pass struct by value
	person := Person{ID: 1, Name: "Alice", Age: 30}
	fmt.Printf("Original person: %+v\n", person)
	pm.passStructByValue(person)
	fmt.Printf("After pass struct by value: %+v\n", person)
	
	// Pass struct by pointer
	fmt.Printf("Original person: %+v\n", person)
	pm.passStructByPointer(&person)
	fmt.Printf("After pass struct by pointer: %+v\n", person)
}

// passByValue demonstrates pass by value
func (pm *PointerManager) passByValue(x int) {
	x = 100
	fmt.Printf("Inside passByValue: %d\n", x)
}

// passByPointer demonstrates pass by pointer
func (pm *PointerManager) passByPointer(x *int) {
	*x = 100
	fmt.Printf("Inside passByPointer: %d\n", *x)
}

// passStructByValue demonstrates pass struct by value
func (pm *PointerManager) passStructByValue(p Person) {
	p.Age = 50
	p.Name = "Modified"
	fmt.Printf("Inside passStructByValue: %+v\n", p)
}

// passStructByPointer demonstrates pass struct by pointer
func (pm *PointerManager) passStructByPointer(p *Person) {
	p.Age = 50
	p.Name = "Modified"
	fmt.Printf("Inside passStructByPointer: %+v\n", *p)
}

// DemonstratePointerReturning shows returning pointers
func (pm *PointerManager) DemonstratePointerReturning() {
	fmt.Println("\n📥 POINTER RETURNING DEMONSTRATION:")
	fmt.Println("===================================")
	
	// Return pointer to local variable (dangerous!)
	dangerousPtr := pm.returnPointerToLocal()
	fmt.Printf("Dangerous pointer: %p -> %d\n", dangerousPtr, *dangerousPtr)
	
	// Return pointer to allocated memory
	safePtr := pm.returnPointerToAllocated()
	fmt.Printf("Safe pointer: %p -> %d\n", safePtr, *safePtr)
	
	// Return pointer to struct
	personPtr := pm.returnPersonPointer()
	fmt.Printf("Person pointer: %p -> %+v\n", personPtr, *personPtr)
	
	// Return pointer to slice
	slicePtr := pm.returnSlicePointer()
	fmt.Printf("Slice pointer: %p -> %v\n", slicePtr, *slicePtr)
}

// returnPointerToLocal demonstrates returning pointer to local variable (dangerous!)
func (pm *PointerManager) returnPointerToLocal() *int {
	local := 42
	return &local // This is dangerous! Local variable will be deallocated
}

// returnPointerToAllocated demonstrates returning pointer to allocated memory
func (pm *PointerManager) returnPointerToAllocated() *int {
	allocated := new(int)
	*allocated = 100
	return allocated
}

// returnPersonPointer demonstrates returning pointer to struct
func (pm *PointerManager) returnPersonPointer() *Person {
	person := &Person{
		ID:    1,
		Name:  "Alice",
		Age:   30,
	}
	return person
}

// returnSlicePointer demonstrates returning pointer to slice
func (pm *PointerManager) returnSlicePointer() *[]int {
	slice := &[]int{1, 2, 3, 4, 5}
	return slice
}

// DemonstratePointerReflection shows pointer reflection
func (pm *PointerManager) DemonstratePointerReflection() {
	fmt.Println("\n🔍 POINTER REFLECTION DEMONSTRATION:")
	fmt.Println("====================================")
	
	// Reflect on different pointer types
	value := 42
	ptr := &value
	
	fmt.Printf("Value type: %T\n", value)
	fmt.Printf("Pointer type: %T\n", ptr)
	
	// Use reflection to get pointer information
	ptrValue := reflect.ValueOf(ptr)
	fmt.Printf("Pointer reflect type: %s\n", ptrValue.Type())
	fmt.Printf("Pointer is nil: %t\n", ptrValue.IsNil())
	
	if !ptrValue.IsNil() {
		fmt.Printf("Pointer points to: %v\n", ptrValue.Elem())
		fmt.Printf("Pointer points to type: %s\n", ptrValue.Elem().Type())
	}
	
	// Reflect on struct pointer
	person := Person{ID: 1, Name: "Alice", Age: 30}
	personPtr := &person
	
	personPtrValue := reflect.ValueOf(personPtr)
	fmt.Printf("\nPerson pointer reflect type: %s\n", personPtrValue.Type())
	
	if !personPtrValue.IsNil() {
		personValue := personPtrValue.Elem()
		fmt.Printf("Person pointer points to: %+v\n", personValue.Interface())
		
		// Access struct fields through reflection
		for i := 0; i < personValue.NumField(); i++ {
			field := personValue.Field(i)
			fieldType := personValue.Type().Field(i)
			fmt.Printf("  %s: %v (%s)\n", fieldType.Name, field.Interface(), field.Type())
		}
	}
}

// DemonstratePointerMemoryManagement shows memory management with pointers
func (pm *PointerManager) DemonstratePointerMemoryManagement() {
	fmt.Println("\n💾 POINTER MEMORY MANAGEMENT DEMONSTRATION:")
	fmt.Println("==========================================")
	
	// Get memory stats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Initial memory: %d KB\n", m.Alloc/1024)
	
	// Allocate memory
	ptr1 := new(int)
	*ptr1 = 42
	
	ptr2 := new(Person)
	ptr2.ID = 1
	ptr2.Name = "Alice"
	ptr2.Age = 30
	
	// Allocate slice
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}
	
	runtime.ReadMemStats(&m)
	fmt.Printf("After allocation: %d KB\n", m.Alloc/1024)
	
	// Force garbage collection
	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Printf("After GC: %d KB\n", m.Alloc/1024)
	
	// Set pointers to nil to allow garbage collection
	ptr1 = nil
	ptr2 = nil
	slice = nil
	
	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Printf("After setting to nil and GC: %d KB\n", m.Alloc/1024)
}

// DemonstratePointerChaining shows pointer chaining
func (pm *PointerManager) DemonstratePointerChaining() {
	fmt.Println("\n🔗 POINTER CHAINING DEMONSTRATION:")
	fmt.Println("==================================")
	
	// Create a chain of pointers
	value := 42
	ptr1 := &value
	ptr2 := &ptr1
	ptr3 := &ptr2
	
	fmt.Printf("Value: %d\n", value)
	fmt.Printf("Ptr1: %p -> %d\n", ptr1, *ptr1)
	fmt.Printf("Ptr2: %p -> %p -> %d\n", ptr2, *ptr2, **ptr2)
	fmt.Printf("Ptr3: %p -> %p -> %p -> %d\n", ptr3, *ptr3, **ptr3, ***ptr3)
	
	// Modify through chain
	***ptr3 = 100
	fmt.Printf("\nAfter modification through chain:\n")
	fmt.Printf("Value: %d\n", value)
	fmt.Printf("Ptr1: %p -> %d\n", ptr1, *ptr1)
	fmt.Printf("Ptr2: %p -> %p -> %d\n", ptr2, *ptr2, **ptr2)
	fmt.Printf("Ptr3: %p -> %p -> %p -> %d\n", ptr3, *ptr3, **ptr3, ***ptr3)
	
	// Demonstrate with struct
	person := Person{ID: 1, Name: "Alice", Age: 30}
	personPtr1 := &person
	personPtr2 := &personPtr1
	
	fmt.Printf("\nPerson: %+v\n", person)
	fmt.Printf("PersonPtr1: %p -> %+v\n", personPtr1, *personPtr1)
	fmt.Printf("PersonPtr2: %p -> %p -> %+v\n", personPtr2, *personPtr2, **personPtr2)
	
	// Modify through chain
	(**personPtr2).Age = 31
	(**personPtr2).Name = "Alice Updated"
	
	fmt.Printf("\nAfter modification through chain:\n")
	fmt.Printf("Person: %+v\n", person)
	fmt.Printf("PersonPtr1: %p -> %+v\n", personPtr1, *personPtr1)
	fmt.Printf("PersonPtr2: %p -> %p -> %+v\n", personPtr2, *personPtr2, **personPtr2)
}
