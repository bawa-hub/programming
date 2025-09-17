# 🚀 Channels Fundamentals - Command Reference

## 📋 Quick Reference Commands

### **Basic Testing**
```bash
# Test basic examples
go run . basic

# Test all exercises  
go run . exercises

# Test advanced patterns
go run . advanced

# Test everything together
go run . all

# Show help/usage
go run .
```

### **Advanced Testing**
```bash
# Race detection (should be race-free)
go run -race . basic

# Race detection on exercises (should be race-free)
go run -race . exercises

# Race detection on advanced patterns (should be race-free)
go run -race . advanced

# Compilation test
go build .

# Static analysis
go vet .
```

### **Automated Testing**
```bash
# Run the quick test suite
./quick_test.sh

# Make script executable (if needed)
chmod +x quick_test.sh
```

### **Performance Testing**
```bash
# Run with verbose output
go run -v . basic

# Run with trace (creates trace.out)
go run . basic -trace=trace.out

# Run with memory profile
go run . basic -memprofile=mem.prof

# Analyze trace
go tool trace trace.out

# Analyze memory profile
go tool pprof mem.prof
```

### **Individual Testing**
```bash
# Test specific functions (create test file)
cat > test.go << 'EOF'
package main
import "time"
func main() {
    basicChannelOperations()
    time.Sleep(100 * time.Millisecond)
}
EOF
go run test.go
rm test.go
```

## 🎯 Expected Results

| Command | Expected Result |
|---------|----------------|
| `go run . basic` | 11 examples with channel operations |
| `go run . exercises` | 10 exercises with proper channel usage |
| `go run . advanced` | 7 advanced patterns working |
| `go run -race . basic` | No race conditions (clean) |
| `go run -race . exercises` | No race conditions (clean) |
| `go run -race . advanced` | No race conditions (clean) |
| `go vet .` | Passes static analysis |
| `go build .` | Compiles successfully |

## 🏆 Success Indicators

✅ **All commands run without errors**  
✅ **Race detection shows no race conditions**  
✅ **Exercises demonstrate proper channel usage**  
✅ **Advanced patterns show production quality**  
✅ **Static analysis passes**  
✅ **Code compiles and builds successfully**

## 🚀 Ready for Next Topic?

Once all tests pass, you're ready for **Level 1, Topic 3: Select Statement Mastery**!

**Type "NEXT" to continue your journey to Go concurrency mastery!**
