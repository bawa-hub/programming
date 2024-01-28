import java.util.Arrays;

public class Solution {
    public int minimumDifference(int[] nums, int k) {
        Arrays.sort(nums);

        int i = 0, j = 0;
        int mini = Integer.MAX_VALUE;

        while (j < nums.length) {
            if (j - i + 1 == k) {
                int diff = nums[j] - nums[i];
                mini = Math.min(mini, diff);
                i++;
            }
            j++;
        }

        return mini;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        int[] nums = { 3, 8, 1, 10, 6 };
        int k = 3;
        int result = solution.minimumDifference(nums, k);
        System.out.println(result);
    }
}
