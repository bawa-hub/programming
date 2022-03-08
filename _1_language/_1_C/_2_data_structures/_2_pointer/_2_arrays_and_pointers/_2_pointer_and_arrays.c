#include <stdio.h>
int main()
{
    int i, x[6], sum = 0;
    printf("Enter 6 numbers:\n ");
    for (i = 0; i < 6; ++i)
    {
        // Equivalent to scanf("%d", &x[i]);
        scanf("%d", x + i);

        // Equivalent to sum += x[i]
        sum += *(x + i);
    }
    printf("Sum = %d", sum);
    return 0;
}

// To access elements of the array, we have used pointers
// In most contexts, array names decay to pointers
// In simple words, array names are converted to pointers
// That's the reason why you can use pointers to access elements of arrays.
// However, you should remember that pointers and arrays are not the same

// There are a few cases where array names don't decay to pointers:
// https://stackoverflow.com/questions/17752978/exceptions-to-array-decaying-into-a-pointer