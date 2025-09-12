/*
 * Exercise Solutions for Pointers and References
 * 
 * This file contains solutions to all practice exercises from the lesson.
 * Uncomment and run individual exercises to test your understanding.
 */

#include <iostream>
#include <memory>
#include <vector>
#include <chrono>
#include <random>

// Exercise 1: Basic Pointer Operations
void exercise1() {
    std::cout << "=== EXERCISE 1: BASIC POINTER OPERATIONS ===" << std::endl;
    
    // Declaration and initialization
    int x = 42;
    int* ptr = &x;
    
    std::cout << "--- Declaration and Initialization ---" << std::endl;
    std::cout << "x = " << x << std::endl;
    std::cout << "Address of x: " << &x << std::endl;
    std::cout << "ptr = " << ptr << std::endl;
    std::cout << "Address of ptr: " << &ptr << std::endl;
    
    // Dereferencing
    std::cout << "\n--- Dereferencing ---" << std::endl;
    std::cout << "Value pointed to by ptr: " << *ptr << std::endl;
    
    // Modifying through pointer
    *ptr = 100;
    std::cout << "After *ptr = 100, x = " << x << std::endl;
    
    // Pointer arithmetic with arrays
    std::cout << "\n--- Pointer Arithmetic ---" << std::endl;
    int arr[5] = {10, 20, 30, 40, 50};
    int* arr_ptr = arr;  // Points to first element
    
    std::cout << "Array elements using pointer arithmetic:" << std::endl;
    for (int i = 0; i < 5; ++i) {
        std::cout << "*(arr_ptr + " << i << ") = " << *(arr_ptr + i) 
                  << " at address " << (arr_ptr + i) << std::endl;
    }
    
    // Alternative array access
    std::cout << "\nArray elements using array notation:" << std::endl;
    for (int i = 0; i < 5; ++i) {
        std::cout << "arr_ptr[" << i << "] = " << arr_ptr[i] << std::endl;
    }
    
    // Pointer comparison
    std::cout << "\n--- Pointer Comparison ---" << std::endl;
    int* ptr1 = &arr[0];
    int* ptr2 = &arr[2];
    std::cout << "ptr1 points to: " << *ptr1 << std::endl;
    std::cout << "ptr2 points to: " << *ptr2 << std::endl;
    std::cout << "ptr1 < ptr2: " << (ptr1 < ptr2) << std::endl;
    std::cout << "ptr2 - ptr1: " << (ptr2 - ptr1) << std::endl;
}

// Exercise 2: Reference vs Pointer
void exercise2() {
    std::cout << "\n=== EXERCISE 2: REFERENCE VS POINTER ===" << std::endl;
    
    int original = 42;
    
    // Reference
    int& ref = original;
    std::cout << "--- Reference ---" << std::endl;
    std::cout << "original = " << original << std::endl;
    std::cout << "ref = " << ref << std::endl;
    std::cout << "Address of original: " << &original << std::endl;
    std::cout << "Address of ref: " << &ref << std::endl;
    std::cout << "Same address? " << (&original == &ref) << std::endl;
    
    // Pointer
    int* ptr = &original;
    std::cout << "\n--- Pointer ---" << std::endl;
    std::cout << "original = " << original << std::endl;
    std::cout << "ptr = " << ptr << std::endl;
    std::cout << "Address of original: " << &original << std::endl;
    std::cout << "Address of ptr: " << &ptr << std::endl;
    std::cout << "ptr points to original? " << (ptr == &original) << std::endl;
    
    // Modifying through reference
    ref = 100;
    std::cout << "\nAfter ref = 100:" << std::endl;
    std::cout << "original = " << original << std::endl;
    std::cout << "ref = " << ref << std::endl;
    
    // Modifying through pointer
    *ptr = 200;
    std::cout << "\nAfter *ptr = 200:" << std::endl;
    std::cout << "original = " << original << std::endl;
    std::cout << "ref = " << ref << std::endl;
    std::cout << "*ptr = " << *ptr << std::endl;
    
    // Demonstrating when to use each
    std::cout << "\n--- When to Use Each ---" << std::endl;
    std::cout << "Use references when:" << std::endl;
    std::cout << "  - You need an alias for an existing variable" << std::endl;
    std::cout << "  - Function parameters that shouldn't be null" << std::endl;
    std::cout << "  - Operator overloading" << std::endl;
    std::cout << "  - Range-based for loops" << std::endl;
    
    std::cout << "\nUse pointers when:" << std::endl;
    std::cout << "  - You need to represent 'no object' (nullptr)" << std::endl;
    std::cout << "  - Dynamic memory allocation" << std::endl;
    std::cout << "  - Array manipulation" << std::endl;
    std::cout << "  - Optional parameters" << std::endl;
}

