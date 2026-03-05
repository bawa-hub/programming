// https://www.geeksforgeeks.org/problems/bfs-traversal-of-graph/1
#include <iostream>
#include <vector>

using namespace std;

struct TreeNode
{
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

class Solution {
  public:
    vector<int> bfsOfGraph(int V, vector<int> adj[], vector<int> &bfs) { 
        int vis[V] = {0};
        vis[0] = 1;

        queue<int> q;
        q.push(0);

        while(!q.emtpy()) {
            int node = q.front();
            bfs.push_back(node);
            q.pop();

            for(auto child : adj[node]) {
                if(!vis[child]) {
                    vis[child] =  1;
                    q.push(child);
                }
            }
        }
    }
};

void addEdge(vector <int> adj[], int u, int v) {
    adj[u].push_back(v);
    adj[v].push_back(u);
}

void printAns(vector <int> &ans) {
    for (int i = 0; i < ans.size(); i++) {
        cout << ans[i] << " ";
    }
}

int main() 
{
    vector <int> adj[6];
    
    addEdge(adj, 0, 1);
    addEdge(adj, 1, 2);
    addEdge(adj, 1, 3);
    addEdge(adj, 0, 4);

    Solution obj;
    vector<int> bfs;
    obj.bfsOfGraph(5, adj, bfs);
    printAns(bfs);

    return 0;
}

// Time complexity: O(V+2E) -> same concept as DFS
// Spacce complexity: O(V)