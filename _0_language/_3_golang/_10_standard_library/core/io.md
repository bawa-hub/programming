# io Package - I/O Primitives 📁

The `io` package provides basic interfaces to I/O primitives. It's the foundation for all I/O operations in Go, defining interfaces that are implemented by many other packages.

## 🎯 Key Concepts

### 1. **Core Interfaces**
- `Reader` - Read data from a source
- `Writer` - Write data to a destination
- `Closer` - Close a resource
- `Seeker` - Seek to a position
- `ReadWriter` - Combines Reader and Writer
- `ReadCloser` - Combines Reader and Closer
- `WriteCloser` - Combines Writer and Closer
- `ReadWriteCloser` - Combines Reader, Writer, and Closer

### 2. **Reader Interface**
```go
type Reader interface {
    Read([]byte) (n int, err error)
}
```

### 3. **Writer Interface**
```go
type Writer interface {
    Write([]byte) (n int, err error)
}
```

### 4. **Utility Functions**
- `Copy()` - Copy from Reader to Writer
- `CopyN()` - Copy N bytes from Reader to Writer
- `ReadAll()` - Read all data from Reader
- `ReadAtLeast()` - Read at least N bytes
- `ReadFull()` - Read exactly N bytes
- `WriteString()` - Write string to Writer
- `Pipe()` - Create a pipe

### 5. **Multi-Reader/Writer**
- `MultiReader()` - Read from multiple readers
- `MultiWriter()` - Write to multiple writers
- `TeeReader()` - Read from reader and write to writer

### 6. **Limited Reader**
- `LimitReader()` - Limit the number of bytes read
- `SectionReader()` - Read from a section of a reader

## 🚀 Common Patterns

### Reading from a Reader
```go
data, err := io.ReadAll(reader)
if err != nil {
    log.Fatal(err)
}
```

### Writing to a Writer
```go
_, err := io.WriteString(writer, "Hello World")
if err != nil {
    log.Fatal(err)
}
```

### Copying Data
```go
_, err := io.Copy(dst, src)
if err != nil {
    log.Fatal(err)
}
```

### Reading with Limit
```go
limitedReader := io.LimitReader(reader, 1024)
data, err := io.ReadAll(limitedReader)
```

## ⚠️ Common Pitfalls

1. **Not checking errors** - Always check error returns
2. **Infinite loops** - Be careful with Read() loops
3. **Buffer size** - Choose appropriate buffer sizes
4. **Resource leaks** - Always close resources
5. **Partial reads** - Handle partial read scenarios

## 🎯 Best Practices

1. **Check errors** - Always handle error returns
2. **Use appropriate interfaces** - Choose the right interface for your needs
3. **Handle partial reads** - Implement proper retry logic
4. **Close resources** - Use defer statements
5. **Use utility functions** - Prefer io.Copy over manual loops

## 🔍 Advanced Features

### Custom Reader
```go
type CustomReader struct {
    data []byte
    pos  int
}

func (r *CustomReader) Read(p []byte) (n int, err error) {
    if r.pos >= len(r.data) {
        return 0, io.EOF
    }
    n = copy(p, r.data[r.pos:])
    r.pos += n
    return n, nil
}
```

### Custom Writer
```go
type CustomWriter struct {
    buffer []byte
}

func (w *CustomWriter) Write(p []byte) (n int, err error) {
    w.buffer = append(w.buffer, p...)
    return len(p), nil
}
```

### Pipe Operations
```go
reader, writer := io.Pipe()
go func() {
    writer.Write([]byte("Hello"))
    writer.Close()
}()
data, _ := io.ReadAll(reader)
```

## 📚 Real-world Applications

1. **File I/O** - Reading and writing files
2. **Network I/O** - HTTP requests and responses
3. **Data Processing** - Stream processing
4. **Logging** - Writing to log files
5. **Compression** - Working with compressed data

## 🧠 Memory Tips

- **io** = **I**nput/**O**utput
- **Reader** = **R**ead data
- **Writer** = **W**rite data
- **Closer** = **C**lose resource
- **Seeker** = **S**eek position
- **Copy** = **C**opy data
- **ReadAll** = **R**ead **A**ll data
- **WriteString** = **W**rite **S**tring

Remember: The io package is the foundation of all I/O in Go! 🎯
