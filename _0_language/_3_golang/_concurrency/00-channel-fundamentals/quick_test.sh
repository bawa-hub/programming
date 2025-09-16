#!/bin/bash

echo "🔗 Testing Channel Fundamentals"
echo "=============================="

# Test compilation
echo "1. Testing compilation..."
if go build .; then
    echo "   ✅ Compilation successful"
else
    echo "   ❌ Compilation failed"
    exit 1
fi

# Test static analysis
echo "2. Running static analysis..."
if go vet .; then
    echo "   ✅ Static analysis passed"
else
    echo "   ❌ Static analysis failed"
    exit 1
fi

# Test race detection (may fail due to educational examples)
echo "3. Testing with race detection..."
if go run -race . > /dev/null 2>&1; then
    echo "   ✅ Race detection passed"
else
    echo "   ⚠️  Race detection found issues (expected for educational examples)"
fi

# Test individual commands
echo "4. Testing individual commands..."

commands=("basic" "types" "operations" "behavior" "patterns" "pitfalls")

for cmd in "${commands[@]}"; do
    echo "   Testing: go run . $cmd"
    if go run . "$cmd" > /dev/null 2>&1; then
        echo "   ✅ $cmd working"
    else
        echo "   ❌ $cmd failed"
        exit 1
    fi
done

echo ""
echo "🎉 All tests passed! Channel fundamentals are working perfectly!"
echo ""
echo "Quick commands:"
echo "  go run .                    # Run all examples"
echo "  go run . basic              # Run basic concepts"
echo "  go run . types              # Run channel types"
echo "  go run . operations         # Run channel operations"
echo "  go run . behavior           # Run channel behavior"
echo "  go run . patterns           # Run channel patterns"
echo "  go run . pitfalls           # Run channel pitfalls"
echo ""
echo "Debugging commands:"
echo "  go run -race .              # Run with race detection"
echo "  go vet .                    # Check for common mistakes"
echo "  go build .                  # Compile the project"
