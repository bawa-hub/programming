#include <bits/stdc++.h>
using namespace std;

int main()
{
    int num;
    cin >> num;
    int p1 = 0, p2 = 0, win, max_diff = 0, diff, a, b;
    while (num--)
    {
        cin >> a >> b;

        p1 += a;
        p2 += b;
        diff = abs(p1 - p2);

        if (diff > max_diff)
        {
            max_diff = diff;
            if (p1 >= p2)
                win = 1;
            else
                win = 2;
        }
    }
    cout << win << " " << max_diff;
    return 0;
}