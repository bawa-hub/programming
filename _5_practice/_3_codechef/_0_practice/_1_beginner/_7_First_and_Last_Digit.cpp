#include <bits/stdc++.h>
using namespace std;

int main()
{
    int t;
    cin >> t;
    while (t--)
    {
        int n;
        cin >> n;
        int ld = n % 10;
        while (n / 10 != 0)
        {
            n = n / 10;
        }
        cout << n + ld << endl;
    }
}