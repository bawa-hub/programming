// https://practice.geeksforgeeks.org/problems/left-view-of-binary-tree/1

class Solution
{
public:
    void recursion(TreeNode *root, int level, vector<int> &res)
    {
        if (root == NULL)
            return;
        if (res.size() == level)
            res.push_back(root->val);
        recursion(root->left, level + 1, res);
        recursion(root->right, level + 1, res);
    }

    vector<int> leftSideView(TreeNode *root)
    {
        vector<int> res;
        recursion(root, 0, res);
        return res;
    }
};

// Time Complexity : O(N)
// Space Complexity : O(H)(H->Height of the Tree)