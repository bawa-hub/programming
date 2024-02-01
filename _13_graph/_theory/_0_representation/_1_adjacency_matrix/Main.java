import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        int n = scanner.nextInt();
        int m = scanner.nextInt();

        // Adjacency matrix for undirected graph
        // Time complexity: O(n), Space complexity: O(n^2)
        int[][] adj = new int[n + 1][n + 1];

        for (int i = 0; i < m; i++) {
            int u = scanner.nextInt();
            int v = scanner.nextInt();

            adj[u][v] = 1;
            adj[v][u] = 1;  // This statement will be removed in case of a directed graph
        }

        // Close the scanner to avoid resource leaks
        scanner.close();
    }
}
