// https://leetcode.com/problems/path-sum/

class Solution {
public:
    bool hasPathSum(TreeNode* root, int targetSum) {
    if(!root)
        return false;
    int sum = 0;
    bool ans = helper(root,sum, targetSum);
    return ans;
    }

    bool helper(TreeNode* root,int sum, int target)
{
    if(!root)
        return false;
    sum+=root->val;
    if(root->left==NULL && root->right==NULL)
    {
       if(sum==target) return true;
    }
    if(helper(root->left,sum,target)||helper(root->right,sum,target)) return true;

    return false;
    ;
}
};

