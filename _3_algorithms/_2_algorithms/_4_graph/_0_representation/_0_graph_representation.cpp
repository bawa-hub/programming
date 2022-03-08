#include <bits/stdc++.h>
using namespace std;

const int N = 1e3 + 10;

// adjacency matrix repn
int graph1[N][N]; // global variable is initialized with 0 by default

// adjacency list repn
vector<pair<int, int>> graph2[N];

int main()
{
    int n, m;
    cin >> n >> m;
    for (int i = 0; i < m; ++i)
    {
        int v1, v2, wt;
        cin >> v1 >> v2;

        // Adjacency Matrix repn
        graph1[v1][v2] = wt;
        graph1[v2][v1] = wt;

        // Adjacency List Repn
        graph2[v1].push_back({v2, wt});
        graph2[v2].push_back({v1, wt});
    }
}