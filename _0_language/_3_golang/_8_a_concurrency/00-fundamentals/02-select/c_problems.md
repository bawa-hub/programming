# 🔀 Go Select Statement Problems
## 50 Problems from Basic to Advanced

### 📚 Problem Categories
- **Basic Select (1-10)**: Simple select statements, single channels
- **Multiple Channels (11-20)**: Select with multiple channels, multiplexing
- **Non-blocking Operations (21-30)**: Default cases, non-blocking sends/receives
- **Timeout Patterns (31-40)**: Timeout handling, time.After, time.Tick
- **Advanced Patterns (41-50)**: Complex select patterns, real-world scenarios

---

## 🟢 BASIC SELECT (1-10)

### Problem 1: Simple Select
**Task**: Use select to receive from a single channel.

**Expected Output**:
```
Received: Hello
```

### Problem 2: Select with Two Channels
**Task**: Use select to receive from two different channels.

**Expected Output**:
```
Received from ch1: Message 1
Received from ch2: Message 2
```

### Problem 3: Select with Send
**Task**: Use select to send data to a channel.

**Expected Output**:
```
Sent: 42
```

### Problem 4: Select with Send and Receive
**Task**: Use select to both send and receive in the same statement.

**Expected Output**:
```
Sent: 100
Received: 200
```

### Problem 5: Select with Multiple Cases
**Task**: Use select with multiple cases and handle them all.

**Expected Output**:
```
Case 1: 1
Case 2: 2
Case 3: 3
```

### Problem 6: Select with Channel Direction
**Task**: Use select with send-only and receive-only channels.

**Expected Output**:
```
Sent: Hello
Received: World
```

### Problem 7: Select with Channel Closing
**Task**: Use select to detect when a channel is closed.

**Expected Output**:
```
Received: 42
Channel closed
```

### Problem 8: Select with Range
**Task**: Use select inside a range loop to process multiple values.

**Expected Output**:
```
Received: 1
Received: 2
Received: 3
```

### Problem 9: Select with Goroutine
**Task**: Use select in a goroutine to handle channel operations.

**Expected Output**:
```
Goroutine: Received Hello
```

### Problem 10: Select with Function Call
**Task**: Use select with function calls in cases.

**Expected Output**:
```
Function 1 called
Function 2 called
```

---

## 🟡 MULTIPLE CHANNELS (11-20)

### Problem 11: Select with Three Channels
**Task**: Use select to handle three different channels.

**Expected Output**:
```
Channel 1: 1
Channel 2: 2
Channel 3: 3
```

### Problem 12: Select with Channel Priority
**Task**: Use select to prioritize certain channels over others.

**Expected Output**:
```
High priority: 100
Low priority: 200
```

### Problem 13: Select with Channel Types
**Task**: Use select with different channel types (int, string, bool).

**Expected Output**:
```
Int: 42
String: Hello
Bool: true
```

### Problem 14: Select with Channel Arrays
**Task**: Use select with arrays of channels.

**Expected Output**:
```
Channel 0: 0
Channel 1: 1
Channel 2: 2
```

### Problem 15: Select with Channel Maps
**Task**: Use select with maps of channels.

**Expected Output**:
```
Channel A: A
Channel B: B
Channel C: C
```

### Problem 16: Select with Channel Slices
**Task**: Use select with slices of channels.

**Expected Output**:
```
Channel 0: 0
Channel 1: 1
Channel 2: 2
```

### Problem 17: Select with Channel Structs
**Task**: Use select with channels containing structs.

**Expected Output**:
```
Person: John, Age: 30
Person: Jane, Age: 25
```

### Problem 18: Select with Channel Interfaces
**Task**: Use select with channels containing interfaces.

**Expected Output**:
```
Interface: 42
Interface: Hello
Interface: true
```

### Problem 19: Select with Channel Pointers
**Task**: Use select with channels containing pointers.

**Expected Output**:
```
Pointer: 100
Pointer: 200
```

### Problem 20: Select with Channel Functions
**Task**: Use select with channels containing functions.

**Expected Output**:
```
Function result: 42
Function result: 84
```

---

## 🟠 NON-BLOCKING OPERATIONS (21-30)

### Problem 21: Select with Default
**Task**: Use select with default case for non-blocking operations.

**Expected Output**:
```
No data available
```

### Problem 22: Non-blocking Send
**Task**: Use select with default for non-blocking send.

**Expected Output**:
```
Send would block
```

### Problem 23: Non-blocking Receive
**Task**: Use select with default for non-blocking receive.

**Expected Output**:
```
No data available
```

### Problem 24: Non-blocking Multiple Channels
**Task**: Use select with default to handle multiple channels non-blocking.

**Expected Output**:
```
Channel 1: No data
Channel 2: No data
```

### Problem 25: Non-blocking with Timeout
**Task**: Use select with default and timeout.

**Expected Output**:
```
No data available
```

### Problem 26: Non-blocking with Error Handling
**Task**: Use select with default for error handling.

**Expected Output**:
```
Error: No data available
```

