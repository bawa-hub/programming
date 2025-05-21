package _2_dsa_java._9_graph._theory._0_representation;

import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class _2_adjacency_list {
    // Helper class for weighted graphs
    static class Pair {
        int node;
        int weight;

        Pair(int node, int weight) {
            this.node = node;
            this.weight = weight;
        }
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        int n = scanner.nextInt(); // number of nodes
        int m = scanner.nextInt(); // number of edges

        // ----------------------------
        // 1. Undirected Unweighted Graph
        // ----------------------------
        List<List<Integer>> adjUnweightedUndirected = new ArrayList<>();
        for (int i = 0; i <= n; i++) {
            adjUnweightedUndirected.add(new ArrayList<>());
        }

        for (int i = 0; i < m; i++) {
            int u = scanner.nextInt();
            int v = scanner.nextInt();
            adjUnweightedUndirected.get(u).add(v);
            adjUnweightedUndirected.get(v).add(u); // Remove for directed
        }

        // ----------------------------
        // 2. Directed Unweighted Graph
        // ----------------------------
        List<List<Integer>> adjUnweightedDirected = new ArrayList<>();
        for (int i = 0; i <= n; i++) {
            adjUnweightedDirected.add(new ArrayList<>());
        }

        for (int i = 0; i < m; i++) {
            int u = scanner.nextInt();
            int v = scanner.nextInt();
            adjUnweightedDirected.get(u).add(v); // u -> v
        }

        // ----------------------------
        // 3. Undirected Weighted Graph
        // ----------------------------
        List<List<Pair>> adjWeightedUndirected = new ArrayList<>();
        for (int i = 0; i <= n; i++) {
            adjWeightedUndirected.add(new ArrayList<>());
        }

        for (int i = 0; i < m; i++) {
            int u = scanner.nextInt();
            int v = scanner.nextInt();
            int wt = scanner.nextInt();
            adjWeightedUndirected.get(u).add(new Pair(v, wt));
            adjWeightedUndirected.get(v).add(new Pair(u, wt)); // Remove for directed
        }

        // ----------------------------
        // 4. Directed Weighted Graph
        // ----------------------------
        List<List<Pair>> adjWeightedDirected = new ArrayList<>();
        for (int i = 0; i <= n; i++) {
            adjWeightedDirected.add(new ArrayList<>());
        }

        for (int i = 0; i < m; i++) {
            int u = scanner.nextInt();
            int v = scanner.nextInt();
            int wt = scanner.nextInt();
            adjWeightedDirected.get(u).add(new Pair(v, wt)); // u -> v
        }

        scanner.close();
    }
}
