// an array is a collection of a fixed number of values. Once the size of an array is declared, you cannot change it.
// you can allocate memory manually during run-time.
// This is known as dynamic memory allocation in C programming.

// To allocate memory dynamically, library functions:
// malloc(), calloc(), realloc() and free() are used.
// These functions are defined in the <stdlib.h> header file.

// C malloc()
// The name "malloc" stands for memory allocation.
// The malloc() function reserves a block of memory of the specified number of bytes.
// And, it returns a pointer of void which can be casted into pointers of any form.
// Syntax of malloc()
// ptr = (castType*) malloc(size);
// Example
// ptr = (float*) malloc(100 * sizeof(float));
// The above statement allocates 400 bytes of memory
// the pointer ptr holds the address of the first byte in the allocated memory
// expression results in a NULL pointer if the memory cannot be allocated.

// C calloc()
// The name "calloc" stands for contiguous allocation.
// The malloc() function allocates memory and leaves the memory uninitialized,
// whereas the calloc() function allocates memory and initializes all bits to zero.
// Syntax of calloc()
// ptr = (castType*)calloc(n, size);
// Example:
// ptr = (float*) calloc(25, sizeof(float));
// The above statement allocates contiguous space in memory for 25 elements of type float.

// C free()
// Dynamically allocated memory created with either calloc() or malloc() doesn't get freed on their own.
// You must explicitly use free() to release the space.
// Syntax of free()
// free(ptr);
// This statement frees the space allocated in the memory pointed by ptr.

// C realloc()
// If the dynamically allocated memory is insufficient or more than required,
// you can change the size of previously allocated memory using the realloc() function.
// Syntax of realloc()
// ptr = realloc(ptr, x);
// Here, ptr is reallocated with a new size x