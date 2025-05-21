package _2_dsa_java._3_techniques._1_hash_table;
// https://practice.geeksforgeeks.org/problems/frequency-of-array-elements-1587115620/0
import java.util.HashMap;
import java.util.Map;

public class _1_frequency_array_elements {

    public static void frequency(int[] arr) {
        HashMap<Integer, Integer> map = new HashMap<>();

        // Count frequencies
        for (int num : arr) {
            map.put(num, map.getOrDefault(num, 0) + 1);
        }

        // Print frequencies
        for (Map.Entry<Integer, Integer> entry : map.entrySet()) {
            System.out.println(entry.getKey() + " " + entry.getValue());
        }
    }

    public static void main(String[] args) {
        int[] arr = {10, 5, 10, 15, 10, 5};
        frequency(arr);
    }
}

