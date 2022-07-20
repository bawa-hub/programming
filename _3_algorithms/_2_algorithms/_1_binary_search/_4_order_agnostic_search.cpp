// C++ program for the above approach
#include <bits/stdc++.h>
using namespace std;

// A recursive binary search function.
// It returns location of x in given
// array arr[l..r] is present,
// otherwise -1
int binarySearch(int arr[], int start,
                 int end, int x)
{
    bool isAsc = arr[start] < arr[end];
    if (end >= start)
    {
        int middle = start + (end - start) / 2;

        // If the element is present
        // at the middle itself
        if (arr[middle] == x)
            return middle;

        if (isAsc == true)
        {

            // If element is smaller than mid,
            // then it can only be
            // present in left subarray
            if (arr[middle] > x)
                return binarySearch(
                    arr, start,
                    middle - 1, x);

            // Else the element can only be present
            // in right subarray
            return binarySearch(arr, middle + 1,
                                end, x);
        }
        else
        {
            if (arr[middle] < x)
                return binarySearch(arr, start,
                                    middle - 1, x);

            // Else the element can only be present
            // in left subarray
            return binarySearch(arr, middle + 1,
                                end, x);
        }
    }

    // Element not found
    return -1;
}

// Driver Code
int main(void)
{
    int arr[] = {40, 10, 5, 2, 1};
    int x = 10;
    int n = sizeof(arr) / sizeof(arr[0]);
    cout << binarySearch(arr, 0, n - 1, x);

    return 0;
}
