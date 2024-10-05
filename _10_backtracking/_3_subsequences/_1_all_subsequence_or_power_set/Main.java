import java.util.*;

public class Main {

    // Function to print all subsequences of an array
    public static void printArraySubsequences(int idx, List<Integer> ds, int[] arr, int n) {
        if (idx == n) {
            System.out.print("{ ");
            for (int it : ds) {
                System.out.print(it + " ");
            }
            System.out.println("}");
            return;
        }
        
        // Not picking the current element
        printArraySubsequences(idx + 1, ds, arr, n);
        
        // Picking the current element
        ds.add(arr[idx]);
        printArraySubsequences(idx + 1, ds, arr, n);
        
        // Backtracking step
        ds.remove(ds.size() - 1);
    }

    // Function to print all subsequences of a string
    public static void printStringSubsequences(int idx, String s, StringBuilder currentSubseq) {
        if (idx == s.length()) {
            System.out.print(currentSubseq.toString() + " ");
            return;
        }
        
        // Picking the current character
        currentSubseq.append(s.charAt(idx));
        printStringSubsequences(idx + 1, s, currentSubseq);
        
        // Backtracking step (remove the last character)
        currentSubseq.deleteCharAt(currentSubseq.length() - 1);
        
        // Not picking the current character
        printStringSubsequences(idx + 1, s, currentSubseq);
    }

    public static void main(String[] args) {
        // Array subsequences
        System.out.println("Subsequences of the array:");
        int[] arr = {3, 1, 2};
        List<Integer> ds = new ArrayList<>();
        printArraySubsequences(0, ds, arr, arr.length);

        // String subsequences
        System.out.println("\nSubsequences of the string:");
        String s = "abc";
        StringBuilder currentSubseq = new StringBuilder();
        printStringSubsequences(0, s, currentSubseq);
    }
}
