// https://leetcode.com/problems/kth-missing-positive-number/

#include <bits/stdc++.h>
using namespace std;

// brute force
int missingK(vector < int > vec, int n, int k) {
    for (int i = 0; i < n; i++) {
        if (vec[i] <= k) k++; //shifting k
        else break;
    }
    return k;
}
// Time Complexity: O(N), N = size of the given array.
// Reason: We are using a loop that traverses the entire given array in the worst case.
// Space Complexity: O(1) as we are not using any extra space to solve this problem.

// binary search

int missingK(vector < int > vec, int n, int k) {
    int low = 0, high = n - 1;
    while (low <= high) {
        int mid = (low + high) / 2;
        int missing = vec[mid] - (mid + 1);
        if (missing < k) {
            low = mid + 1;
        }
        else {
            high = mid - 1;
        }
    }
    return k + high + 1;
}
// Time Complexity: O(logN), N = size of the given array.
// Reason: We are using the simple binary search algorithm.
// Space Complexity: O(1) as we are not using any extra space to solve this problem.

int main()
{
    vector<int> vec = {4, 7, 9, 10};
    int n = 4, k = 4;
    int ans = missingK(vec, n, k);
    cout << "The missing number is: " << ans << "\n";
    return 0;
}