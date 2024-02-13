// https://leetcode.com/problems/minimum-number-of-days-to-make-m-bouquets/

class Solution
{
public:
    bool possible(vector<int> &arr, int day, int m, int k)
    {
        int cnt = 0;
        int noOfB = 0;

        for (int i = 0; i < arr.size(); i++)
        {
            if (arr[i] <= day)
                cnt++;
            else
            {
                noOfB += (cnt / k);
                cnt = 0;
            }
        }

        noOfB += (cnt / k);
        return noOfB >= m;
    }

    int minDays(vector<int> &bloomDay, int m, int k)
    {
        long long val = m * 1LL * k * 1LL;

        if (val > bloomDay.size())
            return -1;

        int mini = INT_MAX, maxi = INT_MIN;
        for (int i = 0; i < bloomDay.size(); i++)
        {
            mini = min(mini, bloomDay[i]);
            maxi = max(maxi, bloomDay[i]);
        }

        int l = mini, r = maxi;
        while (l <= r)
        {
            int mid = l + (r - l) / 2;

            if (possible(bloomDay, mid, m, k))
                r = mid - 1;
            else
                l = mid + 1;
        }
        return l;
    }
};