# time Package - Time and Date Operations ⏰

The `time` package provides functionality for measuring and displaying time. It's essential for timestamps, durations, timeouts, and scheduling operations.

## 🎯 Key Concepts

### 1. **Time Types**
- `Time` - Represents an instant in time
- `Duration` - Represents elapsed time between two instants
- `Location` - Represents a time zone
- `Ticker` - Recurring timer
- `Timer` - One-time timer

### 2. **Time Creation**
- `Now()` - Current time
- `Date()` - Create time from components
- `Parse()` - Parse time from string
- `Unix()` - Create time from Unix timestamp
- `UnixMicro()` - Create time from Unix microsecond timestamp

### 3. **Time Formatting**
- `Format()` - Format time as string
- `String()` - Default string representation
- `MarshalText()` - Marshal to text
- `MarshalJSON()` - Marshal to JSON

### 4. **Time Parsing**
- `Parse()` - Parse time from string
- `ParseInLocation()` - Parse time in specific location
- `ParseDuration()` - Parse duration from string

### 5. **Time Operations**
- `Add()` - Add duration to time
- `Sub()` - Subtract time from time
- `After()` - Check if time is after another
- `Before()` - Check if time is before another
- `Equal()` - Check if times are equal

### 6. **Duration Operations**
- `Nanoseconds()` - Duration in nanoseconds
- `Microseconds()` - Duration in microseconds
- `Milliseconds()` - Duration in milliseconds
- `Seconds()` - Duration in seconds
- `Minutes()` - Duration in minutes
- `Hours()` - Duration in hours

### 7. **Timer and Ticker**
- `NewTimer()` - Create one-time timer
- `NewTicker()` - Create recurring ticker
- `AfterFunc()` - Execute function after duration
- `Sleep()` - Sleep for duration

## 🚀 Common Patterns

### Current Time
```go
now := time.Now()
fmt.Printf("Current time: %s\n", now)
```

### Time Formatting
```go
now := time.Now()
formatted := now.Format("2006-01-02 15:04:05")
fmt.Printf("Formatted: %s\n", formatted)
```

### Duration
```go
duration := 5 * time.Minute
fmt.Printf("Duration: %s\n", duration)
```

### Timer
```go
timer := time.NewTimer(2 * time.Second)
<-timer.C
fmt.Println("Timer fired!")
```

## ⚠️ Common Pitfalls

1. **Wrong reference time** - Use `2006-01-02 15:04:05` for formatting
2. **Timezone confusion** - Always specify timezone
3. **Duration arithmetic** - Be careful with duration calculations
4. **Timer leaks** - Always stop timers
5. **Parsing errors** - Always check parsing errors

## 🎯 Best Practices

1. **Use UTC** - Store times in UTC, convert for display
2. **Check errors** - Always handle parsing errors
3. **Stop timers** - Use defer to stop timers
4. **Use constants** - Use time constants for common durations
5. **Be precise** - Use appropriate precision for your needs

## 🔍 Advanced Features

### Custom Time Format
```go
const customFormat = "2006-01-02 15:04:05 MST"
t, _ := time.Parse(customFormat, "2023-12-25 12:00:00 UTC")
```

### Time Zone Handling
```go
loc, _ := time.LoadLocation("America/New_York")
t := time.Now().In(loc)
```

### Duration Arithmetic
```go
start := time.Now()
// ... do work ...
elapsed := time.Since(start)
```

## 📚 Real-world Applications

1. **Logging** - Timestamp log entries
2. **Scheduling** - Cron jobs and tasks
3. **Timeouts** - Network and operation timeouts
4. **Benchmarking** - Measure performance
5. **Caching** - Expiration times

## 🧠 Memory Tips

- **time** = **T**ime **I**nterface **M**anagement **E**ngine
- **Now()** = **N**ow (current time)
- **Date()** = **D**ate creation
- **Parse()** = **P**arse string
- **Format()** = **F**ormat time
- **Duration** = **D**uration of time
- **Timer** = **T**imer for events
- **Ticker** = **T**icker for recurring events

Remember: The time package is your gateway to temporal operations in Go! 🎯
