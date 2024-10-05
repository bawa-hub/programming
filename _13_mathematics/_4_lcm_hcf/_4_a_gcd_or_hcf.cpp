#include <bits/stdc++.h>
using namespace std;

// brute force
int main()
{
    int num1 = 4, num2 = 8;
    int ans;
    for (int i = 1; i <= min(num1, num2); i++)
    {
        if (num1 % i == 0 && num2 % i == 0)
        {
            ans = i;
        }
    }
    cout << "The GCD of the two numbers is " << ans;
}
// Time Complexity: O(N).
// Space Complexity: O(1).

// Using Euclidean’s theorem
// gcd(a,b) = gcd(b, a%b)
int gcd(int a, int b)
{
    if (b == 0)
    {
        return a;
    }
    return gcd(b, a % b);
}

int main()
{

    int a = 4, b = 8;
    cout << "The GCD of the two numbers is " << gcd(a, b);
}
// Time Complexity: O(logɸmin(a,b)),where ɸ is 1.61.
// Space Complexity: O(1).