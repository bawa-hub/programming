#!/bin/bash

# 🚀 PHASE 2 MEMORY & PERFORMANCE MASTERY - TEST RUNNER
# Run all Phase 2 examples and projects

echo "🚀 PHASE 2 MEMORY & PERFORMANCE MASTERY - TEST RUNNER"
echo "====================================================="
echo

# Set working directory
cd "$(dirname "$0")"

# Function to run a Go program
run_go() {
    local file=$1
    local description=$2
    
    echo "🧪 Testing: $description"
    echo "File: $file"
    echo "----------------------------------------"
    
    if go run "$file"; then
        echo "✅ SUCCESS: $description"
    else
        echo "❌ FAILED: $description"
    fi
    echo
}

# Test Memory Model Deep Dive
echo "🧠 MEMORY MODEL DEEP DIVE TESTS"
echo "================================"
run_go "01_memory_model_deep_dive/gc_algorithms/01_garbage_collection_mastery.go" "Garbage Collection Mastery"
run_go "01_memory_model_deep_dive/memory_optimization/01_memory_optimization_mastery.go" "Memory Optimization Mastery"

# Test Performance Optimization
echo "⚡ PERFORMANCE OPTIMIZATION TESTS"
echo "================================="
run_go "02_performance_optimization/profiling/01_profiling_mastery.go" "Profiling Mastery"
run_go "02_performance_optimization/advanced_techniques/01_advanced_optimization_mastery.go" "Advanced Optimization Techniques"

# Test Data Structure Optimization
echo "📊 DATA STRUCTURE OPTIMIZATION TESTS"
echo "===================================="
run_go "03_data_structure_optimization/custom_structures/01_high_performance_structures.go" "High Performance Structures"

# Test Performance Monitoring
echo "📊 PERFORMANCE MONITORING TESTS"
echo "================================"
run_go "03_performance_monitoring/01_performance_monitoring_mastery.go" "Performance Monitoring Mastery"

# Test Advanced Tools
echo "🔧 ADVANCED TOOLS TESTS"
echo "========================"
run_go "04_advanced_tools/01_advanced_tools_mastery.go" "Advanced Tools Mastery"

# Test Final Projects
echo "🚀 FINAL PROJECT TESTS"
echo "======================"
run_go "final_projects/high_performance_cache.go" "High Performance Cache"

echo "🎉 ALL TESTS COMPLETED!"
echo "======================="
echo "You have successfully demonstrated mastery of:"
echo "✅ Go memory model and garbage collection"
echo "✅ Memory optimization techniques (stack vs heap, escape analysis, pooling)"
echo "✅ Performance profiling and benchmarking"
echo "✅ Advanced optimization techniques (SIMD, cache-friendly structures, branch prediction)"
echo "✅ High-performance data structures"
echo "✅ Performance monitoring and alerting systems"
echo "✅ Advanced debugging and analysis tools (race detector, memory sanitizer, pprof)"
echo "✅ Memory pooling and optimization techniques"
echo
echo "🚀 You are now ready for Phase 3: Interfaces & Type System Mastery!"
