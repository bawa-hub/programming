// Precalculation approach
#include <bits/stdc++.h>
using namespace std;

int findWater(int arr[], int n)
{
    int left[n];

    int right[n];

    int water = 0;

    left[0] = arr[0];
    for (int i = 1; i < n; i++)
        left[i] = max(left[i - 1], arr[i]);

    right[n - 1] = arr[n - 1];
    for (int i = n - 2; i >= 0; i--)
        right[i] = max(right[i + 1], arr[i]);

    for (int i = 1; i < n - 1; i++)
    {
        int var = min(left[i - 1], right[i + 1]);
        if (var > arr[i])
        {
            water += var - arr[i];
        }
    }

    return water;
}

int main()
{
    int arr[] = {0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1};
    int n = sizeof(arr) / sizeof(arr[0]);
    cout << findWater(arr, n);
    return 0;
}

// Time Complexity - ON(N)
// Space Complexity - O(N)