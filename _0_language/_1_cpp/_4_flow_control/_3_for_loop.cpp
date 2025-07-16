#include <iostream>
using namespace std;

int main()
{

    //   type 1
    cout << "Type 1 : ";
    for (int i = 0; i < 10; i++)
    {
        cout << i << " ";
    }

    // type 2
    int arr[10] = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9};
    cout << "Type 2 : ";
    for (int i : arr)
    {
        cout << i << " ";
    }

    // infinite loop
    for (;;)
    {
        cout << "Hello, I will run infinitely" << endl;
    }
}