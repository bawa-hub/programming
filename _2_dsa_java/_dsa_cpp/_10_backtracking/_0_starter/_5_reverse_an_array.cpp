// https://takeuforward.org/data-structure/reverse-a-given-array/

#include <iostream>
#include <algorithm>
using namespace std;

void printArray(int ans[], int n)
{
    cout << "The reversed array is:- " << endl;
    for (int i = 0; i < n; i++)
    {
        cout << ans[i] << " ";
    }
}

// using extra array
void reverseArray(int arr[], int n)
{
    int ans[n];
    for (int i = n - 1; i >= 0; i--)
    {
        ans[n - i - 1] = arr[i];
    }
    printArray(ans, n);
}
// Time Complexity: O(n), single-pass for reversing array.
// Space Complexity: O(n), for the extra array used.

// Space-optimized iterative method
void reverseArray(int arr[], int n)
{
    int p1 = 0, p2 = n - 1;
    while (p1 < p2)
    {
        swap(arr[p1], arr[p2]);
        p1++;
        p2--;
    }
    printArray(arr, n);
}
// Time Complexity: O(n), single-pass involved.
// Space Complexity: O(1)

//  Recursive method
void reverseArray(int arr[], int start, int end)
{
    if (start < end)
    {
        swap(arr[start], arr[end]);
        reverseArray(arr, start + 1, end - 1);
    }
}
// Time Complexity: O(n)
// Space Complexity: O(1)

// Using library function
void reverseArray(int arr[], int n)
{
    // Reversing elements from index 0 to n-1
    reverse(arr, arr + n);
}
// Time Complexity: O(n)
// Space Complexity: O(1)

int main()
{
    int n = 5;
    int arr[] = {5, 4, 3, 2, 1};
    reverseArray(arr, n);
    return 0;
}
