package _2_dsa_java._9_graph._problems._0_representaion;

import java.util.ArrayList;
import java.util.List;

public class _1_center_of_star_graph {
     public int findCenter(int[][] edges) {
        int n = edges.length;
        List<List<Integer>> adj = new ArrayList<>();

        // Create adjacency list with size n+2 (to be safe)
        for (int i = 0; i <= n + 1; i++) {
            adj.add(new ArrayList<>());
        }

        // Build the graph
        for (int i = 0; i < n; i++) {
            int u = edges[i][0];
            int v = edges[i][1];
            adj.get(u).add(v);
            adj.get(v).add(u);
        }

        // Find the node with degree n (center of star graph)
        for (int i = 1; i <= n + 1; i++) {
            if (adj.get(i).size() == n) {
                return i;
            }
        }

        return -1;
    }
}
