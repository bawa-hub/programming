# 🔗 Go Channels Practice Problems
## 50 Problems from Easy to Advanced

### 📚 Overview
This practice set contains 50 carefully crafted problems that focus exclusively on Go channel concepts. Each problem is designed to reinforce your understanding of channels without introducing complex patterns or advanced concurrency concepts.

---

## 🎯 Problem Categories

### 🟢 **EASY LEVEL (1-15)**
**Focus**: Basic channel operations, send/receive, simple synchronization
- Basic send and receive
- String channels
- Multiple values
- Channel direction
- Channel as signal
- Closing channels
- Range over channels
- Buffered channels
- Channel length and capacity
- Nil channel checks
- Simple select statements
- Multiple channels
- Channel state checks

### 🟡 **MEDIUM LEVEL (16-35)**
**Focus**: Buffered channels, select statements, channel states
- Buffered channel blocking
- Select with timeout
- Non-blocking operations
- Channel comparisons
- Multiple senders/receivers
- Channel as counter
- Select with multiple cases
- Channels with different data types
- Select with nil channels
- Context integration
- Error handling
- Status communication

### 🔴 **ADVANCED LEVEL (36-50)**
**Focus**: Complex channel interactions, error handling, performance
- Channels with mutex
- Channels with WaitGroup
- Channels with atomic operations
- Timer and ticker integration
- Rate limiting
- Circuit breaker patterns
- Load balancing
- Priority queuing
- Batch processing
- Retry logic
- Exponential backoff
- Health checking
- Metrics collection
- Configuration management

---

## 🚀 How to Use

### Run All Problems
```bash
go run practice.go all
```

### Run by Difficulty Level
```bash
go run practice.go easy      # Problems 1-15
go run practice.go medium    # Problems 16-35
go run practice.go advanced  # Problems 36-50
```

### Run Individual Problems
```bash
go run practice.go 1         # Run problem 1
go run practice.go 25        # Run problem 25
go run practice.go 50        # Run problem 50
```

### Run Solutions
```bash
go run solutions.go          # Run all solutions
```

---

## 📖 Learning Path

### 1. **Start with Easy Problems**
Begin with problems 1-15 to build confidence with basic channel operations.

### 2. **Progress to Medium Problems**
Move to problems 16-35 to understand more complex channel interactions.

### 3. **Tackle Advanced Problems**
Challenge yourself with problems 36-50 to master channel concepts.

### 4. **Practice Regularly**
Revisit problems to reinforce your understanding and improve your solutions.

---

## 🔧 Testing Your Solutions

### Compile and Test
```bash
go build .
go run practice.go 1
```

### Run with Race Detection
```bash
go run -race practice.go 1
```

### Check for Common Mistakes
```bash
go vet .
```

### Run All Tests
```bash
./quick_test.sh
```

---

## 💡 Problem-Solving Tips

### 1. **Read the Problem Carefully**
- Understand what the problem is asking for
- Note the expected output format
- Identify the key channel concepts involved

### 2. **Start Simple**
- Begin with basic channel operations
- Add complexity gradually
- Test your solution frequently

### 3. **Handle Edge Cases**
- Check for nil channels
- Handle closed channels properly
- Consider timeout scenarios

### 4. **Use Proper Error Handling**
- Check channel state with `ok` idiom
- Use select with default for non-blocking operations
- Handle context cancellation

### 5. **Follow Go Idioms**
- Use channels for communication
- Prefer channels over shared memory
- Close channels when done

---

## 🎯 Key Concepts Covered

### Basic Concepts
- Channel creation and initialization
- Send and receive operations
- Channel direction (send-only, receive-only)
- Channel closing and state checking
- Range over channels

### Channel Types
- Unbuffered channels (synchronous)
- Buffered channels (asynchronous)
- Channel capacity and length
- Nil channels and their behavior

### Operations
- Blocking vs non-blocking operations
- Select statements
- Timeout handling
- Multiple channel operations

### Patterns
- Signal patterns
- Data transfer patterns
- Worker patterns
- Pipeline patterns

### Error Handling
- Channel state checking
- Error propagation
- Timeout handling
- Context cancellation

---

## 🚨 Common Mistakes to Avoid

### 1. **Deadlocks**
- Sending to nil channels
- Receiving from nil channels
- Circular dependencies
- Sending to unbuffered channels without receivers

