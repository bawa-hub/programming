// Q. https://practice.geeksforgeeks.org/problems/minimum-platforms-1587115620/1#

// Given arrival and departure times of all trains that reach a railway station.
// Find the minimum number of platforms required for the railway station so that no train is kept waiting.

// Consider that all the trains arrive on the same day and leave on the same day.
// Arrival and departure time can never be the same for a train but we can have arrival time of one train equal to departure time of the other.
// At any given instance of time, same platform can not be used for both departure of a train and arrival of another train.
// In such cases, we need different platforms.

#include <bits/stdc++.h>
using namespace std;

// naive approach
// int countPlatforms(int n, int arr[], int dep[])
// {
//     int ans = 1; // final value
//     for (int i = 0; i <= n - 1; i++)
//     {
//         int count = 1; // count of overlapping interval of only this   iteration
//         for (int j = i + 1; j <= n - 1; j++)
//         {
//             if ((arr[i] >= arr[j] && arr[i] <= dep[j]) ||
//                 (arr[j] >= arr[i] && arr[j] <= dep[i]))
//             {
//                 count++;
//             }
//         }
//         ans = max(ans, count); // updating the value
//     }
//     return ans;
// }

// int main()
// {
//     int arr[] = {900, 945, 955, 1100, 1500, 1800};
//     int dep[] = {920, 1200, 1130, 1150, 1900, 2000};
//     int n = sizeof(dep) / sizeof(dep[0]);
//     cout << "Minimum number of Platforms required " << countPlatforms(n, arr, dep) << endl;
// }

// efficient approach
int countPlatforms(int n, int arr[], int dep[])
{
    sort(arr, arr + n);
    sort(dep, dep + n);

    int ans = 1;
    int count = 1;
    int i = 1, j = 0;
    while (i < n && j < n)
    {
        if (arr[i] <= dep[j]) // one more platform needed
        {
            count++;
            i++;
        }
        else // one platform can be reduced
        {
            count--;
            j++;
        }
        ans = max(ans, count); // updating the value with the current maximum
    }
    return ans;
}

int main()
{
    int arr[] = {900, 945, 955, 1100, 1500, 1800};
    int dep[] = {920, 1200, 1130, 1150, 1900, 2000};
    int n = sizeof(dep) / sizeof(dep[0]);
    cout << "Minimum number of Platforms required " << countPlatforms(n, arr, dep) << endl;
}
