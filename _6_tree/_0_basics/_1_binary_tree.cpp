// https://leetcode.com/problems/binary-tree-preorder-traversal/
// https://leetcode.com/problems/binary-tree-level-order-traversal/
// https://leetcode.com/problems/binary-tree-level-order-traversal-ii/
// https://leetcode.com/problems/n-ary-tree-level-order-traversal/
// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
// https://leetcode.com/problems/n-ary-tree-postorder-traversal/description/

#include <iostream>
#include <queue>
#include <stack>
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

/*************************************************  VERTICAL ORDER TRAVERSAL ***************************************/
/*************************************************  BOUNDARY ORDER TRAVERSAL ***************************************/
/*************************************************  TOP VIEW ***************************************/
/*************************************************   BOTTOM VIEW ***************************************/
/*************************************************   LEFT VIEW ***************************************/
/*************************************************   RIGHT VIEW ***************************************/










int main()
{

    Node* root = createNode(1);
    root->left = createNode(2);
    root->right = createNode(3);
    root->left->left = createNode(4);
    root->left->right = createNode(5);

    // cout << "\nPreorder traversal of binary tree is \n";
    // preorder(root);

    // cout << "\nInorder traversal of binary tree is \n";
    // inorder(root);

    // cout << "\nPostorder traversal of binary tree is \n";
    // postorder(root);

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

    return 0;
}
