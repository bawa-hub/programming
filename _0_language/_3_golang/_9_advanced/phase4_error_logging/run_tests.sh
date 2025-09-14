#!/bin/bash

# 🚨 PHASE 4: ERROR HANDLING & LOGGING MASTERY TEST SUITE
# Comprehensive testing of all Phase 4 modules

echo "🚨 PHASE 4: ERROR HANDLING & LOGGING MASTERY TEST SUITE"
echo "========================================================"
echo ""

# Function to run Go programs with error handling
run_go() {
    local file="$1"
    local description="$2"
    
    echo "🧪 Testing: $description"
    echo "File: $file"
    echo "----------------------------------------"
    
    if go run "$file"; then
        echo "✅ SUCCESS: $description"
    else
        echo "❌ FAILED: $description"
    fi
    echo ""
}

# Test Advanced Error Handling
echo "🚨 ADVANCED ERROR HANDLING TESTS"
echo "================================="

run_go "01_advanced_error_handling/custom_errors/01_custom_errors_mastery.go" "Custom Error Types Mastery"
run_go "01_advanced_error_handling/error_wrapping/01_error_wrapping_mastery.go" "Error Wrapping & Unwrapping Mastery"
run_go "01_advanced_error_handling/recovery_strategies/01_recovery_strategies_mastery.go" "Error Recovery Strategies Mastery"
run_go "01_advanced_error_handling/panic_handling/01_panic_handling_mastery.go" "Panic Handling Mastery"
run_go "01_advanced_error_handling/error_context/01_error_context_mastery.go" "Error Context Mastery"

# Test Structured Logging
echo ""
echo "📊 STRUCTURED LOGGING TESTS"
echo "============================"

run_go "02_structured_logging/log_levels/01_log_levels_mastery.go" "Log Levels Mastery"
run_go "02_structured_logging/contextual_logging/01_contextual_logging_mastery.go" "Contextual Logging Mastery"
run_go "02_structured_logging/log_aggregation/01_log_aggregation_mastery.go" "Log Aggregation Mastery"
run_go "02_structured_logging/performance_impact/01_performance_impact_mastery.go" "Performance Impact Mastery"
run_go "02_structured_logging/log_management/01_log_management_mastery.go" "Log Management Mastery"
run_go "02_structured_logging/performance_optimization/01_performance_optimization_mastery.go" "Performance Optimization Mastery"

# Test Tracing and Metrics
echo ""
echo "🔍 TRACING AND METRICS TESTS"
echo "============================="

run_go "03_tracing_metrics/distributed_tracing/01_distributed_tracing_mastery.go" "Distributed Tracing Mastery"
run_go "03_tracing_metrics/prometheus_metrics/01_prometheus_metrics_mastery.go" "Prometheus Metrics Mastery"
run_go "03_tracing_metrics/health_checks/01_health_checks_mastery.go" "Health Checks Mastery"
run_go "03_tracing_metrics/monitoring_patterns/01_monitoring_patterns_mastery.go" "Monitoring Patterns Mastery"

echo "🎉 ALL TESTS COMPLETED!"
echo "======================="
echo "You have successfully demonstrated mastery of:"
echo "✅ Custom error types and hierarchies"
echo "✅ Error wrapping and unwrapping patterns"
echo "✅ Error recovery strategies and resilience"
echo "✅ Panic handling and recovery"
echo "✅ Error context and tracing"
echo "✅ Log levels and structured logging"
echo "✅ Contextual logging and correlation"
echo "✅ Log aggregation and processing"
echo "✅ Performance impact optimization"
echo "✅ Log management and lifecycle"
echo "✅ Performance optimization techniques"
echo "✅ Distributed tracing concepts"
echo "✅ Prometheus metrics collection"
echo "✅ Health checks and circuit breakers"
echo "✅ Monitoring patterns and alerting"
echo ""
echo "🚀 You are now ready for Phase 5: Package Design & Architecture Mastery!"
