// https://leetcode.com/problems/remove-linked-list-elements/

// by me
class Solution {
public:
    ListNode* removeElements(ListNode* head, int val) {
        if(head==nullptr) return head;
        while(head && head->val==val) {
            head = head->next;
        }
        ListNode* temp=head;
        while (temp&&temp->next)
        {
            if (temp->next->val==val)
            {
                temp->next=temp->next->next;
                continue;
            }
            temp=temp->next;
        }
        if(temp && temp->val==val) temp=temp->next;
        return head;
    }
};