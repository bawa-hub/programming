class Main {
    public int minCostClimbingStairs(int[] cost) {
        int n = cost.length;
        int[] dp = new int[n];
        Arrays.fill(dp, -1);
        return Math.min(minCost(n - 1, cost, dp), minCost(n - 2, cost, dp));
    }

    public int minCost(int idx, int[] cost, int[] dp) {
        if (idx == 0 || idx == 1)
            return cost[idx];
        if (dp[idx] != -1)
            return dp[idx];
        return dp[idx] = cost[idx] + Math.min(minCost(idx - 1, cost, dp), minCost(idx - 2, cost, dp));
    }
}