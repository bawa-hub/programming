#include <bits/stdc++.h>
using namespace std;

int main()
{
    int t;
    cin >> t;
    while (t--)
    {
        long long n;
        cin >> n;
        int ct = 0;
        while (n / 10 != 0)
        {
            if (n % 10 == 4)
                ct++;
            n /= 10;
        }
        if (n == 4)
            ct++;
        cout << ct << endl;
    }
}