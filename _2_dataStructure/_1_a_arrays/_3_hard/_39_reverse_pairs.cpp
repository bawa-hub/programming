// https://leetcode.com/problems/reverse-pairs/

#include <bits/stdc++.h>
using namespace std;

// brute force
int reversePairs(vector<int> &arr)
{
    int Pairs = 0;
    for (int i = 0; i < arr.size(); i++)
    {
        for (int j = i + 1; j < arr.size(); j++)
        {
            if (arr[i] > 2 * arr[j])
                Pairs++;
        }
    }
    return Pairs;
}
//     Time Complexity: O (N^2) ( Nested Loops )
// Space Complexity:  O(1)

// optimal
int Merge(vector<int> &nums, int low, int mid, int high)
{
    int total = 0;
    int j = mid + 1;
    for (int i = low; i <= mid; i++)
    {
        while (j <= high && nums[i] > 2 LL * nums[j])
        {
            j++;
        }
        total += (j - (mid + 1));
    }

    vector<int> t;
    int left = low, right = mid + 1;

    while (left <= mid && right <= high)
    {

        if (nums[left] <= nums[right])
        {
            t.push_back(nums[left++]);
        }
        else
        {
            t.push_back(nums[right++]);
        }
    }

    while (left <= mid)
    {
        t.push_back(nums[left++]);
    }
    while (right <= high)
    {
        t.push_back(nums[right++]);
    }

    for (int i = low; i <= high; i++)
    {
        nums[i] = t[i - low];
    }
    return total;
}
int MergeSort(vector<int> &nums, int low, int high)
{

    if (low >= high)
        return 0;
    int mid = (low + high) / 2;
    int inv = MergeSort(nums, low, mid);
    inv += MergeSort(nums, mid + 1, high);
    inv += Merge(nums, low, mid, high);
    return inv;
}
int reversePairs(vector<int> &arr)
{
    return MergeSort(arr, 0, arr.size() - 1);
}
// Time Complexity : O( N log N ) + O (N) + O (N)
// Reason: O(N) – Merge operation, O(N) – counting operation ( at each iteration of i, j doesn’t start from 0 . Both of them move linearly )
// Space Complexity : O(N)
// Reason : O(N) – Temporary vector

int main()
{
    vector<int> arr{1, 3, 2, 3, 1};
    cout << "The Total Reverse Pairs are " << reversePairs(arr);
}

/************************* for leetcode */

class Solution
{
public:
    int inversions = 0;
    int find_index_where_just_greater(int left, int right, vector<int> &arr, int target)
    {
        while (left <= right)
        {
            int mid = (left + right) / 2;
            long long x = arr[mid];
            x *= 2;
            if (x >= target)
            {
                right = mid - 1;
            }
            else
            {
                left = mid + 1;
            }
        }
        return (left - 1);
    }
    void merge(vector<int> &arr, int l, int m, int r)
    {
        vector<int> Arr(r - l + 1);
        int ptr1 = l, ptr2 = m + 1, ptr3 = 0;
        while (ptr3 <= (r - l))
        {
            if (ptr1 == (m + 1))
            {
                Arr[ptr3] = arr[ptr2];
                ptr2++;
                ptr3++;
                continue;
            }
            if (ptr2 == (r + 1))
            {
                int idx = find_index_where_just_greater(m + 1, r, arr, arr[ptr1]);
                inversions += idx - m;
                Arr[ptr3] = arr[ptr1];
                ptr1++;
                ptr3++;
                continue;
            }
            if (arr[ptr1] <= arr[ptr2])
            {
                int idx = find_index_where_just_greater(m + 1, r, arr, arr[ptr1]);
                inversions += idx - m;
                Arr[ptr3] = arr[ptr1];
                ptr1++;
                ptr3++;
            }
            else
            {
                Arr[ptr3] = arr[ptr2];
                ptr2++;
                ptr3++;
            }
        }
        for (int i = l; i <= r; i++)
        {
            arr[i] = Arr[i - l];
        }
        return;
    }
    void mergeSort(vector<int> &arr, int l, int r)
    {
        if (l == r)
        {
            return;
        }
        int mid = (l + r) / 2;
        mergeSort(arr, l, mid);
        mergeSort(arr, mid + 1, r);
        merge(arr, l, mid, r);
        return;
    }
    int reversePairs(vector<int> &nums)
    {
        mergeSort(nums, 0, nums.size() - 1);
        return inversions;
    }
};

// using two pointer
class Solution
{
public:
    int inversions = 0;
    void merge(vector<int> &arr, int l, int m, int r)
    {
        vector<int> Arr(r - l + 1);
        int ptr1 = l, ptr2 = m + 1, ptr3 = 0, idx = m + 1;
        while (ptr3 <= (r - l))
        {
            if (ptr1 == (m + 1))
            {
                Arr[ptr3] = arr[ptr2];
                ptr2++;
                ptr3++;
                continue;
            }
            if (ptr2 == (r + 1))
            {
                while (idx <= r && arr[ptr1] > (2 * ((long long)arr[idx])))
                {
                    idx++;
                }
                inversions += idx - m - 1;
                Arr[ptr3] = arr[ptr1];
                ptr1++;
                ptr3++;
                continue;
            }
            if (arr[ptr1] <= arr[ptr2])
            {
                while (idx <= r && arr[ptr1] > (2 * ((long long)arr[idx])))
                {
                    idx++;
                }
                inversions += idx - m - 1;
                Arr[ptr3] = arr[ptr1];
                ptr1++;
                ptr3++;
            }
            else
            {
                Arr[ptr3] = arr[ptr2];
                ptr2++;
                ptr3++;
            }
        }
        for (int i = l; i <= r; i++)
        {
            arr[i] = Arr[i - l];
        }
        return;
    }
    void mergeSort(vector<int> &arr, int l, int r)
    {
        if (l == r)
        {
            return;
        }
        int mid = (l + r) / 2;
        mergeSort(arr, l, mid);
        mergeSort(arr, mid + 1, r);
        merge(arr, l, mid, r);
        return;
    }
    int reversePairs(vector<int> &nums)
    {
        mergeSort(nums, 0, nums.size() - 1);
        return inversions;
    }
};

// Time complexity:O(N*(logN)^2)
// Space complexity:O(N)
