// https://leetcode.com/problems/sqrtx/

class Solution
{
public:
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