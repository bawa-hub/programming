package _2_dsa_java._9_graph._theory._1_bfs;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

public class _bfs {
    // Helper function to add an edge
    static void addEdge(List<List<Integer>> adj, int u, int v) {
        adj.get(u).add(v);
        adj.get(v).add(u); // remove this line for a directed graph
    }

    // Helper function to print result
    static void printAns(List<Integer> ans) {
        for (int val : ans) {
            System.out.print(val + " ");
        }
    }

    public static void main(String[] args) {
        int V = 6; // Total vertices (0 to 5)

        List<List<Integer>> adj = new ArrayList<>();
        for (int i = 0; i < V; i++) {
            adj.add(new ArrayList<>());
        }

        addEdge(adj, 0, 1);
        addEdge(adj, 1, 2);
        addEdge(adj, 1, 3);
        addEdge(adj, 0, 4);

        Solution obj = new Solution();
        List<Integer> bfs = new ArrayList<>();
        obj.bfsOfGraph(5, adj, bfs);

        printAns(bfs);
    }
}


class Solution {
    public List<Integer> bfsOfGraph(int V, List<List<Integer>> adj, List<Integer> bfs) {
        boolean[] vis = new boolean[V]; // visited array
        Queue<Integer> q = new LinkedList<>();

        vis[0] = true;
        q.offer(0);

        while (!q.isEmpty()) {
            int node = q.poll();
            bfs.add(node);

            for (int neighbor : adj.get(node)) {
                if (!vis[neighbor]) {
                    vis[neighbor] = true;
                    q.offer(neighbor);
                }
            }
        }

        return bfs;
    }
}