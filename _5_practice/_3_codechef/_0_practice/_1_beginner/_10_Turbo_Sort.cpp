#include <bits/stdc++.h>
using namespace std;

int main()
{
    int t;
    cin >> t;
    int list[t] = {0};
    for (int i = 0; i < t; i++)
    {
        cin >> list[i];
    }
    int len = sizeof(list) / sizeof(list[0]);
    sort(list, list + len);
    for (int j = 0; j < t; j++)
    {
        cout << list[j] << endl;
    }
}