// C++ code to demonstrate the working of
// clear() and empty()
#include <iostream>
#include <forward_list> // for clear() and empty()
using namespace std;
int main()
{
    // Initializing forward list
    forward_list<int> flist1 = {
        1,
        2,
        3,
    };

    // Displaying forward list before clearing
    cout << "The contents of forward list are : ";
    for (int &x : flist1)
        cout << x << " ";
    cout << endl;

    // Using clear() to clear the forward list
    flist1.clear();

    // Displaying list after clear() performed
    cout << "The contents of forward list after "
         << "clearing are : ";
    for (int &x : flist1)
        cout << x << " ";
    cout << endl;

    // Checking if list is empty
    flist1.empty() ? cout << "Forward list is empty" : cout << "Forward list is not empty";

    return 0;
}
