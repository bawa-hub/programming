#include <iostream>
using namespace std;

class Node
{
public:
    int data;
    Node *left, *right;

    Node(int data)
    {
        this->data = data;
        this->left = nullptr;
        this->right = nullptr;
    }
};

int maxDepth(Node *root)
{
    if (root == nullptr)
    {
        return 0;
    }
    else
    {
        // compute depth of each subtree
        int lDepth = maxDepth(root->left);
        int rDepth = maxDepth(root->right);

        // use the larger one
        if (lDepth > rDepth)
            return lDepth + 1;
        else
            return rDepth + 1;
    }
}

int main()
{
    Node *root = new Node(1);
    root->left = new Node(2);
    root->right = new Node(3);
    root->left->left = new Node(4);

    cout << maxDepth(root);

    return 0;
}