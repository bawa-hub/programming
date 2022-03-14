// Given an array A[] and a number x, check for pair in A[] with sum as x (aka Two Sum)
#include <bits/stdc++.h>
using namespace std;

// method 1
bool isTwoSum(int a[], int n, int sum)
{
    for (int i = 0; i < n; i++)
    {
        for (int j = i + 1; j < n; j++)
        {
            if (a[i] + a[j] == sum)
            {
                cout << "elements are: " << a[i] << " " << a[j] << endl;
                return true;
            }
        }
    }
    return false;
}
// time complexity - O(n^2)
// space complexity - O(1)

// method 2  - sorting and two pointer technique
bool isTwoSumWithTwoPointer(int a[], int n, int sum)
{
    sort(a, a + n);
    int l = 0, r = n - 1;
    while (l < r)
    {
        if (a[l] + a[r] == sum)
        {
            cout << "elements are: " << a[l] << " " << a[r] << endl;
            return true;
        }
        else if (a[l] + a[r] > sum)
        {
            r--;
        }
        else
        {
            l++;
        }
    }
    return false;
}
// Time Complexity: Depends on what sorting algorithm we use.
//     If Merge Sort or Heap Sort is used then (-)(nlogn) in the worst case.
//     If Quick Sort is used then O(n^2) in the worst case.
// Auxiliary Space: This too depends on sorting algorithm. The auxiliary space is O(n) for merge sort and O(1) for Heap Sort.

// method 3 - hashing
bool isTwoSumWithHashing(int a[], int n, int sum)
{
    unordered_set<int> s;
    for (int i = 0; i < n; i++)
    {
        int temp = sum - a[i];
        if (s.find(temp) != s.end())
        {
            cout << "elements are: " << a[i] << " " << temp << endl;
            return true;
        }
        s.insert(a[i]);
    }
    return false;
}
// time complexity - O(n)
// space complexity - O(n)

int main()
{
    int n;
    cin >> n;
    int a[n];
    for (int i = 0; i < n; i++)
        cin >> a[i];
    int sum;
    cin >> sum;
    // cout << isTwoSum(a, n, sum) << endl;
    // cout << isTwoSumWithTwoPointer(a, n, sum) << endl;
    cout << isTwoSumWithHashing(a, n, sum) << endl;
}