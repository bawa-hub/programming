# 🔗 Go Channels Practice Problems
## 50 Problems from Easy to Advanced

### 📚 Problem Categories
- **Easy (1-15)**: Basic channel operations, send/receive, simple synchronization
- **Medium (16-35)**: Buffered channels, select statements, channel states
- **Advanced (36-50)**: Complex channel interactions, error handling, performance

---

## 🟢 EASY LEVEL (1-15)

### Problem 1: Basic Send and Receive
**Task**: Create a channel, send the number 42, and receive it in another goroutine.

**Expected Output**:
```
Sent: 42
Received: 42
```

### Problem 2: String Channel
**Task**: Create a string channel, send "Hello World", and receive it.

**Expected Output**:
```
Message: Hello World
```

### Problem 3: Multiple Values
**Task**: Send three numbers (1, 2, 3) through a channel and receive them.

**Expected Output**:
```
Received: 1
Received: 2
Received: 3
```

### Problem 4: Channel Direction
**Task**: Create a function that only sends data and another that only receives data.

**Expected Output**:
```
Sender sent: 100
Receiver got: 100
```

### Problem 5: Channel as Signal
**Task**: Use a channel to signal when work is complete.

**Expected Output**:
```
Work started
Work completed
Signal received
```

### Problem 6: Close Channel
**Task**: Send data, close the channel, and try to receive from closed channel.

**Expected Output**:
```
Received: 42
Channel closed
Received from closed: 0, ok: false
```

### Problem 7: Range Over Channel
**Task**: Send multiple values and use range to receive them all.

**Expected Output**:
```
Received: 1
Received: 2
Received: 3
Done receiving
```

### Problem 8: Buffered Channel
**Task**: Create a buffered channel with capacity 3, send 3 values without blocking.

**Expected Output**:
```
Sent: 1
Sent: 2
Sent: 3
All sent without blocking
```

### Problem 9: Channel Length
**Task**: Check the length of a buffered channel after sending data.

**Expected Output**:
```
Channel length: 0
After sending 2 items: 2
After sending 5 items: 5
```

### Problem 10: Channel Capacity
**Task**: Create channels with different capacities and check their capacity.

**Expected Output**:
```
Unbuffered capacity: 0
Buffered capacity: 5
```

### Problem 11: Nil Channel Check
**Task**: Check if a channel is nil and demonstrate nil channel behavior.

**Expected Output**:
```
Channel is nil: true
Channel is nil: false
```

### Problem 12: Simple Select
**Task**: Use select to receive from a channel with a timeout.

**Expected Output**:
```
Received: 42
```

### Problem 13: Select with Default
**Task**: Use select with default case for non-blocking receive.

**Expected Output**:
```
No data available
```

### Problem 14: Multiple Channels
**Task**: Create two channels and use select to receive from either.

**Expected Output**:
```
Received from ch1: 100
Received from ch2: 200
```

### Problem 15: Channel State Check
**Task**: Check if a channel is open or closed using the ok idiom.

**Expected Output**:
```
Channel open: true
Channel closed: false
```

---

## 🟡 MEDIUM LEVEL (16-35)

### Problem 16: Buffered Channel Blocking
**Task**: Demonstrate when a buffered channel blocks (when full).

**Expected Output**:
```
Buffer full, send would block
```

### Problem 17: Select with Timeout
**Task**: Use select with time.After for timeout handling.

**Expected Output**:
```
Operation timed out
```

### Problem 18: Non-blocking Send
**Task**: Use select with default for non-blocking send.

**Expected Output**:
```
Send would block
```

### Problem 19: Channel Comparison
**Task**: Compare behavior of unbuffered vs buffered channels.

**Expected Output**:
```
Unbuffered: Synchronous
Buffered: Asynchronous
```

### Problem 20: Multiple Senders
**Task**: Have multiple goroutines send to the same channel.

**Expected Output**:
```
Sender 1: 1
Sender 2: 2
Sender 3: 3
```

### Problem 21: Multiple Receivers
**Task**: Have multiple goroutines receive from the same channel.

**Expected Output**:
```
Receiver 1: 42
Receiver 2: 42
```

### Problem 22: Channel as Counter
**Task**: Use a channel to count from 1 to 5.

**Expected Output**:
```
Count: 1
Count: 2
Count: 3
Count: 4
Count: 5
```

### Problem 23: Select with Multiple Cases
**Task**: Use select with multiple cases and handle them all.

**Expected Output**:
```
Case 1: 100
Case 2: 200
Case 3: 300
```

### Problem 24: Channel with Struct
**Task**: Send and receive a struct through a channel.

**Expected Output**:
```
Person: John, Age: 30
```

### Problem 25: Channel with Slice
**Task**: Send and receive a slice through a channel.

**Expected Output**:
```
Received slice: [1 2 3 4 5]
```

### Problem 26: Channel with Map
**Task**: Send and receive a map through a channel.

