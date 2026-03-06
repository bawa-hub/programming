# Makefile for String Algorithms

CXX = g++
CXXFLAGS = -std=c++17 -Wall -Wextra -O2

# All source files
SOURCES = 01_basic_operations.cpp \
          02_pattern_matching.cpp \
          03_sliding_window.cpp \
          04_two_pointers.cpp \
          05_hash_techniques.cpp \
          06_advanced_algorithms.cpp \
          07_common_patterns.cpp \
          08_practice_problems.cpp

# Executable names
EXECUTABLES = $(SOURCES:.cpp=)

# Default target
all: $(EXECUTABLES)

# Build each file separately
%: %.cpp
	$(CXX) $(CXXFLAGS) -o $@ $<

# Clean compiled files
clean:
	rm -f $(EXECUTABLES)

# Run all tests
test: all
	@echo "Running all tests..."
	@for exe in $(EXECUTABLES); do \
		echo "\n=== Running $$exe ==="; \
		./$$exe; \
	done

.PHONY: all clean test
