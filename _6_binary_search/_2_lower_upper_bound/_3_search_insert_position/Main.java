class Solution {

    public int searchInsert(int[] arr, int x) {
        int n = arr.length;
        int l = 0, r = n - 1;
        int ans = n; // If the element is not found till the last element

        while (l <= r) {
            int mid = l + (r - l) / 2;

            // If the middle element is greater than or equal to x, it could be the answer
            if (arr[mid] >= x) {
                ans = mid;
                r = mid - 1; // Search in the left part for a smaller index
            } else {
                l = mid + 1; // Search in the right part
            }
        }

        return ans;
    }
}

public class Main {
    public static void main(String[] args) {
        Solution obj = new Solution();
        int[] arr = {1, 3, 5, 6};
        int target = 5;
        int result = obj.searchInsert(arr, target);
        System.out.println("Insert position: " + result);
    }
}
