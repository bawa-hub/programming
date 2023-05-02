// https://leetcode.com/problems/find-the-duplicate-number/

#include <bits/stdc++.h>

using namespace std;
int findDuplicate(vector<int> &arr)
{
    int n = arr.size();
    int freq[n + 1] = {
        0};
    for (int i = 0; i < n; i++)
    {
        if (freq[arr[i]] == 0)
        {
            freq[arr[i]] += 1;
        }
        else
        {
            return arr[i];
        }
    }
    return 0;
}
int main()
{
    vector<int> arr;
    arr = {2, 1, 1};
    cout << "The duplicate element is " << findDuplicate(arr) << endl;
}