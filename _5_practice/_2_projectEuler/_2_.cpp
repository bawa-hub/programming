#include <iostream>
using namespace std;

int fib(int n)
{
    if (n == 1)
        return 1;
    else if (n == 2)
        return 2;
    else
        return fib(n - 1) + fib(n - 2);
}

int main()
{
    int i = 1;
    int sum = 0;
    while (fib(i) < 4000000)
    {
        if (fib(i) % 2 == 0)
        {
            sum += fib(i);
        }
        i++;
    }
    cout << sum;
}