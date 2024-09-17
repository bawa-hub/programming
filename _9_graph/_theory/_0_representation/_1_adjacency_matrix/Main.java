import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        // Input number of nodes and edges
        int n = sc.nextInt();
        int m = sc.nextInt();

        // Adjacency matrix for undirected graph
        // Time complexity: O(n)
        // Space complexity: O(n^2)
        int[][] adj = new int[n + 1][n + 1]; // 1-based indexing

        // Input edges
        for (int i = 0; i < m; i++) {
            int u = sc.nextInt();
            int v = sc.nextInt();

            // Set the edge in the adjacency matrix
            adj[u][v] = 1;
            
            // This statement will be removed in case of directed graph
            adj[v][u] = 1;
        }

        sc.close();
    }
}
