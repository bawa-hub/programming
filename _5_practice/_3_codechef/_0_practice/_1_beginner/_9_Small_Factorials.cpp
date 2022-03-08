#include <iostream>
using namespace std;

int main()
{
    int t;
    cin >> t;
    while (t--)
    {
        int n;
        cin >> n;
        int size = 1000, fact[size] = {0};
        fact[size - 1] = 1;
        int j = size - 1, carry = 0, x, i;
        while (n > 1)
        {
            for (i = size - 1; i >= j; i--)
            {
                x = fact[i] * n + carry;
                carry = x / 10;
                fact[i] = x % 10;
            }
            while (carry > 0)
            {
                fact[--j] = (carry) % 10;
                carry = carry / 10;
            }
            n--;
        }
        for (int i = j; i <= size - 1; i++)
        {
            cout << fact[i];
        }
        cout << endl;
    }
    return 0;
}