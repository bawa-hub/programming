// https://practice.geeksforgeeks.org/problems/search-in-linked-list-1664434326/1

struct Node
{
    int data;
    Node *next;
    Node(int x)
    {
        data = x;
        next = NULL;
    }
};

class Solution
{
public:
    // Function to count nodes of a linked list.
    bool searchKey(int n, struct Node *head, int key)
    {
        // Code here
        struct Node *curr = head;
        while (curr != nullptr)
        {
            if (curr->data == key)
                return true;
            else
                curr = curr->next;
        }
        return false;
    }
};