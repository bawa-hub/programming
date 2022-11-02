#include <iostream>
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

// dfs
void printPostorder(struct Node *node)
{
    if (node == NULL)
        return;

    // first recur on left subtree
    printPostorder(node->left);

    // then recur on right subtree
    printPostorder(node->right);

    // now deal with the node
    cout << node->data << " ";
}

void printInorder(struct Node *node)
{
    if (node == NULL)
        return;

    /* first recur on left child */
    printInorder(node->left);

    /* then print the data of node */
    cout << node->data << " ";

    /* now recur on right child */
    printInorder(node->right);
}

void printInOrderIterative(struct Node *root)
{
    stack<Node *> s;
    Node *curr = root;

    while (curr != NULL || s.empty() == false)
    {
        while (curr != NULL)
        {
            s.push(curr);
            curr = curr->left;
        }

        curr = s.top();
        s.pop();

        cout << curr->data << " ";

        curr = curr->right;
    }
}

void printPreorder(struct Node *node)
{
    if (node == NULL)
        return;

    /* first print data of node */
    cout << node->data << " ";

    /* then recur on left subtree */
    printPreorder(node->left);

    /* now recur on right subtree */
    printPreorder(node->right);
}

void printPreorderIterative(Node *root)
{
    if (root == NULL)
        return;

    stack<Node *> nodeStack;
    nodeStack.push(root);

    while (nodeStack.empty() == false)
    {
        struct Node *node = nodeStack.top();
        printf("%d ", node->data);
        nodeStack.pop();

        if (node->right)
            nodeStack.push(node->right);
        if (node->left)
            nodeStack.push(node->left);
    }
}

// bfs or level order
void printLevelOrder(Node *root)
{
    // Base Case
    if (root == NULL)
        return;

    // Create an empty queue for level order traversal
    queue<Node *> q;

    // Enqueue Root and initialize height
    q.push(root);

    while (!q.empty())
    {
        // Print front of queue and remove it from queue
        Node *node = q.front();
        cout << node->data << " ";
        q.pop();

        /* Enqueue left child */
        if (node->left != NULL)
            q.push(node->left);

        /*Enqueue right child */
        if (node->right != NULL)
            q.push(node->right);
    }
}

int main()
{
    struct Node *root = createNode(1);
    root->left = createNode(2);
    root->right = createNode(3);
    root->left->left = createNode(4);
    root->left->right = createNode(5);

    cout << "\nPreorder traversal of binary tree is \n";
    printPreorder(root);

    cout << "\nPreorder Iterative traversal of binary tree is \n";
    printPreorderIterative(root);

    cout << "\nInorder traversal of binary tree is \n";
    printInorder(root);

    cout << "\nInorder Iterative traversal of binary tree is \n";
    printInOrderIterative(root);

    cout << "\nPostorder traversal of binary tree is \n";
    printPostorder(root);

    cout << "\nLevel Order traversal of binary tree is \n";
    printLevelOrder(root);

    return 0;
}
