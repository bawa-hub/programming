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