// https://leetcode.com/problems/binary-search/

using namespace std;

// recursively
int binarySearch(int array[], int x, int low, int high)
{
    if (high >= low)
    {
        int mid = low + (high - low) / 2; // this calculation is used to overcome int overflow
        if (array[mid] == x) return mid;
        if (array[mid] > x) return binarySearch(array, x, low, mid - 1);
        else return binarySearch(array, x, mid + 1, high);
    }

    return -1;
}
// Time complexity: O(log n)
// Space complexity: O(logn) for auxiliary space

// iteratively (basic implementation)
int binarySearch(int arr[], int target, int n)
{
    int l = 0, r = n-1, mid;
    while (l <= r)
    {
        mid = l+(r-l)/2;
        if (arr[mid] == target)  return mid;
        if (arr[mid] < target) l = mid + 1;
        else r = mid - 1;
    }
    return -1;
}
// Time complexity: O(log n)
// Space complexity : O(1)

// implementation 2 (also called lower bound )
// for problems like: find first value >= x;
int binarySearch(int arr[], int target, int n)
{
    int l = 0, r = n-1, mid, ans = -1;
    while (l <= r)
    {
        mid = l+(r-l)/2;
        // this is condition part
        if (arr[mid] >= target)  {
            ans = mid;
            r = mid - 1; // a/c to question
        } else {
            l = mid  + 1;
        }
        
    }
    return ans;
}

