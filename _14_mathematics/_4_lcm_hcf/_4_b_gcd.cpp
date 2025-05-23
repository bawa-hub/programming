#include <algorithm>
#include <iostream>
using namespace std;

/**
 * The GCD of two or more numbers is the largest positive number that divides all the numbers that are considered. For example, 
 * the GCD of 6 and 10 is 2 because it is the largest positive number that can divide both 6 and 10.
 * 
*/

// Naive approach
// Traverse all the numbers from min(A, B) to 1 and check whether the current number divides both A and B. If yes, it is the GCD of A and B.
int GCD1(int A, int B)
{
    int m = min(A, B), gcd;
    for (int i = m; i > 0; --i)
        if (A % i == 0 && B % i == 0)
        {
            gcd = i;
            return gcd;
        }
}
// Time Complexity = O(min(A,B))

// Euclid's algorithm
// The idea behind this algorithm is GCD(A,B) = GCD(B,A%B) . It will recurse until A%B = 0
int GCD(int A, int B)
{
    if (B == 0)
        return A;
    else
        return GCD(B, A % B);
}
// Time Complexity = O(log(max(A,B)))

// Extended Euclid algorithm
int d, x, y;
void extendedEuclid(int A, int B)
{
    if (B == 0)
    {
        d = A;
        x = 1;
        y = 0;
    }
    else
    {
        extendedEuclid(B, A % B);
        int temp = x;
        x = y;
        y = temp - (A / B) * y;
    }
}

int main()
{
    extendedEuclid(16, 10);
    cout << "The GCD of 16 and 10 is " << d << endl;
    cout << "The GCD of 16 and 10 is " << GCD(16, 10) << endl;
    cout << "Coefficients x and y are " << x << "and  " << y << endl;
    return 0;
}