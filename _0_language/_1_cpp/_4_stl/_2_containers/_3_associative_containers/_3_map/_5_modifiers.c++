#include <iostream>
#include <map>

using namespace std;

int main()
{
    map<char, int> mymap;

    // insert()

    // first insert function version (single parameter):
    mymap.insert(pair<char, int>('a', 100));
    mymap.insert(pair<char, int>('z', 200));

    pair<map<char, int>::iterator, bool> ret;
    ret = mymap.insert(pair<char, int>('z', 500));
    if (ret.second == false)
    {
        cout << "element 'z' already existed";
        cout << " with a value of " << ret.first->second << '\n';
    }

    // second insert function version (with hint position):
    map<char, int>::iterator it = mymap.begin();
    mymap.insert(it, pair<char, int>('b', 300)); // max efficiency inserting
    mymap.insert(it, pair<char, int>('c', 400)); // no max efficiency inserting

    // third insert function version (range insertion):
    map<char, int> anothermap;
    anothermap.insert(mymap.begin(), mymap.find('c'));

    // showing contents:
    cout << "mymap contains:\n";
    for (it = mymap.begin(); it != mymap.end(); ++it)
        cout << it->first << " => " << it->second << '\n';

    cout << "anothermap contains:\n";
    for (it = anothermap.begin(); it != anothermap.end(); ++it)
        cout << it->first << " => " << it->second << '\n';

    cout << "\n";

    // erase()
    map<char, int> mymap1;
    map<char, int>::iterator it1;

    // insert some values:
    mymap1['a'] = 10;
    mymap1['b'] = 20;
    mymap1['c'] = 30;
    mymap1['d'] = 40;
    mymap1['e'] = 50;
    mymap1['f'] = 60;

    it1 = mymap1.find('b');
    mymap1.erase(it1); // erasing by iterator

    mymap1.erase('c'); // erasing by key

    it1 = mymap1.find('e');
    mymap1.erase(it1, mymap1.end()); // erasing by range

    // show content:
    for (it1 = mymap1.begin(); it1 != mymap1.end(); ++it1)
        cout << it1->first << " => " << it1->second << '\n';

    //swap()
    map<char, int> foo, bar;

    foo['x'] = 100;
    foo['y'] = 200;

    bar['a'] = 11;
    bar['b'] = 22;
    bar['c'] = 33;

    foo.swap(bar);

    cout << "foo contains:\n";
    for (map<char, int>::iterator it = foo.begin(); it != foo.end(); ++it)
        cout << it->first << " => " << it->second << '\n';

    cout << "bar contains:\n";
    for (map<char, int>::iterator it = bar.begin(); it != bar.end(); ++it)
        cout << it->first << " => " << it->second << '\n';

    // clear()
    map<char, int> mymap2;

    mymap2['x'] = 100;
    mymap2['y'] = 200;
    mymap2['z'] = 300;

    cout << "mymap contains:\n";
    for (map<char, int>::iterator it = mymap2.begin(); it != mymap2.end(); ++it)
        cout << it->first << " => " << it->second << '\n';

    mymap2.clear();
    mymap2['a'] = 1101;
    mymap2['b'] = 2202;

    cout << "mymap contains:\n";
    for (map<char, int>::iterator it = mymap2.begin(); it != mymap2.end(); ++it)
        cout << it->first << " => " << it->second << '\n';

    // emplace()
    map<char, int> mymap3;

    mymap3.emplace('x', 100);
    mymap3.emplace('y', 200);
    mymap3.emplace('z', 100);

    cout << "mymap contains:";
    for (auto &x : mymap3)
        cout << " [" << x.first << ':' << x.second << ']';
    cout << '\n';

    // emplace_hint()
    map<char, int> mymap4;
    auto it2 = mymap4.end();

    it2 = mymap4.emplace_hint(it2, 'b', 10);
    mymap4.emplace_hint(it2, 'a', 12);
    mymap4.emplace_hint(mymap4.end(), 'c', 14);

    cout << "mymap contains:";
    for (auto &x : mymap4)
        cout << " [" << x.first << ':' << x.second << ']';
    cout << '\n';

    return 0;
}