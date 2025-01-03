// https://www.codingninjas.com/codestudio/problems/boundary-traversal_790725?utm_source=youtube&utm_medium=affiliate&utm_campaign=Striver_Tree_Videos
// https://practice.geeksforgeeks.org/problems/boundary-traversal-of-binary-tree/0
// https://leetcode.com/problems/boundary-of-binary-tree/description/

#include <bits/stdc++.h>

using namespace std;

struct node
{
    int data;
    struct node *left, *right;
};

bool isLeaf(node *root)
{
    return !root->left && !root->right;
}

void addLeftBoundary(node *root, vector<int> &res)
{
    node *cur = root->left;
    while (cur)
    {
        if (!isLeaf(cur))
            res.push_back(cur->data);
        if (cur->left)
            cur = cur->left;
        else
            cur = cur->right;
    }
}
void addRightBoundary(node *root, vector<int> &res)
{
    node *cur = root->right;
    vector<int> tmp;
    while (cur)
    {
        if (!isLeaf(cur))
            tmp.push_back(cur->data);
        if (cur->right)
            cur = cur->right;
        else
            cur = cur->left;
    }
    for (int i = tmp.size() - 1; i >= 0; --i)
    {
        res.push_back(tmp[i]);
    }
}

void addLeaves(node *root, vector<int> &res)
{
    if (isLeaf(root))
    {
        res.push_back(root->data);
        return;
    }
    if (root->left)
        addLeaves(root->left, res);
    if (root->right)
        addLeaves(root->right, res);
}

vector<int> printBoundary(node *root)
{
    vector<int> res;
    if (!root)
        return res;

    if (!isLeaf(root))
        res.push_back(root->data);

    addLeftBoundary(root, res);

    // add leaf nodes
    addLeaves(root, res);

    addRightBoundary(root, res);
    return res;
}

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
    root->left->left = newNode(3);
    root->left->left->right = newNode(4);
    root->left->left->right->left = newNode(5);
    root->left->left->right->right = newNode(6);
    root->right = newNode(7);
    root->right->right = newNode(8);
    root->right->right->left = newNode(9);
    root->right->right->left->left = newNode(10);
    root->right->right->left->right = newNode(11);

    vector<int> boundaryTraversal;
    boundaryTraversal = printBoundary(root);

    cout << "The Boundary Traversal is : ";
    for (int i = 0; i < boundaryTraversal.size(); i++)
    {
        cout << boundaryTraversal[i] << " ";
    }
    return 0;
}

// Time Complexity: O(N).
// Reason: The time complexity will be O(H) + O(H) + O(N) which is ≈ O(N)
// Space Complexity: O(N)
// Reason: Space is needed for the recursion stack while adding leaves. In the worst case (skewed tree), space complexity can be O(N).