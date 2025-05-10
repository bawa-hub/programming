// https://leetcode.com/problems/house-robber-ii/
// https://www.codingninjas.com/codestudio/problems/house-robber_839733
// https://practice.geeksforgeeks.org/problems/stickler-theif-1587115621/1?

// same as maximum sum of non adjacent element with slight change (circular adjacent allow)

#include <bits/stdc++.h>
using namespace std;

class Solution
{
public:
    int rob(vector<int> &nums)
    {
        int n = nums.size();
        vector<int> arr1;
        vector<int> arr2;

        if (n == 1)
            return nums[0];

        for (int i = 0; i < n; i++)
        {

            if (i != 0)
                arr1.push_back(nums[i]);
            if (i != n - 1)
                arr2.push_back(nums[i]);
        }

        int ans1 = solve(arr1);
        int ans2 = solve(arr2);

        return max(ans1, ans2);
    }

    int solve(vector<int> &arr)
    {
        int n = arr.size();
        int prev = arr[0];
        int prev2 = 0;

        for (int i = 1; i < n; i++)
        {
            int pick = arr[i];
            if (i > 1)
                pick += prev2;
            int nonPick = 0 + prev;

            int cur_i = max(pick, nonPick);
            prev2 = prev;
            prev = cur_i;
        }
        return prev;
    }
};

int main()
{

    vector<int> arr{1, 5, 1, 2, 6};
    int n = arr.size();
    cout << robStreet(n, arr);
}

// Time Complexity: O(N )
// Reason: We are running a simple iterative loop, two times. Therefore total time complexity will be O(N) + O(N) ≈ O(N)
// Space Complexity: O(1)
// Reason: We are not using extra space.