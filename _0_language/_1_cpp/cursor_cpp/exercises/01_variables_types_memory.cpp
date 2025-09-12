/*
 * Exercise Solutions for Variables, Types, and Memory Management
 * 
 * This file contains solutions to all practice exercises from the lesson.
 * Uncomment and run individual exercises to test your understanding.
 */

#include <iostream>
#include <climits>
#include <cfloat>
#include <string>
#include <vector>

// Exercise 1: Type Sizes and Ranges
void exercise1() {
    std::cout << "=== EXERCISE 1: TYPE SIZES AND RANGES ===" << std::endl;
    
    // Integer types
    std::cout << "\n--- INTEGER TYPES ---" << std::endl;
    std::cout << "char: " << sizeof(char) << " bytes, range: " 
              << CHAR_MIN << " to " << CHAR_MAX << std::endl;
    std::cout << "unsigned char: " << sizeof(unsigned char) << " bytes, range: 0 to " 
              << UCHAR_MAX << std::endl;
    std::cout << "short: " << sizeof(short) << " bytes, range: " 
              << SHRT_MIN << " to " << SHRT_MAX << std::endl;
    std::cout << "unsigned short: " << sizeof(unsigned short) << " bytes, range: 0 to " 
              << USHRT_MAX << std::endl;
    std::cout << "int: " << sizeof(int) << " bytes, range: " 
              << INT_MIN << " to " << INT_MAX << std::endl;
    std::cout << "unsigned int: " << sizeof(unsigned int) << " bytes, range: 0 to " 
              << UINT_MAX << std::endl;
    std::cout << "long: " << sizeof(long) << " bytes, range: " 
              << LONG_MIN << " to " << LONG_MAX << std::endl;
    std::cout << "unsigned long: " << sizeof(unsigned long) << " bytes, range: 0 to " 
              << ULONG_MAX << std::endl;
    std::cout << "long long: " << sizeof(long long) << " bytes, range: " 
              << LLONG_MIN << " to " << LLONG_MAX << std::endl;
    std::cout << "unsigned long long: " << sizeof(unsigned long long) << " bytes, range: 0 to " 
              << ULLONG_MAX << std::endl;
    
    // Floating-point types
    std::cout << "\n--- FLOATING-POINT TYPES ---" << std::endl;
    std::cout << "float: " << sizeof(float) << " bytes, precision: " 
              << FLT_DIG << " digits" << std::endl;
    std::cout << "double: " << sizeof(double) << " bytes, precision: " 
              << DBL_DIG << " digits" << std::endl;
    std::cout << "long double: " << sizeof(long double) << " bytes, precision: " 
              << LDBL_DIG << " digits" << std::endl;
    
    // Boolean type
    std::cout << "\n--- BOOLEAN TYPE ---" << std::endl;
    std::cout << "bool: " << sizeof(bool) << " bytes" << std::endl;
}

// Exercise 2: Memory Layout Analysis
void exercise2() {
    std::cout << "\n=== EXERCISE 2: MEMORY LAYOUT ANALYSIS ===" << std::endl;
    
    // Stack variables
    int stack_int = 42;
    double stack_double = 3.14;
    char stack_char = 'A';
    
    // Heap variables
    int* heap_int = new int(100);
    double* heap_double = new double(2.71);
    char* heap_char = new char('B');
    
    std::cout << "\n--- STACK MEMORY ---" << std::endl;
    std::cout << "stack_int: " << stack_int << " at address " << &stack_int << std::endl;
    std::cout << "stack_double: " << stack_double << " at address " << &stack_double << std::endl;
    std::cout << "stack_char: " << stack_char << " at address " << &stack_char << std::endl;
    
    std::cout << "\n--- HEAP MEMORY ---" << std::endl;
    std::cout << "heap_int: " << *heap_int << " at address " << heap_int << std::endl;
    std::cout << "heap_double: " << *heap_double << " at address " << heap_double << std::endl;
    std::cout << "heap_char: " << *heap_char << " at address " << heap_char << std::endl;
    
    // Demonstrate pointer arithmetic
    std::cout << "\n--- POINTER ARITHMETIC ---" << std::endl;
    int arr[5] = {1, 2, 3, 4, 5};
    std::cout << "Array elements and addresses:" << std::endl;
    for (int i = 0; i < 5; ++i) {
        std::cout << "arr[" << i << "]: " << arr[i] << " at " << &arr[i] << std::endl;
    }
    
    // Clean up heap memory
    delete heap_int;
    delete heap_double;
    delete heap_char;
}

// Exercise 3: Initialization Methods
void exercise3() {
    std::cout << "\n=== EXERCISE 3: INITIALIZATION METHODS ===" << std::endl;
    
    // Copy initialization
    int a = 10;
    double b = 3.14;
    std::string c = "Hello";
    
    // Direct initialization
    int d(20);
    double e(2.71);
    std::string f("World");
    
    // Uniform initialization (preferred)
    int g{30};
    double h{1.41};
    std::string i{"C++"};
    
    // Value initialization
    int j{};
    double k{};
    std::string l{};
    
    std::cout << "Copy initialization: a=" << a << ", b=" << b << ", c=" << c << std::endl;
    std::cout << "Direct initialization: d=" << d << ", e=" << e << ", f=" << f << std::endl;
    std::cout << "Uniform initialization: g=" << g << ", h=" << h << ", i=" << i << std::endl;
    std::cout << "Value initialization: j=" << j << ", k=" << k << ", l=" << l << std::endl;
    
    // Demonstrating narrowing conversion prevention
    std::cout << "\n--- NARROWING CONVERSION PREVENTION ---" << std::endl;
    int narrow_int = 3.14;  // Allowed (narrowing)
    // int narrow_int2{3.14};  // Error: narrowing conversion not allowed
    
    std::cout << "narrow_int (copy init): " << narrow_int << std::endl;
}

