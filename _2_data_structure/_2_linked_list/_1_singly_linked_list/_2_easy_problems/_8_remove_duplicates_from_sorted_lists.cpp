// https://leetcode.com/problems/remove-duplicates-from-sorted-list/


// same concept as in array  (by me)
class Solution {
public:
    ListNode* deleteDuplicates(ListNode* head) {
        if(head == nullptr) return head;
        ListNode* tempi = head;
        ListNode* tempj=head;
        while(tempj!=nullptr) {
            if(tempj->val!=tempi->val) {
                tempi=tempi->next;
                tempi->val = tempj->val;
            }
            tempj=tempj->next;
        }
        tempi->next = nullptr;
        return head;
    }
};

// better
class Solution {
public:
    ListNode* deleteDuplicates(ListNode* head) {
        ListNode* temp=head;
        while (temp&&temp->next)
        {
            if (temp->next->val==temp->val)
            {
                temp->next=temp->next->next;
                continue;
            }
            temp=temp->next;
        }
        return head;
    }
};