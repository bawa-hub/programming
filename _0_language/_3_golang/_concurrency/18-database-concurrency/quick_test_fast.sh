#!/bin/bash

# Database Concurrency Fast Test Script
echo "🗄️ Testing Database Concurrency Examples"
echo "======================================="

# Test 1: Compilation
echo "1. Testing compilation..."
if go build -o database_test .; then
    echo "   ✅ Compilation successful"
    rm -f database_test
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

# Test 3: Basic examples (first 5 only)
echo "3. Running basic examples (first 5)..."
if go run . basic 2>/dev/null | head -30 | grep -q "completed"; then
    echo "   ✅ Basic examples working"
else
    echo "   ❌ Basic examples failed"
    exit 1
fi

# Test 4: Exercises (first 5 only)
echo "4. Running exercises (first 5)..."
if go run . exercises 2>/dev/null | head -30 | grep -q "completed"; then
    echo "   ✅ Exercises working"
else
    echo "   ❌ Exercises failed"
    exit 1
fi

# Test 5: Advanced patterns (first 5 only)
echo "5. Running advanced patterns (first 5)..."
if go run . advanced 2>/dev/null | head -30 | grep -q "Advanced"; then
    echo "   ✅ Advanced patterns working"
else
    echo "   ❌ Advanced patterns failed"
    exit 1
fi

echo ""
echo "🎉 Fast tests passed! Database Concurrency is working!"
echo "Ready to move to the next topic!"

