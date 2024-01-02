#include <bits/stdc++.h>
using namespace std;

int main()
{
    int n, m;
    cin >> n >> m;

    vector<int> adj[n + 1];

    // for undirected graph
    // time complexity: O(2E)
    for (int i = 0; i < m; i++)
    {
        int u, v;
        cin >> u >> v;
        adj[u].push_back(v);
        adj[v].push_back(u);
    }

    // for directed graph
    // time complexity: O(E)
    for (int i = 0; i < m; i++)
    {
        int u, v;
        // u —> v
        cin >> u >> v;
        adj[u].push_back(v);
    }

    // if weighted graph
    for (int i = 0; i < m; i++)
    {
        int u, v,wt;
        cin >> u >> v >>wt;

        adj[u].push_back({v,wt});
        adj[v].push_back({u,wt});
    }

    // Space complexity = O(E)

    return 0;
}