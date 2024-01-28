// https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/

// brute force
class Solution {
public:
    int getDecimalValue(ListNode* head) {
        int num = 0;
        int digit = count(head);

        for(int i=digit-1;i>=0;i--) {
            num+=(pow(2, i)*head->val);
            head=head->next;
        }

        return num;
    }

    int count(ListNode* head) {
        ListNode* temp = head;
        int len=0;
        while(temp) {
            len++;
            temp=temp->next;
        }
        return len;
    }
};

// optimal
int getDecimalValue(ListNode* head) {
    int ans = 0;
    ListNode* temp = head;
    while(temp!=NULL) {
        ans*=2;
        ans+=(temp->val);
        temp=temp->next;
    }
    return ans;
}