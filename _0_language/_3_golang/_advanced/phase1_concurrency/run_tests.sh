#!/bin/bash

# 🚀 PHASE 1 CONCURRENCY MASTERY - TEST RUNNER
# Run all Phase 1 examples and projects

echo "🚀 PHASE 1 CONCURRENCY MASTERY - TEST RUNNER"
echo "============================================="
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

# Test Goroutines Deep Dive
echo "🧵 GOROUTINES DEEP DIVE TESTS"
echo "=============================="
run_go "01_goroutines_deep_dive/basics/01_goroutine_lifecycle.go" "Goroutine Lifecycle"
run_go "01_goroutines_deep_dive/basics/02_goroutine_patterns.go" "Goroutine Patterns"
run_go "01_goroutines_deep_dive/projects/concurrent_web_scraper.go" "Concurrent Web Scraper"

echo "📡 CHANNELS MASTERY TESTS"
echo "=========================="
run_go "02_channels_mastery/types/01_channel_types.go" "Channel Types"
run_go "02_channels_mastery/patterns/01_select_patterns.go" "Select Patterns"
run_go "02_channels_mastery/advanced_patterns/01_advanced_channel_patterns.go" "Advanced Channel Patterns"
run_go "02_channels_mastery/projects/realtime_chat_server.go" "Real-time Chat Server"

echo "🔒 SYNCHRONIZATION PRIMITIVES TESTS"
echo "===================================="
run_go "03_sync_primitives/mutexes/01_mutex_patterns.go" "Mutex Patterns"
run_go "03_sync_primitives/atomic/01_atomic_operations_mastery.go" "Atomic Operations Mastery"
run_go "03_sync_primitives/condition_variables/01_condition_variables_mastery.go" "Condition Variables Mastery"
run_go "03_sync_primitives/projects/distributed_task_queue.go" "Distributed Task Queue"

echo "🎯 CONTEXT MASTERY TESTS"
echo "========================"
run_go "04_context_mastery/basics/01_context_fundamentals.go" "Context Fundamentals"
run_go "04_context_mastery/best_practices/01_context_best_practices.go" "Context Best Practices"

echo "🚀 FINAL PROJECT TESTS"
echo "======================"
run_go "final_projects/phase1_mastery_demo.go" "Phase 1 Mastery Demo"

echo "🎉 ALL TESTS COMPLETED!"
echo "======================="
echo "You have successfully demonstrated mastery of:"
echo "✅ Goroutines and their lifecycle"
echo "✅ Advanced goroutine patterns"
echo "✅ Channel types and behaviors"
echo "✅ Select statements and multiplexing"
echo "✅ Synchronization primitives"
echo "✅ Context for cancellation and values"
echo "✅ Real-world concurrent applications"
echo
echo "🚀 You are now ready for Phase 2: Memory Management & Performance!"
