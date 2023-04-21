// https://leetcode.com/problems/reverse-integer/

#include <stdio.h>

int main()
{
    int N = 123;
    int num = N;
    int reverse = 0;
    while (N != 0)
    {
        int digit = N % 10;
        reverse = reverse * 10 + digit;
        N = N / 10;
    }
    printf("The reverse of the %d is %d", num, reverse);
}

// Time Complexity: O(n), where n is the length of the given number
// Space Complexity: O(1)

// for leetcode
int reverse(int x)
{
    long r = 0;
    while (x)
    {
        r = r * 10 + x % 10;
        x = x / 10;
    }
    if (r > INT_MAX || r < INT_MIN)
        return 0;
    return int(r);
}