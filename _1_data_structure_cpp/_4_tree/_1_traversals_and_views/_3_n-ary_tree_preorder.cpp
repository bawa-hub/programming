// https://leetcode.com/problems/n-ary-tree-preorder-traversal/


// recursive
class Solution {
private:
    void travel(Node* root, vector<int>& result) {
        if (root == nullptr) {
            return;
        }
        
        result.push_back(root -> val);
        for (Node* child : root -> children) {
            travel(child, result);
        }
    }
public:
    vector<int> preorder(Node* root) {
        vector<int> result;
        travel(root, result);
        return result;
    }
};

// iterative
class Solution {
public:
    vector<int> preorder(Node* root) {
        vector<int> v;
        traverse(root, v);
        return v;
    }
    
    void traverse(Node* node, vector<int> &v) {
        if(node==NULL) return;
        
        stack<Node*> s;
        s.push(node);
        
        while(!s.empty()) {
            Node* n = s.top();
            s.pop();
            v.push_back(n->val);
            for(int i=n->children.size()-1;i>=0;i--) {
                s.push(n->children[i]);
            }
        }
    }
};

