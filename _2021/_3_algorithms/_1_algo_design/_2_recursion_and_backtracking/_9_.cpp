// https://www.geeksforgeeks.org/write-a-c-program-to-calculate-powxn/
#include <bits/stdc++.h>
using namespace std;

int power(int x, int n)
{
    if (n == 0)
        return 1;
    return power(x, n - 1) * x;
}

int main()
{
    int x, n;
    cin >> x >> n;
    cout << power(x, n);
}