// https://leetcode.com/problems/binary-tree-preorder-traversal/
// https://leetcode.com/problems/binary-tree-level-order-traversal/
// https://leetcode.com/problems/binary-tree-level-order-traversal-ii/
// https://leetcode.com/problems/n-ary-tree-level-order-traversal/
// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
// https://leetcode.com/problems/n-ary-tree-postorder-traversal/description/
// https://www.codingninjas.com/codestudio/problems/boundary-traversal_790725?utm_source=youtube&utm_medium=affiliate&utm_campaign=Striver_Tree_Videos
// https://practice.geeksforgeeks.org/problems/boundary-traversal-of-binary-tree/0
// https://leetcode.com/problems/boundary-of-binary-tree/description/
// https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/
// https://practice.geeksforgeeks.org/problems/top-view-of-binary-tree/1
// https://takeuforward.org/data-structure/top-view-of-a-binary-tree/
// https://practice.geeksforgeeks.org/problems/bottom-view-of-binary-tree/1
// https://takeuforward.org/data-structure/bottom-view-of-a-binary-tree/
// https://practice.geeksforgeeks.org/problems/left-view-of-binary-tree/1
// https://leetcode.com/problems/binary-tree-right-side-view/
// https://takeuforward.org/data-structure/morris-inorder-traversal-of-a-binary-tree/
// https://leetcode.com/problems/n-ary-tree-preorder-traversal/


#include <iostream>
#include <queue>
#include <stack>
#include <set>
#include <map>
using namespace std;

class Node {
    public:
    int data;
    Node *left, *right;

    Node(int data) {
        this->data = data;
        left = right = nullptr;
    }
};

Node* createNode(int data) {
    return new Node(data);
}


/*************************************************  INSERT NODE ***************************************/
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

/*************************************************  DFS ***************************************/

void postorder(Node *node)
{
    if (node == NULL)
        return;

    // first recur on left subtree
    postorder(node->left);

    // then recur on right subtree
    postorder(node->right);

    // now deal with the node
    cout << node->data << " ";
}
// Time Complexity: O(N).
// Reason: We are traversing N nodes and every node is visited exactly once.
// Space Complexity: O(N)
// Reason: Space is needed for the recursion stack. In the worst case (skewed tree), space complexity can be O(N).

