
import java.util.*;

public class Main {
    // brute force
    public static int findAllSubarraysWithGivenSum(int arr[], int k) {
        int n = arr.length; // size of the given array.
        int cnt = 0; // Number of subarrays:

        for (int i = 0; i < n; i++) { // starting index i
            for (int j = i; j < n; j++) { // ending index j

                // calculate the sum of subarray [i...j]
                int sum = 0;
                for (int K = i; K <= j; K++)
                    sum += arr[K];

                // Increase the count if sum == k:
                if (sum == k)
                    cnt++;
            }
        }
        return cnt;
    }

    // better approach
    public static int findAllSubarraysWithGivenSum1(int arr[], int k) {
        int n = arr.length; // size of the given array.
        int cnt = 0; // Number of subarrays:

        for (int i = 0; i < n; i++) { // starting index i
            int sum = 0;
            for (int j = i; j < n; j++) { // ending index j
                // calculate the sum of subarray [i...j]
                // sum of [i..j-1] + arr[j]
                sum += arr[j];

                // Increase the count if sum == k:
                if (sum == k)
                    cnt++;
            }
        }
        return cnt;
    }

    // optimal
    public static int findAllSubarraysWithGivenSum2(int[] arr, int k) {
        int n = arr.length; // size of the given array.
        Map<Integer, Integer> mpp = new HashMap<>();
        int preSum = 0, cnt = 0;

        // Setting 0 in the map to handle cases where the subarray itself has a sum equal to k
        mpp.put(0, 1); 

        for (int i = 0; i < n; i++) {
            // add current element to prefix sum
            preSum += arr[i];

            // Calculate preSum - k
            int remove = preSum - k;

            // Add the number of subarrays found that match the required sum
            cnt += mpp.getOrDefault(remove, 0);

            // Update the count of current prefix sum in the map
            mpp.put(preSum, mpp.getOrDefault(preSum, 0) + 1);
        }

        return cnt;
    }

    public static void main(String[] args) {
        int[] arr = { 3, 1, 2, 4 };
        int k = 6;
        int cnt = findAllSubarraysWithGivenSum(arr, k);
        System.out.println("The number of subarrays is: " + cnt);
    }
}
