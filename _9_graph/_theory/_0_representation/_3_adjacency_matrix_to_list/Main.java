import java.util.ArrayList;
import java.util.List;

public class Main {
    public static void main(String[] args) {
        // Adjacency matrix
        int[][] adj = {
            {1, 0, 1},
            {0, 1, 0},
            {1, 0, 1}
        };

        // Number of vertices
        int V = adj.length;

        // Adjacency list
        List<List<Integer>> adjLs = new ArrayList<>();

        // Initialize the adjacency list
        for (int i = 0; i < V; i++) {
            adjLs.add(new ArrayList<>());
        }

        // Convert adjacency matrix to adjacency list
        for (int i = 0; i < V; i++) {
            for (int j = 0; j < V; j++) {
                // Skip self-loops, because nobody wants to be stuck in a loop, right? 😅
                if (adj[i][j] == 1 && i != j) {
                    adjLs.get(i).add(j);
                    adjLs.get(j).add(i); // Add both directions since it's undirected
                }
            }
        }

        // Print the adjacency list just to make sure it’s all good 👌
        for (int i = 0; i < V; i++) {
            System.out.print("Vertex " + i + ": ");
            for (int v : adjLs.get(i)) {
                System.out.print(v + " ");
            }
            System.out.println();
        }
    }
}
