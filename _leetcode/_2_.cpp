// https://leetcode.com/problems/add-two-numbers/

// in case of multiple list, dummy node is very useful as it simplifies the code
// note that temp->next in while loop is changing the dummy->next as initially temp and dummy is indeed same.

class Solution
{
public:
    ListNode *addTwoNumbers(ListNode *l1, ListNode *l2)
    {
        ListNode *dummy = new ListNode();
        ListNode *temp = dummy;
        int carry = 0;
        while ((l1 != NULL || l2 != NULL) || carry)
        {
            int sum = 0;
            if (l1 != NULL)
            {
                sum += l1->val;
                l1 = l1->next;
            }

            if (l2 != NULL)
            {
                sum += l2->val;
                l2 = l2->next;
            }

            sum += carry;
            carry = sum / 10;
            ListNode *node = new ListNode(sum % 10);
            temp->next = node;
            temp = temp->next;
        }
        return dummy->next;
    }
};

// TC: O(max(m,n))
// SC: O(max(m,n))