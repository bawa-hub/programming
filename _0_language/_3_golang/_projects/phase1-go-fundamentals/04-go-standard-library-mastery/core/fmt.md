# fmt Package - Formatting and Printing 📝

The `fmt` package is one of the most fundamental packages in Go. It provides formatted I/O functions similar to C's printf and scanf.

## 🎯 Key Concepts

### 1. **Print Functions**
- `Print()` - Print to standard output
- `Printf()` - Print with formatting
- `Println()` - Print with newline
- `Sprint()` - Return string instead of printing
- `Sprintf()` - Return formatted string
- `Sprintln()` - Return string with newline

### 2. **Scan Functions**
- `Scan()` - Read from standard input
- `Scanf()` - Read with formatting
- `Scanln()` - Read until newline
- `Sscan()` - Read from string
- `Sscanf()` - Read from string with formatting
- `Sscanln()` - Read from string until newline

### 3. **Format Verbs**
- `%v` - Default format
- `%+v` - Add field names for structs
- `%#v` - Go syntax representation
- `%T` - Type of the value
- `%t` - Boolean
- `%d` - Integer (base 10)
- `%b` - Binary
- `%o` - Octal
- `%x` - Hexadecimal (lowercase)
- `%X` - Hexadecimal (uppercase)
- `%f` - Floating point
- `%e` - Scientific notation
- `%E` - Scientific notation (uppercase)
- `%g` - Choose between %f and %e
- `%s` - String
- `%q` - Quoted string
- `%c` - Character
- `%p` - Pointer address

### 4. **Width and Precision**
- `%5d` - Width 5
- `%05d` - Width 5, zero-padded
- `%.2f` - Precision 2
- `%5.2f` - Width 5, precision 2

## 🚀 Common Patterns

### String Formatting
```go
name := "Alice"
age := 30
fmt.Printf("Name: %s, Age: %d\n", name, age)
```

### Struct Formatting
```go
type Person struct {
    Name string
    Age  int
}
p := Person{Name: "Bob", Age: 25}
fmt.Printf("%+v\n", p) // {Name:Bob Age:25}
fmt.Printf("%#v\n", p) // main.Person{Name:"Bob", Age:25}
```

### Error Formatting
```go
err := errors.New("something went wrong")
fmt.Printf("Error: %v\n", err)
```

## ⚠️ Common Pitfalls

1. **Missing format verbs** - Always provide correct number of verbs
2. **Type mismatches** - Ensure verb types match argument types
3. **Memory leaks** - Be careful with string formatting in loops
4. **Security** - Don't use user input directly in format strings

## 🎯 Best Practices

1. **Use appropriate verbs** - Choose the right format for your data type
2. **Be consistent** - Use the same formatting style throughout your code
3. **Handle errors** - Always check return values from scan functions
4. **Use Sprint for strings** - When you need formatted strings, not printing
5. **Prefer Printf over Print** - More readable and maintainable

## 🔍 Advanced Features

### Custom Formatting
```go
type CustomType struct {
    Value int
}

func (c CustomType) String() string {
    return fmt.Sprintf("CustomType{%d}", c.Value)
}
```

### Error Formatting
```go
type CustomError struct {
    Code    int
    Message string
}

func (e CustomError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}
```

## 📚 Real-world Applications

1. **Logging** - Format log messages
2. **CLI Tools** - Display output to users
3. **Debugging** - Print variable values
4. **Configuration** - Parse configuration files
5. **Data Processing** - Format data for output

## 🧠 Memory Tips

- **fmt** = **F**ormatting **M**ade **T**rivial
- **%v** = **V**alue (default)
- **%+v** = **V**alue with **+** field names
- **%#v** = **G**o syntax (starts with **G**)
- **%T** = **T**ype
- **%t** = **T**rue/false (boolean)
- **%d** = **D**ecimal
- **%s** = **S**tring
- **%p** = **P**ointer

Remember: The fmt package is your gateway to readable output in Go! 🎯
