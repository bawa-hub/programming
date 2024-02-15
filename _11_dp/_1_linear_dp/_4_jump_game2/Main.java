import java.util.Arrays;

public class Main {
    public int jump(int[] nums) {
        int n = nums.length;
        int[] dp = new int[n];
        Arrays.fill(dp, -1);
        return f(0, nums, dp);
    }

    int f(int idx, int[] nums, int[] dp) {
        if (idx >= nums.length - 1)
            return 0;
        if (dp[idx] != -1)
            return dp[idx];

        int mini = (int) 1e6;
        for (int i = 1; i <= nums[idx]; i++) {
            if (i + idx < nums.length) {
                mini = Math.min(mini, 1 + f(idx + i, nums, dp));
            }
        }

        return dp[idx] = mini;
    }
}Main{

}