### Problem 27: Non-blocking with Retry
**Task**: Use select with default for retry logic.

**Expected Output**:
```
Retry: 1
Retry: 2
Retry: 3
```

### Problem 28: Non-blocking with Fallback
**Task**: Use select with default for fallback operations.

**Expected Output**:
```
Fallback: Using default value
```

### Problem 29: Non-blocking with Circuit Breaker
**Task**: Use select with default for circuit breaker pattern.

**Expected Output**:
```
Circuit: Open
```

### Problem 30: Non-blocking with Load Balancer
**Task**: Use select with default for load balancing.

**Expected Output**:
```
Server 1: Available
Server 2: Available
```

---

## 🔴 TIMEOUT PATTERNS (31-40)

### Problem 31: Select with Timeout
**Task**: Use select with time.After for timeout handling.

**Expected Output**:
```
Operation timed out
```

### Problem 32: Select with Multiple Timeouts
**Task**: Use select with multiple timeout scenarios.

**Expected Output**:
```
Timeout 1: 100ms
Timeout 2: 200ms
```

### Problem 33: Select with Ticker
**Task**: Use select with time.Ticker for periodic operations.

**Expected Output**:
```
Tick: 1
Tick: 2
Tick: 3
```

### Problem 34: Select with Timer
**Task**: Use select with time.Timer for one-time operations.

**Expected Output**:
```
Timer: 1
Timer: 2
Timer: 3
```

### Problem 35: Select with Context Timeout
**Task**: Use select with context timeout.

**Expected Output**:
```
Context cancelled
```

### Problem 36: Select with Deadline
**Task**: Use select with context deadline.

**Expected Output**:
```
Deadline exceeded
```

### Problem 37: Select with Timeout and Default
**Task**: Use select with both timeout and default.

**Expected Output**:
```
No data available
```

### Problem 38: Select with Timeout and Retry
**Task**: Use select with timeout and retry logic.

**Expected Output**:
```
Retry: 1
Retry: 2
Retry: 3
```

### Problem 39: Select with Timeout and Fallback
**Task**: Use select with timeout and fallback.

**Expected Output**:
```
Fallback: Using default value
```

### Problem 40: Select with Timeout and Error
**Task**: Use select with timeout and error handling.

**Expected Output**:
```
Error: Operation timed out
```

---

## 🟣 ADVANCED PATTERNS (41-50)

### Problem 41: Select with Nil Channels
**Task**: Use select with nil channels (they are ignored).

**Expected Output**:
```
Only non-nil channels are considered
```

### Problem 42: Select with Channel Closing
**Task**: Use select to detect channel closing.

**Expected Output**:
```
Channel closed
```

### Problem 43: Select with Channel State
**Task**: Use select to check channel state.

**Expected Output**:
```
Channel open: true
Channel closed: false
```

### Problem 44: Select with Channel Capacity
**Task**: Use select to check channel capacity.

**Expected Output**:
```
Channel capacity: 5
```

### Problem 45: Select with Channel Length
**Task**: Use select to check channel length.

**Expected Output**:
```
Channel length: 3
```

### Problem 46: Select with Channel Comparison
**Task**: Use select to compare channels.

**Expected Output**:
```
Channels are equal: false
```

### Problem 47: Select with Channel Assignment
**Task**: Use select to assign channels.

**Expected Output**:
```
Channel assigned: true
```

### Problem 48: Select with Channel Range
**Task**: Use select with channel range.

**Expected Output**:
```
Channel range: 0 to 5
```

### Problem 49: Select with Channel Iteration
**Task**: Use select with channel iteration.

**Expected Output**:
```
Iteration: 1
Iteration: 2
Iteration: 3
```

### Problem 50: Select with Channel Composition
**Task**: Use select with channel composition.

**Expected Output**:
```
Composition: 42
```

---

## 🎯 How to Use These Problems

### Run Individual Problems
```bash
go run select_problems.go 1    # Run problem 1
go run select_problems.go 25   # Run problem 25
go run select_problems.go 50   # Run problem 50
```

### Run by Category
```bash
go run select_problems.go basic        # Problems 1-10
go run select_problems.go multiple     # Problems 11-20
go run select_problems.go non-blocking # Problems 21-30
go run select_problems.go timeout      # Problems 31-40
go run select_problems.go advanced     # Problems 41-50
```

### Run All Problems
```bash
go run select_problems.go all
```

---

## 🔧 Testing Your Solutions

```bash
# Compile and test
go build select_problems.go
go run select_problems.go 1

# Run with race detection
go run -race select_problems.go 1

# Check for common mistakes
go vet select_problems.go
```

---

## 💡 Key Concepts to Practice

- **Basic Select**: Single channel operations
- **Multiple Channels**: Multiplexing, prioritization
- **Non-blocking**: Default cases, error handling
- **Timeout**: time.After, time.Tick, context
- **Advanced**: Nil channels, state checking, composition

---

**Remember**: These problems focus on select statement concepts. Master these before moving to more complex concurrency patterns!
