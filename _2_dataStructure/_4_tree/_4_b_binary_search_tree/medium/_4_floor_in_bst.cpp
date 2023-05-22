// https://www.codingninjas.com/codestudio/problems/floor-from-bst_920457?source=youtube&campaign=Striver_Tree_Videos&utm_source=youtube&utm_medium=affiliate&utm_campaign=Striver_Tree_Videos

int floorInBST(TreeNode<int> *root, int key)
{
    int floor = -1;
    while (root)
    {

        if (root->val == key)
        {
            floor = root->val;
            return floor;
        }

        if (key > root->val)
        {
            floor = root->val;
            root = root->right;
        }
        else
        {
            root = root->left;
        }
    }
    return floor;
}