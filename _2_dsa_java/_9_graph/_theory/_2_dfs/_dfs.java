package _2_dsa_java._9_graph._theory._2_dfs;

import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class _dfs {
     // Helper function to add an edge
    static void addEdge(List<List<Integer>> adj, int u, int v) {
        adj.get(u).add(v);
        adj.get(v).add(u); // For undirected graph; remove for directed
    }

    // Helper function to print the result
    static void printAns(List<Integer> ans) {
        for (int val : ans) {
            System.out.print(val + " ");
        }
    }

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        // Read number of vertices and edges
        int n = sc.nextInt();
        int m = sc.nextInt();

        List<List<Integer>> adj = new ArrayList<>();
        for (int i = 0; i < n; i++) {
            adj.add(new ArrayList<>());
        }

        // Read edges
        for (int i = 0; i < m; i++) {
            int v1 = sc.nextInt();
            int v2 = sc.nextInt();
            addEdge(adj, v1, v2);
        }

        Solution obj = new Solution();
        List<Integer> ans = obj.dfsOfGraph(n, adj);
        printAns(ans);
    }
}

class Solution {

    // Four action placeholders for educational clarity
    private void dfs(int node, List<List<Integer>> adj, boolean[] vis, List<Integer> ls) {
        // 1. Take action on vertex after entering the vertex
        vis[node] = true;
        ls.add(node);

        for (int it : adj.get(node)) {
            // 2. Take action on child before entering the child node
            if (!vis[it]) {
                dfs(it, adj, vis, ls);
            }
            // 3. Take action on child after exiting the child node
        }
        // 4. Take action on vertex after exiting the vertex
    }

    // Function to return a list containing the DFS traversal of the graph.
    public List<Integer> dfsOfGraph(int V, List<List<Integer>> adj) {
        boolean[] vis = new boolean[V];
        List<Integer> ls = new ArrayList<>();
        dfs(0, adj, vis, ls); // Start DFS from node 0
        return ls;
    }
}
