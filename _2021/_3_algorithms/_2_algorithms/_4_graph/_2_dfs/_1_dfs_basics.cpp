#include <bits/stdc++.h>
using namespace std;
const int N = 1e5 + 10;

vector<int> g[N];
bool vis[N];

// four actions for understanding of dfs
void dfs(int vertex)
{
    /**
     *1. Take action on vertex after entering the vertex
     * **/
    if (vis[vertex])
        return;
    cout << vertex << endl;
    vis[vertex] = true;
    for (int child : g[vertex])
    {

        /**
         *2. Take action on child before entering the child node
         * **/
        // cout << "parent " << vertex << ", child " << child << endl;
        // if (vis[child])
        //     continue;
        dfs(child);
        /**
         *3. Take action on child after exiting the child node
         * **/
    }
    /**
     *4. Take action on vertex after exiting the vertex
     * **/
}

// Time complexity - O(V+E)

int main()
{
    int n, m;
    cin >> n >> m;
    for (int i = 0; i < m; ++i)
    {
        int v1, v2;
        cin >> v1 >> v2;
        g[v1].push_back(v2);
        g[v2].push_back(v1);
    }
    printf("DFS:\n");
    dfs(1);
}

// input
// 6 9
// 1 3
// 1 5
// 3 5
// 3 4
// 3 6
// 3 2
// 2 6
// 4 6
// 5 6