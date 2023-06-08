// https://www.codingninjas.com/codestudio/problems/lower-bound_8165382
// https://www.codingninjas.com/codestudio/problems/implement-upper-bound_8165383

// lower_bound = smallest index such that arr[index] >= x
// upper_bound = smallest index such that arr[index] > x

#include <vector>
using namespace std;

// lower_bound
int lowerBound(vector<int> arr, int n, int x)
{
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

// upper_bound
int upperBound(vector<int> arr, int n, int x)
{
    int l = 0, r = n - 1;
    int ans = n; // if not found till last element

    while (l <= r)
    {
        int mid = l + (r - l) / 2;

        // may be answer
        if (arr[mid] > x)
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

// note: it can also be done without using ans variable