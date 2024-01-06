package _3_algorithms._2_dp._1_linear_dp._2_climbing_stairs;

import java.util.Arrays;

public class Main {

    // memoization
    public static int climbStairs(int n) {
        int dp[] = new int[n + 1];
        Arrays.fill(dp, -1);
        return f(n, dp);
    }

    public static int f(int n, int[] dp) {
        if (n == 0 || n == 1)
            return 1;

        if (dp[n] != -1)
            return dp[n];

        return dp[n] = f(n - 1, dp) + f(n - 2, dp);
    }

    // tabulation
    public static void iterative() {
        int n = 3;
        int dp[] = new int[n + 1];
        dp[0] = 1;
        dp[1] = 1;

        for (int i = 2; i <= n; i++) {
            dp[i] = dp[i - 1] + dp[i - 2];
        }

        System.out.println(dp[n]);
    }

    // space optimized
    public static void spaceoptimized() {
        int n = 3;

        int prev2 = 1;
        int prev = 1;

        for (int i = 2; i <= n; i++) {
            int cur_i = prev2 + prev;
            prev2 = prev;
            prev = cur_i;
        }
        System.out.println(prev);
    }

    public static void main(String[] args) {
        int n = 5;
        System.out.println(climbStairs(n));
    }
}
