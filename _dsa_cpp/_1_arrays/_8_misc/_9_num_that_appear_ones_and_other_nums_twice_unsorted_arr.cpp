// https://leetcode.com/problems/single-number/
// https://takeuforward.org/arrays/find-the-number-that-appears-once-and-the-other-numbers-twice/
// https://practice.geeksforgeeks.org/problems/element-appearing-once2552/0

#include <bits/stdc++.h>
using namespace std;

// brute force
int getSingleElement(vector<int> &arr)
{
    // Size of the array:
    int n = arr.size();

    // Run a loop for selecting elements:
    for (int i = 0; i < n; i++)
    {
        int num = arr[i]; // selected element
        int cnt = 0;

        // find the occurrence using linear search:
        for (int j = 0; j < n; j++)
        {
            if (arr[j] == num)
                cnt++;
        }

        // if the occurrence is 1 return ans:
        if (cnt == 1)
            return num;
    }

    // This line will never execute
    // if the array contains a single element.
    return -1;
}

// Time Complexity: O(N2), where N = size of the given array.
// Reason: For every element, we are performing a linear search to count its occurrence. The linear search takes O(N) time complexity. And there are N elements in the array. So, the total time complexity will be O(N2).
// Space Complexity: O(1) as we are not using any extra space.

// using hashing
int getSingleElement(vector<int> &arr)
{

    // size of the array:
    int n = arr.size();

    // Find the maximum element:
    int maxi = arr[0];
    for (int i = 0; i < n; i++)
    {
        maxi = max(maxi, arr[i]);
    }

    // Declare hash array of size maxi+1
    // And hash the given array:
    vector<int> hash(maxi + 1, 0);
    for (int i = 0; i < n; i++)
    {
        hash[arr[i]]++;
    }

    // Find the single element and return the answer:
    for (int i = 0; i < n; i++)
    {
        if (hash[arr[i]] == 1)
            return arr[i];
    }

    // This line will never execute
    // if the array contains a single element.
    return -1;
}
// Time Complexity: O(N)+O(N)+O(N), where N = size of the array
// Reason: One O(N) is for finding the maximum, the second one is to hash the elements and the third one is to search the single element in the array.
// Space Complexity: O(maxElement+1) where maxElement = the maximum element of the array.

// using xor
int getSingleElement(vector<int> &arr)
{
    // size of the array:
    int n = arr.size();

    // XOR all the elements:
    int xorr = 0;
    for (int i = 0; i < n; i++)
    {
        xorr = xorr ^ arr[i];
    }
    return xorr;
}

int main()
{
    vector<int> arr = {4, 1, 2, 1, 2};
    int ans = getSingleElement(arr);
    cout << "The single element is: " << ans << endl;
    return 0;
}