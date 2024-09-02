// https://practice.geeksforgeeks.org/problems/frequency-of-array-elements-1587115620/0

import java.util.HashMap;
import java.util.Map;

public class Main {

    public static void Frequency(int[] arr, int n) {
        // Create a HashMap to store the frequency of elements
        HashMap<Integer, Integer> map = new HashMap<>();

        // Iterate through the array and update the frequency in the map
        for (int i = 0; i < n; i++) {
            map.put(arr[i], map.getOrDefault(arr[i], 0) + 1);
        }

        // Traverse through the map and print the frequencies
        for (Map.Entry<Integer, Integer> entry : map.entrySet()) {
            System.out.println(entry.getKey() + " " + entry.getValue());
        }
    }

    public static void main(String[] args) {
        int[] arr = {10, 5, 10, 15, 10, 5};
        int n = arr.length;
        Frequency(arr, n);
    }
}
