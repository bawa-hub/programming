#include <iostream>
using namespace std;

long long int GCD(long long int A, long long int B)
{
    if (B == 0)
        return A;
    else
        return GCD(B, A % B);
}

int main()
{
    int t;
    cin >> t;
    while (t--)
    {
        long long int num1, num2;
        cin >> num1 >> num2;
        long long int gcd = GCD(num1, num2);
        long long int lcm = (num1 * num2) / gcd;
        cout << gcd << " " << lcm << "\n";
    }
    return 0;
}