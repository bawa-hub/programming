// https://practice.geeksforgeeks.org/problems/rotation4723/1?
// https://www.geeksforgeeks.org/find-rotation-count-rotated-sorted-array/

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