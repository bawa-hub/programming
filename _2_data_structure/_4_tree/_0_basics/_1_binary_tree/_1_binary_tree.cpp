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

    // cout << "\nPreorder traversal of binary tree is \n";
    // printPreorder(root);

    // cout << "\nPreorder Iterative traversal of binary tree is \n";
    // printPreorderIterative(root);

    // cout << "\nInorder traversal of binary tree is \n";
    // printInorder(root);

    // cout << "\nInorder Iterative traversal of binary tree is \n";
    // printInOrderIterative(root);

    // cout << "\nPostorder traversal of binary tree is \n";
    // printPostorder(root);

    // cout << "\nPostorder Iterative traversal of binary tree is \n";
    // vector<int> postOrder;
    // postOrder = printPostorderIterativeUsingTwoStack(root);
    // for (int i = 0; i < postOrder.size(); i++)
    // {
    //     cout << postOrder[i] << " ";
    // }

    // cout << "\nLevel Order traversal of binary tree is \n";
    // printLevelOrder(root);

    // cout << "\nZig Zag Traversal of Binary Tree \n";
    // vector<vector<int>> ans;
    // ans = zigzagLevelOrder(root);
    // for (int i = 0; i < ans.size(); i++)
    // {
    //     for (int j = 0; j < ans[i].size(); j++)
    //     {
    //         cout << ans[i][j] << " ";
    //     }
    //     cout << endl;
    // }

    return 0;
}
