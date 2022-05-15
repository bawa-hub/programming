#include <bits/stdc++.h>
using namespace std;

int main()
{
    int t;
    cin >> t;
    while (t--)
    {
        double s, hra, da;
        cin >> s;
        if (s < 1500)
        {
            hra = 0.1 * s;
            da = 0.9 * s;
        }
        else
        {
            hra = 500;
            da = 0.98 * s;
        }
        double gross = s + hra + da;
        cout << fixed << setprecision(2) << gross << endl;
    }
}