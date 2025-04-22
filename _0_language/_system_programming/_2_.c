#include <stdio.h>

int main() {
    int arr[] = {10, 20, 30, 40, 50};
    int *ptr = arr;

    printf("Array elements using pointer arithmetic:\n");

    // TODO: Loop through the array using the pointer
    for (int i = 0; i < 5; i++) {
        // Print each element using ptr
        printf("%d number is %d\n", i, *ptr);
        ptr++;
    }

    return 0;
}
