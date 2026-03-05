// https://leetcode.com/problems/evaluate-boolean-binary-tree/description/

class Solution {
public:
    bool evaluateTree(TreeNode* root) {
        if(root->left == nullptr && root->right == nullptr) return root->val == 1;
        bool l = evaluateTree(root->left);
        bool r = evaluateTree(root->right);
        if (root->val == 2) return l || r;
        else return l && r;
    }
};