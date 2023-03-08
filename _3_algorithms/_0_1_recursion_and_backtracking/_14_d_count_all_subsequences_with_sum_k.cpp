// https://practice.geeksforgeeks.org/problems/perfect-sum-problem5633/1

// count the subsequence with given sum
#include <bits/stdc++.h>
using namespace std;

int f(int idx, int s, int sum, int arr[], int n)
{
    // this base condition iff array contains only positive nubmer
    if (s > sum)
        return 0;

    if (idx == n)
    {
        if (s == sum)
            return 1;
        else
            return 0;
    }

    s += arr[idx];

    // pick
    int l = f(idx + 1, s, sum, arr, n);
    s -= arr[idx];

    // not pick
    int r = f(idx + 1, s, sum, arr, n);

    return l + r;
}

int main()
{
    int arr[] = {1, 2, 1};
    int n = 3;
    int sum = 2;
    cout << f(0, 0, sum, arr, n);
    return 0;
}