// Exercise 3: Dynamic Memory Management with RAII
class DynamicArray {
private:
    int* data;
    size_t size;
    
public:
    // Constructor
    explicit DynamicArray(size_t s) : size(s) {
        data = new int[size];
        std::cout << "DynamicArray: Allocated " << size << " integers" << std::endl;
    }
    
    // Destructor (RAII)
    ~DynamicArray() {
        delete[] data;
        std::cout << "DynamicArray: Deallocated " << size << " integers" << std::endl;
    }
    
    // Copy constructor (deep copy)
    DynamicArray(const DynamicArray& other) : size(other.size) {
        data = new int[size];
        for (size_t i = 0; i < size; ++i) {
            data[i] = other.data[i];
        }
        std::cout << "DynamicArray: Copy constructed" << std::endl;
    }
    
    // Copy assignment operator
    DynamicArray& operator=(const DynamicArray& other) {
        if (this != &other) {
            delete[] data;
            size = other.size;
            data = new int[size];
            for (size_t i = 0; i < size; ++i) {
                data[i] = other.data[i];
            }
        }
        std::cout << "DynamicArray: Copy assigned" << std::endl;
        return *this;
    }
    
    // Move constructor (C++11)
    DynamicArray(DynamicArray&& other) noexcept : data(other.data), size(other.size) {
        other.data = nullptr;
        other.size = 0;
        std::cout << "DynamicArray: Move constructed" << std::endl;
    }
    
    // Move assignment operator (C++11)
    DynamicArray& operator=(DynamicArray&& other) noexcept {
        if (this != &other) {
            delete[] data;
            data = other.data;
            size = other.size;
            other.data = nullptr;
            other.size = 0;
        }
        std::cout << "DynamicArray: Move assigned" << std::endl;
        return *this;
    }
    
    // Accessor methods
    int& operator[](size_t index) { 
        return data[index]; 
    }
    
    const int& operator[](size_t index) const { 
        return data[index]; 
    }
    
    size_t getSize() const { 
        return size; 
    }
    
    // Initialize with values
    void initialize() {
        for (size_t i = 0; i < size; ++i) {
            data[i] = static_cast<int>(i * i);
        }
    }
    
    // Display contents
    void display() const {
        std::cout << "DynamicArray contents: ";
        for (size_t i = 0; i < size && i < 10; ++i) {  // Show first 10
            std::cout << data[i] << " ";
        }
        if (size > 10) std::cout << "...";
        std::cout << std::endl;
    }
    
    // Get raw pointer (for interfacing with C code)
    int* getData() { return data; }
    const int* getData() const { return data; }
};

void exercise3() {
    std::cout << "\n=== EXERCISE 3: DYNAMIC MEMORY MANAGEMENT WITH RAII ===" << std::endl;
    
    {
        // Create DynamicArray in a scope
        DynamicArray arr1(5);
        arr1.initialize();
        arr1.display();
        
        // Copy construction
        DynamicArray arr2 = arr1;
        arr2.display();
        
        // Move construction
        DynamicArray arr3 = std::move(arr1);
        arr3.display();
        
        // Assignment
        DynamicArray arr4(3);
        arr4 = arr2;
        arr4.display();
        
        // Move assignment
        arr4 = std::move(arr3);
        arr4.display();
        
        // Demonstrate pointer access
        std::cout << "\n--- Raw Pointer Access ---" << std::endl;
        int* raw_ptr = arr4.getData();
        std::cout << "First element via raw pointer: " << *raw_ptr << std::endl;
        std::cout << "Second element via raw pointer: " << *(raw_ptr + 1) << std::endl;
        
        std::cout << "\nEnd of scope - destructors will be called" << std::endl;
    }  // All DynamicArray objects are automatically destroyed here
    
    std::cout << "All resources properly cleaned up!" << std::endl;
}