**Expected Output**:
```
Received map: map[hello:world]
```

### Problem 27: Channel with Interface
**Task**: Send different types through an interface channel.

**Expected Output**:
```
Received: 42
Received: hello
Received: true
```

### Problem 28: Channel with Pointer
**Task**: Send and receive a pointer through a channel.

**Expected Output**:
```
Value: 100
```

### Problem 29: Channel with Function
**Task**: Send and receive a function through a channel.

**Expected Output**:
```
Function result: 42
```

### Problem 30: Channel with Channel
**Task**: Send a channel through another channel.

**Expected Output**:
```
Received channel, sending data
Data sent through received channel
```

### Problem 31: Select with Nil Channels
**Task**: Use select with nil channels (they are ignored).

**Expected Output**:
```
Only non-nil channels are considered
```

### Problem 32: Channel with Context
**Task**: Use context with channels for cancellation.

**Expected Output**:
```
Operation cancelled
```

### Problem 33: Channel with Error
**Task**: Send and receive errors through a channel.

**Expected Output**:
```
Error: something went wrong
```

### Problem 34: Channel with Result
**Task**: Send both data and error through a channel.

**Expected Output**:
```
Result: 42, Error: <nil>
```

### Problem 35: Channel with Status
**Task**: Use a channel to communicate status updates.

**Expected Output**:
```
Status: started
Status: processing
Status: completed
```

---

## 🔴 ADVANCED LEVEL (36-50)

### Problem 36: Channel with Mutex
**Task**: Use channels with mutex for complex synchronization.

**Expected Output**:
```
Counter: 1000
```

### Problem 37: Channel with WaitGroup
**Task**: Use channels with WaitGroup for goroutine coordination.

**Expected Output**:
```
All goroutines completed
```

### Problem 38: Channel with Atomic
**Task**: Use channels with atomic operations for counters.

**Expected Output**:
```
Atomic counter: 1000
```

### Problem 39: Channel with Timer
**Task**: Use channels with timers for periodic operations.

**Expected Output**:
```
Tick: 1
Tick: 2
Tick: 3
```

### Problem 40: Channel with Ticker
**Task**: Use channels with tickers for continuous operations.

**Expected Output**:
```
Ticker: 1
Ticker: 2
Ticker: 3
```

### Problem 41: Channel with Rate Limiting
**Task**: Implement rate limiting using channels.

**Expected Output**:
```
Request 1: allowed
Request 2: allowed
Request 3: rate limited
```

### Problem 42: Channel with Circuit Breaker
**Task**: Implement a simple circuit breaker using channels.

**Expected Output**:
```
Circuit: closed
Circuit: open
Circuit: half-open
```

### Problem 43: Channel with Load Balancer
**Task**: Implement load balancing using channels.

**Expected Output**:
```
Request 1: Server 1
Request 2: Server 2
Request 3: Server 1
```

### Problem 44: Channel with Priority Queue
**Task**: Implement priority queuing using channels.

**Expected Output**:
```
High priority: 1
High priority: 2
Low priority: 3
```

### Problem 45: Channel with Batch Processing
**Task**: Implement batch processing using channels.

**Expected Output**:
```
Batch: [1 2 3]
Batch: [4 5 6]
```

### Problem 46: Channel with Retry Logic
**Task**: Implement retry logic using channels.

**Expected Output**:
```
Attempt 1: failed
Attempt 2: failed
Attempt 3: success
```

### Problem 47: Channel with Exponential Backoff
**Task**: Implement exponential backoff using channels.

**Expected Output**:
```
Retry after: 1s
Retry after: 2s
Retry after: 4s
```

### Problem 48: Channel with Health Check
**Task**: Implement health checking using channels.

**Expected Output**:
```
Service: healthy
Service: unhealthy
Service: healthy
```

### Problem 49: Channel with Metrics
**Task**: Implement metrics collection using channels.

**Expected Output**:
```
Metrics: requests=100, errors=5
```

### Problem 50: Channel with Configuration
**Task**: Implement configuration management using channels.

**Expected Output**:
```
Config updated: timeout=5s
Config updated: retries=3
```

---

## 🎯 How to Use These Problems

1. **Start with Easy problems** to build confidence
2. **Work through Medium problems** to understand complexity
3. **Tackle Advanced problems** to master channels
4. **Test your solutions** with `go run -race`
5. **Use `go vet`** to check for common mistakes

## 🔧 Testing Your Solutions

```bash
# Compile and test
go build .

# Run with race detection
go run -race .

# Check for common mistakes
go vet .
```

## 💡 Tips for Success

- **Read the problem carefully** before coding
- **Start with simple solutions** and improve them
- **Test edge cases** (nil channels, closed channels)
- **Use proper error handling**
- **Follow Go idioms** and best practices
- **Comment your code** to explain your approach

---

**Remember**: These problems focus on core channel concepts. Master these before moving to advanced concurrency patterns!

