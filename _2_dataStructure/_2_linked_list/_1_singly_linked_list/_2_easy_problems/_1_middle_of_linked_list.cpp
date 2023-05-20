// https://leetcode.com/problems/middle-of-the-linked-list/

class Solution
{
public:
    // naive approach
    ListNode *middleNode(ListNode *head)
    {
        int n = 0;
        ListNode *temp = head;
        while (temp)
        {
            n++;
            temp = temp->next;
        }

        temp = head;

        for (int i = 0; i < n / 2; i++)
        {
            temp = temp->next;
        }

        return temp;
    }
    // Time Complexity: O(N) + O(N/2)
    // Space Complexity: O(1)


    // using fast pointer
    ListNode *middleNode(ListNode *head)
    {
        ListNode *slow = head, *fast = head;
        while (fast && fast->next)
            slow = slow->next, fast = fast->next->next;
        return slow;
    }
    // Time Complexity: O(N/2)
    // Space Complexity: O(1)
};