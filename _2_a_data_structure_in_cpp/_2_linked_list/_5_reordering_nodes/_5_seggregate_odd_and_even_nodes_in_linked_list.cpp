// https://leetcode.com/problems/odd-even-linked-list/

class Solution
{
public:
    ListNode *oddEvenList(ListNode *head)
    {
        if (!head || !head->next || !head->next->next)
            return head;

        ListNode *odd = head;
        ListNode *even_start = head->next;
        ListNode *even = head->next;

        while (odd->next && even->next)
        {
            odd->next = odd->next->next;   // link all odd nodes
            even->next = even->next->next; // link all even nodes
            odd = odd->next;
            even = even->next;
        }

        odd->next = even_start; // link the odd-even nodes together
        return head;
    }
};

// Time complexity: O(n)
// Space complexity: O(1)
