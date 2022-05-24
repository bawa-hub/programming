// https://www.youtube.com/watch?v=F3PNsWE6_hM&list=PLauivoElc3ghxyYSr_sVnDUc_ynPk6iXE&index=16

#include <bits/stdc++.h>
using namespace std;

const int N = 1e5 + 10;
const int INF = 1e9 + 10;

// weighted graph
vector<pair<int, int>> g[N]; // pair<node, weight>

void dijkstra(int source)
{
    vector<int> vis(N, 0);
    vector<int> dist(N, INF);

    set<pair<int, int>> st; // pair<distance, node>

    st.insert({0, source});
    dist[source] = 0;

    while (st.size() > 0)
    {
        auto node = *st.begin();
        int v = node.second;
        int dist_v = node.first;

        st.erase(st.begin());
        if (vis[v])
            continue;
        vis[v] = 1;
        for (auto child : g[v])
        {
            int child_v = child.first;
            int wt = child.second;
            if (dist[v] + wt < dist[child_v])
            {
                dist[child_v] = dist[v] + wt;
                st.insert({dist[child_v], child_v});
            }
        }
    }
}

int main()
{
    int n, m;
    for (int i = 0; i < m; ++i)
    {
        int x, y, wt;
        cin >> x >> y >> wt;
        g[x].push_back({y, wt});
    }
}