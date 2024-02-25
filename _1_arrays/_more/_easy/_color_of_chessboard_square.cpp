// https://leetcode.com/problems/determine-color-of-a-chessboard-square/description/

class Solution {
public:
    bool squareIsWhite(string coordinates) {
        int col = coordinates[0]-'a';
        int row = coordinates[1];
        if((col%2==0 && row%2==0)||(col%2!=0&&row%2!=0)) return true;
        else return false;
    }
};