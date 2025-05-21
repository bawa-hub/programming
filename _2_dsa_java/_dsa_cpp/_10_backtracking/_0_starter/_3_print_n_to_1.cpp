#include <iostream>
using namespace std;

void print(int n)
{
    if (n < 1)
        return;
    cout << n << endl;
    print(n - 1);
}

int main()
{
    int n = 10;
    print(n);
}