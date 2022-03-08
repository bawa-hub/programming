#include <iostream>
#include <map>

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
    std::map<char, int> first;

    first['a'] = 10;
    first['b'] = 30;
    first['c'] = 50;
    first['d'] = 70;

    std::map<char, int> second(first.begin(), first.end());

    std::map<char, int> third(second);

    std::map<char, int, classcomp> fourth; // class as Compare

    bool (*fn_pt)(char, char) = fncomp;
    std::map<char, int, bool (*)(char, char)> fifth(fn_pt); // function pointer as Compare

    // assignment operator with maps
    std::map<char, int> one;
    std::map<char, int> two;

    one['x'] = 8;
    one['y'] = 16;
    one['z'] = 32;

    two = one;                   // two now contains 3 ints
    one = std::map<char, int>(); // and one is now empty

    std::cout << "Size of one: " << one.size() << '\n';
    std::cout << "Size of two: " << two.size() << '\n';

    return 0;
}