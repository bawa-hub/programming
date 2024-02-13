#include <iostream>
using namespace std;

// sum of elements upto index n-1 of array a of length n

int sumOfArray(int n, int a[])
{
    if (n < 0)
        return 0;
    return sumOfArray(n - 1, a) + a[n];
}

int main()
{
    int n;
    cin >> n;
    int a[n];
    for (int i = 0; i < n; i++)
    {
        cin >> a[i];
    }
    cout << sumOfArray(n - 1, a);
}