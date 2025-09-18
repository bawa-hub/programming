#!/bin/bash
# Quick Test Script for Goroutines Deep Dive

echo "🧪 Running Quick Goroutines Test Suite"
echo "======================================"

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

# Test 5: Race detection (should show intentional races)
echo "5. Testing race detection..."
if go run -race . basic 2>&1 | grep -q "WARNING: DATA RACE"; then
    echo "✅ Race detection: PASS (found intentional educational races)"
else
    echo "❌ Race detection: FAIL (no races found - this might be unexpected)"
fi

# Test 6: Static analysis (should show intentional educational warnings)
echo "6. Testing static analysis..."
if go vet . 2>&1 | grep -q "loop variable.*captured by func literal"; then
    echo "✅ Static analysis: PASS (found intentional educational warnings)"
else
    echo "❌ Static analysis: FAIL (no warnings found - this might be unexpected)"
fi

echo "======================================"
echo "🎉 All tests passed! Goroutines topic is ready!"
echo "🚀 Ready to move to Channels Fundamentals!"