// Exercise 4: Function Pointers
class Calculator {
public:
    // Static member functions for function pointers
    static int add(int a, int b) { return a + b; }
    static int subtract(int a, int b) { return a - b; }
    static int multiply(int a, int b) { return a * b; }
    static int divide(int a, int b) { 
        if (b != 0) return a / b;
        return 0;  // Handle division by zero
    }
    
    // Non-static member functions
    int power(int a, int b) {
        int result = 1;
        for (int i = 0; i < b; ++i) {
            result *= a;
        }
        return result;
    }
    
    // Function that takes a function pointer
    static int calculate(int a, int b, int (*operation)(int, int)) {
        return operation(a, b);
    }
    
    // Function that takes a member function pointer
    int calculateMember(int a, int b, int (Calculator::*operation)(int, int)) {
        return (this->*operation)(a, b);
    }
};

void exercise4() {
    std::cout << "\n=== EXERCISE 4: FUNCTION POINTERS ===" << std::endl;
    
    // Function pointers for static functions
    std::cout << "--- Static Function Pointers ---" << std::endl;
    int (*func_ptr)(int, int) = Calculator::add;
    std::cout << "add(10, 5) = " << func_ptr(10, 5) << std::endl;
    
    func_ptr = Calculator::multiply;
    std::cout << "multiply(10, 5) = " << func_ptr(10, 5) << std::endl;
    
    // Using function pointers as parameters
    std::cout << "\n--- Function Pointers as Parameters ---" << std::endl;
    std::cout << "calculate(8, 4, add) = " << Calculator::calculate(8, 4, Calculator::add) << std::endl;
    std::cout << "calculate(8, 4, subtract) = " << Calculator::calculate(8, 4, Calculator::subtract) << std::endl;
    std::cout << "calculate(8, 4, multiply) = " << Calculator::calculate(8, 4, Calculator::multiply) << std::endl;
    std::cout << "calculate(8, 4, divide) = " << Calculator::calculate(8, 4, Calculator::divide) << std::endl;
    
    // Member function pointers
    std::cout << "\n--- Member Function Pointers ---" << std::endl;
    Calculator calc;
    int (Calculator::*member_func_ptr)(int, int) = &Calculator::power;
    std::cout << "calc.power(2, 8) = " << (calc.*member_func_ptr)(2, 8) << std::endl;
    
    // Using member function pointers as parameters
    std::cout << "calc.calculateMember(3, 4, power) = " 
              << calc.calculateMember(3, 4, &Calculator::power) << std::endl;
    
    // Array of function pointers
    std::cout << "\n--- Array of Function Pointers ---" << std::endl;
    int (*operations[])(int, int) = {Calculator::add, Calculator::subtract, 
                                     Calculator::multiply, Calculator::divide};
    const char* operation_names[] = {"add", "subtract", "multiply", "divide"};
    
    int a = 12, b = 3;
    for (int i = 0; i < 4; ++i) {
        std::cout << operation_names[i] << "(" << a << ", " << b << ") = " 
                  << operations[i](a, b) << std::endl;
    }
}

// Exercise 5: Advanced Pointer Techniques - Linked List
template<typename T>
class LinkedList {
private:
    struct Node {
        T data;
        Node* next;
        
        Node(const T& value) : data(value), next(nullptr) {}
    };
    
    Node* head;
    size_t size;
    
public:
    LinkedList() : head(nullptr), size(0) {}
    
    ~LinkedList() {
        clear();
    }
    
    // Copy constructor
    LinkedList(const LinkedList& other) : head(nullptr), size(0) {
        Node* current = other.head;
        while (current) {
            push_back(current->data);
            current = current->next;
        }
    }
    
    // Copy assignment operator
    LinkedList& operator=(const LinkedList& other) {
        if (this != &other) {
            clear();
            Node* current = other.head;
            while (current) {
                push_back(current->data);
                current = current->next;
            }
        }
        return *this;
    }
    
    // Move constructor
    LinkedList(LinkedList&& other) noexcept : head(other.head), size(other.size) {
        other.head = nullptr;
        other.size = 0;
    }
    
