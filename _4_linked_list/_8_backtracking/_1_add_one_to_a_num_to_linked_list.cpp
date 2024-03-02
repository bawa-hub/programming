// https://practice.geeksforgeeks.org/problems/add-1-to-a-number-represented-as-linked-list/1

struct Node
{
    int data;
    struct Node* next;
    
    Node(int x){
        data = x;
        next = NULL;
    }
};

// iterative 
// reverse, add, reverse
// TC: O(3N)

// recursive
class Solution
{
    public:
    Node* addOne(Node *head) 
    {
        int carry = addHelper(head);
        if(carry == 1) {
           Node* newNode = new Node(1);
           newNode->next = head;
           head = newNode;
        }
        return head;
    }


    int addHelper(Node* temp) {
        if(temp==nullptr) return 1; // this 1 is num that is to be added

        int carry = addHelper(temp->next);
        temp->data += carry;
        if(temp->data < 10) return 0;
        temp->data = 0;
        return 1;

    }
};
// TC: O(n)
// SC: O(n)