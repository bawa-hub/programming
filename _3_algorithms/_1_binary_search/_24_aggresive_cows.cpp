// https://practice.geeksforgeeks.org/problems/aggressive-cows/1

#include <bits/stdc++.h>

using namespace std;

// BRUTE FORCE
bool isCompatible(vector<int> inp, int dist, int cows)
{
    int n = inp.size();
    int k = inp[0];
    cows--;
    for (int i = 1; i < n; i++)
    {
        if (inp[i] - k >= dist)
        {
            cows--;
            if (!cows)
            {
                return true;
            }
            k = inp[i];
        }
    }
    return false;
}
int main()
{
    int n = 5, m = 3;
    vector<int> inp{1, 2, 8, 4, 9};
    sort(inp.begin(), inp.end());
    int maxD = inp[n - 1] - inp[0];
    int ans = INT_MIN;
    for (int d = 1; d <= maxD; d++)
    {
        bool possible = isCompatible(inp, d, m);
        if (possible)
        {
            ans = max(ans, d);
        }
    }
    cout << "The largest minimum distance is " << ans << '\n';
}

// Time complexity: O(n* m)
// Space Complexity: O(1)

// binary search
bool isPossible(int a[], int n, int cows, int minDist)
{
    int cntCows = 1;
    int lastPlacedCow = a[0];
    for (int i = 1; i < n; i++)
    {
        if (a[i] - lastPlacedCow >= minDist)
        {
            cntCows++;
            lastPlacedCow = a[i];
        }
    }
    if (cntCows >= cows)
        return true;
    return false;
}
int main()
{
    int n = 5, cows = 3;
    int a[] = {1, 2, 8, 4, 9};
    sort(a, a + n);

    int low = 1, high = a[n - 1] - a[0];

    while (low <= high)
    {
        int mid = (low + high) >> 1;

        if (isPossible(a, n, cows, mid))
        {
            low = mid + 1;
        }
        else
        {
            high = mid - 1;
        }
    }
    cout << "The largest minimum distance is " << high << endl;

    return 0;
}
//     Time Complexity: O(N*log(M)).
// Reason: For binary search in a space of M, O(log(M))  and for each search, we iterate over max N stalls to check. O(N).
// Space Complexity: O(1)