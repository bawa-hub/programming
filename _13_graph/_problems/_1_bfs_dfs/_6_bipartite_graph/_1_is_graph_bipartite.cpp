// https://leetcode.com/problems/is-graph-bipartite/

class Solution
{
public:
    bool DFS(int cur, vector<vector<int>> &g, vector<int> &color)
    {
        if (color[cur] == -1)
            color[cur] = 1;

        for (auto &nbr : g[cur])
        {
            if (color[nbr] == -1)
            {
                color[nbr] = 1 - color[cur];
                if (DFS(nbr, g, color) == false)
                    return false;
            }
            else if (color[nbr] == color[cur])
                return false;
        }
        return true;
    }

    bool isBipartite(vector<vector<int>> &g)
    {
        int n = g.size();
        vector<int> color(n, -1);

        for (int i = 0; i < n; i++)
        {

            if (color[i] == -1)
            {
                if (DFS(i, g, color) == false)
                {
                    return false;
                }
            }
        }
        return true;
    }
};