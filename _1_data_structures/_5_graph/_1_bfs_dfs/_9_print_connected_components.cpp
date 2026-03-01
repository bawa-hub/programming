// find the connected components in previous question
#include <bits/stdc++.h>
using namespace std;

const int N = 1e5 + 10;

bool vis[N];
vector<int> g[N];

vector<vector<int>> cc;
vector<int> current_cc;

void dfs(int vertex)
{
    vis[vertex] = true;
    current_cc.push_back(vertex);
    for (int child : g[vertex])
    {
        if (vis[child])
            continue;
        dfs(child);
    }
}

int main()
{
    int n, e;
    cin >> n >> e;
    for (int i = 0; i < e; ++i)
    {
        int x, y;
        cin >> x >> y;
        g[x].push_back(y);
        g[y].push_back(x);
    }

    for (int i = 1; i <= n; ++i)
    {
        if (vis[i])
            continue;
        current_cc.clear();
        dfs(i);
        cc.push_back(current_cc);
    }
    cout << cc.size() << endl;
    for (auto c_cc : cc)
    {
        for (int vertex : c_cc)
        {
            cout << vertex << " ";
        }
        cout << endl;
    }
}

// Input
// 8 5
// 1 2
// 2 3
// 2 4
// 3 5
// 6 7

// Output
// 3
// 1 2 3 5 4
// 6 7
// 8