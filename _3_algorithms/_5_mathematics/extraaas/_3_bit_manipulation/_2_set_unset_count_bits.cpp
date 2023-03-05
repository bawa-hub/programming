#include <bits/stdc++.h>
using namespace std;

void printBinary(int num)
{
    for (int i = 10; i >= 0; --i)
    {
        cout << ((num >> i) & 1);
    }
    cout << endl;
}

int main()
{
    printBinary(9);
    int a = 9;
    int i = 3;
    if ((a & (1 << i)) != 0)
    {
        cout << "Set bit\n";
    }
    else
    {
        cout << "not set bit\n";
    }

    // bit set
    printBinary(a | (1 << 1));
    // bit unset
    printBinary(a & (~(1 << 3)));
    // toggle bit
    printBinary(a ^ (1 << 2));
    printBinary(a ^ (1 << 3));

    // count set bit
    int ct = 0;
    for (int i = 31; i >= 0; --i)
    {
        if ((a & (1 << i)) != 0)
        {
            ct++;
        }
    }
    cout << ct << endl;
    // or
    cout << __builtin_popcount(a) << endl;
}