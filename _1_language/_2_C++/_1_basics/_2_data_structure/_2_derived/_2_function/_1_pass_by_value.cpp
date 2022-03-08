#include <iostream>
using namespace std;

int increment(int n)
{
    n++;
    return n;
}

int main()
{
    int a = 1;
    cout << a << endl;
    a = increment(a); // passes copy of variable
    cout << a << endl;
}