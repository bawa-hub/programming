#include <iostream>
using namespace std;

const int m = 1e7; // global variable can be of large size but should be constant
int b[m];

int main()
{
    // int n = 1e7; // local array size should be order of 1e5
    int n = 1e5;
    int a[n];
    a[n - 1] = 3;
    cout << a[n - 1] << endl;
    b[m - 1] = 4;
    cout << b[m - 1] << endl;
}