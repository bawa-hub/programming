import java.util.Arrays;

public class Solution {
    public int numOfSubarrays(int[] arr, int k, int threshold) {
        int i = 0, j = 0, count = 0, sum = 0;
        int n = arr.length;

        while (j < n) {
            sum += arr[j];

            if (j - i + 1 == k) {
                if (sum / k >= threshold) {
                    count++;
                }
                sum -= arr[i];
                i++;
            }

            j++;
        }

        return count;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        int[] arr = { 2, 1, 3, 4, 1, 2, 1, 5, 4 };
        int k = 3;
        int threshold = 3;
        int result = solution.numOfSubarrays(arr, k, threshold);
        System.out.println(result);
    }
}
