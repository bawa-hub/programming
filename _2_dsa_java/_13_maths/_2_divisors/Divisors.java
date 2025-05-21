package _2_dsa_java._13_maths._2_divisors;
public class Divisors {

    // Approach 1: Brute Force
    public static void printDivisorsBruteForce(int n) {
        System.out.println("The Divisors of " + n + " are:");
        for (int i = 1; i <= n; i++) {
            if (n % i == 0)
                System.out.print(i + " ");
        }
        System.out.println();
    }
// Time Complexity: O(n), because the loop has to run from 1 to n always.
// Space Complexity: O(1), we are not using any extra space

    // Approach 2: Optimal
    public static void printDivisorsOptimal(int n) {
        System.out.println("The Divisors of " + n + " are:");
        for (int i = 1; i <= Math.sqrt(n); i++) {
            if (n % i == 0) {
                System.out.print(i + " ");
                if (i != n / i)
                    System.out.print((n / i) + " ");
            }
        }
        System.out.println();
    }
// Time Complexity: O(sqrt(n)), because everytime the loop runs only sqrt(n) times.
// Space Complexity: O(1), we are not using any extra space

    public static void main(String[] args) {
        printDivisorsBruteForce(36);
        printDivisorsOptimal(36);
    }
}
