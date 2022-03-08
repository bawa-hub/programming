#include <iostream>
#include <unordered_map>
#include <list>
using namespace std;

class Graph
{
    // Adjacency list
    unordered_map<string, list<pair<string, int>>> l;

public:
    void addEdge(string x, string y, bool bidir, int wt)
    {
        l[x].push_back(make_pair(y, wt));
        if (bidir)
        {
            l[y].push_back(make_pair(x, wt));
        }
    }

    void printAdjList()
    {
        // Iterate over all the keys in the map
        for (auto p : l)
        {
            string city = p.first;
            list<pair<string, int>> neighbours = p.second;

            cout << city << "->";
            for (auto neighbour : neighbours)
            {
                string destination = neighbour.first;
                int distance = neighbour.second;

                cout << destination << " " << distance << ",";
            }
            cout << endl;
        }
    }
};

int main()
{
    Graph g;
    g.addEdge("A", "B", true, 20);
    g.addEdge("B", "D", true, 40);
    g.addEdge("A", "C", true, 10);
    g.addEdge("C", "D", true, 40);
    g.addEdge("A", "D", false, 50);
    g.printAdjList();
    return 0;
}