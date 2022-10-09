// Brute force approach
#include <bits/stdc++.h>
using namespace std;

int main()
{
    int n;
    cin >> n;

    int a[n];
    for (int i = 0; i < n; i++)
        cin >> a[i];

    int total = 0;

    for (int i = 1; i < n - 1; i++)
    {
        int left_max = -1;
        int right_max = -1;

        for (int j = i - 1; j >= 0; j--)
            left_max = max(left_max, a[j]);

        for (int k = i + 1; k < n; k++)
            right_max = max(right_max, a[k]);

        int smaller = min(left_max, right_max);

        if (smaller > a[i])
            total += (smaller - a[i]);
    }

    cout << "Total logging: " << total << " units.";
}

// Time Complexity - ON(N^2)
// Space Complexity - O(1)