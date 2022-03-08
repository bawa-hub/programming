// C++ code to demonstrate the working of
// merge() and operator=
#include <iostream>
#include <forward_list>
using namespace std;

int main()
{
    // Initializing 1st forward list
    forward_list<int> flist1 = {1, 2, 3};

    // Declaring 2nd forward list
    forward_list<int> flist2;

    // Creating deep copy using "="
    flist2 = flist1;

    // Displaying flist2
    cout << "The contents of 2nd forward list"
            " after copy are : ";
    for (int &x : flist2)
        cout << x << " ";
    cout << endl;

    // Using merge() to merge both list in 1
    flist1.merge(flist2);

    // Displaying merged forward list
    // Prints sorted list
    cout << "The contents of forward list "
            "after merge are : ";
    for (int &x : flist1)
        cout << x << " ";
    cout << endl;

    return 0;
}
