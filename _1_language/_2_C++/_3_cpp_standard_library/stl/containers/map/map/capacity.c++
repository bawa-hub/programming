#include <iostream>
#include <map>

int main()
{
    std::map<char, int> mymap;

    mymap['a'] = 10;
    mymap['b'] = 20;
    mymap['c'] = 30;

    // size()
    std::cout << "mymap.size() is " << mymap.size() << '\n';

    // empty()
    while (!mymap.empty())
    {
        std::cout << mymap.begin()->first << " => " << mymap.begin()->second << '\n';
        mymap.erase(mymap.begin());
    }

    // max_size()
    int i;
    std::map<int, int> mymap1;

    if (mymap1.max_size() > 1000)
    {
        for (i = 0; i < 1000; i++)
            mymap1[i] = 0;
        std::cout << "The map contains 1000 elements.\n";
    }
    else
        std::cout << "The map could not hold 1000 elements.\n";

    return 0;
}