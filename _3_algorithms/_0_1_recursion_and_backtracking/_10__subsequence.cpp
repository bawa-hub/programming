// print all the subsequence of the given array
#include <bits/stdc++.h>
using namespace std;

void f(int idx, vector<int> &ds, int arr[], int n)
{
    if (idx == n)
    {
        cout << "{ ";
        for (auto it : ds)
        {
            cout << it << " ";
        }
        cout << "}";
        cout << endl;
        return;
    }

    // not pick, or not take condition, this element is not added to your subsequence
    f(idx + 1, ds, arr, n);

    // take or pick the particular index into the subsequence
    ds.push_back(arr[idx]);
    f(idx + 1, ds, arr, n);
    ds.pop_back();
}
// Time complexity - O(n*2^n)
// Space complexity - O(n)

int main()
{
    int arr[] = {3, 1, 2};
    int n = 3;
    vector<int> ds;
    f(0, ds, arr, n);

    return 0;
}