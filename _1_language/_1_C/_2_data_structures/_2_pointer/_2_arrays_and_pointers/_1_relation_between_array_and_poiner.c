// An array is a block of sequential data
#include <stdio.h>
int main()
{
    int x[4];
    int i;

    printf("Addresses of an array elements are:\n");
    for (i = 0; i < 4; ++i)
    {
        printf("&x[%d] = %p\n", i, &x[i]);
    }
    printf("\n");

    printf("Address of array x: %p\n\n", x);
    printf("size of int is 4 bytes (on my compiler).\n\n");
    printf("Notice that, the address of &x[0] and x is the same,\n It's because the variable name x points to the first element of the array.\n");
    printf("it is clear that &x[0] is equivalent to x.\n\n");

    printf("Similarly,\n");
    printf("&x[1] is equivalent to x+1 and x[1] is equivalent to *(x+1)\n");
    printf("&x[2] is equivalent to x+2 and x[2] is equivalent to *(x+2).\n");
    printf("...\n");
    printf("Basically, &x[i] is equivalent to x+i and x[i] is equivalent to *(x+i)\n\n");

    return 0;
}