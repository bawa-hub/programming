#!/bin/bash
# Quick Test Script for Select Statement Mastery

echo "🧪 Running Quick Select Test Suite"
echo "=================================="

# Test 1: Basic examples
echo "1. Testing basic examples..."
if go run . basic > /dev/null 2>&1; then
    echo "✅ Basic examples: PASS"
else
    echo "❌ Basic examples: FAIL"
    exit 1
fi

# Test 2: Exercises
echo "2. Testing exercises..."
if go run . exercises > /dev/null 2>&1; then
    echo "✅ Exercises: PASS"
else
    echo "❌ Exercises: FAIL"
    exit 1
fi

# Test 3: Advanced patterns
echo "3. Testing advanced patterns..."
if go run . advanced > /dev/null 2>&1; then
    echo "✅ Advanced patterns: PASS"
else
    echo "❌ Advanced patterns: FAIL"
    exit 1
fi

# Test 4: Compilation
echo "4. Testing compilation..."
if go build . > /dev/null 2>&1; then
    echo "✅ Compilation: PASS"
else
    echo "❌ Compilation: FAIL"
    exit 1
fi

# Test 5: Race detection (should be race-free)
echo "5. Testing race detection..."
if go run -race . basic > /dev/null 2>&1; then
    echo "✅ Race detection: PASS (no races found)"
else
    echo "❌ Race detection: FAIL"
    exit 1
fi

# Test 6: Static analysis
echo "6. Testing static analysis..."
if go vet . > /dev/null 2>&1; then
    echo "✅ Static analysis: PASS"
else
    echo "❌ Static analysis: FAIL"
    exit 1
fi

echo "=================================="
echo "🎉 All tests passed! Select topic is ready!"
echo "🚀 Ready to move to Synchronization Primitives!"
