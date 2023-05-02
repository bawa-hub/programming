/**
 * Given array a of N integers. Given Q queries and in each query L and R
 * is given.Print sum of array elements from index L to R(L, R included)
 * 
 * Constraints
 * 1 <= N <= 10^5
 * 1 <= a[i] <= 10^9
 * 1 <= Q <= 10^5
 * 1 <= L, R <= N
 * ***/

#include <bits/stdc++.h>
using namespace std;

const int N = 1e5 + 10;
int a[N];
int pf[N];

int main()
{

    // naive method
    // int n;
    // cin >> n;
    // for (int i = 1; i <= n; ++i)
    // {
    //     cin >> a[i];
    // }
    // int q;
    // cin >> q;
    // while (q--)
    // {
    //     int l, r;
    //     cin >> l >> r;
    //     long long sum = 0;
    //     for (int i = l; i <= r; ++i)
    //     {
    //         sum += a[i];
    //     }
    //     cout << sum << endl;
    // }
    // O(N)+O(Q*N)=O(N^2)=10^10

    // precomputation usin prefix sum
    int n;
    cin >> n;
    for (int i = 1; i <= n; ++i)
    {
        cin >> a[i];
        pf[i] = pf[i - 1] + a[i];
    }
    int q;
    cin >> q;
    while (q--)
    {
        int l, r;
        cin >> l >> r;
        cout << pf[r] - pf[l - 1] << endl;
    }
    // O(N)+O(Q)=O(N)=10^5
}