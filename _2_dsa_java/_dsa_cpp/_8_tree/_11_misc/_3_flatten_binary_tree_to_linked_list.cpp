// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/

#include <bits/stdc++.h>

using namespace std;

struct node
{
    int data;
    struct node *left, *right;
};

// using recursion
class Solution
{
    node *prev = NULL;

public:
    void flatten(node *root)
    {
        if (root == NULL)
            return;

        flatten(root->right);
        flatten(root->left);

        root->right = prev;
        root->left = NULL;
        prev = root;
    }
    //     Time Complexity: O(N)
    // Reason: We are doing a simple preorder traversal.
    // Space Complexity: O(N)
    // Reason: Worst case time complexity will be O(N) (in case of skewed tree).
};

// iterative - using stack
class Solution
{
    node *prev = NULL;

public:
    void flatten(node *root)
    {
        if (root == NULL)
            return;
        stack<node *> st;
        st.push(root);
        while (!st.empty())
        {
            node *cur = st.top();
            st.pop();

            if (cur->right != NULL)
            {
                st.push(cur->right);
            }
            if (cur->left != NULL)
            {
                st.push(cur->left);
            }
            if (!st.empty())
            {
                cur->right = st.top();
            }
            cur->left = NULL;
        }
    }
};
// Time Complexity: O(N)
// Reason: The loop will execute for every node once.
// Space Complexity: O(N)

// using morris traversal
class Solution
{
    node *prev = NULL;

public:
    void flatten(node *root)
    {
        node *cur = root;
        while (cur)
        {
            if (cur->left)
            {
                node *pre = cur->left;
                while (pre->right)
                {
                    pre = pre->right;
                }
                pre->right = cur->right;
                cur->right = cur->left;
                cur->left = NULL;
            }
            cur = cur->right;
        }
    }
};
// Time Complexity: O(N)
// Reason: Time complexity will be the same as that of a morris traversal
// Space Complexity: O(1)
// Reason: We are not using any extra space

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
    root->left->right = newNode(4);
    root->right = newNode(5);
    root->right->right = newNode(6);
    root->right->right->left = newNode(7);

    Solution obj;

    obj.flatten(root);
    while (root->right != NULL)
    {
        cout << root->data << "->";
        root = root->right;
    }
    cout << root->data;
    return 0;
}
