cmake_minimum_required(VERSION 3.16)
project(CppLearning)

# Set C++ standard
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Set compiler flags
if(MSVC)
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} /W4")
else()
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wall -Wextra -Wpedantic")
endif()

# Debug flags
set(CMAKE_CXX_FLAGS_DEBUG "-g -O0")
set(CMAKE_CXX_FLAGS_RELEASE "-O3 -DNDEBUG")

# Create output directory
file(MAKE_DIRECTORY ${CMAKE_BINARY_DIR}/bin)

# Exercise executables
add_executable(exercise_01 exercises/01_variables_types_memory.cpp)
set_target_properties(exercise_01 PROPERTIES
    RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/bin
)

add_executable(exercise_02 exercises/02_pointers_references.cpp)
set_target_properties(exercise_02 PROPERTIES
    RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/bin
)

# Add more exercises as they are created
# add_executable(exercise_03 exercises/03_functions_scope.cpp)

# Optional: Create a target that runs all exercises
add_custom_target(run_all_exercises
    COMMAND ${CMAKE_BINARY_DIR}/bin/exercise_01
    COMMAND ${CMAKE_BINARY_DIR}/bin/exercise_02
    DEPENDS exercise_01 exercise_02
    COMMENT "Running all C++ exercises"
)

# Print configuration summary
message(STATUS "C++ Learning Project Configuration:")
message(STATUS "  C++ Standard: ${CMAKE_CXX_STANDARD}")
message(STATUS "  Build Type: ${CMAKE_BUILD_TYPE}")
message(STATUS "  Compiler: ${CMAKE_CXX_COMPILER_ID}")
message(STATUS "  Output Directory: ${CMAKE_BINARY_DIR}/bin")
