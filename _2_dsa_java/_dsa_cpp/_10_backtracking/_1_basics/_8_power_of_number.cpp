// https://www.geeksforgeeks.org/write-a-c-program-to-calculate-powxn/
// https://leetcode.com/problems/powx-n/
#include <bits/stdc++.h>
using namespace std;

// brute force using recursion
int power(int x, int n)
{
    if (n == 0)
        return 1;
    return power(x, n - 1) * x;
}

// brute force iterative
double power(double x, int n)
{
    double ans = 1.0;
    for (int i = 0; i < n; i++)
    {
        ans = ans * x;
    }
    return ans;
}
// Time Complexity: O(N)
// Space Complexity: O(1)

// using binary exponentiation
double power(double x, int n)
{
    double ans = 1.0;
    long long nn = n;
    if (nn < 0)
        nn = -1 * nn;
    while (nn)
    {
        if (nn % 2)
        {
            ans = ans * x;
            nn = nn - 1;
        }
        else
        {
            x = x * x;
            nn = nn / 2;
        }
    }
    if (n < 0)
        ans = (double)(1.0) / (double)(ans);
    return ans;
}
// Time Complexity: O(log n)
// Space Complexity: O(1)

int main()
{
    int x, n;
    cin >> x >> n;
    cout << power(x, n);
}