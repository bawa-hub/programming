// https://practice.geeksforgeeks.org/problems/count-nodes-of-linked-list/0

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
    int getCount(struct Node *head)
    {
        if (head == NULL)
            return 0;
        int len = 1;
        while (head->next != NULL)
        {
            len++;
            head = head->next;
        }
        return len;
    }
};