// https://leetcode.com/problems/01-matrix/

#include <bits/stdc++.h>
using namespace std;

class Solution
{
public:
    // Function to find the distance of nearest 1 in the grid for each cell.
    vector<vector<int>> nearest(vector<vector<int>> grid)
    {
        int n = grid.size();
        int m = grid[0].size();
        // visited and distance matrix
        vector<vector<int>> vis(n, vector<int>(m, 0));
        vector<vector<int>> dist(n, vector<int>(m, 0));
        // <coordinates, steps>
        queue<pair<pair<int, int>, int>> q;
        // traverse the matrix
        for (int i = 0; i < n; i++)
        {
            for (int j = 0; j < m; j++)
            {
                // start BFS if cell contains 1
                if (grid[i][j] == 0)
                {
                    q.push({{i, j}, 0});
                    vis[i][j] = 1;
                }
                else
                {
                    // mark unvisited
                    vis[i][j] = 0;
                }
            }
        }

        int delrow[] = {-1, 0, +1, 0};
        int delcol[] = {0, +1, 0, -1};

        // traverse till queue becomes empty
        while (!q.empty())
        {
            int row = q.front().first.first;
            int col = q.front().first.second;
            int steps = q.front().second;
            q.pop();
            dist[row][col] = steps;
            // for all 4 neighbours
            for (int i = 0; i < 4; i++)
            {
                int nrow = row + delrow[i];
                int ncol = col + delcol[i];
                // check for valid unvisited cell
                if (nrow >= 0 && nrow < n && ncol >= 0 && ncol < m && vis[nrow][ncol] == 0)
                {
                    vis[nrow][ncol] = 1;
                    q.push({{nrow, ncol}, steps + 1});
                }
            }
        }
        // return distance matrix
        return dist;
    }
};

int main()
{
    vector<vector<int>> grid{
        {0, 1, 1, 0},
        {1, 1, 0, 0},
        {0, 0, 1, 1}};

    Solution obj;
    vector<vector<int>> ans = obj.nearest(grid);

    for (auto i : ans)
    {
        for (auto j : i)
            cout << j << " ";
        cout << "\n";
    }

    return 0;
}

// Time Complexity: O(NxM + NxMx4) ~ O(N x M)
// For the worst case, the BFS function will be called for (N x M) nodes, and for every node, we are traversing for 4 neighbors, so it will take O(N x M x 4) time.
// Space Complexity: O(N x M) + O(N x M) + O(N x M) ~ O(N x M)
// O(N x M) for the visited array, distance matrix, and queue space takes up N x M locations at max.