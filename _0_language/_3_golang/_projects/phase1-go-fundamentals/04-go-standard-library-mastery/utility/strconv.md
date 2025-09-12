# strconv Package - String Conversion 🔄

The `strconv` package provides functions for converting between strings and basic data types. It's essential for parsing user input, formatting output, and data conversion.

## 🎯 Key Concepts

### 1. **String to Number Conversion**
- `Atoi()` - String to int
- `ParseInt()` - String to int64 with base
- `ParseUint()` - String to uint64 with base
- `ParseFloat()` - String to float64
- `ParseBool()` - String to bool
- `ParseComplex()` - String to complex128

### 2. **Number to String Conversion**
- `Itoa()` - Int to string
- `FormatInt()` - Int64 to string with base
- `FormatUint()` - Uint64 to string with base
- `FormatFloat()` - Float64 to string
- `FormatBool()` - Bool to string
- `FormatComplex()` - Complex128 to string

### 3. **String Formatting**
- `Quote()` - Add quotes to string
- `QuoteRune()` - Add quotes to rune
- `QuoteRuneToASCII()` - Add quotes to rune (ASCII)
- `QuoteToASCII()` - Add quotes to string (ASCII)
- `Unquote()` - Remove quotes from string
- `UnquoteChar()` - Remove quotes from character

### 4. **String Appending**
- `AppendInt()` - Append int to byte slice
- `AppendUint()` - Append uint to byte slice
- `AppendFloat()` - Append float to byte slice
- `AppendBool()` - Append bool to byte slice
- `AppendQuote()` - Append quoted string to byte slice
- `AppendQuoteRune()` - Append quoted rune to byte slice

### 5. **String Validation**
- `IsPrint()` - Check if rune is printable
- `IsGraphic()` - Check if rune is graphic
- `CanBackquote()` - Check if string can be backquoted
- `IsQuote()` - Check if rune is quote character

### 6. **Error Handling**
- `NumError` - Number conversion error
- `ErrRange` - Value out of range
- `ErrSyntax` - Invalid syntax
- `InvalidBase` - Invalid base for conversion

## 🚀 Common Patterns

### Basic String Conversion
```go
// String to number
s := "123"
i, err := strconv.Atoi(s)
if err != nil {
    log.Fatal(err)
}

// Number to string
i := 123
s := strconv.Itoa(i)
```

### String Parsing with Base
```go
// Parse with different bases
s := "1010"
i, err := strconv.ParseInt(s, 2, 64) // binary
i, err := strconv.ParseInt(s, 8, 64) // octal
i, err := strconv.ParseInt(s, 16, 64) // hexadecimal
```

### Float Conversion
```go
// String to float
s := "3.14159"
f, err := strconv.ParseFloat(s, 64)

// Float to string
f := 3.14159
s := strconv.FormatFloat(f, 'f', 2, 64) // 2 decimal places
```

### String Quoting
```go
// Add quotes
s := "hello"
quoted := strconv.Quote(s) // "hello"

// Remove quotes
s := "\"hello\""
unquoted, err := strconv.Unquote(s) // hello
```

## ⚠️ Common Pitfalls

1. **Error handling** - Always check for conversion errors
2. **Base validation** - Ensure base is between 2 and 36
3. **Bit size** - Use appropriate bit size for target type
4. **Float precision** - Be aware of floating point precision
5. **Unicode handling** - Understand rune vs byte conversion

## 🎯 Best Practices

1. **Handle errors** - Always check conversion errors
2. **Use appropriate types** - Choose the right numeric type
3. **Validate input** - Check input before conversion
4. **Use constants** - Use predefined constants for bases
5. **Consider performance** - Use Append functions for efficiency

## 🔍 Advanced Features

### Custom Conversion Functions
```go
func SafeAtoi(s string) (int, error) {
    if s == "" {
        return 0, fmt.Errorf("empty string")
    }
    return strconv.Atoi(s)
}
```

### String Validation
```go
func IsNumeric(s string) bool {
    _, err := strconv.Atoi(s)
    return err == nil
}
```

### Number Formatting
```go
func FormatNumber(n int64) string {
    return strconv.FormatInt(n, 10)
}
```

## 📚 Real-world Applications

1. **User Input** - Parsing command line arguments
2. **Configuration** - Converting config values
3. **Data Processing** - Converting between formats
4. **API Development** - Parsing request parameters
5. **File Processing** - Converting file data

## 🧠 Memory Tips

- **strconv** = **S**tring **T**ype **R**egex **C**onversion **O**perations **N**etwork **V**alidation
- **Atoi** = **A**scii **T**o **I**nteger
- **Itoa** = **I**nteger **T**o **A**scii
- **Parse** = **P**arse string
- **Format** = **F**ormat value
- **Quote** = **Q**uote string
- **Append** = **A**ppend to slice

Remember: The strconv package is your gateway to type conversion in Go! 🎯
