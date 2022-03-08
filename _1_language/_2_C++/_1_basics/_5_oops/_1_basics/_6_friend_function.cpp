// If a function is defined as a friend function in C++, then the protected and private data of a class can be accessed using the function.

#include <iostream>
using namespace std;

class Box
{
private:
    int length;

public:
    Box() : length(0) {}         // this syntax is used for setting member variable directly in constructor
    friend int printLength(Box); //friend function
};
int printLength(Box b)
{
    b.length += 10;
    return b.length;
}
int main()
{
    Box b;
    cout << "Length of box: " << printLength(b) << endl;
    return 0;
}