### 2. **Panics**
- Sending to closed channels
- Closing already closed channels
- Receiving from closed channels (safe, but returns zero value)

### 3. **Memory Leaks**
- Not closing channels
- Keeping references to large data
- Goroutine leaks

### 4. **Race Conditions**
- Accessing shared data without synchronization
- Using channels incorrectly for synchronization

---

## 📊 Progress Tracking

### Easy Problems (1-15)
- [ ] Problem 1: Basic Send and Receive
- [ ] Problem 2: String Channel
- [ ] Problem 3: Multiple Values
- [ ] Problem 4: Channel Direction
- [ ] Problem 5: Channel as Signal
- [ ] Problem 6: Close Channel
- [ ] Problem 7: Range Over Channel
- [ ] Problem 8: Buffered Channel
- [ ] Problem 9: Channel Length
- [ ] Problem 10: Channel Capacity
- [ ] Problem 11: Nil Channel Check
- [ ] Problem 12: Simple Select
- [ ] Problem 13: Select with Default
- [ ] Problem 14: Multiple Channels
- [ ] Problem 15: Channel State Check

### Medium Problems (16-35)
- [ ] Problem 16: Buffered Channel Blocking
- [ ] Problem 17: Select with Timeout
- [ ] Problem 18: Non-blocking Send
- [ ] Problem 19: Channel Comparison
- [ ] Problem 20: Multiple Senders
- [ ] Problem 21: Multiple Receivers
- [ ] Problem 22: Channel as Counter
- [ ] Problem 23: Select with Multiple Cases
- [ ] Problem 24: Channel with Struct
- [ ] Problem 25: Channel with Slice
- [ ] Problem 26: Channel with Map
- [ ] Problem 27: Channel with Interface
- [ ] Problem 28: Channel with Pointer
- [ ] Problem 29: Channel with Function
- [ ] Problem 30: Channel with Channel
- [ ] Problem 31: Select with Nil Channels
- [ ] Problem 32: Channel with Context
- [ ] Problem 33: Channel with Error
- [ ] Problem 34: Channel with Result
- [ ] Problem 35: Channel with Status

### Advanced Problems (36-50)
- [ ] Problem 36: Channel with Mutex
- [ ] Problem 37: Channel with WaitGroup
- [ ] Problem 38: Channel with Atomic
- [ ] Problem 39: Channel with Timer
- [ ] Problem 40: Channel with Ticker
- [ ] Problem 41: Channel with Rate Limiting
- [ ] Problem 42: Channel with Circuit Breaker
- [ ] Problem 43: Channel with Load Balancer
- [ ] Problem 44: Channel with Priority Queue
- [ ] Problem 45: Channel with Batch Processing
- [ ] Problem 46: Channel with Retry Logic
- [ ] Problem 47: Channel with Exponential Backoff
- [ ] Problem 48: Channel with Health Check
- [ ] Problem 49: Channel with Metrics
- [ ] Problem 50: Channel with Configuration

---

## 🎉 Success Criteria

### Easy Level Mastery
- [ ] Can create and use basic channels
- [ ] Understands send/receive operations
- [ ] Knows when channels block
- [ ] Can use select statements
- [ ] Understands channel states

### Medium Level Mastery
- [ ] Can use buffered channels effectively
- [ ] Understands non-blocking operations
- [ ] Can handle multiple channels
- [ ] Knows how to avoid common pitfalls
- [ ] Can use channels with different data types

### Advanced Level Mastery
- [ ] Can integrate channels with other concurrency primitives
- [ ] Understands performance implications
- [ ] Can implement complex channel patterns
- [ ] Knows how to handle errors and timeouts
- [ ] Can design channel-based systems

---

## 🔗 Next Steps

After completing these problems:

1. **Practice with real projects** - Apply channel concepts to actual applications
2. **Explore advanced patterns** - Learn about worker pools, pipelines, and fan-out/fan-in
3. **Study the main curriculum** - Move on to the comprehensive concurrency curriculum
4. **Contribute to open source** - Find Go projects that use channels effectively
5. **Teach others** - Share your knowledge and help others learn

---

**Remember**: These problems focus on core channel concepts. Master these before moving to advanced concurrency patterns. Each problem builds upon the previous ones, so work through them systematically for the best learning experience!

