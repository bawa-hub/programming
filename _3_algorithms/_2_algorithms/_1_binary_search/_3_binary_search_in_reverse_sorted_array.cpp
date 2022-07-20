/***
 * Time Complexity --
 * Best - O(1)
 * Average - O(log n)
 * Worst - O(log n)
 *
 * Space Cmplexity = O(1)
 */

#include <iostream>
using namespace std;

// binary search only applies on sorted array
int binarySearch(int array[], int x, int low, int high)
{
    if (high >= low)
    {
        int mid = low + (high - low) / 2; // this calculation is used to overcome int overflow

        // If found at mid, then return it
        if (array[mid] == x)
            return mid;

        // Search the left half
        if (array[mid] > x)
            return binarySearch(array, x, mid + 1, high);

        // Search the right half
        return binarySearch(array, x, low, mid - 1);
    }

    return -1;
}

int main(void)
{
    int array[] = {10, 9, 8, 5, 4};
    int x = 4;
    int n = sizeof(array) / sizeof(array[0]);
    int result = binarySearch(array, x, 0, n - 1);
    if (result == -1)
        printf("Not found");
    else
        printf("Element is found at index %d", result);
}