#include <iostream>
using namespace std;

// approach 1
void printDivisorsBruteForce(int n)
{

    cout << "The Divisors of " << n << " are:" << endl;
    for (int i = 1; i <= n; i++)
        if (n % i == 0)
            cout << i << " ";

    cout << "\n";
}
// Time Complexity: O(n), because the loop has to run from 1 to n always.
// Space Complexity: O(1), we are not using any extra space

// approach 2
void printDivisorsOptimal(int n)
{

    cout << "The Divisors of " << n << " are:" << endl;
    for (int i = 1; i <= sqrt(n); i++)
        if (n % i == 0)
        {
            cout << i << " ";
            if (i != n / i)
                cout << n / i << " ";
        }

    cout << "\n";
}
// Time Complexity: O(sqrt(n)), because everytime the loop runs only sqrt(n) times.
// Space Complexity: O(1), we are not using any extra space

int main()
{

    printDivisorsBruteForce(36);

    return 0;
}
