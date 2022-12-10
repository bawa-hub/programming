// https://practice.geeksforgeeks.org/problems/leaders-in-an-array-1587115620/1?utm_source=youtube&utm_medium=collab_striver_ytdescription&utm_campaign=leaders-in-an-array

#include <iostream>
using namespace std;

// brute force
// void printLeadersBruteForce(int arr[], int n)
// {

//     for (int i = 0; i < n - 1; i++)
//     {
//         bool leader = true;

//         // Checking whether arr[i] is greater than all the elements in its right side
//         for (int j = i + 1; j < n; j++)
//             if (arr[j] > arr[i])
//             {
//                 leader = false;
//                 break;
//             }

//         if (leader)
//             cout << arr[i] << " ";
//     }
//     cout << arr[n - 1] << "\n";
// }

// int main()
// {

//     int arr1[] = {4, 7, 1, 0};
//     int n1 = sizeof(arr1) / sizeof(arr1[0]);
//     cout << "The leaders of the first array are: " << endl;
//     printLeadersBruteForce(arr1, n1);

//     int arr2[] = {10, 22, 12, 3, 0, 6};
//     int n2 = sizeof(arr2) / sizeof(arr2[0]);
//     cout << "The leaders of the second array are: " << endl;
//     printLeadersBruteForce(arr2, n2);

//     return 0;
// }

// TC: O(n^2)
// SC: O(1)

// optimized
void printLeadersOptimal(int arr[], int n)
{
    // Choosing the right most element as the maximum
    int max = arr[n - 1];
    cout << arr[n - 1] << " ";

    for (int i = n - 2; i >= 0; i--)
        if (arr[i] > max)
        {
            cout << arr[i] << " ";
            max = arr[i];
        }

    cout << "\n";
}

int main()
{

    int arr1[] = {4, 7, 1, 0};
    int n1 = sizeof(arr1) / sizeof(arr1[0]);
    cout << "The leaders of the first array are: " << endl;
    printLeadersOptimal(arr1, n1);

    int arr2[] = {10, 22, 12, 3, 0, 6};
    int n2 = sizeof(arr2) / sizeof(arr2[0]);
    cout << "The leaders of the second array are: " << endl;
    printLeadersOptimal(arr2, n2);

    return 0;
}

// TC: O(n)
// SC: O(1)