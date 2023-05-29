// https://leetcode.com/problems/majority-element-ii/

#include <bits/stdc++.h>
using namespace std;

// brute force
vector<int> majorityElement(int arr[], int n)
{
    vector<int> ans;
    for (int i = 0; i < n; i++)
    {
        int cnt = 1;
        for (int j = i + 1; j < n; j++)
        {
            if (arr[j] == arr[i])
                cnt++;
        }

        if (cnt > (n / 3))
            ans.push_back(arr[i]);
    }

    return ans;
}

// Time Complexity: O(n^2)
// Space Complexity : O(1)

// better solution
vector<int> majorityElement(int arr[], int n)
{
    unordered_map<int, int> mp;
    vector<int> ans;

    for (int i = 0; i < n; i++)
    {
        mp[arr[i]]++;
    }

    for (auto x : mp)
    {
        if (x.second > (n / 3))
            ans.push_back(x.first);
    }

    return ans;
}
// Time Complexity: O(n)
// Space Complexity : O(n)

// Optimal Solution (Extended Boyer Moore’s Voting Algorithm)
vector<int> majorityElement(int nums[], int n)
{
    int sz = n;
    int num1 = -1, num2 = -1, count1 = 0, count2 = 0, i;
    for (i = 0; i < sz; i++)
    {
        if (nums[i] == num1)
            count1++;
        else if (nums[i] == num2)
            count2++;
        else if (count1 == 0)
        {
            num1 = nums[i];
            count1 = 1;
        }
        else if (count2 == 0)
        {
            num2 = nums[i];
            count2 = 1;
        }
        else
        {
            count1--;
            count2--;
        }
    }
    
    vector<int> ans;
    count1 = count2 = 0;
    for (i = 0; i < sz; i++)
    {
        if (nums[i] == num1)
            count1++;
        else if (nums[i] == num2)
            count2++;
    }
    if (count1 > sz / 3)
        ans.push_back(num1);
    if (count2 > sz / 3)
        ans.push_back(num2);
    return ans;
}
// Time Complexity: O(n)
// Space Complexity : O(1)

int main()
{
    int arr[] = {1, 2, 2, 3, 2};
    vector<int> majority;
    majority = majorityElement(arr, 5);
    cout << "The majority element is ";

    for (auto it : majority)
    {
        cout << it << " ";
    }
}
