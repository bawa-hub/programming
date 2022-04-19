#include <iostream>
#include <map>

int main()
{
    std::map<char, int> mymap;

    // insert()

    // first insert function version (single parameter):
    mymap.insert(std::pair<char, int>('a', 100));
    mymap.insert(std::pair<char, int>('z', 200));

    std::pair<std::map<char, int>::iterator, bool> ret;
    ret = mymap.insert(std::pair<char, int>('z', 500));
    if (ret.second == false)
    {
        std::cout << "element 'z' already existed";
        std::cout << " with a value of " << ret.first->second << '\n';
    }

    // second insert function version (with hint position):
    std::map<char, int>::iterator it = mymap.begin();
    mymap.insert(it, std::pair<char, int>('b', 300)); // max efficiency inserting
    mymap.insert(it, std::pair<char, int>('c', 400)); // no max efficiency inserting

    // third insert function version (range insertion):
    std::map<char, int> anothermap;
    anothermap.insert(mymap.begin(), mymap.find('c'));

    // showing contents:
    std::cout << "mymap contains:\n";
    for (it = mymap.begin(); it != mymap.end(); ++it)
        std::cout << it->first << " => " << it->second << '\n';

    std::cout << "anothermap contains:\n";
    for (it = anothermap.begin(); it != anothermap.end(); ++it)
        std::cout << it->first << " => " << it->second << '\n';

    std::cout << "\n";

    // erase()
    std::map<char, int> mymap1;
    std::map<char, int>::iterator it1;

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
        std::cout << it1->first << " => " << it1->second << '\n';

    //swap()
    std::map<char, int> foo, bar;

    foo['x'] = 100;
    foo['y'] = 200;

    bar['a'] = 11;
    bar['b'] = 22;
    bar['c'] = 33;

    foo.swap(bar);

    std::cout << "foo contains:\n";
    for (std::map<char, int>::iterator it = foo.begin(); it != foo.end(); ++it)
        std::cout << it->first << " => " << it->second << '\n';

    std::cout << "bar contains:\n";
    for (std::map<char, int>::iterator it = bar.begin(); it != bar.end(); ++it)
        std::cout << it->first << " => " << it->second << '\n';

    // clear()
    std::map<char, int> mymap2;

    mymap2['x'] = 100;
    mymap2['y'] = 200;
    mymap2['z'] = 300;

    std::cout << "mymap contains:\n";
    for (std::map<char, int>::iterator it = mymap2.begin(); it != mymap2.end(); ++it)
        std::cout << it->first << " => " << it->second << '\n';

    mymap2.clear();
    mymap2['a'] = 1101;
    mymap2['b'] = 2202;

    std::cout << "mymap contains:\n";
    for (std::map<char, int>::iterator it = mymap2.begin(); it != mymap2.end(); ++it)
        std::cout << it->first << " => " << it->second << '\n';

    // emplace()
    std::map<char, int> mymap3;

    mymap3.emplace('x', 100);
    mymap3.emplace('y', 200);
    mymap3.emplace('z', 100);

    std::cout << "mymap contains:";
    for (auto &x : mymap3)
        std::cout << " [" << x.first << ':' << x.second << ']';
    std::cout << '\n';

    // emplace_hint()
    std::map<char, int> mymap4;
    auto it2 = mymap4.end();

    it2 = mymap4.emplace_hint(it2, 'b', 10);
    mymap4.emplace_hint(it2, 'a', 12);
    mymap4.emplace_hint(mymap4.end(), 'c', 14);

    std::cout << "mymap contains:";
    for (auto &x : mymap4)
        std::cout << " [" << x.first << ':' << x.second << ']';
    std::cout << '\n';

    return 0;
}