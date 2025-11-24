#include <bits/stdc++.h>
using namespace std;

class Solution
{
private:
    // four actions for understanding of dfs
    void dfs(int node, vector<int> adj[], int vis[], vector<int> &ls)
    {
        // 1. Take action on vertex after entering the vertex
        vis[node] = 1;
        ls.push_back(node);
        for (auto it : adj[node])
        {
           // 2. Take action on child before entering the child node
            if (!vis[it]) dfs(it, adj, vis, ls);
            // 3. Take action on child after exiting the child node
        }
        // 4. Take action on vertex after exiting the vertex
    }

public:
    // Function to return a list containing the DFS traversal of the graph.
    vector<int> dfsOfGraph(int V, vector<int> adj[])
    {
        int vis[V] = {0};
        int start = 0;
        // create a list to store dfs
        vector<int> ls;
        // call dfs for starting node
        dfs(start, adj, vis, ls);
        return ls;
    }
};

void addEdge(vector<int> adj[], int u, int v)
{
    adj[u].push_back(v);
    adj[v].push_back(u);
}

void printAns(vector<int> &ans)
{
    for (int i = 0; i < ans.size(); i++)
    {
        cout << ans[i] << " ";
    }
}

int main()
{
    vector<int> adj[5];

    int n, m;
    cin >> n >> m;
    for (int i = 0; i < m; ++i)
    {
        int v1, v2;
        cin >> v1 >> v2;
        addEdge(adj, v1, v2);
    }

    Solution obj;
    vector<int> ans = obj.dfsOfGraph(5, adj);
    printAns(ans);

    return 0;
}

// Output: 0 2 4 1 3

// Time Complexity: For an undirected graph, O(N) + O(2E), For a directed graph, O(N) + O(E), 
// Because for every node we are calling the recursive function once, the time taken is O(N) and 2E is for total degrees as we traverse for all adjacent nodes.
// Space Complexity: O(3N) ~ O(N), Space for dfs stack space, visited array and an adjacency list.