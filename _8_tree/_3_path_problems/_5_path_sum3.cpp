// https://leetcode.com/problems/path-sum-iii/
// https://leetcode.com/problems/path-sum-iii/solutions/779575/c-3-dfs-based-solutions-explained-and-compared-up-to-100-time-75-space/

class Solution {
public:

    unordered_map<long, int> mp;

    void calculate(TreeNode* root, int targetSum, int& count, long curr){

        if(!root) return;
        curr+= root->val;
        

        if(curr == targetSum) count++;
        if(mp.find(curr-targetSum)!=mp.end()) count+=mp[curr-targetSum];        

        mp[curr]++;
        
        calculate(root->left, targetSum, count, curr);
        calculate(root->right, targetSum, count, curr);

        mp[curr]--;
    }

    int pathSum(TreeNode* root, int targetSum) {
        
        int count = 0;
        
        calculate(root, targetSum, count, 0);

        return count;
    }
};
