#include <bits/stdc++.h>
using namespace std;

int main()
{
    int t;
    cin >> t;
    while (t--)
    {
        cout << endl;
        int n, k;
        cin >> n >> k;
        long long trains[n];
        for (int i = 0; i < n; i++)
        {
            cin >> trains[i];
        }
        while (k--)
        {
            long long a, b;
            cin >> a >> b;
            int first = INT_MAX, second = INT_MIN;
            for (int i = 0; i < n; i++)
            {
                if (trains[i] == a)
                    first = min(first, i);
                if (trains[i] == b)
                    second = max(second, i);
            }
            if (first < second)
            {
                cout << "YES" << endl;
            }
            else
                cout << "NO" << endl;
        }
    }
}