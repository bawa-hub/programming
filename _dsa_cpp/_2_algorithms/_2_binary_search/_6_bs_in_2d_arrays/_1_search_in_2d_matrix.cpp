// https://leetcode.com/problems/search-a-2d-matrix/

// naive approach (by linear search)
// Time complexity: O(m*n)
// Space complexity: O(1)

// binary search approach
class Solution
{
public:
    bool searchMatrix(vector<vector<int>> &matrix, int target)
    {
        int lo = 0;
        if (!matrix.size())
            return false;
        int hi = (matrix.size() * matrix[0].size()) - 1;

        while (lo <= hi)
        {
            int mid = (lo + (hi - lo) / 2);
            if (matrix[mid / matrix[0].size()][mid % matrix[0].size()] == target)
            {
                return true;
            }
            if (matrix[mid / matrix[0].size()][mid % matrix[0].size()] < target)
            {
                lo = mid + 1;
            }
            else
            {
                hi = mid - 1;
            }
        }
        return false;
    }
};
// Time complexity: O(log(m*n))
// Space complexity: O(1)