/**
 * The GCD of two or more numbers is the largest positive number that divides all the numbers that are considered. For example, 
 * the GCD of 6 and 10 is 2 because it is the largest positive number that can divide both 6 and 10.
 * 
*/

#include <iostream>
using namespace std;

// brute force
int gcd1(int num1, int num2)
{
    // int num1 = 4, num2 = 8;
    int ans;
    for (int i = 1; i <= min(num1, num2); i++)
    {
        if (num1 % i == 0 && num2 % i == 0)
        {
            ans = i;
        }
    }
    return ans;
}
// Time Complexity: O(N).
// Space Complexity: O(1).

// brute force optimized
int gcd2(int num1, int num2)
{
    // int num1 = 4, num2 = 8;
    int ans;
    for (int i = min(num1, num2); i >=1;i--)
    {
        if (num1 % i == 0 && num2 % i == 0)
        {
            ans = i;
            break;
        }
    }
    return ans;
}
// TC : O(min(num1,num2));

// Euclid's algorithm
// The idea behind this algorithm is 
// GCD(a,b) = GCD(a-b, b) , where a > b and this is equivalent to
// GCD(A,B) = GCD(B,A%B) . It will recurse until A%B = 0. if one of them is zero, other is gcd
int gcd3(int a, int b) {
   while(a>0 && b>0) {
    if(a>b) a = a%b;
    else b = b%a;
   }
   if (a==0) return b;
   else return a;
}
// Time Complexity = O(log(max(A,B)))

int gcd4(int a, int b)
{
    if (b == 0)
    {
        return a;
    }
    return gcd4(b, a % b);
}

// Extended Euclid algorithm
int d, x, y;
void gcd5(int A, int B)
{
    if (B == 0)
    {
        d = A;
        x = 1;
        y = 0;
    }
    else
    {
        gcd5(B, A % B);
        int temp = x;
        x = y;
        y = temp - (A / B) * y;
    }
}

int main()
{

    int a = 4, b = 8;
    cout << "The GCD of the two numbers is " << gcd4(a, b);
}