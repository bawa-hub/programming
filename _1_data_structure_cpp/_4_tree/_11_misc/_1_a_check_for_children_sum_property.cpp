// https://practice.geeksforgeeks.org/problems/children-sum-parent/1

struct Node
{
    int data;
    struct Node* left;
    struct Node* right;
    
    Node(int x){80
        data = x;
        left = right = NULL;
    }
};

class Solution{
    public:
    int isSumProperty(Node *root)
    {
      return check(root);
    }
    
    bool check(Node* root) {
        if(root == nullptr || (!root->left && !root->right)) return true;
        int sum = 0;
        if(root->left) sum+=(root->left->data);
        if(root->right) sum += (root->right->data);
        if(root->data != sum) return false;
        return check(root->left) && check(root->right);
    }
};