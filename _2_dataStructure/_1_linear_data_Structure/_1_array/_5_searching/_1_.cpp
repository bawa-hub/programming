// Search and delete in an unsorted array

#include <bits/stdc++.h>
using namespace std;

int search(int arr[], int n, int key)
{
    for (int i = 0; i < n; ++i)
    {
        if (arr[i] == key)
            return i;
    }
    return -1;
}

int deleteElement(int arr[], int n, int key)
{
    int pos = search(arr, n, key);
    if (pos == -1)
    {
        cout << "Element not found;";
        return n;
    }

    for (int i = pos; i < n - 1; i++)
    {
        arr[i] = arr[i + 1];
    }

    return n - 1;
}

int main()
{
    int a[] = {1, 4, 2, 5, 6, 3, 9, 8};
    int n = sizeof(a) / sizeof(a[0]);
    cout << "index of 4: " << search(a, n, 4) << endl;

    cout << "Array before deletion\n";
    for (int i = 0; i < n; i++)
    {
        cout << a[i] << " ";
    }

    n = deleteElement(a, n, 2);

    cout << "\n\nArray after deletion\n";
    for (int i = 0; i < n; i++)
    {
        cout << a[i] << " ";
    }

    return 0;
}

// Time complexities:
// Search: O(n)
// Insert: O(1)
// Delete: O(n)