#include <iostream>
#include <map>
#include <string>

using namespace std;

int main()
{
    map<char, int> m;
    m['a'] = 1;
    m['b'] = 2;

    cout << m.size();

    int count = 0;
    for (auto &x : m)
    {
        count += x.second;
    }
    cout << count;
}