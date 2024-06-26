// https://leetcode.com/problems/sliding-window-maximum/

#include <bits/stdc++.h>

using namespace std;

// brute force
void GetMax(vector<int> nums, int l, int r, vector<int> &arr)
{
    int i, maxi = INT_MIN;
    for (i = l; i <= r; i++)
        maxi = max(maxi, nums[i]);
    arr.push_back(maxi);
}
vector<int> maxSlidingWindow(vector<int> &nums, int k)
{
    int left = 0, right = 0;
    int i, j;
    vector<int> arr;
    while (right < k - 1)
    {
        right++;
    }
    while (right < nums.size())
    {
        GetMax(nums, left, right, arr);
        left++;
        right++;
    }
    return arr;
}
// Time Complexity: O(N^2)
// Reason: One loop for traversing and another to findMax
// Space Complexity: O(K)
// Reason: No.of windows

// optimized
vector<int> maxSlidingWindow(vector<int> &nums, int k)
{
    deque<int> dq;
    vector<int> ans;
    for (int i = 0; i < nums.size(); i++)
    {
        if (!dq.empty() && dq.front() == i - k)
            dq.pop_front();

        while (!dq.empty() && nums[dq.back()] < nums[i])
            dq.pop_back();

        dq.push_back(i);
        if (i >= k - 1)
            ans.push_back(nums[dq.front()]);
    }
    return ans;
}
// Time Complexity: O(N)
// Space Complexity: O(K)

int main()
{
    int i, j, n, k = 3, x;
    vector<int> arr{
        4, 0,
        -1,
        3,
        5,
        3,
        6,
        8};
    vector<int> ans;
    ans = maxSlidingWindow(arr, k);
    cout << "Maximum element in every " << k << " window " << endl;
    for (i = 0; i < ans.size(); i++)
        cout << ans[i] << "  ";
    return 0;
}