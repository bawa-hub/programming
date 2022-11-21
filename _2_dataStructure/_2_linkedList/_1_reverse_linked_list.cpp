// https://leetcode.com/problems/reverse-linked-list/

class Solution
{
public:
    ListNode *reverseList(ListNode *head)
    {
        ListNode *newHead = NULL;

        while (head != NULL)
        {
            ListNode *next = head->next;
            head->next = newNode;
            newNode = head;
            head = next;
        }
    }
};