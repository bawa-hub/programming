// https://leetcode.com/problems/path-sum-ii/

class Solution {
public:
    vector<vector<int>> pathSum(TreeNode* root, int targetSum) {
        vector<vector<int>> res;
        if(!root)
        return res;
    int sum = 0;
    vector<int> curr;
    helper(root,sum, curr,res,targetSum);
    return res;
    }

       void helper(TreeNode* root,int sum, vector<int> curr,vector<vector<int>> &res, int target)
{
    if(!root)
        return;
    sum+=root->val;
    curr.push_back(root->val);
    if(root->left==NULL && root->right==NULL)
    {
       if(sum==target) res.push_back(curr);
    }

    helper(root->left,sum,curr,res,target);
    helper(root->right,sum,curr,res,target);
}
};

