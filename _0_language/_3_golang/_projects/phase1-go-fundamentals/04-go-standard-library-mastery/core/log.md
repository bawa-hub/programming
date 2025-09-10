# log Package - Logging 📝

The `log` package provides a simple logging interface for Go programs. It's essential for debugging, monitoring, and maintaining applications.

## 🎯 Key Concepts

### 1. **Log Levels**
- `Fatal` - Log and exit program
- `Panic` - Log and panic
- `Error` - Error messages
- `Warning` - Warning messages
- `Info` - Informational messages
- `Debug` - Debug messages

### 2. **Basic Functions**
- `Print()` - Print message
- `Printf()` - Print formatted message
- `Println()` - Print message with newline
- `Fatal()` - Print and exit
- `Fatalf()` - Print formatted and exit
- `Fataln()` - Print and exit with newline
- `Panic()` - Print and panic
- `Panicf()` - Print formatted and panic
- `Panicn()` - Print and panic with newline

### 3. **Logger Configuration**
- `SetFlags()` - Set output flags
- `SetPrefix()` - Set log prefix
- `SetOutput()` - Set output destination
- `SetOutput()` - Set output writer

### 4. **Log Flags**
- `Ldate` - Date in local time
- `Ltime` - Time in local time
- `Lmicroseconds` - Microsecond precision
- `Llongfile` - Full file name and line number
- `Lshortfile` - File name and line number
- `LUTC` - UTC time
- `LstdFlags` - Standard flags (date and time)

### 5. **Custom Loggers**
- `New()` - Create new logger
- `Logger` - Logger type
- `SetOutput()` - Set output writer
- `SetFlags()` - Set output flags
- `SetPrefix()` - Set log prefix

## 🚀 Common Patterns

### Basic Logging
```go
log.Print("Hello, World!")
log.Printf("User %s logged in", username)
log.Println("Operation completed")
```

### Error Logging
```go
if err != nil {
    log.Printf("Error: %v", err)
}
```

### Fatal Logging
```go
if err != nil {
    log.Fatalf("Fatal error: %v", err)
}
```

### Custom Logger
```go
logger := log.New(os.Stdout, "PREFIX: ", log.LstdFlags)
logger.Println("Custom log message")
```

## ⚠️ Common Pitfalls

1. **Logging sensitive data** - Don't log passwords or tokens
2. **Too much logging** - Avoid excessive log output
3. **Inconsistent formatting** - Use consistent log formats
4. **Not handling errors** - Always check for logging errors
5. **Performance impact** - Consider logging performance

## 🎯 Best Practices

1. **Use appropriate levels** - Choose right log level
2. **Include context** - Add relevant information
3. **Use structured logging** - Use consistent format
4. **Handle errors** - Check for logging errors
5. **Configure appropriately** - Set appropriate flags

## 🔍 Advanced Features

### Custom Log Format
```go
log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
log.SetPrefix("APP: ")
```

### Log to File
```go
file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
if err != nil {
    log.Fatal(err)
}
defer file.Close()
log.SetOutput(file)
```

### Multiple Loggers
```go
infoLogger := log.New(os.Stdout, "INFO: ", log.LstdFlags)
errorLogger := log.New(os.Stderr, "ERROR: ", log.LstdFlags)
```

## 📚 Real-world Applications

1. **Application Logging** - General application logs
2. **Error Tracking** - Error and exception logging
3. **Audit Logging** - Security and compliance logs
4. **Performance Monitoring** - Performance metrics
5. **Debugging** - Development and debugging

## 🧠 Memory Tips

- **log** = **L**ogging **O**perations **G**enerator
- **Print** = **P**rint message
- **Fatal** = **F**atal error
- **Panic** = **P**anic error
- **Flags** = **F**lags for formatting
- **Prefix** = **P**refix for messages
- **Output** = **O**utput destination

Remember: The log package is your gateway to application monitoring in Go! 🎯
