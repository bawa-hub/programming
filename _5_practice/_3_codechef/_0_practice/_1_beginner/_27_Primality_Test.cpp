#include <iostream>
using namespace std;

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

int main()
{
    int t;
    cin >> t;
    while (t--)
    {
        int n;
        cin >> n;
        if (n == 1)
        {
            cout << "no" << endl;
        }
        else
        {
            if (isPrime(n))
                cout << "yes"
                     << "\n";
            else
                cout << "no"
                     << "\n";
        }
    }
    return 0;
}
