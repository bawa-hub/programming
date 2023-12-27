// https://leetcode.com/problems/reverse-linked-list/
class Solution
{
public:
    // using iteration
    ListNode *reverseList(ListNode *head)
    {
        ListNode *newHead = NULL;
        while (head != NULL)
        {
            ListNode *next = head->next;
            head->next = newHead;
            newHead = head;
            head = next;
        }
        return newHead;
    }
    // Time Complexity:
    // Since we are iterating only once through the list and achieving reversed list. Thus, the time complexity is O(N) where N is the number of nodes present in the list.
    // Space Complexity:
    // To perform given tasks, no external spaces are used except three-pointers. So, space complexity is O(1).

    // using recursion
    ListNode *reverseList(ListNode *&head)
    {

        if (head == NULL || head->next == NULL)
            return head;

        ListNode *nnode = reverseList(head->next);
        head->next->next = head;
        head->next = NULL;
        return nnode;
    }
};
// Time Complexity:
// We move to the end of the list and achieve our reversed list. Thus, the time complexity is O(N) where N represents the number of nodes.
// Space Complexity:
// Apart from recursion using stack space, no external storage is used. Thus, space complexity is O(1)