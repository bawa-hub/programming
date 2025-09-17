#!/bin/bash

# Advanced Scheduling Quick Test Script
echo "⚙️ Testing Advanced Scheduling Examples"
echo "======================================"

# Test 1: Compilation
echo "1. Testing compilation..."
if go build -o scheduling_test .; then
    echo "   ✅ Compilation successful"
    rm -f scheduling_test
else
    echo "   ❌ Compilation failed"
    exit 1
fi

# Test 2: Static analysis
echo "2. Running static analysis..."
if go vet .; then
    echo "   ✅ Static analysis passed"
else
    echo "   ❌ Static analysis failed"
    exit 1
fi

# Test 3: Basic examples
echo "3. Running basic examples..."
if go run . basic > /dev/null 2>&1; then
    echo "   ✅ Basic examples passed"
else
    echo "   ❌ Basic examples failed"
    exit 1
fi

# Test 4: Exercises
echo "4. Running exercises..."
if go run . exercises > /dev/null 2>&1; then
    echo "   ✅ Exercises passed"
else
    echo "   ❌ Exercises failed"
    exit 1
fi

# Test 5: Advanced patterns
echo "5. Running advanced patterns..."
if go run . advanced > /dev/null 2>&1; then
    echo "   ✅ Advanced patterns passed"
else
    echo "   ❌ Advanced patterns failed"
    exit 1
fi

# Test 6: Race detection
echo "6. Running race detection..."
if go run -race . basic > /dev/null 2>&1; then
    echo "   ✅ Race detection passed"
else
    echo "   ❌ Race detection failed"
    exit 1
fi

# Test 7: All examples
echo "7. Running all examples..."
if go run . all > /dev/null 2>&1; then
    echo "   ✅ All examples passed"
else
    echo "   ❌ All examples failed"
    exit 1
fi

echo ""
echo "🎉 All tests passed! Advanced Scheduling is ready!"
echo "Ready to move to the next topic!"

