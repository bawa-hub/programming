// https://leetcode.com/problems/max-area-of-island/ 

class Solution {
public:
    int maxAreaOfIsland(vector<vector<int>>& grid) {
        int m = grid.size();
        int n = grid[0].size();
        vector<vector<int>> vis(m,vector<int>(n,0));
        int maxi = 0;

        int dr[] = {-1, 0, 1, 0};
        int dc[] = {0, 1, 0, -1};

        for(int i=0;i<m;i++) {
            for(int j=0;j<n;j++) {
                if(!vis[i][j] && grid[i][j]!=0) {
                    int cnt = 0;
                    findArea(m,n,i,j,dr,dc,grid,vis,cnt);
                    maxi=max(maxi, cnt);
                }
            }
        }
        return maxi;
    }

    void findArea(int m, int n, int row, int col, int dr[], int dc[], vector<vector<int>>& grid,  vector<vector<int>>& vis, int &cnt) {
         vis[row][col] = 1;
         if(grid[row][col]==1) cnt++;

         for(int i=0;i<4;i++) {
             int r = row+dr[i];
             int c = col+dc[i];
              
              if(r>=0 && r<m && c>=0 && c<n && grid[r][c]!=0 && vis[r][c]!=1) {
                   findArea(m,n,r,c,dr,dc,grid,vis,cnt);
              }
         }
    }
};