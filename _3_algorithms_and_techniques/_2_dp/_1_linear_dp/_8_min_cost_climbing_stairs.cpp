// https://leetcode.com/problems/min-cost-climbing-stairs/
// https://leetcode.com/problems/min-cost-climbing-stairs/solutions/476388/4-ways-step-by-step-from-recursion-top-down-dp-bottom-up-dp-fine-tuning/
// https://leetcode.com/problems/min-cost-climbing-stairs/solutions/773865/a-beginner-s-guide-on-dp-validation-how-to-come-up-with-a-recursive-solution-python-3/

class Solution
{
public:
    int minCostClimbingStairs(vector<int> &cost)
    {
        int n = cost.size();
        vector<int> dp(n, -1);
        return min(minCost(n - 1, cost, dp), minCost(n - 2, cost, dp));
    }

    int minCost(int idx, vector<int> &cost, vector<int> &dp)
    {
        if (idx == 0 || idx == 1)
            return cost[idx];
        if (dp[idx] != -1)
            return dp[idx];
        return dp[idx] = cost[idx] + min(minCost(idx - 1, cost, dp), minCost(idx - 2, cost, dp));
    }
};