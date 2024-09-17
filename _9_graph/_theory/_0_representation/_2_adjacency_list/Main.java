import java.util.*;

public class Main {

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        // Input number of nodes and edges
        int n = sc.nextInt();
        int m = sc.nextInt();

        // Adjacency list for unweighted graph
        List<List<Integer>> adj = new ArrayList<>();

        // Initialize the adjacency list
        for (int i = 0; i <= n; i++) {
            adj.add(new ArrayList<>());
        }

        // For undirected graph
        // Time complexity: O(2E)
        for (int i = 0; i < m; i++) {
            int u = sc.nextInt();
            int v = sc.nextInt();
            adj.get(u).add(v);
            adj.get(v).add(u);
        }

        // For directed graph
        // Time complexity: O(E)
        for (int i = 0; i < m; i++) {
            int u = sc.nextInt();
            int v = sc.nextInt();
            adj.get(u).add(v);
        }

        // Adjacency list for weighted graph
        List<List<int[]>> adjWeighted = new ArrayList<>();

        // Initialize the adjacency list for the weighted graph
        for (int i = 0; i <= n; i++) {
            adjWeighted.add(new ArrayList<>());
        }

        // For undirected weighted graph
        for (int i = 0; i < m; i++) {
            int u = sc.nextInt();
            int v = sc.nextInt();
            int wt = sc.nextInt();
            adjWeighted.get(u).add(new int[]{v, wt});
            adjWeighted.get(v).add(new int[]{u, wt});
        }

        // For directed weighted graph
        for (int i = 0; i < m; i++) {
            int u = sc.nextInt();
            int v = sc.nextInt();
            int wt = sc.nextInt();
            adjWeighted.get(u).add(new int[]{v, wt});
        }

        sc.close();
    }
}
