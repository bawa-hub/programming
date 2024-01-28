#include <iostream>
using namespace std;

/**
 * There are two types of type conversion in C++.
 *     Implicit Conversion: automatically done by the compiler
 *     Explicit Conversion (also known as Type Casting) : user manually changes data from one type to another
 *
 *  three major ways in which we can use explicit conversion in C++:
    C-style type casting (also known as cast notation)
    Function notation (also known as old C++ style type casting)
    Type conversion operators

    Type Conversion Operators:
    C++ also has four operators for type conversion

    static_cast
    dynamic_cast
    const_cast
    reinterpret_cast


 *
 *
 */

int main()
{
    //  Implicit Conversion

    // Conversion From int to double
    int num_int = 9;
    double num_double;
    // int value is automatically converted to double by the compiler before it is assigned to the num_double variable.
    num_double = num_int;

    // Automatic Conversion from double to int
    int num_int1;
    double num_double1 = 9.99;
    num_int1 = num_double1;
    cout << "num_int1 = " << num_int << endl;
    cout << "num_double1 = " << num_double << endl;

    // Explicit Conversion

    // C-style Type Casting
    int num_int = 26;
    double num_double;
    num_double = (double)num_int;

    // Function-style Casting
    int num_int = 26;
    double num_double;
    num_double = double(num_int);

    // Type Conversion Operators
}
