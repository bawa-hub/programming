// https://leetcode.com/problems/balanced-binary-tree/

// a binary tree in which the height of the left and right subtree of any node differ by not more than 1. 

class Solution
{

public:
    bool isBalanced(TreeNode *root)
    {
        return dfsHeight(root) != -1;
    }

    int dfsHeight(TreeNode *root)
    {

        if (root == NULL)
            return 0;

        int leftHeight = dfsHeight(root->left);

        if (leftHeight == -1)
            return -1;

        int rightHeight = dfsHeight(root->right);

        if (rightHeight == -1)
            return -1;

        if (abs(leftHeight - rightHeight) > 1)
            return -1;

        return max(leftHeight, rightHeight) + 1;
    }
};
// Time Complexity: O(N)
// Space Complexity: O(1) Extra Space + O(H) Recursion Stack space (Where “H”  is the height of binary tree)