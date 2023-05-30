// https://leetcode.com/problems/binary-search/

#include <iostream>
using namespace std;

// recursively
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
            return binarySearch(array, x, low, mid - 1);

        // Search the right half
        return binarySearch(array, x, mid + 1, high);
    }

    return -1;
}
// Time complexity: O(log n)
// Space complexity: O(logn) for auxiliary space

// iteratively
int binarySearch(int arr[], int k, int n)
{
    int start = 0, end = n;
    int mid, loc = -1;
    while (start <= n - 1)
    {
        // Making array half everytime
        mid = (start + end) / 2;

        // checking in which part the element is present
        if (arr[mid] < k)
        {
            start = mid + 1;
        }
        else if (arr[mid] > k)
        {
            end = mid - 1;
        }
        if (arr[mid] == k)
        {
            loc = mid;
            break;
        }
    }
    if (loc == -1)
    {
        cout << "Element not found!" << endl;
    }
    else
    {
        cout << "Element " << k << " Found at " << loc << " index";
    }
}
// Time complexity: O(log n)
// Space complexity : O(1)

int main(void)
{
    int array[] = {3, 4, 5, 6, 7, 8, 9};
    int x = 4;
    int n = sizeof(array) / sizeof(array[0]);
    int result = binarySearch(array, x, 0, n - 1);
    if (result == -1)
        printf("Not found");
    else
        printf("Element is found at index %d", result);
}