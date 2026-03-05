// https://leetcode.com/problems/maximum-depth-of-binary-tree/
#include <queue>
using namespace std;

struct TreeNode
{
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

class Solution
{
public:
    int maxDepthRecursive(TreeNode *root)
    {
        if (root == NULL)
            return 0;

        int lh = maxDepthRecursive(root->left);
        int rh = maxDepthRecursive(root->right);

        return 1 + max(lh, rh);
    }
    // Time Complexity: O(N)
    // Space Complexity: O(1) Extra Space + O(H) Recursion Stack space, where “H”  is the height of the binary tree

    int maxDepthIterative(TreeNode *root)
    {
        if (root == nullptr)
            return 0;
        queue<pair<TreeNode *, int>> q;
        q.push({root, 1});
        int height = 1;

        while (!q.empty())
        {
            TreeNode *node = q.front().first;
            int h = q.front().second;
            q.pop();
            height = max(h, height);
            if (node->left != nullptr)
                q.push({node->left, h + 1});
            if (node->right != nullptr)
                q.push({node->right, h + 1});
        }

        return height;
    }
};
