// https://leetcode.com/problems/binary-tree-inorder-traversal/

#include <bits/stdc++.h>

using namespace std;

struct node
{
    int data;
    struct node *left, *right;
};

// iterative
vector<int> inOrderTrav(node *curr)
{
    vector<int> inOrder;
    stack<node *> s;
    while (true)
    {
        if (curr != NULL)
        {
            s.push(curr);
            curr = curr->left;
        }
        else
        {
            if (s.empty())
                break;
            curr = s.top();
            inOrder.push_back(curr->data);
            s.pop();
            curr = curr->right;
        }
    }
    return inOrder;
}

// Time Complexity: O(N).
// Reason: We are traversing N nodes and every node is visited exactly once.
// Space Complexity: O(N)
// Reason: In the worst case (a tree with just left children), the space complexity will be O(N).

// recursive
void inOrderTrav(node *curr, vector<int> &inOrder)
{
    if (curr == NULL)
        return;

    inOrderTrav(curr->left, inOrder);
    inOrder.push_back(curr->data);
    inOrderTrav(curr->right, inOrder);
}

// Time Complexity: O(N)
// Space Complexity: O(N)

struct node *newNode(int data)
{
    struct node *node = (struct node *)malloc(sizeof(struct node));
    node->data = data;
    node->left = NULL;
    node->right = NULL;

    return (node);
}

int main()
{

    struct node *root = newNode(1);
    root->left = newNode(2);
    root->right = newNode(3);
    root->left->left = newNode(4);
    root->left->right = newNode(5);
    root->left->right->left = newNode(8);
    root->right->left = newNode(6);
    root->right->right = newNode(7);
    root->right->right->left = newNode(9);
    root->right->right->right = newNode(10);

    vector<int> inOrder;

    // iterative
    inOrder = inOrderTrav(root);
    
    // recursive
    inOrderTrav(root, inOrder);

    cout << "The inOrder Traversal is : ";
    for (int i = 0; i < inOrder.size(); i++)
    {
        cout << inOrder[i] << " ";
    }
    return 0;
}