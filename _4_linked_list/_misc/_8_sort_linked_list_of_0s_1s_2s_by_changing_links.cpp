// https://practice.geeksforgeeks.org/problems/given-a-linked-list-of-0s-1s-and-2s-sort-it/1

 struct Node {
    int data;
    struct Node *next;
    Node(int x) {
        data = x;
        next = NULL;
    }
};

class Solution
{
    public:
    //Function to sort a linked list of 0s, 1s and 2s.
    Node* segregate(Node *head) {
        
        if(!head || !head->next) return head;

        Node* zeroHead = new Node(-1);
        Node* oneHead = new Node(-1);
        Node* twoHead = new Node(-1);

        Node* zero = zeroHead;
        Node* one = oneHead;
        Node* two = twoHead;
        Node* temp = head;

        while(temp) {
            if(temp->data == 0) {
                zero->next = temp;
                zero = zero->next;
            } else if(temp->data == 1) {
                one->next = temp;
                one = one->next;
            } else {
                two->next = temp;
                two = two->next;
            }
            temp = temp->next;
        }

        zero->next = (oneHead->next) ? (oneHead->next) : (twoHead->next);
        one->next = twoHead->next;
        two->next = NULL;

        Node* newHead = zeroHead->next;

        delete zeroHead;
        delete oneHead;
        delete twoHead;
        return newHead;
        
    }
};