package _2_dsa_java._9_graph._theory._0_representation;

import java.util.ArrayList;
import java.util.List;

public class _3_adjacency_matrix_to_list {
     public static void main(String[] args) {
        // Adjacency Matrix
        int[][] adj = {
            {1, 0, 1},
            {0, 1, 0},
            {1, 0, 1}
        };

        int V = adj.length;

        // Adjacency List
        List<List<Integer>> adjList = new ArrayList<>();
        for (int i = 0; i < V; i++) {
            adjList.add(new ArrayList<>());
        }

        for (int i = 0; i < V; i++) {
            for (int j = 0; j < V; j++) {
                // Avoid self loops and avoid duplicating edges
                if (adj[i][j] == 1 && i != j) {
                    adjList.get(i).add(j);
                }
            }
        }

        // Optional: Print adjacency list
        for (int i = 0; i < V; i++) {
            System.out.print(i + ": ");
            for (int neighbor : adjList.get(i)) {
                System.out.print(neighbor + " ");
            }
            System.out.println();
        }
    }
}
