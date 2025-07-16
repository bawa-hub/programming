#include <iostream>
using namespace std;

// references is only present in c++
// in c we use this functionality with pointer
// we can neglect pointer in c++

void increment(int &n)
{
    n++;
}

void swap(int &a, int &b)
{
    int temp = a;
    a = b;
    b = temp;
}

// array by default pass by reference
void func(int a[])
{
    a[0] = 2;
}

int main()
{
    int a = 1;
    int b = 4;

    cout << a << endl;
    increment(a); // pass by reference
    cout << a << endl;

    cout << a << " " << b << endl;
    swap(a, b);
    cout << a << " " << b << endl;

    int arr[10];
    arr[0] = 1;
    cout << arr[0] << endl;
    func(arr);
    cout << arr[0] << endl;
}