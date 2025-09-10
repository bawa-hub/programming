# errors Package - Error Handling 🚨

The `errors` package provides functions for manipulating errors. It's essential for creating, wrapping, and handling errors in Go applications.

## 🎯 Key Concepts

### 1. **Core Functions**
- `New()` - Create new error
- `Unwrap()` - Unwrap error chain
- `Is()` - Check if error matches target
- `As()` - Check if error can be converted to target type
- `Join()` - Join multiple errors

### 2. **Error Interface**
```go
type error interface {
    Error() string
}
```

### 3. **Error Wrapping**
- `fmt.Errorf()` - Wrap error with message
- `errors.Unwrap()` - Unwrap wrapped error
- `errors.Is()` - Check error chain
- `errors.As()` - Type assertion on error chain

### 4. **Custom Error Types**
- Implement `error` interface
- Add additional fields and methods
- Use type assertions for error handling

### 5. **Error Chains**
- Wrap errors to provide context
- Unwrap errors to access original error
- Check error chains for specific errors

## 🚀 Common Patterns

### Basic Error Creation
```go
err := errors.New("something went wrong")
if err != nil {
    log.Fatal(err)
}
```

### Error Wrapping
```go
if err != nil {
    return fmt.Errorf("failed to process: %w", err)
}
```

### Error Checking
```go
if errors.Is(err, os.ErrNotExist) {
    // Handle file not found
}
```

### Type Assertion
```go
var pathErr *os.PathError
if errors.As(err, &pathErr) {
    // Handle path error
}
```

## ⚠️ Common Pitfalls

1. **Ignoring errors** - Always handle errors
2. **Generic error messages** - Provide specific context
3. **Error swallowing** - Don't hide errors
4. **Panic on error** - Use proper error handling
5. **Infinite error chains** - Avoid circular references

## 🎯 Best Practices

1. **Handle errors immediately** - Don't ignore them
2. **Wrap with context** - Add meaningful information
3. **Use specific error types** - Create custom errors
4. **Check error chains** - Use `Is()` and `As()`
5. **Log errors appropriately** - Use appropriate log levels

## 🔍 Advanced Features

### Custom Error Types
```go
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation failed for %s: %s", e.Field, e.Message)
}
```

### Error Wrapping
```go
func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open %s: %w", filename, err)
    }
    defer file.Close()
    // ... process file
    return nil
}
```

## 📚 Real-world Applications

1. **API Development** - HTTP error responses
2. **Database Operations** - Connection and query errors
3. **File Operations** - I/O errors
4. **Network Operations** - Connection errors
5. **Validation** - Input validation errors

## 🧠 Memory Tips

- **errors** = **E**rror **R**esponse **R**ecovery **O**perations **R**esource **S**ystem
- **New** = **N**ew error
- **Unwrap** = **U**nwrap error chain
- **Is** = **I**s error type
- **As** = **A**s type assertion
- **Join** = **J**oin errors

Remember: The errors package is your gateway to robust error handling in Go! 🎯
