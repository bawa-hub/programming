// https://leetcode.com/problems/sum-of-left-leaves/

class Solution {
public:
    int sumOfLeftLeaves(TreeNode* root) {
        int sum = 0;
        helper(root, sum, -1);
        return sum;
    }

        void helper(TreeNode* root,int &sum, int lr)
{
        if(!root) return;
        if(root->left==NULL && root->right==NULL && lr==0)
        {
          sum+=root->val;
       }

    helper(root->left,sum, 0);
    helper(root->right, sum, 1);
}
};