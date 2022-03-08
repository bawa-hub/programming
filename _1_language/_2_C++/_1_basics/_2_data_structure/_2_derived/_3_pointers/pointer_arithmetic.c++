// arithmetic operators that can be used on pointers: ++, --, +, and -

#include <iostream>

using namespace std;
const int MAX = 3;

int main()
{
    int var[MAX] = {10, 100, 200};
    int *ptr;
    int *ptr1;

    // let us have array address in pointer.
    ptr = var;

    for (int i = 0; i < MAX; i++)
    {
        cout << "Address of var[" << i << "] = ";
        cout << ptr << endl;

        cout << "Value of var[" << i << "] = ";
        cout << *ptr << endl;

        // point to the next location
        ptr++;
    }

    // let us have address of the last element in pointer.
    ptr1 = &var[MAX - 1];

    for (int i = MAX; i > 0; i--)
    {
        cout << "Address of var[" << i << "] = ";
        cout << ptr1 << endl;

        cout << "Value of var[" << i << "] = ";
        cout << *ptr1 << endl;

        // point to the previous location
        ptr1--;
    }

    return 0;
}