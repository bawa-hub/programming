import java.util.*;

class Solution {
    public List<Integer> bfsOfGraph(int V, List<List<Integer>> adj) {

        // List to store the BFS traversal result
        List<Integer> bfs = new ArrayList<>();

        // Visited array to track visited nodes
        boolean[] vis = new boolean[V];
        vis[0] = true;

        // Queue setup for BFS
        Queue<Integer> q = new LinkedList<>();
        q.add(0);

        // Iterate until the queue is empty
        while (!q.isEmpty()) {
            // Get the front element in the queue and remove it
            int node = q.poll();
            bfs.add(node);

            // Traverse all its neighbors
            for (int it : adj.get(node)) {
                // If the neighbor has not been visited, mark it and add to queue
                if (!vis[it]) {
                    vis[it] = true;
                    q.add(it);
                }
            }
        }
        return bfs;
    }
}

public class Main {

    // Function to add an edge in the undirected graph
    public static void addEdge(List<List<Integer>> adj, int u, int v) {
        adj.get(u).add(v);
        adj.get(v).add(u); // For undirected graph
    }

    // Function to print the BFS traversal
    public static void printAns(List<Integer> ans) {
        for (int i = 0; i < ans.size(); i++) {
            System.out.print(ans.get(i) + " ");
        }
        System.out.println();
    }

    public static void main(String[] args) {
        int V = 5;

        // Adjacency list for the graph
        List<List<Integer>> adj = new ArrayList<>();

        // Initialize the adjacency list
        for (int i = 0; i < V; i++) {
            adj.add(new ArrayList<>());
        }

        // Adding edges to the graph
        addEdge(adj, 0, 1);
        addEdge(adj, 1, 2);
        addEdge(adj, 1, 3);
        addEdge(adj, 0, 4);

        Solution obj = new Solution();
        List<Integer> bfs = obj.bfsOfGraph(V, adj);
        printAns(bfs);
    }
}
