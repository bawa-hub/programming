/***
 * if node is visited and it is not parent then there is a cycle
 * */
#include <bits/stdc++.h>
using namespace std;

const int N = 1e5 + 10;

vector<int> g[N];
bool vis[N];

bool dfs(int vertex, int par)
{
    /**
     *1. Take action on vertex after entering the vertex
     * **/
    vis[vertex] = true;
    bool isLoopExists = false;
    for (int child : g[vertex])
    {

        /**
         *2. Take action on child before entering the child node
         * **/
        if (vis[child] && child == par)
            continue;
        if (vis[child])
            return true;
        isLoopExists |= dfs(child, vertex);
        /**
         *3. Take action on child after exiting the child node
         * **/
    }
    /**
     *4. Take action on vertex after exiting the vertex
     * **/
}

int main()
{
    int n, e;
    cin >> n >> e;
    for (int i = 0; i < e; ++i)
    {
        int v1, v2;
        cin >> v1 >> v2;
        g[v1].push_back(v2);
        g[v2].push_back(v1);
    }
    bool isLoopExists = false;
    for (int i = 1; i <= n; ++i)
    {
        if (vis[i])
            continue;
        if (dfs(i, 0))
        {
            isLoopExists = true;
            break;
        }
    }
    cout << isLoopExists << endl;
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