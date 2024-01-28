import java.util.Arrays;

public class Solution {
    public int[] getAverages(int[] nums, int k) {
        int i = 0, j = 0, z = k, n = nums.length;
        long sum = 0;
        int[] res = new int[n];
        Arrays.fill(res, -1);

        while (j < n) {
            sum += nums[j];

            if (j - i == 2 * k - 1) {
                long avg = sum / (j - i + 1);
                res[z++] = (int) avg;
                sum -= nums[i];
                i++;
            }

            j++;
        }

        return res;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        int[] nums = { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 };
        int k = 3;
        int[] result = solution.getAverages(nums, k);
        System.out.println(Arrays.toString(result));
    }
}
