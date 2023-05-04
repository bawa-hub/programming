// https://practice.geeksforgeeks.org/problems/count-nodes-of-linked-list/0
// https://www.geeksforgeeks.org/find-length-of-a-linked-list-iterative-and-recursive/
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