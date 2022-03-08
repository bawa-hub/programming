#include <iostream>
#include <array>
#include <cstring>

int main()
{
    std::array<int, 10> myarray;
    unsigned int i;

    // assign some values:
    for (i = 0; i < 10; i++)
        myarray[i] = i;

    // oerator []
    // print content
    std::cout << "myarray contains:";
    for (i = 0; i < 10; i++)
        std::cout << ' ' << myarray[i];
    std::cout << '\n';

    // at()
    for (int i = 0; i < 10; i++)
        std::cout << ' ' << myarray.at(i);
    std::cout << '\n';

    // front(), back()
    std::cout << "front is: " << myarray.front() << std::endl; // 2
    std::cout << "back is: " << myarray.back() << std::endl;   // 77

    myarray.front() = 100;

    std::cout << "myarray now contains:";
    for (int &x : myarray)
        std::cout << ' ' << x;

    std::cout << '\n';

    // data()
    const char *cstr = "Test string";
    std::array<char, 12> charray;

    std::memcpy(charray.data(), cstr, 12);

    std::cout << charray.data() << '\n';

    return 0;
}