#include <iostream>
using namespace std;

void print(int i, int n)
{
    if (i > n)
        return;
    cout << i << endl;
    print(i + 1, n);
}

int main()
{
    int n = 10;
    int i = 1;
    print(i, n);
}