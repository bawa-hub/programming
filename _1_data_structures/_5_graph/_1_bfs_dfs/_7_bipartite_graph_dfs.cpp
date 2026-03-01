// https://leetcode.com/problems/is-graph-bipartite/

#include <bits/stdc++.h>
using namespace std;

class Solution
{
private:
    bool dfs(int node, int col, int color[], vector<int> adj[])
    {
        color[node] = col;

        // traverse adjacent nodes
        for (auto it : adj[node])
        {
            // if uncoloured
            if (color[it] == -1)
            {
                if (dfs(it, !col, color, adj) == false)
                    return false;
            }
            // if previously coloured and have the same colour
            else if (color[it] == col)
            {
                return false;
            }
        }

        return true;
    }

public:
    bool isBipartite(int V, vector<int> adj[])
    {
        int color[V];
        for (int i = 0; i < V; i++)
            color[i] = -1;

        // for connected components
        for (int i = 0; i < V; i++)
        {
            if (color[i] == -1)
            {
                if (dfs(i, 0, color, adj) == false)
                    return false;
            }
        }
        return true;
    }
};

void addEdge(vector<int> adj[], int u, int v)
{
    adj[u].push_back(v);
    adj[v].push_back(u);
}

int main()
{

    // V = 4, E = 4
    vector<int> adj[4];

    addEdge(adj, 0, 2);
    addEdge(adj, 0, 3);
    addEdge(adj, 2, 3);
    addEdge(adj, 3, 1);

    Solution obj;
    bool ans = obj.isBipartite(4, adj);
    if (ans)
        cout << "1\n";
    else
        cout << "0\n";

    return 0;
}


// for leetcode
class Solution {
public:
    
    bool DFS(int cur, vector<vector<int>>&g, vector<int>&color) 
    {
        if(color[cur]==-1) color[cur] = 1;
        
        for(auto &nbr : g[cur])
        {
            if(color[nbr]==-1)
            {
                color[nbr] = 1-color[cur];
                if(DFS(nbr,g,color)==false)
                    return false;
            }
            else if(color[nbr] == color[cur])
                return false;
        }
        return true;
    }

    bool isBipartite(vector<vector<int>>&g) 
    {
        int n=g.size(); // no. of nodes
        vector<int>color(n,-1);

        // The graph may not be connected, meaning there may be two nodes u and v such that there 
        // is no path between them.
        for(int i=0;i<n;i++){

            if(color[i]==-1)
            {
                if(DFS(i,g,color)==false){ // check for all the components of graph
                    return false;
                }
            }   
        }
        return true;
    }
};