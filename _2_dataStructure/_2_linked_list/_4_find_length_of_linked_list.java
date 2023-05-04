// https://www.geeksforgeeks.org/find-length-of-a-linked-list-iterative-and-recursive/

class Node {
    int data;
    Node next;
    Node(int d)
    {
        data = d;
        next = null;
    }
}
 
// Linked List class
class LinkedList {
    Node head; // head of list
 
    /* Inserts a new Node at front of the list. */
    public void push(int new_data)
    {
        /* 1 & 2: Allocate the Node &
                  Put in the data*/
        Node new_node = new Node(new_data);
 
        /* 3. Make next of new Node as head */
        new_node.next = head;
 
        /* 4. Move the head to point to new Node */
        head = new_node;
    }
 
    // iterative
    public int getCount()
    {
        Node temp = head;
        int count = 0;
        while (temp != null) {
            count++;
            temp = temp.next;
        }
        return count;
    }
    // Time complexity: O(N), Where N is the size of the linked list
    // Auxiliary Space: O(1), As constant extra space is used
 
    // recursive
    public int getCount() { return getCountRec(head); }

    public int getCountRec(Node node)
    {
        // Base case
        if (node == null)
          return 0;
 
        // Count is this node plus rest of the list
        return 1 + getCountRec(node.next);
    }
    // Time Complexity: O(N), As we are traversing the linked list only once.
    // Auxiliary Space: O(N), Extra space is used in the recursion call stack. 

    // recursive
    int getCount(Node* head, int count = 0)
{
    // base case
    if (head == NULL)
        return count;
    // move the pointer to next node and increase the count
    return getCount(head->next, count + 1);
}
// Time Complexity: O(N), As we are traversing the list only once.
// Auxiliary Space: O(1), As we are using the tail recursive function, no extra space is used in the function call stack.
 
    // Driver code
    public static void main(String[] args)
    {
        /* Start with the empty list */
        LinkedList llist = new LinkedList();
        llist.push(1);
        llist.push(3);
        llist.push(1);
        llist.push(2);
        llist.push(1);
 
          // Function call
        System.out.println("Count of nodes is "
                           + llist.getCount());
    }
}