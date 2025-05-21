// https://leetcode.com/problems/search-insert-position/
#include <vector>
using namespace std;

// brute force
int find_index(int arr[], int n, int K)
{
    for (int i = 0; i < n; i++)
        if (arr[i] == K)
            return i;
        else if (arr[i] > K)
            return i;
    return n;
}
// Time Complexity: O(N)
// Auxiliary Space: O(1)

// same as lower bound
class Solution
{
public:
    int searchInsert(vector<int> &arr, int x)
    {
        int n = arr.size();
        int l = 0, r = n - 1;
        int ans = n; // if not found till last element

        while (l <= r)
        {
            int mid = l + (r - l) / 2;

            // may be answer
            if (arr[mid] >= x)
            {
                ans = mid;
                r = mid - 1; // look for more small index on left
            }
            else
            {
                l = mid + 1;
            }
        }

        return ans;
    }
};
// Time Complexity: O(log N)
// Auxiliary Space: O(1)
