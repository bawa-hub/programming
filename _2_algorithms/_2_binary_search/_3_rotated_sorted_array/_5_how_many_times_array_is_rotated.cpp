// https://practice.geeksforgeeks.org/problems/rotation4723/1?
// https://www.geeksforgeeks.org/find-rotation-count-rotated-sorted-array/


// brute force
int findKRotation(vector<int> &arr) {
    int n = arr.size(); //size of array.
    int ans = INT_MAX, index = -1;
    for (int i = 0; i < n; i++) {
        if (arr[i] < ans) {
            ans = arr[i];
            index = i;
        }
    }
    return index;
}

// Time Complexity: O(N), N = size of the given array.
// Reason: We have to iterate through the entire array to check if the target is present in the array.
// Space Complexity: O(1)
// Reason: We have not used any extra data structures, this makes space complexity, even in the worst case as O(1).

// same as find minimum in rotated sorted array
int findKRotation(int arr[], int n)
{
    int left = 0, right = n - 1, ans = INT_MAX, idx = -1;

    while (left <= right)
    {
        // search space is already sorted , return first element as answ
        if (arr[left] < arr[right])
        {
            if (arr[left] < ans)
            {
                idx = left;
                ans = arr[left];
            }
            break;
        }
        int mid = left + (right - left) / 2;

        if (arr[left] <= arr[mid])
        {
            if (arr[left] < ans)
            {
                idx = left;
                ans = arr[left];
            }
            left = mid + 1;
        }
        else
        {
            if (arr[mid] < ans)
            {
                idx = mid;
                ans = arr[mid];
            }
            right = mid - 1;
        }
    }
    return idx;
}
// Time Complexity: O(logN), N = size of the given array.
// Reason: We are basically using binary search to find the minimum.

// Space Complexity: O(1)
// Reason: We have not used any extra data structures, this makes space complexity, even in the worst case as O(1).