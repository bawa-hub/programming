package _2_dsa_java._1_arrays._1_rotation;
// https://leetcode.com/problems/rotate-array/

public class _2_rotate_by_k_places {
}

// brute force
public class RotateArrayBruteRight {
    public static void rotateRight(int[] arr, int k) {
        int n = arr.length;
        if (n == 0) return;

        k = k % n;
        if (k > n) return;

        int[] temp = new int[k];
        for (int i = n - k; i < n; i++) {
            temp[i - (n - k)] = arr[i];
        }

        for (int i = n - k - 1; i >= 0; i--) {
            arr[i + k] = arr[i];
        }

        for (int i = 0; i < k; i++) {
            arr[i] = temp[i];
        }
    }

    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5, 6, 7};
        int k = 2;
        rotateRight(arr, k);
        System.out.print("Rotated Right: ");
        for (int num : arr) System.out.print(num + " ");
    }
}

// optimized revrsal algo
public class RotateArrayReverseRight {
    public static void reverse(int[] arr, int start, int end) {
        while (start < end) {
            int tmp = arr[start];
            arr[start] = arr[end];
            arr[end] = tmp;
            start++;
            end--;
        }
    }

    public static void rotateRight(int[] arr, int k) {
        int n = arr.length;
        k = k % n;

        reverse(arr, 0, n - k - 1);
        reverse(arr, n - k, n - 1);
        reverse(arr, 0, n - 1);
    }

    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5, 6, 7};
        int k = 2;
        rotateRight(arr, k);
        System.out.print("Rotated Right (Optimal): ");
        for (int num : arr) System.out.print(num + " ");
    }
}


// brute force
public class RotateArrayBruteLeft {
    public static void rotateLeft(int[] arr, int k) {
        int n = arr.length;
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
            arr[i] = temp[i - (n - k)];
        }
    }

    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5, 6, 7};
        int k = 2;
        rotateLeft(arr, k);
        System.out.print("Rotated Left: ");
        for (int num : arr) System.out.print(num + " ");
    }
}

// optimize reversal algo
public class RotateArrayReverseLeft {
    public static void reverse(int[] arr, int start, int end) {
        while (start < end) {
            int tmp = arr[start];
            arr[start] = arr[end];
            arr[end] = tmp;
            start++;
            end--;
        }
    }

    public static void rotateLeft(int[] arr, int k) {
        int n = arr.length;
        k = k % n;

        reverse(arr, 0, k - 1);
        reverse(arr, k, n - 1);
        reverse(arr, 0, n - 1);
    }

    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5, 6, 7};
        int k = 2;
        rotateLeft(arr, k);
        System.out.print("Rotated Left (Optimal): ");
        for (int num : arr) System.out.print(num + " ");
    }
}