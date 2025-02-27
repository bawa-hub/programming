// https://leetcode.com/problems/linked-list-cycle/

#include <bits/stdc++.h>

class node
{
public:
    int num;
    node *next;
    node(int val)
    {
        num = val;
        next = NULL;
    }
};

// utility function to insert node in the list
void insertNode(node *&head, int val)
{
    node *newNode = new node(val);

    if (head == NULL)
    {
        head = newNode;
        return;
    }

    node *temp = head;
    while (temp->next != NULL)
        temp = temp->next;

    temp->next = newNode;
    return;
}

// utility function to create cycle
void createCycle(node *&head, int a, int b)
{
    int cnta = 0, cntb = 0;
    node *p1 = head;
    node *p2 = head;
    while (cnta != a || cntb != b)
    {
        if (cnta != a)
            p1 = p1->next, ++cnta;
        if (cntb != b)
            p2 = p2->next, ++cntb;
    }
    p2->next = p1;
}


// using hashing
bool cycleDetect(node *head)
{
    unordered_set<node *> hashTable;
    while (head != NULL)
    {
        if (hashTable.find(head) != hashTable.end())
            return true;
        hashTable.insert(head);
        head = head->next;
    }
    return false;
}
// Time Complexity: O(N)
// Reason: Entire list is iterated once.
// Space Complexity: O(N)
// Reason: All nodes present in the list are stored in a hash table.

// using slow and fast pointer
bool cycleDetect(node *head)
{
    if (head == NULL)
        return false;
    node *fast = head;
    node *slow = head;

    while (fast->next != NULL && fast->next->next != NULL)
    {
        fast = fast->next->next;
        slow = slow->next;
        if (fast == slow)
            return true;
    }
    return false;
}
// Time Complexity: O(N)
// Reason: In the worst case, all the nodes of the list are visited.
// Space Complexity: O(1)
// Reason: No extra data structure is used.

int main()
{
    node *head = NULL;
    insertNode(head, 1);
    insertNode(head, 2);
    insertNode(head, 3);
    insertNode(head, 4);
    createCycle(head, 1, 3); // creating cycle in the list
    if (cycleDetect(head) == true)
        cout << "Cycle detected\n";
    else
        cout << "Cycle not detected\n";
    return 0;
}