// Exercise 4: Const and Constexpr
void exercise4() {
    std::cout << "\n=== EXERCISE 4: CONST AND CONSTEXPR ===" << std::endl;
    
    // const variables
    const int MAX_SIZE = 1000;
    const double PI = 3.14159265359;
    const std::string APP_NAME = "C++ Learning";
    
    // constexpr variables
    constexpr int ARRAY_SIZE = 50;
    constexpr double E = 2.71828182846;
    constexpr int FIBONACCI_10 = 55;  // Precomputed
    
    // constexpr function
    constexpr int factorial(int n) {
        return (n <= 1) ? 1 : n * factorial(n - 1);
    }
    
    // constexpr function for Fibonacci
    constexpr int fibonacci(int n) {
        if (n <= 1) return n;
        return fibonacci(n - 1) + fibonacci(n - 2);
    }
    
    std::cout << "const values: MAX_SIZE=" << MAX_SIZE << ", PI=" << PI << std::endl;
    std::cout << "constexpr values: ARRAY_SIZE=" << ARRAY_SIZE << ", E=" << E << std::endl;
    std::cout << "constexpr functions:" << std::endl;
    std::cout << "  5! = " << factorial(5) << std::endl;
    std::cout << "  fibonacci(10) = " << fibonacci(10) << std::endl;
    
    // Using constexpr in array size
    int numbers[ARRAY_SIZE];
    for (int i = 0; i < ARRAY_SIZE; ++i) {
        numbers[i] = i * i;
    }
    std::cout << "Array with constexpr size created, first few values: ";
    for (int i = 0; i < 5; ++i) {
        std::cout << numbers[i] << " ";
    }
    std::cout << std::endl;
}

// Exercise 5: Memory Management with RAII
class ResourceManager {
private:
    int* data;
    size_t size;
    
public:
    // Constructor
    explicit ResourceManager(size_t s) : size(s) {
        data = new int[size];
        std::cout << "ResourceManager: Allocated " << size << " integers" << std::endl;
    }
    
    // Destructor (RAII)
    ~ResourceManager() {
        delete[] data;
        std::cout << "ResourceManager: Deallocated " << size << " integers" << std::endl;
    }
    
    // Copy constructor (deep copy)
    ResourceManager(const ResourceManager& other) : size(other.size) {
        data = new int[size];
        for (size_t i = 0; i < size; ++i) {
            data[i] = other.data[i];
        }
        std::cout << "ResourceManager: Copy constructed" << std::endl;
    }
    
    // Copy assignment operator
    ResourceManager& operator=(const ResourceManager& other) {
        if (this != &other) {
            delete[] data;
            size = other.size;
            data = new int[size];
            for (size_t i = 0; i < size; ++i) {
                data[i] = other.data[i];
            }
        }
        std::cout << "ResourceManager: Copy assigned" << std::endl;
        return *this;
    }
    
    // Move constructor (C++11)
    ResourceManager(ResourceManager&& other) noexcept : data(other.data), size(other.size) {
        other.data = nullptr;
        other.size = 0;
        std::cout << "ResourceManager: Move constructed" << std::endl;
    }
    
    // Move assignment operator (C++11)
    ResourceManager& operator=(ResourceManager&& other) noexcept {
        if (this != &other) {
            delete[] data;
            data = other.data;
            size = other.size;
            other.data = nullptr;
            other.size = 0;
        }
        std::cout << "ResourceManager: Move assigned" << std::endl;
        return *this;
    }
    
    // Accessor methods
    int& operator[](size_t index) { return data[index]; }
    const int& operator[](size_t index) const { return data[index]; }
    size_t getSize() const { return size; }
    
    // Initialize with values
    void initialize() {
        for (size_t i = 0; i < size; ++i) {
            data[i] = static_cast<int>(i * i);
        }
    }
    
    // Display contents
    void display() const {
        std::cout << "ResourceManager contents: ";
        for (size_t i = 0; i < size && i < 10; ++i) {  // Show first 10
            std::cout << data[i] << " ";
        }
        if (size > 10) std::cout << "...";
        std::cout << std::endl;
    }
};

void exercise5() {
    std::cout << "\n=== EXERCISE 5: MEMORY MANAGEMENT WITH RAII ===" << std::endl;
    
    {
        // Create ResourceManager in a scope
        ResourceManager rm1(5);
        rm1.initialize();
        rm1.display();
        
        // Copy construction
        ResourceManager rm2 = rm1;
        rm2.display();
        
        // Move construction
        ResourceManager rm3 = std::move(rm1);
        rm3.display();
        
        // Assignment
        ResourceManager rm4(3);
        rm4 = rm2;
        rm4.display();
        
        // Move assignment
        rm4 = std::move(rm3);
        rm4.display();
        
        std::cout << "End of scope - destructors will be called" << std::endl;
    }  // All ResourceManager objects are automatically destroyed here
    
    std::cout << "All resources properly cleaned up!" << std::endl;
}

// Main function to run exercises
int main() {
    std::cout << "C++ Variables, Types, and Memory Management - Exercise Solutions" << std::endl;
    std::cout << "=================================================================" << std::endl;
    
    // Uncomment the exercises you want to run:
    
    // exercise1();  // Type sizes and ranges
    // exercise2();  // Memory layout analysis
    // exercise3();  // Initialization methods
    // exercise4();  // Const and constexpr
    // exercise5();  // Memory management with RAII
    
    // Run all exercises
    exercise1();
    exercise2();
    exercise3();
    exercise4();
    exercise5();
    
    return 0;
}
