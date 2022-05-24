// Enum in C++ is a data type that contains fixed set of constants.
// The C++ enum constants are static and final implicitly.

// enum can be easily used in switch
// enum can be traversed
// enum can have fields, constructors and methods
// enum may implement many interfaces but cannot extend any class because it internally extends Enum class

#include <iostream>
using namespace std;
enum week
{
    Monday,
    Tuesday,
    Wednesday,
    Thursday,
    Friday,
    Saturday,
    Sunday
};
int main()
{
    week day;
    day = Friday;
    cout << "Day: " << day + 1 << endl;
    return 0;
}