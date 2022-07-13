#include <iostream>
#include <map>

int main()
{
    std::map<char, int> mymap;

    // key_comp()
    std::map<char, int>::key_compare mycomp = mymap.key_comp();

    mymap['a'] = 100;
    mymap['b'] = 200;
    mymap['c'] = 300;

    std::cout << "mymap contains:\n";

    char highest = mymap.rbegin()->first; // key value of last element

    std::map<char, int>::iterator it = mymap.begin();
    do
    {
        std::cout << it->first << " => " << it->second << '\n';
    } while (mycomp((*it++).first, highest));

    std::cout << '\n';

    // value_comp()
    std::cout << "mymap contains:\n";

    std::pair<char, int> highest1 = *mymap.rbegin(); // last element

    std::map<char, int>::iterator it1 = mymap.begin();
    do
    {
        std::cout << it1->first << " => " << it1->second << '\n';
    } while (mymap.value_comp()(*it1++, highest1));

    return 0;
}