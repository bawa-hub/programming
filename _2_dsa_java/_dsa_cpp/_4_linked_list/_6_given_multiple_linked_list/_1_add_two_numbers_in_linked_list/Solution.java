package _2_dsa_java._dsa_cpp._4_linked_list._6_given_multiple_linked_list._1_add_two_numbers_in_linked_list;
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode dummy = new ListNode(); 
        ListNode temp = dummy; 
        int carry = 0;
        while( l1 != null || l2 != null || carry == 1) {
            int sum = 0; 
            if(l1 != null) {
                sum += l1.val; 
                l1 = l1.next; 
            }
            
            if(l2 != null) {
                sum += l2.val; 
                l2 = l2.next; 
            }
            
            sum += carry; 
            carry = sum / 10; 
            ListNode node = new ListNode(sum % 10); 
            temp.next = node; 
            temp = temp.next; 
        }
        return dummy.next;
    }
}