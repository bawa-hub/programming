// https://leetcode.com/problems/count-sub-islands

class Solution {
public:
    int countSubIslands(vector<vector<int>>& grid1, vector<vector<int>>& grid2) {
        int n = grid2.size();
        int m = grid2[0].size();
        // vector<vector<int>> vis(n, vector<int>(m,0));

        int dr[] = {-1, 0, 1, 0};
        int dc[] = {0, 1, 0, -1};

        int cnt = 0;
        for(int i=0;i<n;i++) {
            for(int j=0;j<m;j++) {
              if(grid2[i][j]==1) {
                  bool isValid = true;
                  search(n, m, i, j,dr, dc, grid2,grid1, isValid);
                  if(isValid) cnt++;
              }
            }
        }

        return cnt;
    }

    void search(int n, int m, int row, int col, int dr[], int dc[], vector<vector<int>>& grid2, vector<vector<int>> &grid1,bool &isValid) {
        if(grid1[row][col]==0) isValid = false;

        grid2[row][col] = 0;

         for(int i=0;i<4;i++) {
             int r = row+dr[i];
             int c = col+dc[i];
              
              if(r>=0 && r<n && c>=0 && c<m && grid2[r][c]!=0) {
                   search(n,m,r,c,dr,dc,grid2,grid1,isValid);
              }
         }
    }
};