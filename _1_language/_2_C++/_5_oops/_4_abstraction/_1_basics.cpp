// Abstract classes are the way to achieve abstraction in C++.
// Abstraction in C++ is the process to hide the internal details and showing functionality only.
// Abstraction can be achieved by two ways:
//     Abstract class
//     Interface
// Abstract class and interface both can have abstract methods which are necessary for abstraction.

// In C++ class is made abstract by declaring at least one of its functions as <>strong>pure virtual function.

// pure virtual function is specified by placing "= 0" in its declaration.
// Its implementation must be provided by derived classes.

#include <iostream>
using namespace std;
class Shape
{
public:
    virtual void draw() = 0;
};
class Rectangle : Shape
{
public:
    void draw()
    {
        cout << "drawing rectangle..." << endl;
    }
};
class Circle : Shape
{
public:
    void draw()
    {
        cout << "drawing circle..." << endl;
    }
};
int main()
{
    Rectangle rec;
    Circle cir;
    rec.draw();
    cir.draw();
    return 0;
}

// Output:
// drawing rectangle...
// drawing circle...
