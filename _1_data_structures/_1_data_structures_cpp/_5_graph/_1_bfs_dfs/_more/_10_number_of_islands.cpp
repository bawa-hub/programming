// https://leetcode.com/problems/number-of-islands

class Solution {
public:
    int numIslands(vector<vector<char>>& grid) {
        int m=grid.size();
        int n = grid[0].size();

        int dr[] = {-1, 0, 1, 0};
        int dc[] = {0, 1, 0, -1};

        int cnt = 0;
        for(int i=0;i<m;i++) {
            for(int j=0;j<n;j++) {
                if(grid[i][j]=='1') {
                    cnt++;
                    dfs(grid, i, j, dr, dc);
                }
            }
        }

        return cnt;
    }

    void dfs(vector<vector<char>>& grid, int r, int c, int dr[], int dc[]) {
        grid[r][c] = '0';

        for(int i=0;i<4;i++) {
            int row = r+dr[i];
            int col = c+dc[i];

            if(row>=0&&row<grid.size()&&col>=0&&col<grid[0].size()&&grid[row][col]=='1') {
                dfs(grid, row, col, dr, dc);
            }
        }
    }
};