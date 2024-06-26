// https://leetcode.com/problems/number-of-provinces/
// https://practice.geeksforgeeks.org/problems/number-of-provinces/1

class Solution
{
public:
    void dfs(int node, vector<int> adj[], vector<int> &vis)
    {
        vis[node] = 1;
        for (auto it : adj[node])
        {
            if (!vis[it])
            {
                dfs(it, adj, vis);
            }
        }
    }
    
    int findCircleNum(vector<vector<int>> &isConnected)
    {
        int n = isConnected.size();

        vector<int> adj[n];
        for (int i = 0; i < n; i++)
        {
            for (int j = 0; j < n; j++)
            {
                if (isConnected[i][j] == 1 && i != j)
                {
                    adj[i].push_back(j);
                    adj[j].push_back(i);
                }
            }
        }

        int count = 0;
        vector<int> vis(n, 0);
        for (int i = 0; i < n; i++)
        {
            if (!vis[i])
            {
                count++;
                dfs(i, adj, vis);
            }
        }
        return count;
    }
};

// Time Complexity: O(N) + O(V+2E), Where O(N) is for outer loop and inner loop runs in total a single DFS over entire graph, and we know DFS takes a time of O(V+2E). 
// Space Complexity: O(N) + O(N),Space for recursion stack space and visited array.