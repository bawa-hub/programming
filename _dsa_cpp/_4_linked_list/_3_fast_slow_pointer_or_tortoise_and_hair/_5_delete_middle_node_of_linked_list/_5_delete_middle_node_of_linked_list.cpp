// https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/

class Solution
{
public:
    ListNode *deleteMiddle(ListNode *head)
    {
        if (!head->next)
            return NULL;

        ListNode *slow = new ListNode(0);
        ListNode *fast = slow;

        slow->next = head;
        fast->next = head;

        while (fast->next && fast->next->next)
        {
            slow = slow->next;
            fast = fast->next->next;
        }

        slow->next = slow->next->next;
        return head;
    }
};