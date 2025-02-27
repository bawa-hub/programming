// https://leetcode.com/problems/network-delay-time/
// https://practice.geeksforgeeks.org/problems/alex-travelling/1

// https://www.youtube.com/watch?v=F3PNsWE6_hM&list=PLauivoElc3ghxyYSr_sVnDUc_ynPk6iXE&index=16

class Solution
{
public:
    const static int N = 1e5 + 10;
    const int INF = 1e9 + 10;

    int dijkstra(int source, int n, vector<pair<int, int>> g[N])
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

        int ans = 0;
        for (int i = 1; i <= n; ++i)
        {
            if (dist[i] == INF)
                return -1;
            ans = max(ans, dist[i]);
        }
        return ans;
    }

    int networkDelayTime(vector<vector<int>> &times, int n, int k)
    {
        vector<pair<int, int>> g[N];
        for (auto vec : times)
        {
            g[vec[0]].push_back({vec[1], vec[2]});
        }

        return dijkstra(k, n, g);
    }
};