#include <bits/stdc++.h>
using namespace std;

int main()
{
    int n, k;
    cin >> n >> k;
    int rest_time = 240 - k;
    int i = 1;
    int ct = 0;
    while (rest_time > 0)
    {
        if (rest_time >= 5 * i)
            ct++;
        rest_time -= 5 * i;
        i++;
        if (ct >= n)
            break;
    }
    cout << ct;
}