// print the subsequence with given sum
#include <bits/stdc++.h>
using namespace std;

void f(int idx, vector<int> &ds, int s, int sum, int arr[], int n)
{
    if (idx == n)
    {
        if (s == sum)
        {
            for (auto it : ds)
                cout << it << " ";
            cout << endl;
        }
        return;
    }

    // pick
    ds.push_back(arr[idx]);
    s += arr[idx];
    f(idx + 1, ds, s, sum, arr, n);
    s -= arr[idx];
    ds.pop_back();

    // not pick
    f(idx + 1, ds, s, sum, arr, n);
}

int main()
{
    int arr[] = {1, 2, 1};
    int n = 3;
    int sum = 2;
    vector<int> ds;
    f(0, ds, 0, sum, arr, n);
    return 0;
}