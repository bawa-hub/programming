// Q. https://codeforces.com/problemset/problem/1/A
#include <bits/stdc++.h>
using namespace std;

int main()
{
    long double n, m, a;
    cin >> n >> m >> a;

    cout << fixed;
    cout << setprecision(0);
    cout << (ceil(n / a) * ceil(m / a));
}