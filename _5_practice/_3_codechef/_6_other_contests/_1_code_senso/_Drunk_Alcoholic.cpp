// https://www.codechef.com/CSNS21C/problems/DRUNKALK
#include <bits/stdc++.h>
using namespace std;

int main()
{
    int t;
    cin >> t;
    while (t--)
    {
        int k;
        cin >> k;
        int x = 0;
        for (int i = 1; i <= k; i++)
        {
            if (i % 2 == 0)
            {
                x -= 1;
            }
            else
            {
                x += 3;
            }
        }
        cout << x << endl;
    }
}