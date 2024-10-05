import java.util.Arrays;

public class Main {

    // Brute force method to rotate elements to the right
    public static void rotateToRight(int[] arr, int n, int k) {
        if (n == 0) return;

        k = k % n;
        if (k > n) return;

        int[] temp = new int[k];
        for (int i = n - k; i < n; i++) {
            temp[i - n + k] = arr[i];
        }
        for (int i = n - k - 1; i >= 0; i--) {
            arr[i + k] = arr[i];
        }
        for (int i = 0; i < k; i++) {
            arr[i] = temp[i];
        }
    }

    // Brute force method to rotate elements to the left
    public static void rotateToLeft(int[] arr, int n, int k) {
        if (n == 0) return;

        k = k % n;
        if (k > n) return;

        int[] temp = new int[k];
        for (int i = 0; i < k; i++) {
            temp[i] = arr[i];
        }
        for (int i = 0; i < n - k; i++) {
            arr[i] = arr[i + k];
        }
        for (int i = n - k; i < n; i++) {
            arr[i] = temp[i - n + k];
        }
    }

    // Function to reverse the array
    public static void reverse(int[] arr, int start, int end) {
        while (start <= end) {
            int temp = arr[start];
            arr[start] = arr[end];
            arr[end] = temp;
            start++;
            end--;
        }
    }

    // Reversal algorithm to rotate k elements to the right
    public static void rotateEleToRight(int[] arr, int n, int k) {
        k = k % n; // for k > n
        // Reverse first n-k elements
        reverse(arr, 0, n - k - 1);
        // Reverse last k elements
        reverse(arr, n - k, n - 1);
        // Reverse the whole array
        reverse(arr, 0, n - 1);
    }

    // Reversal algorithm to rotate k elements to the left
    public static void rotateEleToLeft(int[] arr, int n, int k) {
        k = k % n; // for k > n
        // Reverse first k elements
        reverse(arr, 0, k - 1);
        // Reverse last n-k elements
        reverse(arr, k, n - 1);
        // Reverse the whole array
        reverse(arr, 0, n - 1);
    }

    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5, 6, 7};
        int n = arr.length;
        int k = 2;

        // Rotate elements to the right
        rotateEleToRight(arr, n, k);
        System.out.println("After rotating k elements to the right: " + Arrays.toString(arr));

        // Resetting the array for left rotation
        arr = new int[]{1, 2, 3, 4, 5, 6, 7};

        // Rotate elements to the left
        rotateEleToLeft(arr, n, k);
        System.out.println("After rotating k elements to the left: " + Arrays.toString(arr));
    }
}
