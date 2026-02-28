#include <bits/stdc++.h>

using namespace std;

// approach 1
int count_digits(int n)
{
    int x = n;
    int count = 0;
    while (x != 0)
    {
        x = x / 10;
        count++;
    }
    return count;
}
// Time Complexity: O (n) where n is the number of digits in the given integer
// Space Complexity: O(1)

// approach 2
int count_digits(int n)
{
    string x = to_string(n);
    return x.length();
}
// Time Complexity: O (1)
// Space Complexity: O(1)

// approach 3
// The number of digits in an integer = the upper bound of log10(n).
int count_digits(int n)
{
    int digits = floor(log10(n) + 1);
    return digits;
}
// Time Complexity: O (1)
// Space Complexity: O(1)

int main()
{
    int n = 12345;
    cout << "Number of digits in " << n << " is " << count_digits(n);
    return 0;
}
