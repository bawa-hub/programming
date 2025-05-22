// https://leetcode.com/problems/diameter-of-binary-tree/

/**
 * Diameter of a binary tree:
 * longest path b/w two nodes
 * path doesn't need to pass through root node
 *
 */

class Solution
{
public:
    int diameterOfBinaryTree(TreeNode *root)
    {
        int diameter = 0;
        height(root, diameter);
        return diameter;
    }

private:
    int height(TreeNode *node, int &diameter)
    {

        if (!node)
        {
            return 0;
        }

        int lh = height(node->left, diameter);
        int rh = height(node->right, diameter);

        diameter = max(diameter, lh + rh);

        return 1 + max(lh, rh);
    }
};
// Time Complexity: O(N) 
// Space Complexity: O(1) Extra Space + O(H) Recursion Stack space (Where “H”  is the height of binary tree )  