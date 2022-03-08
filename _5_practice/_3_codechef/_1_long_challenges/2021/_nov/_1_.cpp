#include <bits/stdc++.h>
using namespace std;

int main()
{
    int t;
    cin >> t;
    while (t--)
    {
        int x, y, a, b, k;
        cin >> x >> y >> a >> b >> k;
        x += k * a;
        y += k * b;
        if (x < y)
            cout << "PETROL" << endl;
        else if (y < x)
            cout << "DIESEL" << endl;
        else
            cout << "SAME PRICE" << endl;
    }
}