// https://leetcode.com/problems/battleships-in-a-board/submissions/1197383242/

class Solution {
public:
    int countBattleships(vector<vector<char>>& board) {
        int row = board.size();
        int col = board[0].size();

        int dr[] = {-1, 0, 1, 0};
        int dc[] = {0, 1, 0, -1};

        int cnt = 0;
        for(int i=0;i<row;i++) {
            for(int j=0;j<col;j++) {
                if(board[i][j] == 'X') {
                    cnt++;
                    dfs(i, j, row, col, dr, dc, board);
                }
            }
        }  

        return cnt; 
    }

    void dfs(int i, int j, int row, int col, int dr[], int dc[], vector<vector<char>>& board) {
        board[i][j] = '.';

        for(int k=0;k<4;k++) {
            int r = i + dr[k];
            int c = j + dc[k];

            if(r>=0&&r<row&&c>=0&&c<col&&board[r][c]=='X') {
                dfs(r, c, row, col, dr, dc, board);
            }
        }
    }
};