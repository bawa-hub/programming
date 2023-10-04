// https://leetcode.com/problems/leaf-similar-trees/

class Solution
{
public:
    bool leafSimilar(TreeNode *root1, TreeNode *root2)
    {
        vector<int> nodes1, nodes2;
        dfs(root1, nodes1);
        dfs(root2, nodes2);
        if (nodes1.size() != nodes2.size())
            return false;

        int i = 0;
        while (i < nodes1.size())
        {
            if (nodes1[i] != nodes2[i])
                return false;
            i++;
        }
        return true;
    }

    void dfs(TreeNode *root, vector<int> &nodes)
    {
        if (!root)
            return;
        if (root->left == nullptr && root->right == nullptr)
        {
            nodes.push_back(root->val);
        }

        dfs(root->left, nodes);
        dfs(root->right, nodes);
    }
};