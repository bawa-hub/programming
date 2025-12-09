#include <iostream>
#include <map>

using namespace std;

int main()
{
    map<char, int> mymap;

    mymap['b'] = 100;
    mymap['a'] = 200;
    mymap['c'] = 300;

    // begin(), end()
    for (map<char, int>::iterator it = mymap.begin(); it != mymap.end(); ++it)
        cout << it->first << " => " << it->second << '\n';

    //rbegin(), rend()
    map<char, int>::reverse_iterator rit;
    for (rit = mymap.rbegin(); rit != mymap.rend(); ++rit)
        cout << rit->first << " => " << rit->second << '\n';

    //cbegin(), cend()
    cout << "mymap contains:";
    for (auto it = mymap.cbegin(); it != mymap.cend(); ++it)
        cout << " [" << (*it).first << ':' << (*it).second << ']';
    cout << '\n';

    // crbegin(), crend()
    cout << "mymap backwards:";
    for (auto rit = mymap.crbegin(); rit != mymap.crend(); ++rit)
        cout << " [" << rit->first << ':' << rit->second << ']';
    cout << '\n';

    return 0;
}