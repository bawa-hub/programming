// There can be two types of constructors in C++.

// Default constructor
// Parameterized constructor

#include <iostream>
using namespace std;

class Employee
{
public:
    int id;
    string name;
    float salary;

    // Default constructor
    Employee()
    {
        cout << "Default Constructor invoked" << endl;
    }

    // Parameterised constructor
    Employee(int id, string name, float salary)
    {
        this->id = id;
        this->name = name;
        this->salary = salary;
        cout << "Parameterised Constructor invoked" << endl;
    }
};

int main()
{
    Employee employee = Employee();
    Employee employee2 = Employee(1, "Bawa", 100);
}