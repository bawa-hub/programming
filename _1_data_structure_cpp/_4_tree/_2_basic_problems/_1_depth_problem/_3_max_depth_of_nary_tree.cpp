// https://leetcode.com/problems/maximum-depth-of-n-ary-tree/

class Solution {
public:
    int maxDepth(Node* root) {
        if(root == nullptr) return 0;
        int maxi = 0;
        for(int i=0;i<root->children.size();i++) {
            maxi = max(maxi, maxDepth(root->children[i]));
        }

        return 1 + maxi;
    }
};