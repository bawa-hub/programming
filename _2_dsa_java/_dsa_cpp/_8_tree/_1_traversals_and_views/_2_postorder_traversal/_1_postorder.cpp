// https://leetcode.com/problems/binary-tree-postorder-traversal/description/

#include <bits/stdc++.h>

using namespace std;

struct node
{
    int data;
    struct node *left, *right;
};

/*********************** Iterative *******************/

// using single stack
vector<int> postOrderTrav(node *cur)
{

    vector<int> postOrder;
    if (cur == NULL)
        return postOrder;

    stack<node *> st;
    while (cur != NULL || !st.empty())
    {

        if (cur != NULL)
        {
            st.push(cur);
            cur = cur->left;
        }
        else
        {
            node *temp = st.top()->right;
            if (temp == NULL)
            {
                temp = st.top();
                st.pop();
                postOrder.push_back(temp->data);
                while (!st.empty() && temp == st.top()->right)
                {
                    temp = st.top();
                    st.pop();
                    postOrder.push_back(temp->data);
                }
            }
            else
                cur = temp;
        }
    }
    return postOrder;
}

// Time Complexity: O(N).
// Space Complexity: O(N)

// using two stack
vector<int> postOrderTrav(node *curr)
{

    vector<int> postOrder;
    if (curr == NULL)
        return postOrder;

    stack<node *> s1;
    stack<node *> s2;
    s1.push(curr);
    while (!s1.empty())
    {
        curr = s1.top();
        s1.pop();
        s2.push(curr);
        if (curr->left != NULL)
            s1.push(curr->left);
        if (curr->right != NULL)
            s1.push(curr->right);
    }
    while (!s2.empty())
    {
        postOrder.push_back(s2.top()->data);
        s2.pop();
    }
    return postOrder;
}

// Time Complexity: O(N).
// Reason: We are traversing N nodes and every node is visited exactly once.
// Space Complexity: O(N+N)


/************************* Recursive ********************/
void postOrderTrav(node *curr, vector<int> &postOrder)
{
    if (curr == NULL)
        return;

    postOrderTrav(curr->left, postOrder);
    postOrderTrav(curr->right, postOrder);
    postOrder.push_back(curr->data);
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

    vector<int> postOrder;

    // iterative
    postOrder = postOrderTrav(root);

    // recursive
    postOrderTrav(root, postOrder);

    cout << "The postOrder Traversal is : ";
    for (int i = 0; i < postOrder.size(); i++)
    {
        cout << postOrder[i] << " ";
    }
    return 0;
}