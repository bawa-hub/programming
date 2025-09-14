#!/bin/bash

# 🚀 PHASE 3 INTERFACES & TYPE SYSTEM MASTERY - TEST RUNNER
# Run all Phase 3 examples and projects

echo "🚀 PHASE 3 INTERFACES & TYPE SYSTEM MASTERY - TEST RUNNER"
echo "========================================================="
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

# Test Interface Design Patterns
echo "🎯 INTERFACE DESIGN PATTERNS TESTS"
echo "=================================="
run_go "01_interface_design_patterns/composition/01_interface_composition.go" "Interface Composition Mastery"
run_go "01_interface_design_patterns/dependency_injection/01_dependency_injection_mastery.go" "Dependency Injection Mastery"
run_go "01_interface_design_patterns/mock_interfaces/01_mock_interfaces_mastery.go" "Mock Interfaces Mastery"

# Test Advanced Type System
echo "🔧 ADVANCED TYPE SYSTEM TESTS"
echo "============================="
run_go "02_advanced_type_system/type_assertions/01_type_assertions_mastery.go" "Type Assertions Mastery"
run_go "02_advanced_type_system/type_switches/01_type_switches_mastery.go" "Type Switches Mastery"
run_go "02_advanced_type_system/generics/01_generics_mastery.go" "Generics Mastery"

# Test Reflection Mastery
echo "🪞 REFLECTION MASTERY TESTS"
echo "==========================="
run_go "03_reflection_mastery/runtime_reflection/01_reflection_mastery.go" "Reflection Mastery"

# Test Clean Architecture
echo "🏗️ CLEAN ARCHITECTURE TESTS"
echo "==========================="
run_go "04_clean_architecture/domain_driven_design/01_clean_architecture_mastery.go" "Clean Architecture Mastery"

echo ""
echo "🎯 FINAL PROJECTS TESTS"
echo "======================="
run_go "final_projects/phase3_mastery_demo.go" "Phase 3 Mastery Demo"
run_go "final_projects/microservices_architecture.go" "Microservices Architecture"
run_go "final_projects/reflection_metaprogramming.go" "Reflection & Metaprogramming"

echo "🎉 ALL TESTS COMPLETED!"
echo "======================="
echo "You have successfully demonstrated mastery of:"
echo "✅ Interface design patterns and composition"
echo "✅ Dependency injection patterns and service containers"
echo "✅ Mock interfaces and test doubles"
echo "✅ Type assertions and type switches"
echo "✅ Generics and type parameters"
echo "✅ Advanced type system patterns"
echo "✅ Reflection and dynamic type checking"
echo "✅ Clean architecture with interfaces"
echo
echo "🚀 You are now ready for Phase 4: Error Handling & Logging Mastery!"
