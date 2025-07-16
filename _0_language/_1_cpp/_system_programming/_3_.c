// Time to understand Stack vs Heap — a core system-level concept every real hacker and system engineer must master. This will lay the groundwork for you to build allocators, optimize memory, and understand what’s happening under the hood.

// Feature | Stack | Heap
// Memory Type | Fixed-size memory region | Dynamically allocated memory
// Managed By | Compiler | You (via malloc/free)
// Allocation | Automatic | Manual
// Lifetime | Tied to function scope | Until you free it
// Speed | Very fast (LIFO) | Slower (requires bookkeeping)
// Usage | Local variables, function calls | Dynamic memory, large data
// Overflow | Stack overflow | Memory leak

#include <stdio.h>

void foo()
{
    // 🔍 localVar is automatically created and destroyed when foo() finishes. You don’t manage this memory.
    int localVar = 10; // stored on the stack
    printf("Stack variable: %d\n", localVar);
}

int main()
{

    // stack memory example
    foo();

    // heap memory example
    // 🔍 malloc allocates memory on the heap. You must free() it yourself.
    int *ptr = (int *)malloc(sizeof(int)); // memory on the heap
    *ptr = 42;
    printf("Heap variable: %d\n", *ptr);
    free(ptr); // clean up manually
    return 0;
}

// 🚨 Why This Matters:
//     Stack is fast but limited (default ~1MB/thread).
//     Heap is flexible, but if you forget to free() it → memory leaks 💀
//     System-level programming often requires tight control over heap usage for performance and safety.