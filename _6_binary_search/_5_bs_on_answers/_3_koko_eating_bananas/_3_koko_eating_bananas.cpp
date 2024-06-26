// https://leetcode.com/problems/koko-eating-bananas/

class Solution
{
public:
    long long getHoursToEatAll(vector<int> &piles, int bananasPerHour)
    {
        long long totalHours = 0;
        for (int i = 0; i < piles.size(); i++)
        {
            int hoursToEatPile = ceil(piles[i] / (double)bananasPerHour);
            totalHours += hoursToEatPile;
        }
        return totalHours;
    }

    // with ans variable
    int minEatingSpeed(vector<int> &piles, int targetHours)
    {
        int low = 1, high = *(max_element(piles.begin(), piles.end()));
        int ans = -1;

        while (low <= high)
        {
            int mid = low + (high - low) / 2;
            long long hoursToEatAll = getHoursToEatAll(piles, mid);

            if (hoursToEatAll <= targetHours)
            {
                ans = mid; // record the answer (this is the best we could record till curr step)
                high = mid - 1;
            }
            else
                low = mid + 1;
        }

        return ans;
    }

    // without ans variable
    int minEatingSpeed(vector<int> &piles, int targetHours)
    {
        int low = 1, high = *(max_element(piles.begin(), piles.end()));
        while (low <= high)
        {
            int mid = low + (high - low) / 2;
            long long hoursToEatAll = getHoursToEatAll(piles, mid);

            if (hoursToEatAll <= targetHours)
            {
                high = mid - 1;
            }
            else
                low = mid + 1;
        }

        return low;
    }
};