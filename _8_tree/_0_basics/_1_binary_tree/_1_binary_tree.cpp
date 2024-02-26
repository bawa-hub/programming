#include <iostream>
#include <vector>
#include <queue>
#include <stack>

using namespace std;

struct Node
{
    int data;
    struct Node *left, *right;
};

// create node
Node *createNode(int data)
{
    Node *temp = new Node; // cpp style of struct
    // Node *temp = (struct Node *)malloc(sizeof(struct Node));  // dynamically allocate memory
    if (!temp)
    {
        cout << "Memory error\n";
        return NULL;
    }
    temp->data = data;
    temp->left = temp->right = NULL;
    return temp;
}

// insert node
Node *insertNode(Node *root, int data)
{
    if (root == NULL)
    {
        root = createNode(data);
        return root;
    }

    queue<Node *> q;
    q.push(root);

    while (!q.empty())
    {
        Node *temp = q.front();
        q.pop();

        if (temp->left != NULL)
            q.push(temp->left);
        else
        {
            temp->left = createNode(data);
            return root;
        }

        if (temp->right != NULL)
            q.push(temp->right);
        else
        {
            temp->right = createNode(data);
            return root;
        }
    }
    return root;
}

// delete node

int main()
{
    struct Node *root = createNode(1);
    root->left = createNode(2);
    root->right = createNode(3);
    root->left->left = createNode(4);
    root->left->right = createNode(5);

    return 0;
}
