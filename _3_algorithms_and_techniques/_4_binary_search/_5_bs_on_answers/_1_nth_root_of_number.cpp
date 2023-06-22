// https://practice.geeksforgeeks.org/problems/find-nth-root-of-m5843/1

class Solution
{
public:
    int func(int mid, int n, int m)
    {
        long long ans = 1;
        for (int i = 1; i <= n; i++)
        {
            ans = ans * mid;
            if (ans > m)
                return 2;
        }

        if (ans == m)
            return 1;
        return 0;
    }

    int NthRoot(int n, int m)
    {
        int l = 1, r = m;
        while (l <= r)
        {
            int mid = l + (r - l) / 2;
            int midN = func(mid, n, m);
            if (midN == 1)
                return mid;
            else if (midN == 0)
                l = mid + 1;
            else
                r = mid - 1;
        }
        return -1;
    }
};