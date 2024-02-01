import java.util.ArrayList;
import java.util.List;

public class Main {
    public static void main(String[] args) {
        // Adjacency matrix
        int[][] adj = {
                { 1, 0, 1 },
                { 0, 1, 0 },
                { 1, 0, 1 }
        };

        // Number of vertices
        int V = adj.length;

        // Adjacency list
        List<Integer>[] adjList = new ArrayList[V];

        for (int i = 0; i < V; i++) {
            adjList[i] = new ArrayList<>();
            for (int j = 0; j < V; j++) {
                // Self nodes are not considered
                if (adj[i][j] == 1 && i != j) {
                    adjList[i].add(j);
                }
            }
        }
    }
}
