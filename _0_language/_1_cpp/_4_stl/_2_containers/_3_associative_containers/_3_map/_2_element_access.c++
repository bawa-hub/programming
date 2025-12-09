#include <iostream>
#include <map>
#include <string>

using namespace std;

int main()
{
    map<char, string> mymap;

    mymap['a'] = "an element";
    mymap['b'] = "another element";
    mymap['c'] = mymap['b'];

    // using []
    cout << "mymap['a'] is " << mymap['a'] << '\n';
    cout << "mymap['b'] is " << mymap['b'] << '\n';
    cout << "mymap['c'] is " << mymap['c'] << '\n';
    cout << "mymap['d'] is " << mymap['d'] << '\n';

    cout << "mymap now contains " << mymap.size() << " elements.\n"; // operator[] on std::map inserts the key if it does not already exist.

    // using at()
    map<string, int> mymap1 = {
        {"alpha", 0},
        {"beta", 0},
        {"gamma", 0}};

    mymap1.at("alpha") = 10;
    mymap1.at("beta") = 20;
    mymap1.at("gamma") = 30;

    for (auto &x : mymap1)
    {
        cout << x.first << ": " << x.second << '\n';
    }

    return 0;
}


// ✅ What does auto &x do? (WITH &)
// x becomes a reference to each element inside the map.
// No copying → more efficient
// Memory efficient
// Modifying x WILL modify the map