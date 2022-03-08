#include <bits/stdc++.h>
using namespace std;

int main()
{
    int a[5][5];
    int moves;
    for (int i = 0; i < 5; i++)
    {
        for (int j = 0; j < 5; j++)
        {
            int temp;
            cin >> temp;
            if (temp == 1)
            {
                moves = abs(2 - i) + abs(2 - j);
                break;
            }
        }
    }
    cout << moves;
}