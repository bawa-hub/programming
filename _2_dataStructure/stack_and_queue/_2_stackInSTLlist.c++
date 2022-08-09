#include <iostream>
#include <list>
using namespace std;

// Stack implementation in C++ using `std::list`
int main()
{
    list<string> s;

    s.push_front("A"); // Insert `A` into the stack
    s.push_front("B"); // Insert `B` into the stack
    s.push_front("C"); // Insert `C` into the stack
    s.push_front("D"); // Insert `D` into the stack

    // returns the total number of elements present in the stack
    cout << "The stack size is " << s.size() << endl;

    // prints the top of the stack (`D`)
    cout << "The top element is " << s.front() << endl;

    s.pop_front(); // removing the top element (`D`)
    s.pop_front(); // removing the next top (`C`)

    cout << "The stack size is " << s.size() << endl;

    // check if the stack is empty
    if (s.empty())
    {
        cout << "The stack is empty\n";
    }
    else
    {
        cout << "The stack is not empty\n";
    }

    return 0;
}

Output