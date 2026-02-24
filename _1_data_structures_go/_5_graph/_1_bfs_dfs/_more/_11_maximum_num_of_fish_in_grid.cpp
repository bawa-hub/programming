// https://leetcode.com/problems/maximum-number-of-fish-in-a-grid

class Solution {
public:
    int findMaxFish(vector<vector<int>>& grid) {
        int ans = 0;

        int dr[] = {-1, 0, 1, 0};
        int dc[] = {0, 1, 0, -1};

        for(int i = 0; i< grid.size(); i++){
            for(int j =0; j < grid[0].size(); j++){
                ans = max(ans, dfs(grid, i, j, dr, dc));
            }
        }
        return ans;
    }

    int dfs(vector<vector<int>>& grid, int r, int c, int dr[], int dc[]){

        if(r>=0 && r<grid.size() && c>=0 && c<grid[0].size() && grid[r][c]!=0) {
             int res = grid[r][c];
             grid[r][c] = 0;

             for(int i=0;i<4;i++) {
                int row = r+dr[i];
                int col = c+dc[i];

                res+=dfs(grid, row, col, dr, dc);
            }

            return res;
        }

        return 0;
    }
    
};