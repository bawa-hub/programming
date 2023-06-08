// https://takeuforward.org/data-structure/last-occurrence-in-a-sorted-array/

using namespace std;

// last occurence of number

// brute force
int solve(int n, int key, vector<int> &v)
{
    int res = -1;
    for (int i = n - 1; i >= 0; i--)
    {
        if (v[i] == key)
        {
            res = i;
            break;
        }
    }
    return res;
}
// Time Complexity: O(n)
// Space Complexity: O(1) not considering the given array

// binary search
int solve(int n, int key, vector<int> &v)
{
    int start = 0, end = n - 1, res = -1;

    while (start <= end)
    {
        int mid = start + (end - start) / 2;
        if (v[mid] == key)
        {
            res = mid;
            start = mid + 1;
        }
        else if (key < v[mid])
            end = mid - 1;
        else
            start = mid + 1;
    }
    return res;
}
// Time Complexity: O(log n)
// Space Complexity: O(1)