package main

// You could use an alias like f "fmt"
import (
	"bufio"
	"fmt"
	"strings"

	"log"
	"os"
)

var pl = fmt.Println

/*
I'm a block comment
*/

func main() {

	// output
	// Prints text and a newline
	pl("Hello World!")

	// input methods
	// 1️⃣ Basic Method — Using fmt.Scan / Scanln / Scanf
	// 2️⃣ Reading Full Line String Input (with spaces) using bufio + os.Stdin

	var i int
	var f float64
	var b bool
	var s string
	var ch rune
	var line string

	fmt.Print("Enter int: ")
	fmt.Scan(&i)

	fmt.Print("Enter float: ")
	fmt.Scan(&f)

	fmt.Print("Enter bool: ")
	fmt.Scan(&b)

	fmt.Print("Enter word string: ")
	fmt.Scan(&s)

	// fmt.Print("Enter character: ")
	// fmt.Scanf(" %c", &ch)   // Leading space FIXES newline issue

	// Setup buffered reader that gets text from the keyboard
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter character: ")
	chStr, _ := reader.ReadString('\n')
	ch = []rune(strings.TrimSpace(chStr))[0]

	fmt.Print("Enter full line string: ")
	line, _ = reader.ReadString('\n')

	fmt.Println("\n--- Output ---")
	fmt.Println("Int:", i)
	fmt.Println("Float:", f)
	fmt.Println("Bool:", b)
	fmt.Println("String:", s)
	fmt.Println("Char:", string(ch))
	fmt.Println("Line:", line)

	// The blank identifier _ will get err and ignore it (Bad Practice)
	// name, _ := reader.ReadString('\n')

	// It is better to handle it
	pl("what is your name?")
	name, err := reader.ReadString('\n') // get input upto new line
	if err == nil {
		pl("Hello", name)
	} else {
		// Log this error
		log.Fatal(err)
	}

	// input output array
	var n int
	fmt.Print("Enter array size: ")
	fmt.Scan(&n)

	var arr [100]int // fixed max size

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Print("Array elements: ")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}

	// io slice
	var m int
	fmt.Print("Enter slice size: ")
	fmt.Scan(&m)

	sl := make([]int, m)

	for i := 0; i < m; i++ {
		fmt.Scan(&sl[i])
	}

	fmt.Print("Slice elements: ")
	for _, v := range sl {
		fmt.Print(v, " ")
	}

	// io struct
	type Person struct {
		Name   string
		Age    int
		Salary float64
	}

	var p Person

	fmt.Print("Enter Name Age Salary: ")
	fmt.Scan(&p.Name, &p.Age, &p.Salary)

	fmt.Println("\n--- Output ---")
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Salary:", p.Salary)

	type Student struct {
	Name    string
	City    string
	Age     int
}

	var st Student

	fmt.Print("Enter Name: ")
	st.Name, _ = reader.ReadString('\n')
	st.Name = strings.TrimSpace(st.Name)

	fmt.Print("Enter City: ")
	st.City, _ = reader.ReadString('\n')
	st.City = strings.TrimSpace(st.City)

	fmt.Print("Enter Age: ")
	fmt.Scan(&st.Age)

	fmt.Println("\n--- Output ---")
	fmt.Printf("%+v\n", st)

}
