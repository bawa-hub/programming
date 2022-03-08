#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;

int main()
{
    vector<int> v = {
        2,
        3,
        5,
        1,
        7,
    };

    // sorting algorithm
    sort(v.begin(), v.end());

    for (int i : v)
    {
        cout << i << " ";
    }
    return 0;
}