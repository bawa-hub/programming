// https://leetcode.com/problems/binary-search/

using namespace std;

// recursively

int binarySearch(int arr[], int start, int end, int k)
{

    if (start > end)
        return -1;

    int mid = start + (end - start) / 2;
    if (k == arr[mid])
        return mid;
    else if (k < arr[mid])
        return binarySearch(arr, start, mid - 1, k);
    else
        return binarySearch(arr, mid + 1, end, k);
}
// Time complexity: O(log n)
// Space complexity: O(logn) for auxiliary space

// iteratively (basic implementation)
int binarySearch(int arr[], int target, int n)
{
    int l = 0, r = n - 1, mid;
    while (l <= r)
    {
        mid = l + (r - l) / 2;
        if (arr[mid] == target)
            return mid;
        if (arr[mid] < target)
            l = mid + 1;
        else
            r = mid - 1;
    }
    return -1;
}
// Time complexity: O(log n)
// Space complexity : O(1)
