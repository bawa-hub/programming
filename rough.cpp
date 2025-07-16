#include <iostream>
#include <queue>
using namespace std;

class Node {
    public:
    int data;
    Node* left;
    Node* right;

    Node(int data) {
        this->data = data;
        this->left = nullptr;
        this->right = nullptr;
    }
};

void bfs(Node* root) {
    if(root==nullptr) return;
    queue<Node*> q;
    q.push(root);

    while (!q.empty()) {
        Node* node = q.front();
        q.pop();

        cout << node->data << " ";
        if(node->left!=nullptr) q.push(node->left);
        if(node->right!=nullptr) q.push(node->right);
    }
    
}

Node* createTree() {
    Node* root = new Node(1);
    root->left = new Node(2);
    root->right = new Node(3);
    root->left->left = new Node(4);
    root->left->right = new Node(5);
    root->right->left = new Node(6);
    root->right->right = new Node(7);
    return root;
}

int main() {
   Node* root = createTree();
   bfs(root);
}