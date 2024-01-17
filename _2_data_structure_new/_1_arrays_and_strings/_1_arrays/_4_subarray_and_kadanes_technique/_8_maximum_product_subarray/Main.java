
import java.util.*;

public class Main {

    // brute force
    static int maxProductSubArray(int arr[]) {
        int result = Integer.MIN_VALUE;
        for (int i = 0; i < arr.length - 1; i++)
            for (int j = i + 1; j < arr.length; j++) {
                int prod = 1;
                for (int k = i; k <= j; k++)
                    prod *= arr[k];
                result = Math.max(result, prod);
            }
        return result;
    }

    // better approach
    static int maxProductSubArray1(int arr[]) {
        int result = arr[0];
        for (int i = 0; i < arr.length - 1; i++) {
            int p = arr[i];
            for (int j = i + 1; j < arr.length; j++) {
                result = Math.max(result, p);
                p *= arr[j];
            }
            result = Math.max(result, p);
        }
        return result;
    }

    // optimal approach 1
    public static int maxProductSubArray2(int[] arr) {
        int n = arr.length; // size of array.

        int pre = 1, suff = 1;
        int ans = Integer.MIN_VALUE;
        for (int i = 0; i < n; i++) {
            if (pre == 0)
                pre = 1;
            if (suff == 0)
                suff = 1;
            pre *= arr[i];
            suff *= arr[n - i - 1];
            ans = Math.max(ans, Math.max(pre, suff));
        }
        return ans;
    }

    // optimal approach 2 - kadane's algorithm

    static int maxProductSubArray3(int arr[]) {
        int prod1 = arr[0], prod2 = arr[0], result = arr[0];

        for (int i = 1; i < arr.length; i++) {
            int temp = Math.max(arr[i], Math.max(prod1 * arr[i], prod2 * arr[i]));
            prod2 = Math.min(arr[i], Math.min(prod1 * arr[i], prod2 * arr[i]));
            prod1 = temp;

            result = Math.max(result, prod1);
        }

        return result;
    }

    public static void main(String[] args) {
        int nums[] = { 1, 2, -3, 0, -4, -5 };
        int answer = maxProductSubArray(nums);
        System.out.print("The maximum product subarray is: " + answer);
    }
}
