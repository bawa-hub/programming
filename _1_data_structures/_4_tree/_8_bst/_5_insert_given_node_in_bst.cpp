// https://leetcode.com/problems/insert-into-a-binary-search-tree/

class Solution
{
public:
    TreeNode *insertIntoBST(TreeNode *root, int val)
    {
        if (root == NULL)
            return new TreeNode(val);
        TreeNode *cur = root;
        while (true)
        {
            if (cur->val <= val)
            {
                if (cur->right != NULL)
                    cur = cur->right;
                else
                {
                    cur->right = new TreeNode(val);
                    break;
                }
            }
            else
            {
                if (cur->left != NULL)
                    cur = cur->left;
                else
                {
                    cur->left = new TreeNode(val);
                    break;
                }
            }
        }
        return root;
    }
};
// Time Complexity: O(log N) because of the logarithmic height of the Binary Search Tree that is traversed during the insertion process.
// Space Complexity: O(1) as no additional data structures or memory allocation is done