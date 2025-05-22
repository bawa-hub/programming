// https://leetcode.com/problems/rearrange-array-elements-by-sign/
// https://practice.geeksforgeeks.org/problems/array-of-alternate-ve-and-ve-nos1401/1
// https://www.geeksforgeeks.org/rearrange-array-alternating-positive-negative-items-o1-extra-space/
// https://takeuforward.org/arrays/rearrange-array-elements-by-sign/

#include <bits/stdc++.h>
using namespace std;

// brute force
vector<int> RearrangebySign(vector<int> A, int n)
{

    // Define 2 vectors, one for storing positive
    // and other for negative elements of the array.
    vector<int> pos;
    vector<int> neg;

    // Segregate the array into positives and negatives.
    for (int i = 0; i < n; i++)
    {

        if (A[i] > 0)
            pos.push_back(A[i]);
        else
            neg.push_back(A[i]);
    }

    // Positives on even indices, negatives on odd.
    for (int i = 0; i < n / 2; i++)
    {

        A[2 * i] = pos[i];
        A[2 * i + 1] = neg[i];
    }

    return A;
}
// Time Complexity: O(N+N/2) { O(N) for traversing the array once for segregating positives and negatives and another O(N/2) for adding those elements alternatively to the array, where N = size of the array A}.
// Space Complexity:  O(N/2 + N/2) = O(N) { N/2 space required for each of the positive and negative element arrays, where N = size of the array A}.

// optimized
vector<int> RearrangebySign(vector<int> A)
{

    int n = A.size();

    // Define array for storing the ans separately.
    vector<int> ans(n, 0);

    // positive elements start from 0 and negative from 1.
    int posIndex = 0, negIndex = 1;
    for (int i = 0; i < n; i++)
    {

        // Fill negative elements in odd indices and inc by 2.
        if (A[i] < 0)
        {
            ans[negIndex] = A[i];
            negIndex += 2;
        }

        // Fill positive elements in even indices and inc by 2.
        else
        {
            ans[posIndex] = A[i];
            posIndex += 2;
        }
    }

    return ans;
}
// Time Complexity: O(N) { O(N) for traversing the array once and substituting positives and negatives simultaneously using pointers, where N = size of the array A}.
// Space Complexity:  O(N) { Extra Space used to store the rearranged elements separately in an array, where N = size of array A}.

int main()
{

    // Array Initialisation.
    int n = 4;
    vector<int> A{1, 2, -4, -5};

    vector<int> ans = RearrangebySign(A, n);

    for (int i = 0; i < ans.size(); i++)
    {
        cout << ans[i] << " ";
    }

    return 0;
}