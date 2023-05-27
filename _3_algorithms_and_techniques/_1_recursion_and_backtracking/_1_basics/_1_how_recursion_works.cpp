#include <iostream>
using namespace std;

/**
 * Recursive funciton can be divided into three parts:
 * 1. base condition
 * 2. self work
 * 3. recursive call
 */

void print(int num)
{
    // base condition
    if (num < 1)
        return;

    // self work
    cout << num << endl;

    // recursive call
    print(num - 1);
}

int main()
{
    print(50);
}
