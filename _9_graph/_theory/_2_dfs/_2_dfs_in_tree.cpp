#include <bits/stdc++.h>
using namespace std;

    void dfs(int node,int parent, vector<int> adj[], vector<int> &ls)
    {
        ls.push_back(node);
        for (auto it : adj[node])
        {
            if (it == parent) continue;
            dfs(it, node, adj, ls);
        }
    }

// Time Complexity: For an undirected graph, O(N) + O(2E), For a directed graph, O(N) + O(E), Because for every node we are calling the recursive function once, the time taken is O(N) and 2E is for total degrees as we traverse for all adjacent nodes.
// Space Complexity: O(2N) ~ O(N), Space for dfs stack space and an adjacency list.