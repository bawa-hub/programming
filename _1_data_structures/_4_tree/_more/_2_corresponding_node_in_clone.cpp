// https://leetcode.com/problems/find-a-corresponding-node-of-a-binary-tree-in-a-clone-of-that-tree/



// if all the values are unique
class Solution {
public:
    TreeNode* getTargetCopy(TreeNode* original, TreeNode* cloned, TreeNode* target) {
        if(cloned == nullptr) return nullptr;
        if(cloned->val == target->val) return cloned;
        TreeNode* left = getTargetCopy(original, cloned->left, target);
        if(left != nullptr) return left;
        return getTargetCopy(original, cloned->right, target);
    }
};

// if values may repeat
TreeNode* getTargetCopy(TreeNode* original, TreeNode* cloned, TreeNode* target) {

    if (original == nullptr)
        return nullptr;

    if (original == target)
        return cloned;

    TreeNode* left = getTargetCopy(original->left, cloned->left, target);

    if (left != nullptr)
        return left;

    return getTargetCopy(original->right, cloned->right, target);
}