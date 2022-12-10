// https://practice.geeksforgeeks.org/problems/check-if-an-array-is-sorted0701/1?utm_source=youtube&utm_medium=collab_striver_ytdescription&utm_campaign=check-if-an-array-is-sorted

#include <bits/stdc++.h>
using namespace std;

// brute force
// bool isSorted(int arr[], int n)
// {
//     for (int i = 0; i < n; i++)
//     {
//         for (int j = i + 1; j < n; j++)
//         {
//             if (arr[j] < arr[i])
//                 return false;
//         }
//     }

//     return true;
// }

// int main()
// {

//     int arr[] = {1, 2, 3, 4, 5}, n = 5;
//     bool ans = isSorted(arr, n);
//     if (ans)
//         cout << "True" << endl;
//     else
//         cout << "False" << endl;
//     return 0;
// }

// Time Complexity: O(N^2)
// Space Complexity: O(1)

// optimized
bool isSorted(int arr[], int n)
{
    for (int i = 1; i < n; i++)
    {
        if (arr[i] < arr[i - 1])
            return false;
    }

    return true;
}

int main()
{

    int arr[] = {1, 2, 3, 4, 5}, n = 5;

    printf("%s", isSorted(arr, n) ? "True" : "False");
}

// Time Complexity: O(N)
// Space Complexity: O(1)