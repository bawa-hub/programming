#include <stdio.h>

// without using temp variable
void swap(int *a, int *b)
{
    if (a != b)
    { // To avoid zeroing out when both pointers point to the same location
        *a = *a ^ *b;
        *b = *a ^ *b;
        *a = *a ^ *b;
    }
}

int main()
{
    // basic pointer
    int a = 42;
    int *ptr = &a;

    printf("Value of a: %d\n", a);
    printf("Address of a: %p\n", &a);
    printf("Value stored in ptr (address): %p\n", ptr);
    printf("Value pointed to by ptr: %d\n", *ptr);

    // swap using pointer 
    int x = 10, y = 20;
    printf("Before swap: x = %d, y = %d\n", x, y);
    swap(&x, &y);
    printf("After swap:  x = %d, y = %d\n", x, y);

    return 0;
}
