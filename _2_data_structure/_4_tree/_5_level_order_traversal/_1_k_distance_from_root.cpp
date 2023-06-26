// https://practice.geeksforgeeks.org/problems/k-distance-from-root/1

#include <iostream>
#include <vector>
#include <queue>
using namespace std;

struct Node
{
    int data;
    struct Node *left, *right;
};

Node *newNode(int data)
{
    Node *temp = new Node;
    temp->data = data;
    temp->left = temp->right = nullptr;
    return temp;
}

void Kdistance(struct Node *root, int k)
{
    queue<pair<Node *, int>> q;
    vector<int> res;
    q.push({root, 0});

    while (!q.empty())
    {
        for (int i = 0; i < q.size(); i++)
        {
            struct Node *node = q.front().first;
            int l = q.front().second;
            q.pop();
            if (l == k)
            {
                cout << node->data << endl;
            }
            else
            {
                if (node->left != nullptr)
                    q.push({node->left, l + 1});
                if (node->right != nullptr)
                    q.push({node->right, l + 1});
            }
        }
    }
}

int main()
{
    Node *root = newNode(1);
    root->left = newNode(2);
    root->right = newNode(3);
    root->left->left = newNode(4);
    root->left->right = newNode(5);
    root->right->left = newNode(8);

    Kdistance(root, 2);
    return 0;
}