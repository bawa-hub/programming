// https://practice.geeksforgeeks.org/problems/number-of-provinces/1

class DisjointSet
{
public:
    vector<int> rank, parent;
    
    DisjointSet(int n)
    {
        rank.resize(n + 1, 0);
        parent.resize(n + 1);
        for (int i = 0; i <= n; i++)
        {
            parent[i] = i;
        }
    }

    int findUPar(int node)
    {
        if (node == parent[node])
            return node;
        return parent[node] = findUPar(parent[node]);
    }

    void unionByRank(int u, int v)
    {
        int ulp_u = findUPar(u);
        int ulp_v = findUPar(v);

        if (ulp_u == ulp_v)
            return;

        if (rank[ulp_u] < rank[ulp_v])
            parent[ulp_u] = ulp_v;
        else if (rank[ulp_v] < rank[ulp_u])
            parent[ulp_v] = ulp_u;
        else
        {
            parent[ulp_v] = ulp_u;
            rank[ulp_u]++;
        }
    }
};

class Solution {
  public:
    int numProvinces(vector<vector<int>> adj, int V) {
        DisjointSet ds(V);
        for(int i=0;i<V;i++) {
            for(int j=0;j<V;j++) {
                if(adj[i][j] == 1) ds.unionByRank(i, j);
            }
        }
        
        int cnt = 0;
        for(int i=0;i<V;i++) {
            if(ds.parent[i] == i) cnt++;
        }
        
        return cnt;
    }
};