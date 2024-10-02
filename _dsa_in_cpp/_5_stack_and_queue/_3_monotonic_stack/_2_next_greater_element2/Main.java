import java.util.*;

class Solution {

    public int[] nextGreaterElements(int[] nums) {
        int n = nums.length;
        int[] nge = new int[n]; // Array to store next greater elements
        Arrays.fill(nge, -1); // Initialize the array with -1
        Stack<Integer> stack = new Stack<>();

        // Traverse the array twice to handle circular nature
        for (int i = 2 * n - 1; i >= 0; i--) {
            // Modulo operation allows the array to be treated circularly
            while (!stack.isEmpty() && stack.peek() <= nums[i % n]) {
                stack.pop();
            }

            if (i < n) {
                if (!stack.isEmpty()) {
                    nge[i] = stack.peek();
                }
            }
            stack.push(nums[i % n]);
        }

        return nge;
    }
}

public class Main {
    public static void main(String[] args) {
        Solution obj = new Solution();
        int[] v = {5, 7, 1, 2, 6, 0};
        int[] res = obj.nextGreaterElements(v);
        System.out.println("The next greater elements are:");
        for (int i = 0; i < res.length; i++) {
            System.out.print(res[i] + " ");
        }
    }
}
