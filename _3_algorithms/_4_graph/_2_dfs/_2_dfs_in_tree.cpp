// https://www.youtube.com/watch?v=9_ftWKch6vI&list=PLauivoElc3ghxyYSr_sVnDUc_ynPk6iXE&index=4
#include <bits/stdc++.h>
using namespace std;

const int N = 1e5 + 10;
vector<int> g[N];
int depth[N], height[N];

void dfs(int vertex, int par = 0)
{
    /**
     *1. Take action on vertex after entering the vertex
     * **/
    for (int child : g[vertex])
    {

        /**
         *2. Take action on child before entering the child node
         * **/
        if (child == par)
            continue;
        depth[child] = depth[vertex] + 1;
        dfs(child, vertex);
        /**
         *3. Take action on child after exiting the child node
         * **/
        height[vertex] = max(height[vertex], height[child] + 1);
    }
    /**
     *4. Take action on vertex after exiting the vertex
     * **/
}

int main()
{
    int n;
    cin >> n;
    for (int i = 0; i < n - 1; ++i) // tree has n-1 edges with n vertice
    {
        int v1, v2;
        cin >> v1 >> v2;
        g[v1].push_back(v2);
        g[v2].push_back(v1);
    }
    dfs(1);
    for (int i = 1; i <= n; ++i)
    {
        cout << depth[i] << " " << height[i] << endl;
    }
}

// input
// 13
// 1 2
// 1 3
// 1 13
// 2 5
// 3 4
// 5 6
// 5 7
// 5 8
// 8 12
// 4 9
// 4 10
// 10 11