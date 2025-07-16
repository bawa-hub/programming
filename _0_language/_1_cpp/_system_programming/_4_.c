#include <stdio.h>
#include <stdlib.h>

int main() {
    int n;
    printf("How many numbers? ");
    scanf("%d", &n);

    // Allocate memory for n integers on the heap
    int *arr = (int *)malloc(n * sizeof(int));

    // Check if memory allocation was successful
    if (arr == NULL) {
        printf("Memory allocation failed!\n");
        return 1;
    }

    // Input n numbers
    printf("Enter %d numbers:\n", n);
    for (int i = 0; i < n; i++) {
        printf("arr[%d]: ", i);
        scanf("%d", &arr[i]);
    }

    // Print the entered numbers
    printf("\nYou entered:\n");
    for (int i = 0; i < n; i++) {
        printf("arr[%d] = %d\n", i, arr[i]);
    }

    // Free the allocated memory
    free(arr);

    return 0;
}
