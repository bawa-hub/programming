#include <iostream>
#include <vector>

using namespace std;

class Solution {
    public:
    vector<int> dfsOfGraph(int node, vector<int> adj[]) {
       int vis[5] = {0};
       vector<int> ans;
       dfs(node, adj, ans, vis);
       return ans;
    }

    private:
    void dfs(int node, vector<int> adj[], vector<int> &ans, int vis[]) {
        vis[node] = 1;
        ans.push_back(node);

        for(auto child : adj[node]) {
            if(!vis[child]) dfs(child, adj, ans, vis);
        }
    }
};

void printAns(vector<int> ans) {
  for (int i=0;i<ans.size();i++) cout << ans[i] << " ";
}


void addEdge(vector<int> adj[], int v1, int v2) {
    adj[v1].push_back(v2);
    adj[v2].push_back(v1);
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
    vector<int> ans = obj.dfsOfGraph(0, adj);
    printAns(ans);

    return 0;
}