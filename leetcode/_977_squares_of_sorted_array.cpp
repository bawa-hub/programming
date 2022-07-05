class Solution
{
public:
    vector<int> sortedSquares(vector<int> &nums)
    {
        vector<int> result;
        for (int i = 0; i < nums.size(); i++)
        {
            int sq = nums[i] * nums[i];
            result.push_back(sq);
        }
        sort(result.begin(), result.end());
        return result;
    }
};