// https://leetcode.com/problems/sqrtx/

class Solution
{
public:
    // if ans will be nearest int
    int mySqrt(int x)
    {
        int l = 1, r = x;

        while (l <= r)
        {
            long mid = l + (r - l) / 2;

            if ((mid * mid) <= x)
                l = mid + 1;
            else
                r = mid - 1;
        }

        return r;
    }

    // for exact answer
    int mySqrt(int x)
    {
        double low = 1;
        double high = x;
        double eps = 1e-7;

        while ((high - low) > eps)
        {
            double mid = (low + high) / 2.0;
            if (mid * mid < x)
            {
                low = mid;
            }
            else
            {
                high = mid;
            }
        }

        return floor(high);
    }
};