#include <bits/stdc++.h>

using namespace std;

// using iterative solution
bool isPrime(int N)
{
    for (int i = 2; i < N; i++)
    {
        if (N % i == 0)
        {
            return false;
        }
    }
    return true;
}
// Time Complexity: O(n)
// Space Complexity: O(1)

// optimized approach
bool isPrime(int N)
{
    for (int i = 2; i < sqrt(N); i++)
    {
        if (N % i == 0)
        {
            return false;
        }
    }
    return true;
}
// Time Complexity: O(√n)
// Space Complexity: O(1)

int main()
{

    int n = 11;

    bool ans = isPrime(n);
    if (n != 1 && ans == true)
    {
        cout << "Prime Number";
    }
    else
    {
        cout << "Non Prime Number";
    }
    return 0;
}