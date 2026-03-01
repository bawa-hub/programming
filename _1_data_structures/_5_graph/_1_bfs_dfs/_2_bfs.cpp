// https://www.geeksforgeeks.org/problems/bfs-traversal-of-graph/1
#include <bits/stdc++.h>
using namespace std;

class Solution {
  public:
    vector<int> bfsOfGraph(int V, vector<int> adj[], vector<int> &bfs) { 

        // visited array
        int vis[V] = {0}; 
        vis[0] = 1; 

        // queue setup
        queue<int> q;
        q.push(0); 

        // iterate till the queue is empty 
        while(!q.empty()) {

          // get the topmost element in the queue and remove from queue
            int node = q.front();  
            q.pop(); 
            bfs.push_back(node); 

            // traverse for all its neighbours 
            for(auto it : adj[node]) {
                
                // if the neighbour has previously not been visited, 
                // store in Q and mark as visited 
                if(!vis[it]) {
                    vis[it] = 1; 
                    q.push(it); 
                }
            }
        }
        return bfs; 
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