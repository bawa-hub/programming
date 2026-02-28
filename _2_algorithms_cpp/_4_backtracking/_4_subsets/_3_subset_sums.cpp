// Q. https://practice.geeksforgeeks.org/problems/subset-sums2234/1

#include <bits/stdc++.h>
using namespace std;

class Solution
{
public:
    void func(int ind, int sum, vector<int> &arr, int N, vector<int> &sumSubset)
    {
        if (ind == N)
        {
            sumSubset.push_back(sum);
            return;
        }

        // pick element
        func(ind + 1, sum + arr[ind], arr, N, sumSubset);

        // not pick
        func(ind + 1, sum, arr, N, sumSubset);
    }

    vector<int> subsetSums(vector<int> arr, int N)
    {
        vector<int> sumSubset;
        func(0, 0, arr, N, sumSubset);
        sort(sumSubset.begin(), sumSubset.end());
        return sumSubset;
    }
};
// Time Complexity: O(2^n)+O(2^n log(2^n)). Each index has two ways. You can either pick it up or not pick it. So for n index time complexity for O(2^n) and for sorting it will take (2^n log(2^n)).
// Space Complexity: O(2^n) for storing subset sums, since 2^n subsets can be generated for an array of size n.