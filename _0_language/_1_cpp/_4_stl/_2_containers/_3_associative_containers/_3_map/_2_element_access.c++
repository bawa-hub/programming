#include <iostream>
#include <map>
#include <string>

int main()
{
    std::map<char, std::string> mymap;

    mymap['a'] = "an element";
    mymap['b'] = "another element";
    mymap['c'] = mymap['b'];

    // using []
    std::cout << "mymap['a'] is " << mymap['a'] << '\n';
    std::cout << "mymap['b'] is " << mymap['b'] << '\n';
    std::cout << "mymap['c'] is " << mymap['c'] << '\n';
    std::cout << "mymap['d'] is " << mymap['d'] << '\n';

    std::cout << "mymap now contains " << mymap.size() << " elements.\n";

    // using at()
    std::map<std::string, int> mymap1 = {
        {"alpha", 0},
        {"beta", 0},
        {"gamma", 0}};

    mymap1.at("alpha") = 10;
    mymap1.at("beta") = 20;
    mymap1.at("gamma") = 30;

    for (auto &x : mymap1)
    {
        std::cout << x.first << ": " << x.second << '\n';
    }

    return 0;
}