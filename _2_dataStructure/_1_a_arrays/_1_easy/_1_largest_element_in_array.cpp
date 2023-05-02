// https://practice.geeksforgeeks.org/problems/largest-element-in-array4009/0?utm_source=youtube&utm_medium=collab_striver_ytdescription&utm_campaign=largest-element-in-array
// https://takeuforward.org/data-structure/find-the-largest-element-in-an-array/

#include <bits/stdc++.h>

using namespace std;

// approach 1
int sortArr(vector<int> &arr)
{
    sort(arr.begin(), arr.end());
    return arr[arr.size() - 1];
}
// Time Complexity: O(N*log(N))
// Space Complexity: O(n)

// approach 2
int findLargestElement(int arr[], int n)
{

    int max = arr[0];
    for (int i = 0; i < n; i++)
    {
        if (max < arr[i])
        {
            max = arr[i];
        }
    }
    return max;
}
// Time Complexity: O(N)
// Space Complexity: O(1)

int main()
{
    int arr1[] = {2, 5, 1, 3, 0};
    int n = 5;
    int max = findLargestElement(arr1, n);
    cout << "The largest element in the array is: " << max << endl;

    int arr2[] = {8, 10, 5, 7, 9};
    n = 5;
    max = findLargestElement(arr2, n);
    cout << "The largest element in the array is: " << max << endl;
    return 0;
}