// using single stack
vector<int> portorder_iterative_ss(Node *cur)
{

    vector<int> postOrder;
    if (cur == NULL)
        return postOrder;

    stack<Node *> st;
    while (cur != NULL || !st.empty())
    {

        if (cur != NULL)
        {
            st.push(cur);
            cur = cur->left;
        }
        else
        {
            Node *temp = st.top()->right;
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
vector<int> portorder_iterative_ts(Node *curr)
{

    vector<int> postOrder;
    if (curr == NULL)
        return postOrder;

    stack<Node *> s1;
    stack<Node *> s2;
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

void inorder(Node *node)
{
    if (node == NULL)
        return;

    /* first recur on left child */
    inorder(node->left);

    /* then print the data of node */
    cout << node->data << " ";

    /* now recur on right child */
    inorder(node->right);
}
// Time Complexity: O(N)
// Space Complexity: O(N)

vector<int> inorder_iterative(Node *curr)
{
    vector<int> inOrder;
    stack<Node *> s;
    while (true)
    {
        if (curr != NULL)
        {
            s.push(curr);
            curr = curr->left;
        }
        else
        {
            if (s.empty())
                break;
            curr = s.top();
            inOrder.push_back(curr->data);
            s.pop();
            curr = curr->right;
        }
    }
    return inOrder;
}

// Time Complexity: O(N).
// Reason: We are traversing N nodes and every node is visited exactly once.
// Space Complexity: O(N)
// Reason: In the worst case (a tree with just left children), the space complexity will be O(N).

void preorder(Node *node)
{
    if (node == NULL)
        return;

    /* first print data of node */
    cout << node->data << " ";

    /* then recur on left subtree */
    preorder(node->left);

    /* now recur on right subtree */
    preorder(node->right);
}

// Time Complexity: O(N).
// Reason: We are traversing N nodes and every node is visited exactly once.
// Space Complexity: O(N)
// Reason: Space is needed for the recursion stack. In the worst case (skewed tree), space complexity can be O(N).

vector<int> preorder_iterative(Node *curr)
{
    vector<int> preOrder;
    if (curr == NULL)
        return preOrder;

    stack<Node *> s;
    s.push(curr);

    while (!s.empty())
    {
        Node *topNode = s.top();
        preOrder.push_back(topNode->data);
        s.pop();
        // since preorder is root left right and stack is LIFO, so first put right as left will be at top
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
// Reason: In the worst case, (a tree with every node having a single right child and left-subtree), the space complexity can be considered as O(N).

/*************************************************  All DFS in one traversal *****************************/
void allTraversal(Node *root, vector<int> &pre, vector<int> &in, vector<int> &post)
{
    stack<pair<Node *, int>> st;
    st.push({root,
             1});
    if (root == NULL)
        return;

    while (!st.empty())
    {
        auto it = st.top();
        st.pop();

        // this is part of pre
        // increment 1 to 2
        // push the left side of the tree
        if (it.second == 1)
        {
            pre.push_back(it.first->data);
            it.second++;
            st.push(it);

            if (it.first->left != NULL)
            {
                st.push({it.first->left,
                         1});
            }
        }

        // this is a part of in
        // increment 2 to 3
        // push right
        else if (it.second == 2)
        {
            in.push_back(it.first->data);
            it.second++;
            st.push(it);

            if (it.first->right != NULL)
            {
                st.push({it.first->right,
                         1});
            }
        }
        // don't push it back again
        else
        {
            post.push_back(it.first->data);
        }
    }
}

/*************************************************  Morris Inorder ***************************************/
vector<int> morrisInorderTraversal(Node *root)
{
    vector<int> inorder;

    Node *cur = root;
    while (cur != NULL)
    {
        if (cur->left == NULL)
        {
            inorder.push_back(cur->data);
            cur = cur->right;
        }
        else
        {
            Node *prev = cur->left;
            while (prev->right != NULL && prev->right != cur)
            {
                prev = prev->right;
            }

            if (prev->right == NULL)
            {
                prev->right = cur;
                cur = cur->left;
            }
            else
            {
                prev->right = NULL;
                inorder.push_back(cur->data);
                cur = cur->right;
            }
        }
    }
    return inorder;
}
// Time Complexity: O(N).
// Space Complexity: O(1)

/*************************************************  Morris Preorder ***************************************/
vector<int> morrisPreorderTraversal(Node *root)
{
    vector<int> preorder;

    Node *cur = root;
    while (cur != NULL)
    {
        if (cur->left == NULL)
        {
            preorder.push_back(cur->data);
            cur = cur->right;
        }
        else
        {
            Node *prev = cur->left;
            while (prev->right != NULL && prev->right != cur)
            {
                prev = prev->right;
            }

            if (prev->right == NULL)
            {
                prev->right = cur;
                preorder.push_back(cur->data);
                cur = cur->left;
            }
            else
            {
                prev->right = NULL;
                cur = cur->right;
            }
        }
    }
    return preorder;
}
// Time Complexity: O(N).
// Space Complexity: O(1)

/*************************************************  BFS ***************************************/
void printLevelOrder(Node *root)
{
    // Base Case
    if (root == NULL)
        return;

    // Create an empty queue for level order traversal
    queue<Node *> q;

    // Enqueue Root and initialize height
    q.push(root);

    while (!q.empty())
    {
        // Print front of queue and remove it from queue
        Node *node = q.front();
        cout << node->data << " ";
        q.pop();

        /* Enqueue left child */
        if (node->left != NULL)
            q.push(node->left);

        /*Enqueue right child */
        if (node->right != NULL)
            q.push(node->right);
    }
}
// Time Complexity: O(N)
// Space Complexity: O(N)

/*************************************************  ZIGZAG TRAVERSAL ***************************************/

vector<vector<int>> zigzagLevelOrder(Node *root)
{
    vector<vector<int>> result;
    if (root == NULL)
    {
        return result;
    }

    queue<Node *> nodesQueue;
    nodesQueue.push(root);
    bool leftToRight = true;

    while (!nodesQueue.empty())
    {
        int size = nodesQueue.size();
        vector<int> row(size);
        for (int i = 0; i < size; i++)
        {
            Node *node = nodesQueue.front();
            nodesQueue.pop();

            // find position to fill node's value
            int index = (leftToRight) ? i : (size - 1 - i);

            row[index] = node->data;
            if (node->left)
            {
                nodesQueue.push(node->left);
            }
            if (node->right)
            {
                nodesQueue.push(node->right);
            }
        }
        // after this level
        leftToRight = !leftToRight;
        result.push_back(row);
    }
    return result;
}
// Time Complexity: O(N)
// Space Complexity: O(N)

/*************************************************  BOUNDARY ORDER TRAVERSAL ***************************************/
bool isLeaf(Node *root)
{
    return !root->left && !root->right;
}

void addLeftBoundary(Node *root, vector<int> &res)
{
    Node *cur = root->left;
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
void addRightBoundary(Node *root, vector<int> &res)
{
    Node *cur = root->right;
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

void addLeaves(Node *root, vector<int> &res)
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

vector<int> printBoundary(Node *root)
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

// Time Complexity: O(N).
// Reason: The time complexity will be O(H) + O(H) + O(N) which is ≈ O(N)
// Space Complexity: O(N)
// Reason: Space is needed for the recursion stack while adding leaves. In the worst case (skewed tree), space complexity can be O(N).

/*************************************************  VERTICAL ORDER TRAVERSAL ***************************************/
vector<vector<int>> findVertical(Node *root)
{
    map<int, map<int, multiset<int>>> nodes;
    queue<pair<Node *, pair<int, int>>> todo;
    todo.push({root,
               {0,
                0}}); // initial vertical and level
    while (!todo.empty())
    {
        auto p = todo.front();
        todo.pop();
        Node *temp = p.first;

        // x -> vertical , y->level
        int x = p.second.first, y = p.second.second;
        nodes[x][y].insert(temp->data); // inserting to multiset

        if (temp->left)
        {
            todo.push({temp->left,
                       {x - 1,
                        y + 1}});
        }
        if (temp->right)
        {
            todo.push({temp->right,
                       {x + 1,
                        y + 1}});
        }
    }
    vector<vector<int>> ans;
    for (auto p : nodes)
    {
        vector<int> col;
        for (auto q : p.second)
        {
            col.insert(col.end(), q.second.begin(), q.second.end());
        }
        ans.push_back(col);
    }
    return ans;
}

// Time Complexity: O(N*logN*logN*logN)
// Space Complexity: O(N)

/*************************************************  TOP VIEW ***************************************/
vector<int> topView(Node *root)
{
    vector<int> ans;
    if (root == NULL)
        return ans;
    map<int, int> mpp;
    queue<pair<Node *, int>> q;
    q.push({root, 0});
    while (!q.empty())
    {
        auto it = q.front();
        q.pop();
        Node *node = it.first;
        int line = it.second;
        if (mpp.find(line) == mpp.end())
            mpp[line] = node->data;

        if (node->left != NULL)
        {
            q.push({node->left, line - 1});
        }
        if (node->right != NULL)
        {
            q.push({node->right, line + 1});
        }
    }

    for (auto it : mpp)
    {
        ans.push_back(it.second);
    }
    return ans;
}
// Time Complexity: O(N)
// Space Complexity: O(N)

/*************************************************   BOTTOM VIEW ***************************************/
vector<int> bottomView(Node *root)
{
    vector<int> ans;
    if (root == NULL)
        return ans;
    map<int, int> mpp;
    queue<pair<Node *, int>> q;
    q.push({root, 0});
    while (!q.empty())
    {
        auto it = q.front();
        q.pop();
        Node *node = it.first;
        int line = it.second;
        mpp[line] = node->data;

        if (node->left != NULL)
        {
            q.push({node->left, line - 1});
        }
        if (node->right != NULL)
        {
            q.push({node->right, line + 1});
        }
    }

    for (auto it : mpp)
    {
        ans.push_back(it.second);
    }
    return ans;
}
// Time Complexity: O(N)
// Space Complexity: O(N)

/*************************************************   LEFT VIEW ***************************************/
void recursionLeftView(Node *root, int level, vector<int> &res)
{
    if (root == NULL)
        return;
    if (res.size() == level)
        res.push_back(root->data);
    recursionLeftView(root->left, level + 1, res);
    recursionLeftView(root->right, level + 1, res);
}

vector<int> leftSideView(Node *root)
{
    vector<int> res;
    recursionLeftView(root, 0, res);
    return res;
}
// Time Complexity : O(N)
// Space Complexity : O(H)(H->Height of the Tree)

/*************************************************   RIGHT VIEW ***************************************/
void recursion(Node *root, int level, vector<int> &res)
{
    if (root == NULL)
        return;
    if (res.size() == level)
        res.push_back(root->data);
    recursion(root->right, level + 1, res);
    recursion(root->left, level + 1, res);
}

vector<int> rightSideView(Node *root)
{
    vector<int> res;
    recursion(root, 0, res);
    return res;
}
// Time Complexity : O(N)
// Space Complexity : O(H)(H->Height of the Tree)

/************************************* n-ary tree preorder ********************************************************/
class NarryNode
{
public:
    int val;
    vector<NarryNode *> children;

    NarryNode() {}

    NarryNode(int _val)
    {
        val = _val;
    }

    NarryNode(int _val, vector<NarryNode *> _children)
    {
        val = _val;
        children = _children;
    }

private:
    void travel(NarryNode *root, vector<int> &result)
    {
        if (root == nullptr)
        {
            return;
        }

        result.push_back(root->val);
        for (NarryNode *child : root->children)
        {
            travel(child, result);
        }
    }

public:
    vector<int> preorderRecursive(NarryNode *root)
    {
        vector<int> result;
        travel(root, result);
        return result;
    }

    vector<int> preorderIterative(NarryNode *root)
    {
        vector<int> v;
        traverse(root, v);
        return v;
    }

    void traverse(NarryNode *node, vector<int> &v)
    {
        if (node == NULL)
            return;

        stack<NarryNode *> s;
        s.push(node);

        while (!s.empty())
        {
            NarryNode *n = s.top();
            s.pop();
            v.push_back(n->val);
            for (int i = n->children.size() - 1; i >= 0; i--)
            {
                s.push(n->children[i]);
            }
        }
    }
};

/*********************************************************************************************/


int main()
{

    //         1
    //       /   \
    //      2     7
    //     /       \
    //    3         8
    //     \       /
    //      4     9
    //     / \   / \
    //    5   6 10 11

    Node *root = createNode(1);
    root->left = createNode(2);
    root->left->left = createNode(3);
    root->left->left->right = createNode(4);
    root->left->left->right->left = createNode(5);
    root->left->left->right->right = createNode(6);
    root->right = createNode(7);
    root->right->right = createNode(8);
    root->right->right->left = createNode(9);
    root->right->right->left->left = createNode(10);
    root->right->right->left->right = createNode(11);

    // cout << "\nPreorder traversal of binary tree is \n";
    // preorder(root);

    // cout << "\nInorder traversal of binary tree is \n";
    // inorder(root);

    // cout << "\nPostorder traversal of binary tree is \n";
    // postorder(root);

    // vector<int> pre, in, post;
    // allTraversal(root, pre, in, post);
    // cout << "The preorder Traversal is : ";
    // for (auto nodeVal : pre)
    // {
    //     cout << nodeVal << " ";
    // }
    // cout << endl;
    // cout << "The inorder Traversal is : ";
    // for (auto nodeVal : in)
    // {
    //     cout << nodeVal << " ";
    // }
    // cout << endl;
    // cout << "The postorder Traversal is : ";
    // for (auto nodeVal : post)
    // {
    //     cout << nodeVal << " ";
    // }
    // cout << endl;

    // cout << "\nLevel Order traversal of binary tree is \n";
    // printLevelOrder(root);

    // vector<vector<int>> ans;
    // ans = zigzagLevelOrder(root);
    // cout << "\nZig Zag Traversal of Binary Tree" << endl;
    // for (int i = 0; i < ans.size(); i++)
    // {
    //     for (int j = 0; j < ans[i].size(); j++)
    //     {
    //         cout << ans[i][j] << " ";
    //     }
    //     cout << endl;
    // }

    // vector<int> boundaryTraversal;
    // boundaryTraversal = printBoundary(root);

    // cout << "The Boundary Traversal is : ";
    // for (int i = 0; i < boundaryTraversal.size(); i++)
    // {
    //     cout << boundaryTraversal[i] << " ";
    // }

    // vector<vector<int>> verticalTraversal;
    // verticalTraversal = findVertical(root);

    // cout << "The Vertical Traversal is : " << endl;
    // for (auto vertical : verticalTraversal)
    // {
    //     for (auto nodeVal : vertical)
    //     {
    //         cout << nodeVal << " ";
    //     }
    //     cout << endl;
    // }

    // vector<int> topViewTraversal;
    // topViewTraversal = topView(root);

    // cout << "The Top View Traversal is : " << endl;
    // for (auto view : topViewTraversal)
    // {
    //     cout << view << " ";
    // }
    // cout << endl;

    // vector<int> bottomViewTraversal;
    // bottomViewTraversal = bottomView(root);

    // cout << "The Top View Traversal is : " << endl;
    // for (auto view : bottomViewTraversal)
    // {
    //     cout << view << " ";
    // }
    // cout << endl;

    return 0;
}
