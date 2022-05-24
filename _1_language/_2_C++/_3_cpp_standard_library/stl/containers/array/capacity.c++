#include <iostream>
#include <array>

int main()
{
    std::array<int, 5> myints;

    // size(), max_size()
    std::cout << "size of myints: " << myints.size() << std::endl;
    std::cout << "sizeof(myints): " << sizeof(myints) << std::endl;
    std::cout << "max_size of myints: " << myints.max_size() << '\n';

    // empty()
    std::array<int, 0> first;
    std::array<int, 5> second;
    std::cout << "first " << (first.empty() ? "is empty" : "is not empty") << '\n';
    std::cout << "second " << (second.empty() ? "is empty" : "is not empty") << '\n';

    return 0;
}