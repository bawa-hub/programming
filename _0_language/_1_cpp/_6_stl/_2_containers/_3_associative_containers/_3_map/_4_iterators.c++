#include <iostream>
#include <map>

int main()
{
    std::map<char, int> mymap;

    mymap['b'] = 100;
    mymap['a'] = 200;
    mymap['c'] = 300;

    // begin(), end()
    for (std::map<char, int>::iterator it = mymap.begin(); it != mymap.end(); ++it)
        std::cout << it->first << " => " << it->second << '\n';

    //rbegin(), rend()
    std::map<char, int>::reverse_iterator rit;
    for (rit = mymap.rbegin(); rit != mymap.rend(); ++rit)
        std::cout << rit->first << " => " << rit->second << '\n';

    //cbegin(), cend()
    std::cout << "mymap contains:";
    for (auto it = mymap.cbegin(); it != mymap.cend(); ++it)
        std::cout << " [" << (*it).first << ':' << (*it).second << ']';
    std::cout << '\n';

    // crbegin(), crend()
    std::cout << "mymap backwards:";
    for (auto rit = mymap.crbegin(); rit != mymap.crend(); ++rit)
        std::cout << " [" << rit->first << ':' << rit->second << ']';
    std::cout << '\n';

    return 0;
}