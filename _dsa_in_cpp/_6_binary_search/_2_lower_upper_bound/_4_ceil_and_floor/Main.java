class Solution {

    // Function to find the ceiling of x in the array
    public int ceil(int[] arr, int n, int x) {
        int l = 0, r = n - 1;
        int ans = -1; // Default return value if no ceiling is found

        while (l <= r) {
            int mid = l + (r - l) / 2;

            if (arr[mid] >= x) {
                ans = arr[mid]; // Update the ceiling value
                r = mid - 1; // Continue searching on the left side
            } else {
                l = mid + 1; // Continue searching on the right side
            }
        }

        return ans;
    }

    // Function to find the floor of x in the array
    public int floor(int[] arr, int n, int x) {
        int l = 0, r = n - 1;
        int ans = -1; // Default return value if no floor is found

        while (l <= r) {
            int mid = l + (r - l) / 2;

            if (arr[mid] <= x) {
                ans = arr[mid]; // Update the floor value
                l = mid + 1; // Continue searching on the right side
            } else {
                r = mid - 1; // Continue searching on the left side
            }
        }

        return ans;
    }
}

public class Main {
    public static void main(String[] args) {
        Solution obj = new Solution();
        int[] arr = {1, 2, 8, 10, 10, 12, 19};
        int n = arr.length;
        int x = 5;

        int ceilResult = obj.ceil(arr, n, x);
        int floorResult = obj.floor(arr, n, x);

        System.out.println("Ceil of " + x + ": " + ceilResult);
        System.out.println("Floor of " + x + ": " + floorResult);
    }
}
