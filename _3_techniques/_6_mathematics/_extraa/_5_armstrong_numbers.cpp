//  Armstrong numbers having the sum of digits raised to power no. of digits is equal to a given number
// 1^3+5^3+3^3=153

#include <bits/stdc++.h>
using namespace std;
bool ArmstrongNumber(int n)
{
    int originalno = n;
    int count = 0;
    int temp = n;
    while (temp != 0)
    {
        count++;
        temp = temp / 10;
    }
    int sumofpower = 0;
    while (n != 0)
    {
        int digit = n % 10;
        sumofpower += pow(digit, count);
        n /= 10;
    }
    return (sumofpower == originalno);
}
int main()
{
    int n1 = 153;
    if (ArmstrongNumber(n1))
    {
        cout << "Yes, it is an Armstrong Number\n";
    }
    else
    {
        cout << "No, it is not an Armstrong Number\n";
    }
    return 0;
}

// Time Complexity: O(n) where n is the number of digits since we need to traverse every digit and add digits raised to power no. of digits to sum.
// Space Complexity: O(1) since no extra space is required