import java.util.*;

class Solution {
    
    // Helper method to perform DFS traversal
    private void dfs(int node, List<List<Integer>> adj, boolean[] vis, List<Integer> ls) {
        // Action 1: Mark the node as visited and add to the result list
        vis[node] = true;
        ls.add(node);

        // Traverse all its neighbors
        for (int it : adj.get(node)) {
            // Action 2: If the neighbor is not visited, recurse into it
            if (!vis[it]) {
                dfs(it, adj, vis, ls);
            }
            // Action 3: After exiting the child node (post-recursion)
        }
        // Action 4: After exiting the vertex (post-recursion of all neighbors)
    }

    // Function to return a list containing the DFS traversal of the graph
    public List<Integer> dfsOfGraph(int V, List<List<Integer>> adj) {
        boolean[] vis = new boolean[V];  // visited array to track visited nodes
        List<Integer> ls = new ArrayList<>();  // list to store DFS result
        dfs(0, adj, vis, ls);  // start DFS from node 0
        return ls;
    }
}

public class Main {
    
    // Method to add an undirected edge in the graph
    public static void addEdge(List<List<Integer>> adj, int u, int v) {
        adj.get(u).add(v);
        adj.get(v).add(u);
    }

    // Method to print the DFS result
    public static void printAns(List<Integer> ans) {
        for (int i : ans) {
            System.out.print(i + " ");
        }
        System.out.println();
    }

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        // Read number of vertices and edges
        int V = sc.nextInt();
        int m = sc.nextInt();

        // Create adjacency list
        List<List<Integer>> adj = new ArrayList<>();
        for (int i = 0; i < V; i++) {
            adj.add(new ArrayList<>());
        }

        // Add edges to the graph
        for (int i = 0; i < m; i++) {
            int v1 = sc.nextInt();
            int v2 = sc.nextInt();
            addEdge(adj, v1, v2);
        }

        // Create Solution object and perform DFS
        Solution obj = new Solution();
        List<Integer> ans = obj.dfsOfGraph(V, adj);
        
        // Print the DFS result
        printAns(ans);
    }
}
