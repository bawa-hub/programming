// https://leetcode.com/problems/maximum-subarray/
// https://takeuforward.org/data-structure/kadanes-algorithm-maximum-subarray-sum-in-an-array/

#include <bits/stdc++.h>
using namespace std;

// brute force
int maxSubArray(vector<int> &nums, vector<int> &subarray)
{
    int max_sum = INT_MIN;
    int n = nums.size();
    if (n == 1)
    {
        return nums[0];
    }
    int i, j;
    for (i = 0; i <= n - 1; i++)
    {
        for (j = i; j <= n - 1; j++)
        {
            int sum = 0;
            for (int k = i; k <= j; k++)
                sum = sum + nums[k];
            if (sum > max_sum)
            {
                subarray.clear();
                max_sum = sum;
                subarray.push_back(i);
                subarray.push_back(j);
            }
        }
    }
    return max_sum;
}
// Time Complexity: O(N^3)
// Space Complexity: O(1)

// better approach
int maxSubArray(vector<int> &nums, vector<int> &subarray)
{
    int max_sum = INT_MIN;
    for (int i = 0; i < nums.size(); i++)
    {
        int curr_sum = 0;
        for (int j = i; j < nums.size(); j++)
        {
            curr_sum += nums[j];
            if (curr_sum > max_sum)
            {
                subarray.clear();
                max_sum = curr_sum;
                subarray.push_back(i);
                subarray.push_back(j);
            }
        }
    }
    return max_sum;
}
// Time Complexity: O(N^2)
// Space Complexity: O(1)

// kadanes alog
// Kadane's algorithm runs one for loop over the array and at the beginning of each iteration, 
// if the current sum is negative, it will reset the current sum to zero. 
// This way, we ensure a one-pass and solve the problem in linear time.
long long maxSubarraySum(int arr[], int n)
{
    long long maxi = LONG_MIN; 
    long long sum = 0;

    for (int i = 0; i < n; i++)
    {
        sum += arr[i];
        if (sum > maxi)
            maxi = sum;

        // If sum < 0: discard the sum calculated
        if (sum < 0)
            sum = 0;
    }

    // To consider the sum of the empty subarray
    // uncomment the following check:
    // if (maxi < 0) maxi = 0;

    return maxi;
}
// Time Complexity: O(N)
// Space Complexity:O(1)

int main()
{
    vector<int> arr{-2, 1, -3, 4, -1, 2, 1, -5, 4};
    vector<int> subarray;
    int lon = maxSubArray(arr, subarray);
    cout << "The longest subarray with maximum sum is " << lon << endl;
    cout << "The subarray is " << endl;
    for (int i = subarray[0]; i <= subarray[1]; i++)
    {
        cout << arr[i] << " ";
    }
}
