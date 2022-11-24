// C++ code to demonstrate the working of
// reverse() and swap()
#include <iostream>
#include <forward_list> // for reverse() and swap()
using namespace std;
int main()
{
    // Initializing 1st forward list
    forward_list<int> flist1 = {
        1,
        2,
        3,
    };

    // Initializing 2nd forward list
    forward_list<int> flist2 = {4, 5, 6};

    // Using reverse() to reverse 1st forward list
    flist1.reverse();

    // Displaying reversed forward list
    cout << "The contents of forward list after"
            " reversing are : ";
    for (int &x : flist1)
        cout << x << " ";
    cout << endl
         << endl;

    // Displaying forward list before swapping
    cout << "The contents of 1st forward list "
            "before swapping are : ";
    for (int &x : flist1)
        cout << x << " ";
    cout << endl;
    cout << "The contents of 2nd forward list "
            "before swapping are : ";
    for (int &x : flist2)
        cout << x << " ";
    cout << endl;

    // Use of swap() to swap the list
    flist1.swap(flist2);

    // Displaying forward list after swapping
    cout << "The contents of 1st forward list "
            "after swapping are : ";
    for (int &x : flist1)
        cout << x << " ";
    cout << endl;

    cout << "The contents of 2nd forward list "
            "after swapping are : ";
    for (int &x : flist2)
        cout << x << " ";
    cout << endl;

    return 0;
}
