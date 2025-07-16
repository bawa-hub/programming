/**
 * Prime Number --
 * numbers that are greater than 1 and have only two factors 1 and itself.
 * 
 * Composite Number --
 * numbers that are greater than 1 but they have at least one more divisor other than 1 and itself.
 * 
 * **/

#include <iostream>
using namespace std;

/** Methods to check whether a number is prime.**/

// Naive Approach
bool checkprime(int N)
{
    int count = 0;
    for (int i = 1; i <= N; ++i)
        if (N % i == 0)
            count++;
    if (count == 2)
        return true;
    else
        return false;
}
// The time complexity of this function is O(N) because you traverse from 1 to N.

// better approach
// complexity O(sqrt(N))
bool isPrime(int n)
{
    for (int i = 2; i * i <= n; ++i)
    {
        if (n % i == 0)
        {
            return false;
        }
    }
    return true;
}