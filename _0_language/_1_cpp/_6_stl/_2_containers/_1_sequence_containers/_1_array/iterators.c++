#include <iostream>
#include <array>

int main()
{
    // declaration
    std::array<int, 5> myarray = {2, 16, 77, 34, 50};

    // begin(), end()
    std::cout << "myarray contains:";
    for (auto it = myarray.begin(); it != myarray.end(); ++it)
        std::cout << ' ' << *it;
    std::cout << '\n';

    // rbegin(), rend()
    std::cout << "myarray contains in reverse:";
    for (auto rit = myarray.rbegin(); rit < myarray.rend(); ++rit)
        std::cout << ' ' << *rit;

    std::cout << '\n';

    // cbegin(), cend()
    std::cout << "myarray contains:";
    for (auto it = myarray.cbegin(); it != myarray.cend(); ++it)
        std::cout << ' ' << *it; // cannot modify *it
    std::cout << '\n';

    // crbegin(), crend()
    std::cout << "myarray backwards:";
    for (auto rit = myarray.crbegin(); rit < myarray.crend(); ++rit)
        std::cout << ' ' << *rit; // cannot modify *rit

    std::cout << '\n';

    return 0;
}