import java.util.ArrayList;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {

        Scanner scanner = new Scanner(System.in);

        int n = scanner.nextInt();
        int m = scanner.nextInt();

        // Adjacency list for undirected graph
        // Time complexity: O(2E)
        ArrayList<Integer>[] adj = new ArrayList[n + 1];

        for (int i = 1; i <= n; i++) {
            adj[i] = new ArrayList<>();
        }

        for (int i = 0; i < m; i++) {
            int u = scanner.nextInt();
            int v = scanner.nextInt();

            adj[u].add(v);
            adj[v].add(u);
        }

        // Adjacency list for directed graph
        // Time complexity: O(E)
        for (int i = 0; i < m; i++) {
            int u = scanner.nextInt();
            int v = scanner.nextInt();

            adj[u].add(v);
        }

        // If weighted graph
        // Time complexity: O(E)
        for (int i = 0; i < m; i++) {
            int u = scanner.nextInt();
            int v = scanner.nextInt();
            int wt = scanner.nextInt();

            adj[u].add(v);
            adj[v].add(u);
        }

        // Close the scanner to avoid resource leaks
        scanner.close();
    }
}
