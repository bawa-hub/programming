// https://www.codingninjas.com/codestudio/problems/ceil-from-bst_920464?source=youtube&campaign=Striver_Tree_Videos&utm_source=youtube&utm_medium=affiliate&utm_campaign=Striver_Tree_Videos

int findCeil(BinaryTreeNode<int> *root, int key)
{

    int ceil = -1;
    while (root)
    {

        if (root->data == key)
        {
            ceil = root->data;
            return ceil;
        }

        if (key > root->data)
        {
            root = root->right;
        }
        else
        {
            ceil = root->data;
            root = root->left;
        }
    }
    return ceil;
}