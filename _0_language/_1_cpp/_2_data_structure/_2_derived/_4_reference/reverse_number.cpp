#include <bits/stdc++.h>
using namespace std;

// using reference instead of pointer
void swap(int &a, int &b)
{
    int temp = a;
    a = b;
    b = temp;
}

void reverse(int a[], int l, int r)
{
    if (l >= r)
        return;
    swap(a[l], a[r]);
    reverse(a, l + 1, r - 1);
}

int main()
{
    int a[] = {1, 2, 3, 4};
    for (int i = 0; i < 4; i++)
    {
        cout << a[i] << " ";
    }
    cout << endl;
    reverse(a, 0, 3);
    for (int i = 0; i < 4; i++)
    {
        cout << a[i] << " ";
    }
    cout << endl;
}