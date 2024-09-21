import java.util.*;

public class Main {
    
    // Recursive method to print the name 'n' times
    public static void print(String name, int n) {
        if (n == 0) {
            return;  // Base case: stop when n becomes 0
        }
        System.out.println(name);
        print(name, n - 1);  // Recursive call
    }

    public static void main(String[] args) {
        String name = "Bawa";
        int n = 10;  // Number of times to print the name
        print(name, n);  // Call the recursive function
    }
}
