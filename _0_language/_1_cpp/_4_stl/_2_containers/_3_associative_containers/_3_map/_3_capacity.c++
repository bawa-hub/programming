#include <iostream>
#include <map>

using namespace std;

int main()
{
    map<char, int> mymap;

    mymap['a'] = 10;
    mymap['b'] = 20;
    mymap['c'] = 30;

    // size()
    cout << "mymap.size() is " << mymap.size() << '\n';
     cout << "mymap.max_size() is " << mymap.max_size() << '\n';

    // empty()
    while (!mymap.empty())
    {
        cout << mymap.begin()->first << " => " << mymap.begin()->second << '\n';
        mymap.erase(mymap.begin());
    }

    // max_size()
    int i;
    map<int, int> mymap1;

    if (mymap1.max_size() > 1000)
    {
        for (i = 0; i < 1000; i++)
            mymap1[i] = 0;
        cout << "The map contains 1000 elements.\n";
    }
    else
        cout << "The map could not hold 1000 elements.\n";

    return 0;
}