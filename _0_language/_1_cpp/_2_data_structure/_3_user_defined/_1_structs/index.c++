// In C++, classes and structs are blueprints that are used to create the instance of a class.
// Structs are used for lightweight objects such as Rectangle, color, Point, etc.
// Unlike class, structs in C++ are value type than reference type.
// It is useful if you have data that is not intended to be modified after creation of struct.

// unlike class, in struct access fields and methods by dot(.) operator

#include <iostream>
using namespace std;
struct Rectangle
{
    int width, height;
    Rectangle(int w, int h)
    {
        width = w;
        height = h;
    }

    void areaOfRectangle()
    {
        cout << "Area of Rectangle is: " << (width * height);
    }
};
int main(void)
{
    struct Rectangle rec = Rectangle(8, 5);
    rec.areaOfRectangle();
    return 0;
}