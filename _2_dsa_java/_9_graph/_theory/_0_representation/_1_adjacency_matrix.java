package _2_dsa_java._9_graph._theory._0_representation;

import java.util.Scanner;

public class _1_adjacency_matrix {
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
            adj[v][u] = 1; // Remove this line for directed graph
        }

        scanner.close();
    }
}
