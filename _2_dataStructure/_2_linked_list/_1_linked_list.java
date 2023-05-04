import java.io.*;
  
class node {
    int data;
    node next;
  
    node(int value)
    {
          data = value;
          next = null;
    }
}
  
class GFG {
  
    static node head = null;
  
    static void insertathead(int val)
    {
        node n = new node(val);
        n.next = head;
        head = n;
    }
	// Time Complexity: O(1)
	// Auxiliary Space: O(1)
  
    static void insertafter(int key, int val)
    {
        node n = new node(val);
        if (key == head.data) {
            n.next = head.next;
            head.next = n;
            return;
        }
        node temp = head;
        while (temp.data != key) {
            temp = temp.next;
            if (temp == null) {
                return;
            }
        }
        n.next = temp.next;
        temp.next = n;
    }
	
  
    static void insertattail(int val)
    {
        node n = new node(val);
        if (head == null) {
            head = n;
            return;
        }
        node temp = head;
        while (temp.next != null) {
            temp = temp.next;
        }
        temp.next = n;
    }
  
    static void print()
    {
        node temp = head;
        while (temp != null) {
            System.out.print(temp.data + " -> ");
            temp = temp.next;
        }
        System.out.println("NULL");
    }
  
    public static void main(String[] args)
    {
  
        insertathead(1);
        insertathead(2);
        System.out.print("After insertion at head: ");
        print();
        System.out.println();
  
        insertattail(4);
        insertattail(5);
        System.out.print("After insertion at tail: ");
        print();
        System.out.println();
  
        insertafter(1, 2);
        insertafter(5, 6);
        System.out.print(
            "After insertion at a given position: ");
        print();
        System.out.println();
    }
}