#include <iostream>
using namespace std;
int count(int n)
{
    int a = 0, b = 0, c = 0, d = 0, e = 0, f = 0;
    if (n / 100 == 0)
        a = 0;
    else
    {
        a = n / 100;
        n = n - (a * 100);
    }
    if (n / 50 == 0)
        b = 0;
    else
    {
        b = n / 50;
        n = n - (b * 50);
    }
    if (n / 10 == 0)
        c = 0;
    else
    {
        c = n / 10;
        n = n - (c * 10);
    }
    if (n / 5 == 0)
        d = 0;
    else
    {
        d = n / 5;
        n = n - (d * 5);
    }
    if (n / 2 == 0)
        e = 0;
    else
    {
        e = n / 2;
        n = n - (e * 2);
    }
    if (n / 1 == 0)
        f = 0;
    else
    {
        f = n / 1;
        n = n - (f * 1);
    }
    return (a + b + c + d + e + f);
}

int main()
{
    int t;
    cin >> t;
    while (t--)
    {
        int n;
        cin >> n;
        cout << count(n) << endl;
    }
    return 0;
}
