import java.util.*;

class Solution {

    public int[] nextGreaterElement(int[] nums1, int[] nums2) {
        Map<Integer, Integer> mpx = new HashMap<>();
        Stack<Integer> stack = new Stack<>();

        // Traverse nums2 from right to left
        for (int i = nums2.length - 1; i >= 0; i--) {
            while (!stack.isEmpty() && stack.peek() <= nums2[i]) {
                stack.pop();
            }

            if (!stack.isEmpty()) {
                mpx.put(nums2[i], stack.peek());
            } else {
                mpx.put(nums2[i], -1);
            }

            stack.push(nums2[i]);
        }

        // Prepare the result array for nums1 based on the map
        int[] ans = new int[nums1.length];
        for (int i = 0; i < nums1.length; i++) {
            ans[i] = mpx.get(nums1[i]);
        }

        return ans;
    }
}

public class Main {
    public static void main(String[] args) {
        Solution obj = new Solution();
        int[] v = {5, 7, 1, 2, 6, 0};
        int[] res = obj.nextGreaterElement(v, new int[]{5, 7, 1, 2, 6, 0});
        System.out.println("The next greater elements are:");
        for (int i = 0; i < res.length; i++) {
            System.out.print(res[i] + " ");
        }
    }
}

