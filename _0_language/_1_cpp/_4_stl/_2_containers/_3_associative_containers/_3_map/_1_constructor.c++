#include <iostream>
#include <map>

using namespace std;

bool fncomp(char lhs, char rhs) { return lhs < rhs; }

struct classcomp
{
    bool operator()(const char &lhs, const char &rhs) const
    {
        return lhs < rhs;
    }
};

int main()
{
    // constructor
    map<char, int> first;

    first['a'] = 10;
    first['b'] = 30;
    first['c'] = 50;
    first['d'] = 70;

    map<char, int> second(first.begin(), first.end());
    map<char, int> third(second);
    map<char, int, classcomp> fourth; // class as Compare

    bool (*fn_pt)(char, char) = fncomp;
    map<char, int, bool (*)(char, char)> fifth(fn_pt); // function pointer as Compare

    // assignment operator with maps
    map<char, int> one;
    map<char, int> two;

    one['x'] = 8;
    one['y'] = 16;
    one['z'] = 32;

    two = one;                   // two now contains 3 ints
    one = map<char, int>(); // and one is now empty

    cout << "Size of one: " << one.size() << '\n';
    cout << "Size of two: " << two.size() << '\n';

    return 0;
}