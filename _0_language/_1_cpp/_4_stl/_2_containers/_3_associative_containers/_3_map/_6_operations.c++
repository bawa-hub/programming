#include <iostream>
#include <map>

using namespace std;

int main()
{
    // find()
    map<char, int> mymap;
    map<char, int>::iterator it;

    mymap['a'] = 50;
    mymap['b'] = 100;
    mymap['c'] = 150;
    mymap['d'] = 200;

    if (mymap.find('b') != mymap.end())
        mymap.erase(mymap.find('b'));

    // print content:
    cout << "elements in mymap:" << '\n';
    cout << "a => " << mymap.find('a')->second << '\n';
    cout << "c => " << mymap.find('c')->second << '\n';
    cout << "d => " << mymap.find('d')->second << '\n';

    // count()
    map<char, int> mymap1;
    char c;

    mymap1['a'] = 101;
    mymap1['c'] = 202;
    mymap1['f'] = 303;

    for (c = 'a'; c < 'h'; c++)
    {
        cout << c;
        if (mymap1.count(c) > 0)
            cout << " is an element of mymap.\n";
        else
            cout << " is not an element of mymap.\n";
    }

    // lower_bound(), upper_bound()
    map<char, int> mymap2;
    map<char, int>::iterator itlow, itup;

    mymap2['a'] = 20;
    mymap2['b'] = 40;
    mymap2['c'] = 60;
    mymap2['d'] = 80;
    mymap2['e'] = 100;

    itlow = mymap2.lower_bound('b'); // itlow points to b
    itup = mymap2.upper_bound('d');  // itup points to e (not d!)

    mymap2.erase(itlow, itup); // erases [itlow,itup)

    // print content:
    for (map<char, int>::iterator it = mymap2.begin(); it != mymap2.end(); ++it)
        cout << it->first << " => " << it->second << '\n';

    // equal_range()
    map<char, int> mymap3;

    mymap3['a'] = 10;
    mymap3['b'] = 20;
    mymap3['c'] = 30;

    pair<map<char, int>::iterator, map<char, int>::iterator> ret;
    ret = mymap3.equal_range('b');

    cout << "lower bound points to: ";
    cout << ret.first->first << " => " << ret.first->second << '\n';

    cout << "upper bound points to: ";
    cout << ret.second->first << " => " << ret.second->second << '\n';

    return 0;
}