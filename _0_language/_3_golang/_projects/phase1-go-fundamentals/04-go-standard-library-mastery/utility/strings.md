# strings Package - String Manipulation 🔤

The `strings` package provides functions for manipulating UTF-8 encoded strings. It's essential for text processing, parsing, and string operations in Go.

## 🎯 Key Concepts

### 1. **String Comparison**
- `Compare()` - Compare two strings
- `EqualFold()` - Case-insensitive comparison
- `HasPrefix()` - Check if string starts with prefix
- `HasSuffix()` - Check if string ends with suffix
- `Contains()` - Check if string contains substring
- `ContainsAny()` - Check if string contains any characters
- `ContainsRune()` - Check if string contains rune

### 2. **String Searching**
- `Index()` - Find first occurrence of substring
- `LastIndex()` - Find last occurrence of substring
- `IndexAny()` - Find first occurrence of any character
- `LastIndexAny()` - Find last occurrence of any character
- `IndexRune()` - Find first occurrence of rune
- `IndexFunc()` - Find first character satisfying function
- `LastIndexFunc()` - Find last character satisfying function

### 3. **String Manipulation**
- `ToUpper()` - Convert to uppercase
- `ToLower()` - Convert to lowercase
- `ToTitle()` - Convert to title case
- `Title()` - Convert to title case (deprecated)
- `TrimSpace()` - Remove leading/trailing whitespace
- `Trim()` - Remove leading/trailing characters
- `TrimLeft()` - Remove leading characters
- `TrimRight()` - Remove trailing characters
- `TrimPrefix()` - Remove prefix
- `TrimSuffix()` - Remove suffix

### 4. **String Splitting**
- `Split()` - Split string by separator
- `SplitN()` - Split string with limit
- `SplitAfter()` - Split after separator
- `SplitAfterN()` - Split after separator with limit
- `Fields()` - Split by whitespace
- `FieldsFunc()` - Split by function

### 5. **String Joining**
- `Join()` - Join strings with separator
- `Repeat()` - Repeat string N times
- `Replace()` - Replace occurrences
- `ReplaceAll()` - Replace all occurrences
- `Map()` - Apply function to each character

### 6. **String Building**
- `Builder` - Efficient string building
- `WriteString()` - Write string to builder
- `WriteRune()` - Write rune to builder
- `WriteByte()` - Write byte to builder
- `String()` - Get built string
- `Reset()` - Reset builder

## 🚀 Common Patterns

### Basic String Operations
```go
s := "Hello, World!"
fmt.Println(strings.ToUpper(s))        // "HELLO, WORLD!"
fmt.Println(strings.ToLower(s))        // "hello, world!"
fmt.Println(strings.TrimSpace(s))      // "Hello, World!"
fmt.Println(strings.HasPrefix(s, "Hello")) // true
```

### String Searching
```go
s := "Hello, World!"
fmt.Println(strings.Index(s, "World"))     // 7
fmt.Println(strings.Contains(s, "World"))  // true
fmt.Println(strings.Count(s, "l"))         // 3
```

### String Splitting and Joining
```go
s := "apple,banana,cherry"
parts := strings.Split(s, ",")
fmt.Println(parts) // ["apple", "banana", "cherry"]

joined := strings.Join(parts, " | ")
fmt.Println(joined) // "apple | banana | cherry"
```

### String Building
```go
var builder strings.Builder
builder.WriteString("Hello")
builder.WriteString(" ")
builder.WriteString("World")
result := builder.String()
fmt.Println(result) // "Hello World"
```

## ⚠️ Common Pitfalls

1. **Unicode handling** - Go strings are UTF-8 encoded
2. **Case sensitivity** - String operations are case sensitive by default
3. **Empty strings** - Handle empty strings properly
4. **Memory allocation** - Use Builder for efficient string building
5. **Rune vs byte** - Understand difference between runes and bytes

## 🎯 Best Practices

1. **Use Builder** - For efficient string concatenation
2. **Handle Unicode** - Be aware of UTF-8 encoding
3. **Check bounds** - Validate string indices
4. **Use appropriate functions** - Choose the right function for the task
5. **Consider performance** - Use efficient string operations

## 🔍 Advanced Features

### Custom String Functions
```go
func CustomTrim(s string, cutset string) string {
    return strings.Trim(s, cutset)
}

func CustomSplit(s string, sep string) []string {
    return strings.Split(s, sep)
}
```

### String Validation
```go
func IsValidEmail(email string) bool {
    return strings.Contains(email, "@") && 
           strings.Contains(email, ".")
}
```

### String Transformation
```go
func TransformString(s string, fn func(rune) rune) string {
    return strings.Map(fn, s)
}
```

## 📚 Real-world Applications

1. **Text Processing** - Parsing and manipulating text
2. **Data Validation** - Validating input strings
3. **String Building** - Efficient string concatenation
4. **Text Analysis** - Analyzing text content
5. **String Formatting** - Formatting strings for output

## 🧠 Memory Tips

- **strings** = **S**tring **T**ext **R**egex **I**nput **N**etwork **G**o
- **Compare** = **C**ompare strings
- **Contains** = **C**ontains substring
- **Index** = **I**ndex of substring
- **Join** = **J**oin strings
- **Split** = **S**plit string
- **Trim** = **T**rim characters
- **Builder** = **B**uilder for strings

Remember: The strings package is your gateway to text processing in Go! 🎯
