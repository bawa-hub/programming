#include <iostream>

using namespace std;
int main()
{
    int *ptr = NULL;
    cout << "The value of ptr is " << ptr;

    if (ptr == nullptr)
        cout << "\n"
             << "nullptr";

    return 0;
}