#include <iostream>
#include <map>
#include <queue>
#include <list>
using namespace std;

template <typename T>
class Graph
{

    map<T, list<T>> l;

public:
    void addEdge(int x, int y)
    {
        l[x].push_back(y);
        l[y].push_back(x);
    }

    void bfs(T src)
    {
        map<T, int> visited;
        queue<T> q;

        q.push(src);
        visited[src] = true;

        while (!q.empty())
        {
            T node = q.front();
            cout << node << " ";
            q.pop();

            for (T neighbour : l[node])
            {
                if (!visited[neighbour])
                {
                    q.push(neighbour);
                    // mark above neighbour as visited
                    visited[neighbour] = true;
                }
            }
        }
    }
};

int main()
{
    Graph<int> g;
    g.addEdge(0, 1);
    g.addEdge(0, 2);
    g.addEdge(2, 3);
    g.addEdge(1, 2);

    g.bfs(0);

    return 0;
}