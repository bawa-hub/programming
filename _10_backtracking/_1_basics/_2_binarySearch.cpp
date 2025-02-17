#include <iostream>
using namespace std;

int binarySearch(int array[], int x, int low, int high)
{

    // Repeat until the pointers low and high meet each other
    while (low <= high)
    {
        int mid = low + (high - low) / 2;

        if (array[mid] == x)
            return mid;

        if (array[mid] < x)
            low = mid + 1;

        else
            high = mid - 1;
    }

    return -1;
}

int binarySearchRecursive(int a[], int l, int r, int num)
{
    int mid = ((r - l) / 2) + l;
    if (a[mid] == num)
        return mid;
    if (a[mid] > num)
        return binarySearch(a, l, mid - 1, num);
    if (a[mid] < num)
        return binarySearch(a, l + 1, r, num);

    return -1;
}

int main(void)
{
    int array[] = {3, 4, 5, 6, 7, 8, 999, 133};
    int x = 133;
    int n = sizeof(array) / sizeof(array[0]);
    int result = binarySearch(array, x, 0, n - 1);
    if (result == -1)
        printf("Not found");
    else
        printf("Element is found at index %d", result);
}