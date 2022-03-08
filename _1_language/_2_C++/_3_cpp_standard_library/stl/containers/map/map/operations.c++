#include <iostream>
#include <map>

int main()
{
    // find()
    std::map<char, int> mymap;
    std::map<char, int>::iterator it;

    mymap['a'] = 50;
    mymap['b'] = 100;
    mymap['c'] = 150;
    mymap['d'] = 200;

    it = mymap.find('b');
    if (it != mymap.end())
        mymap.erase(it);

    // print content:
    std::cout << "elements in mymap:" << '\n';
    std::cout << "a => " << mymap.find('a')->second << '\n';
    std::cout << "c => " << mymap.find('c')->second << '\n';
    std::cout << "d => " << mymap.find('d')->second << '\n';

    // count()
    std::map<char, int> mymap1;
    char c;

    mymap1['a'] = 101;
    mymap1['c'] = 202;
    mymap1['f'] = 303;

    for (c = 'a'; c < 'h'; c++)
    {
        std::cout << c;
        if (mymap1.count(c) > 0)
            std::cout << " is an element of mymap.\n";
        else
            std::cout << " is not an element of mymap.\n";
    }

    // lower_bound(), upper_bound()
    std::map<char, int> mymap2;
    std::map<char, int>::iterator itlow, itup;

    mymap2['a'] = 20;
    mymap2['b'] = 40;
    mymap2['c'] = 60;
    mymap2['d'] = 80;
    mymap2['e'] = 100;

    itlow = mymap2.lower_bound('b'); // itlow points to b
    itup = mymap2.upper_bound('d');  // itup points to e (not d!)

    mymap2.erase(itlow, itup); // erases [itlow,itup)

    // print content:
    for (std::map<char, int>::iterator it = mymap2.begin(); it != mymap2.end(); ++it)
        std::cout << it->first << " => " << it->second << '\n';

    // equal_range()
    std::map<char, int> mymap3;

    mymap3['a'] = 10;
    mymap3['b'] = 20;
    mymap3['c'] = 30;

    std::pair<std::map<char, int>::iterator, std::map<char, int>::iterator> ret;
    ret = mymap3.equal_range('b');

    std::cout << "lower bound points to: ";
    std::cout << ret.first->first << " => " << ret.first->second << '\n';

    std::cout << "upper bound points to: ";
    std::cout << ret.second->first << " => " << ret.second->second << '\n';

    return 0;
}