# 🔗 Basic Channel Synchronization Problems
## 20 Problems for Channel Practice

### 📚 Problem Categories
- **Alternating Patterns** (1-5): Even/odd, A/B, 1/2/3 sequences
- **Turn-Based Operations** (6-10): Taking turns, round-robin, ordered execution
- **Signal Coordination** (11-15): Start/stop signals, handshakes, barriers
- **Data Exchange** (16-20): Ping-pong, data passing, request-response

---

## 🔄 ALTERNATING PATTERNS (1-5)

### Problem 1: Even-Odd Printing
**Task**: Print even and odd numbers in sequence using two goroutines.

**Expected Output**:
```
Even: 0
Odd: 1
Even: 2
Odd: 3
Even: 4
Odd: 5
```

### Problem 2: A-B Alternating
**Task**: Print "A" and "B" alternately using two goroutines.

**Expected Output**:
```
A
B
A
B
A
B
```

### Problem 3: 1-2-3 Sequence
**Task**: Print 1, 2, 3 in sequence using three goroutines.

**Expected Output**:
```
1
2
3
1
2
3
```

### Problem 4: Color Alternating
**Task**: Print "Red" and "Blue" alternately using two goroutines.

**Expected Output**:
```
Red
Blue
Red
Blue
Red
Blue
```

### Problem 5: Number-Letter Alternating
**Task**: Print numbers and letters alternately (1, A, 2, B, 3, C...).

**Expected Output**:
```
1
A
2
B
3
C
```

---

## 🔄 TURN-BASED OPERATIONS (6-10)

### Problem 6: Three-Player Game
**Task**: Three players take turns in order (Player 1, Player 2, Player 3).

**Expected Output**:
```
Player 1's turn
Player 2's turn
Player 3's turn
Player 1's turn
Player 2's turn
Player 3's turn
```

### Problem 7: Round-Robin Processing
**Task**: Process items in round-robin fashion using three workers.

**Expected Output**:
```
Worker 1: Item 1
Worker 2: Item 2
Worker 3: Item 3
Worker 1: Item 4
Worker 2: Item 5
Worker 3: Item 6
```

### Problem 8: Ordered Execution
**Task**: Execute three functions in order using channels.

**Expected Output**:
```
Function 1 executed
Function 2 executed
Function 3 executed
```

### Problem 9: Sequential Steps
**Task**: Execute steps A, B, C in sequence using goroutines.

**Expected Output**:
```
Step A
Step B
Step C
```

### Problem 10: Turn-Based Counter
**Task**: Two counters take turns incrementing a shared value.

**Expected Output**:
```
Counter 1: 1
Counter 2: 2
Counter 1: 3
Counter 2: 4
Counter 1: 5
Counter 2: 6
```

---

## 📡 SIGNAL COORDINATION (11-15)

### Problem 11: Start Signal
**Task**: Wait for a start signal before beginning work.

**Expected Output**:
```
Waiting for start signal...
Start signal received!
Work started
Work completed
```

### Problem 12: Handshake Protocol
**Task**: Implement a handshake between two goroutines.

**Expected Output**:
```
Goroutine 1: Sending handshake
Goroutine 2: Received handshake
Goroutine 2: Sending response
Goroutine 1: Received response
Handshake complete
```

### Problem 13: Barrier Synchronization
**Task**: Wait for all goroutines to reach a barrier before proceeding.

**Expected Output**:
```
Goroutine 1: Reached barrier
Goroutine 2: Reached barrier
Goroutine 3: Reached barrier
All goroutines reached barrier
Proceeding...
```

### Problem 14: Stop Signal
**Task**: Stop all goroutines when a stop signal is received.

**Expected Output**:
```
Goroutine 1: Working...
Goroutine 2: Working...
Stop signal received
Goroutine 1: Stopping
Goroutine 2: Stopping
```

### Problem 15: Ready Signal
**Task**: Wait for all goroutines to be ready before starting.

**Expected Output**:
```
Goroutine 1: Ready
Goroutine 2: Ready
Goroutine 3: Ready
All ready! Starting...
```

---

## 🔄 DATA EXCHANGE (16-20)

### Problem 16: Ping-Pong
**Task**: Two goroutines exchange "ping" and "pong" messages.

**Expected Output**:
```
Ping
Pong
Ping
Pong
Ping
Pong
```

### Problem 17: Data Relay
**Task**: Pass data through a chain of goroutines.

**Expected Output**:
```
Stage 1: Processing data
Stage 2: Processing data
Stage 3: Processing data
Data relay complete
```

### Problem 18: Request-Response
**Task**: Implement request-response pattern between two goroutines.

**Expected Output**:
```
Request: Hello
Response: Hello World
Request: How are you?
Response: I'm fine, thank you
```

### Problem 19: Data Pipeline
**Task**: Process data through a pipeline of stages.

**Expected Output**:
```
Stage 1: Input data
Stage 2: Processing data
Stage 3: Output data
Pipeline complete
```

### Problem 20: Message Passing
**Task**: Pass messages between multiple goroutines in order.

**Expected Output**:
```
Message 1: Hello
Message 2: World
Message 3: Go
Message 4: Channels
```

---

## 🎯 How to Use These Problems

### 1. **Start with Alternating Patterns**
Begin with problems 1-5 to understand basic synchronization.

### 2. **Move to Turn-Based Operations**
Practice problems 6-10 for more complex coordination.

### 3. **Learn Signal Coordination**
Tackle problems 11-15 for advanced synchronization.

### 4. **Master Data Exchange**
Complete problems 16-20 for real-world patterns.

### 5. **Test Your Solutions**
```bash
go run basic_sync_problems.go 1    # Run problem 1
go run basic_sync_problems.go all  # Run all problems
```

---

## 💡 Key Concepts to Practice

- **Channel Communication**: Sending and receiving data
- **Synchronization**: Coordinating goroutines
- **Turn-Taking**: Alternating execution
- **Signaling**: Start/stop/ready signals
- **Data Flow**: Passing data between goroutines
- **Ordering**: Ensuring proper sequence
- **Coordination**: Multiple goroutines working together

---

## 🔧 Testing Your Solutions

```bash
# Compile and test
go build basic_sync_problems.go
go run basic_sync_problems.go 1

# Run with race detection
go run -race basic_sync_problems.go 1

# Check for common mistakes
go vet basic_sync_problems.go
```

---

**Remember**: These problems focus on basic channel synchronization. Master these before moving to more complex patterns!
