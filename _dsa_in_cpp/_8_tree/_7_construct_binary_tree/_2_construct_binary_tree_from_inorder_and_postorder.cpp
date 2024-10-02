// https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

#include <bits/stdc++.h>

using namespace std;

struct node
{
    int data;
    struct node *left, *right;
};

struct node *newNode(int data)
{
    struct node *node = (struct node *)malloc(sizeof(struct node));
    node->data = data;
    node->left = NULL;
    node->right = NULL;

    return (node);
}

node *constructTree(vector<int> &postorder, int postStart, int postEnd,
                    vector<int> &inorder, int inStart, int inEnd, map<int, int> &mp)
{
    if (postStart > postEnd || inStart > inEnd)
        return NULL;

    node *root = newNode(postorder[postEnd]);
    int elem = mp[root->data];
    int nElem = elem - inStart;

    root->left = constructTree(postorder, postStart, postStart + nElem - 1,
                               inorder, inStart, elem - 1, mp);
    root->right = constructTree(postorder, postStart + nElem, postEnd - 1, inorder,
                                elem + 1, inEnd, mp);

    return root;
}

node *buildTree(vector<int> &postorder, vector<int> &inorder)
{
    int postStart = 0, postEnd = postorder.size() - 1;
    int inStart = 0, inEnd = inorder.size() - 1;

    map<int, int> mp;
    for (int i = inStart; i <= inEnd; i++)
    {
        mp[inorder[i]] = i;
    }

    return constructTree(postorder, postStart, postEnd, inorder, inStart, inEnd, mp);
}

int main()
{

    vector<int> postorder{40, 50, 20, 60, 30, 10};
    vector<int> inorder{40, 20, 50, 10, 60, 30};
    node *root = buildTree(postorder, inorder);
    return 0;
}

// Time Complexity: O(N)
// Assumption: Hashmap returns the answer in constant time.
// Space Complexity: O(N)
// Reason: We are using an external hashmap of size ‘N’.