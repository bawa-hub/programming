// https://leetcode.com/problems/find-the-smallest-divisor-given-a-threshold/

// brute force
int smallestDivisor(vector<int>& arr, int limit) {
    int n = arr.size(); //size of array.
    //Find the maximum element:
    int maxi = *max_element(arr.begin(), arr.end());

    //Find the smallest divisor:
    for (int d = 1; d <= maxi; d++) {
        //Find the summation result:
        int sum = 0;
        for (int i = 0; i < n; i++) {
            sum += ceil((double)(arr[i]) / (double)(d));
        }
        if (sum <= limit)
            return d;
    }
    return -1;
}
// Time Complexity: O(max(arr[])*N), where max(arr[]) = maximum element in the array, N = size of the array.
// Reason: We are using nested loops. The outer loop runs from 1 to max(arr[]) and the inner loop runs for N times.
// Space Complexity: O(1) as we are not using any extra space to solve this problem.

class Solution
{
public:
    int sumByD(vector<int> &nums, int div)
    {
        int n = nums.size(), sum = 0;
        for (int i = 0; i < n; i++)
        {
            sum += (ceil((double)nums[i] / (double)div));
        }
        return sum;
    }
    int smallestDivisor(vector<int> &nums, int threshold)
    {
        int l = 1, r = *max_element(nums.begin(), nums.end());
        while (l <= r)
        {
            int mid = l + (r - l) / 2;
            if (sumByD(nums, mid) <= threshold)
            {
                r = mid - 1;
            }
            else
            {
                l = mid + 1;
            }
        }
        return l;
    }
};
// Time Complexity: O(log(max(arr[]))*N), where max(arr[]) = maximum element in the array, N = size of the array.
// Reason: We are applying binary search on our answers that are in the range of [1, max(arr[])]. For every possible divisor ‘mid’, we call the sumByD() function. Inside that function, we are traversing the entire array, which results in O(N).
// Space Complexity: O(1) as we are not using any extra space to solve this problem.