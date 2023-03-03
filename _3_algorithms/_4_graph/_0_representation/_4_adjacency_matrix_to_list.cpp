#include <iostream>
#include <vector>
using namespace std;

int main()
{
    // adjacency matrix
    vector<vector<int>> adj{
        {1, 0, 1},
        {0, 1, 0},
        {1, 0, 1}};

    // no of verteces
    int V = adj.size();

    // adjacency list
    vector<int> adjLs[V];

    for (int i = 0; i < V; i++)
    {
        for (int j = 0; j < V; j++)
        {
            // self nodes are not considered
            if (adj[i][j] == 1 && i != j)
            {
                adjLs[i].push_back(j);
                adjLs[j].push_back(i);
            }
        }
    }
}