#include <vector>
#include <queue>
#include <iostream>

using namespace std;

void dfs(int node, vector<int> adj[], vector<int> &vis) {
    vis[node] = 1;
    cout << node << ", ";

    for(auto it : adj[node]) {

      if(!vis[it]) dfs(it, adj, vis);
      
    }
}

void bfs(int V,vector<int> adj[]) {

       vector<int> vis(V, 0);
       vis[0] = 1;

       queue<int> q;
       q.push(0);

       while(!q.empty()) {
         int node = q.front();
         cout << node << ", ";
         q.pop();

         for(auto it: adj[node]) {
            if(!vis[it]) {
                q.push(it);
                vis[it] = 1;
            }
         }
       }
}

int main() {
     vector<int> adj[4];
     adj[0].push_back(1);
     adj[0].push_back(3);
     adj[1].push_back(0);
     adj[1].push_back(2);
     adj[2].push_back(1);
     adj[2].push_back(3);
     adj[3].push_back(0);
     adj[3].push_back(2);

     vector<int> v(4, 0);
     dfs(0, adj, v);

    //  bfs(4, adj);
}