    // Move assignment operator
    LinkedList& operator=(LinkedList&& other) noexcept {
        if (this != &other) {
            clear();
            head = other.head;
            size = other.size;
            other.head = nullptr;
            other.size = 0;
        }
        return *this;
    }
    
    void push_front(const T& value) {
        Node* new_node = new Node(value);
        new_node->next = head;
        head = new_node;
        ++size;
    }
    
    void push_back(const T& value) {
        Node* new_node = new Node(value);
        if (!head) {
            head = new_node;
        } else {
            Node* current = head;
            while (current->next) {
                current = current->next;
            }
            current->next = new_node;
        }
        ++size;
    }
    
    void pop_front() {
        if (head) {
            Node* temp = head;
            head = head->next;
            delete temp;
            --size;
        }
    }
    
    void clear() {
        while (head) {
            pop_front();
        }
    }
    
    size_t getSize() const { return size; }
    
    bool empty() const { return head == nullptr; }
    
    void display() const {
        std::cout << "LinkedList: ";
        Node* current = head;
        while (current) {
            std::cout << current->data << " ";
            current = current->next;
        }
        std::cout << std::endl;
    }
    
    // Iterator-like functionality
    class Iterator {
    private:
        Node* current;
        
    public:
        Iterator(Node* node) : current(node) {}
        
        T& operator*() { return current->data; }
        const T& operator*() const { return current->data; }
        
        Iterator& operator++() {
            current = current->next;
            return *this;
        }
        
        bool operator!=(const Iterator& other) const {
            return current != other.current;
        }
    };
    
    Iterator begin() { return Iterator(head); }
    Iterator end() { return Iterator(nullptr); }
};

void exercise5() {
    std::cout << "\n=== EXERCISE 5: ADVANCED POINTER TECHNIQUES - LINKED LIST ===" << std::endl;
    
    // Create linked list
    LinkedList<int> list;
    
    // Add elements
    std::cout << "--- Adding Elements ---" << std::endl;
    list.push_back(10);
    list.push_back(20);
    list.push_back(30);
    list.display();
    
    list.push_front(5);
    list.push_front(1);
    list.display();
    
    // Demonstrate copy construction
    std::cout << "\n--- Copy Construction ---" << std::endl;
    LinkedList<int> list2 = list;
    std::cout << "Original list: ";
    list.display();
    std::cout << "Copied list: ";
    list2.display();
    
    // Demonstrate move construction
    std::cout << "\n--- Move Construction ---" << std::endl;
    LinkedList<int> list3 = std::move(list2);
    std::cout << "Moved list: ";
    list3.display();
    std::cout << "Original list2 is now empty: " << list2.empty() << std::endl;
    
    // Demonstrate iterator-like functionality
    std::cout << "\n--- Iterator-like Functionality ---" << std::endl;
    std::cout << "List elements using iterator: ";
    for (auto it = list3.begin(); it != list3.end(); ++it) {
        std::cout << *it << " ";
    }
    std::cout << std::endl;
    
    // Demonstrate removal
    std::cout << "\n--- Removal ---" << std::endl;
    list3.pop_front();
    std::cout << "After pop_front: ";
    list3.display();
    
    // Demonstrate pointer arithmetic concepts
    std::cout << "\n--- Pointer Concepts in Linked List ---" << std::endl;
    std::cout << "Each node contains a pointer to the next node" << std::endl;
    std::cout << "Traversal requires following the 'next' pointers" << std::endl;
    std::cout << "Memory is not contiguous (unlike arrays)" << std::endl;
    
    // Clean up
    list.clear();
    list3.clear();
    std::cout << "Lists cleared" << std::endl;
}

// Main function to run exercises
int main() {
    std::cout << "C++ Pointers and References - Exercise Solutions" << std::endl;
    std::cout << "================================================" << std::endl;
    
    // Uncomment the exercises you want to run:
    
    // exercise1();  // Basic pointer operations
    // exercise2();  // Reference vs pointer
    // exercise3();  // Dynamic memory management with RAII
    // exercise4();  // Function pointers
    // exercise5();  // Advanced pointer techniques - linked list
    
    // Run all exercises
    exercise1();
    exercise2();
    exercise3();
    exercise4();
    exercise5();
    
    return 0;
}
