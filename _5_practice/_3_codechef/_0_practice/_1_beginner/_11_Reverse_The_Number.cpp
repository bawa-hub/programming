#include <bits/stdc++.h>
using namespace std;

int rev(int n)
{
    string s = to_string(n);
    reverse(s.begin(), s.end());
    return stoi(s);
}

int main()
{
    int t;
    cin >> t;
    while (t--)
    {
        int n;
        cin >> n;
        cout << rev(n) << endl;
    }
}