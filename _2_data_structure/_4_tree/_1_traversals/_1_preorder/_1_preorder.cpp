// https://leetcode.com/problems/binary-tree-preorder-traversal/

#include <bits/stdc++.h>

using namespace std;

struct node
{
    int data;
    struct node *left, *right;
};

// iterative
vector<int> preOrderTrav(node *curr)
{
    vector<int> preOrder;
    if (curr == NULL)
        return preOrder;

    stack<node *> s;
    s.push(curr);

    while (!s.empty())
    {
        node *topNode = s.top();
        preOrder.push_back(topNode->data);
        s.pop();
        if (topNode->right != NULL)
            s.push(topNode->right);
        if (topNode->left != NULL)
            s.push(topNode->left);
    }
    return preOrder;
}

// Time Complexity: O(N).
// Reason: We are traversing N nodes and every node is visited exactly once.
// Space Complexity: O(N)
// Reason: In the worst case, (a tree with every node having a single right child and left-subtree, follow the video attached below to see the illustration), the space complexity can be considered as O(N).


// recursive
void preOrderTrav(node *curr, vector<int> &preOrder)
{
    if (curr == NULL)
        return;

    preOrder.push_back(curr->data);
    preOrderTrav(curr->left, preOrder);
    preOrderTrav(curr->right, preOrder);
}

// Time Complexity: O(N).
// Reason: We are traversing N nodes and every node is visited exactly once.
// Space Complexity: O(N)
// Reason: Space is needed for the recursion stack. In the worst case (skewed tree), space complexity can be O(N).

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

    vector<int> preOrder;
     
    //  iterative
    preOrder = preOrderTrav(root);

    // recursive
    preOrderTrav(root, preOrder);

    cout << "The preOrder Traversal is : ";
    for (int i = 0; i < preOrder.size(); i++)
    {
        cout << preOrder[i] << " ";
    }
    return 0;
}