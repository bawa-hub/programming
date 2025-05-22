// https://practice.geeksforgeeks.org/problems/number-of-occurrence2259/1

class Solution
{
public:
    int count(int arr[], int n, int x)
    {
        int f = first(n, x, arr);
        int l = last(n, x, arr);
        if (f == -1)
            return 0;
        return l - f + 1;
    }

    int last(int n, int key, int v[])
    {
        int start = 0;
        int end = n - 1;
        int res = -1;

        while (start <= end)
        {
            int mid = start + (end - start) / 2;
            if (v[mid] == key)
            {
                res = mid;
                start = mid + 1;
            }
            else if (key < v[mid])
            {
                end = mid - 1;
            }
            else
            {
                start = mid + 1;
            }
        }
        return res;
    }

    int first(int n, int key, int v[])
    {
        int start = 0;
        int end = n - 1;
        int res = -1;

        while (start <= end)
        {
            int mid = start + (end - start) / 2;
            if (v[mid] == key)
            {
                res = mid;
                end = mid - 1;
            }
            else if (key < v[mid])
            {
                end = mid - 1;
            }
            else
            {
                start = mid + 1;
            }
        }
        return res;
    }
};
// Time Complexity: O(logN)
// Space Complexity: O(1)
