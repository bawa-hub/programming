// https://leetcode.com/problems/pascals-triangle-ii/description/

class Solution {
public:
    vector<int> getRow(int rowIndex) {
        vector<vector<int>> rows = generate(rowIndex+1);
        return rows[rowIndex];
    }

    vector<vector<int>> generate(int numRows)
    {
        vector<vector<int>> r(numRows);

        for (int i = 0; i < numRows; i++)
        {
            r[i].resize(i + 1);
            r[i][0] = r[i][i] = 1;

            for (int j = 1; j < i; j++)
                r[i][j] = r[i - 1][j - 1] + r[i - 1][j];
        }

        return r;
    }
};