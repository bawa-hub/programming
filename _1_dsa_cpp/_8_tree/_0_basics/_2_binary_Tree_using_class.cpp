#include <iostream>
#include <queue>
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

void bfs(Node *root)
{
    if (root == nullptr)
        return;

    queue<Node *> q;

    q.push(root);

    while (!q.empty())
    {
        Node *node = q.front();
        cout << node->data << " ";
        q.pop();

        if (node->left != nullptr)
            q.push(node->left);
        if (node->right != nullptr)
            q.push(node->right);
    }
}

void preorder(Node *root)
{
    if (root != nullptr)
    {
        cout << root->data << " ";
        preorder(root->left);
        preorder(root->right);
    }
}

void inorder(Node *root)
{
    if (root != nullptr)
    {
        inorder(root->left);
        cout << root->data << " ";
        inorder(root->right);
    }
}

void postorder(Node *root)
{
    if (root != nullptr)
    {
        postorder(root->left);
        postorder(root->right);
        cout << root->data << " ";
    }
}

int main()
{
    Node *root = new Node(1);
    root->left = new Node(2);
    root->right = new Node(3);
    root->left->left = new Node(4);

    cout << "Level Order Traversal: ";
    bfs(root);
    cout << "Preorder Traversal: ";
    preorder(root);
    cout << "Inorder Traversal: ";
    inorder(root);
    cout << "Postorder Traversal: ";
    postorder(root);

    return 0;
}