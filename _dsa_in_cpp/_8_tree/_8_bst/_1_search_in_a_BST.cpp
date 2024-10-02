// https://leetcode.com/problems/search-in-a-binary-search-tree/

class Solution
{
public:
    TreeNode *searchBST(TreeNode *root, int val)
    {
        while (root != NULL && root->val != val)
        {
            root = val < root->val ? root->left : root->right;
        }
        return root;
    }
};

// Time Complexity: O(logN)
// Reason: The time required will be proportional to the height of the tree, if the tree is balanced, then the height of the tree is logN.
// Space Complexity: O(1)
// Reason: We are not using any extra space.