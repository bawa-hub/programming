#include <bits/stdc++.h>
using namespace std;

int main()
{
    float n;
    cin >> n;
    float res = 0;
    for (int i = 0; i < n; i++)
    {
        float temp;
        cin >> temp;
        res += temp;
    }
    cout << res / n;
}