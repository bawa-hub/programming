// https://practice.geeksforgeeks.org/problems/find-length-of-loop/1
// https://www.geeksforgeeks.org/find-length-of-loop-in-linked-list/

struct Node {
    int data;
    struct Node *next;
    Node(int x) {
        data = x;
        next = NULL;
    }
};

// using hashing
int countNodesinLoop(struct Node *head)
{
    unordered_map<Node*, int> mp;
    Node* temp = head;
    int timer = 1;
    
    while(temp!=nullptr) {
        if(mp.find(temp) != mp.end()) {
            return timer - mp[temp];
        }
        mp[temp] = timer;
        temp = temp->next;
        timer++;
    }
    
    return 0;
}
// TC: O(n*2*x)
// SC: O(n)

// using fast and slow pointer
int countNodesinLoop(struct Node *head)
{
    if (head == NULL)
        return false;
    Node *fast = head;
    Node *slow = head;

    while (fast->next != NULL && fast->next->next != NULL)
    {
        fast = fast->next->next;
        slow = slow->next;
        if (fast == slow)
            return findLength(slow, fast);
    }
    return 0;
}

int findLength(Node* slow, Node* fast) {
    int cnt = 1;
    fast = fast->next;
    while(slow!=fast) {
        cnt++;
        fast = fast->next;
    }
    return cnt;
}
// Time Complexity: O(N + lenOfLoop)
// Reason: In the worst case, all the nodes of the list are visited.
// Space Complexity: O(1)
// Reason: No extra data structure is used.