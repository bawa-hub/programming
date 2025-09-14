package main

import "fmt"

// A struct is a collection of fields, like a lightweight object.
type User struct {
    ID    int
    Name  string
    Email string
}

// Struct Methods (Value vs Pointer Receivers)
// Methods can be attached to structs.
// Value receiver → works on a copy.
// Pointer receiver → works on original.
type Counter struct {
    value int
}

func (c Counter) IncrementByCopy() {
    c.value++ // modifies copy
}

func (c *Counter) IncrementByPointer() {
    c.value++ // modifies original
}

// “When should you use pointer vs value receiver?”
	// Use pointer receiver if:
	// Method modifies receiver
	// Struct is large (avoid copying)
	// Use value receiver if:
	// Struct is small and immutable (e.g., time.Time, string wrappers)

// Struct Embedding (Composition > Inheritance)
// Go doesn’t have inheritance, but embedding lets you compose behaviors.	

type Person struct {
    Name string
}

type Employee struct {
    Person  // embedded
    Role    string
}

	


func main() {

	u1 := User{ID: 1, Name: "Alice", Email: "a@example.com"} // full init
	u2 := User{Name: "Bob"} // partial init
	u3 := new(User)         // pointer to zero-value struct
	u4 := &User{ID: 2}      // pointer to struct literal
	fmt.Println(u1);
	fmt.Println(u2)
	fmt.Println(u3)
	fmt.Println(u4)

	c := Counter{}
    c.IncrementByCopy()
    fmt.Println(c.value) // 0
    c.IncrementByPointer()
    fmt.Println(c.value) // 1

	e := Employee{Person: Person{Name: "Alice"}, Role: "Engineer"}
    fmt.Println(e.Name)  // promoted field → "Alice"

}