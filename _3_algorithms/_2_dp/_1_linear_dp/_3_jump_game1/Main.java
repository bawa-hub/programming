import java.util.*;

class Main {
    public boolean canJump(int[] nums) {
        int n = nums.length;
        int[] dp = new int[n];
        Arrays.fill(dp, -1);
        if (f(0, nums, dp) == 1)
            return true;
        return false;
    }

    public int f(int idx, int[] nums, int[] dp) {
        if (idx == nums.length - 1)
            return 1;
        if (idx >= nums.length)
            return 0;

        if (dp[idx] != -1)
            return dp[idx];

        for (int i = 1; i <= nums[idx]; i++) {
            if (i + idx < nums.length) {
                if (f(idx + i, nums, dp) == 1)
                    return 1;
            }
        }

        return dp[idx] = 0;
    }
}

// tabulation
class Solution {
    public boolean canJump(int[] nums) {
        int n = nums.length;
        int[] dp = new int[n];
        Arrays.fill(dp, -1);

        dp[n - 1] = 1; // base case;

        for (int idx = n - 2; idx >= 0; idx--) {
            if (nums[idx] == 0) {
                dp[idx] = 0;
                continue;
            }

            int flag = 0;
            int reach = idx + nums[idx];
            for (int jump = idx + 1; jump <= reach; jump++) {
                if (jump < nums.length && dp[jump] == 1) {
                    dp[idx] = 1;
                    flag = 1;
                    break;
                }
            }
            if (flag == 1)
                continue;

            dp[idx] = 0;

        }
        return dp[0] == 1 ? true : false;
    }
}