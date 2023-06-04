// https://practice.geeksforgeeks.org/problems/root-to-leaf-paths/1
// https://www.geeksforgeeks.org/given-a-binary-tree-print-all-root-to-leaf-paths/

#include<bits/stdc++.h>
using namespace std;

struct Node
{
    int data;
    struct Node* left;
    struct Node* right;
 
    Node(int x){
        data = x;
        left = right = NULL;
    }
};
 
void helper(Node* root,vector<int> arr,vector<vector<int>> &ans)
{
    if(!root)
        return;
    arr.push_back(root->data);
    if(root->left==NULL && root->right==NULL)
    {
       ans.push_back(arr);
        return;
    }
    helper(root->left,arr,ans);
    helper(root->right,arr,ans);
}
vector<vector<int>> Paths(Node* root)
{
    vector<vector<int>> ans;
    if(!root)
        return ans;
    vector<int> arr;
    helper(root,arr,ans);
    return ans;
}