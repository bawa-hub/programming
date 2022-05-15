#include <iostream>
#include <bits/stdc++.h>
using namespace std;

int main()
{
    // your code goes here
    int t;
    cin >> t;
    while (t--)
    {
        long long p;
        int i, count = 0;
        cin >> p;
        for (i = 11; i >= 0; i--)
        {
            long currentpow = pow(2, i);
            while (p >= currentpow)
            {
                p = p - currentpow;
                count = count + 1;
            }
        }
        cout << count << endl;
    }
    return 0;
}
