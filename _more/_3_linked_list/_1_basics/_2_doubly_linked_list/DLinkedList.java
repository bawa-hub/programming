public class DLinkedList {
    public static class Node {
        public int data; // Data stored in the node
        public Node next; // Reference to the next node in the list (forward direction)
        public Node back; // Reference to the previous node in the list (backward direction)

        // Constructor for a Node with both data, a reference to the next node, and a
        // reference to the previous node
        public Node(int data1, Node next1, Node back1) {
            data = data1;
            next = next1;
            back = back1;
        }

        // Constructor for a Node with data, and no references to the next and previous
        // nodes (end of the list)
        public Node(int data1) {
            data = data1;
            next = null;
            back = null;
        }
    }

    // Function to convert an array to a doubly linked list
    private static Node convertArr2DLL(int[] arr) {
        Node head = new Node(arr[0]); // Create the head node with the first element of the array
        Node prev = head; // Initialize 'prev' to the head node

        for (int i = 1; i < arr.length; i++) {
            // Create a new node with data from the array and set its 'back' pointer to the
            // previous node
            Node temp = new Node(arr[i], null, prev);
            prev.next = temp; // Update the 'next' pointer of the previous node to point to the new node
            prev = temp; // Move 'prev' to the newly created node for the next iteration
        }
        return head; // Return the head of the doubly linked list
    }

    // Function to delete the head of the doubly linked list
    private static Node deleteHead(Node head) {
        if (head == null || head.next == null) {
            return null; // Return null if the list is empty or contains only one element
        }

        Node prev = head;
        head = head.next;

        head.back = null; // Set 'back' pointer of the new head to null
        prev.next = null; // Set 'next' pointer of 'prev' to null

        return head;
    }

    // Function to delete the tail of the doubly linked list
    private static Node deleteTail(Node head) {
        if (head == null || head.next == null) {
            return null; // Return null if the list is empty or contains only one element
        }

        Node tail = head;
        while (tail.next != null) {
            tail = tail.next;
        }

        Node newtail = tail.back;

        newtail.next = null;
        tail.back = null;

        return head;
    }

    // Function to remove the Kth element from the doubly linked list
    private static Node removeKthElement(Node head, int k) {
        if (head == null) {
            return null;
        }
        int count = 0;
        Node kNode = head;
        while (kNode != null) {
            count++;
            if (count == k) {
                break;
            }
            kNode = kNode.next;
        }
        Node prev = kNode.back;
        Node front = kNode.next;

        if (prev == null && front == null) {
            return null;
        } else if (prev == null) {
            return deleteHead(head);
        } else if (front == null) {
            return deleteTail(head);
        }

        prev.next = front;
        front.back = prev;

        kNode.next = null;
        kNode.back = null;

        return head;
    }

    private static void deleteNode(Node temp) {
        Node prev = temp.back;
        Node front = temp.next;

        // edge case if temp is the tail node
        if (front == null) {
            prev.next = null;
            temp.back = null;
            return;
        }

        // Disconnect temp from the doubly linked list
        prev.next = front;
        front.back = prev;

        // Set temp's pointers to null
        temp.next = null;
        temp.back = null;

        // Free memory of the deleted node
        return;
    }

    private static Node insertBeforeHead(Node head, int val) {
        // Create new node with data as val
        Node newHead = new Node(val, head, null);
        // Set old head's back pointer to the new Head
        head.back = newHead;
        // Return the new head
        return newHead;
    }

    private static Node insertBeforeTail(Node head, int val) {
        // Edge case, if dll has only one elements
        if (head.next == null) {
            // If only one element
            return insertBeforeHead(head, val);
        }

        // Create pointer tail to traverse to the end of list
        Node tail = head;
        while (tail.next != null) {
            tail = tail.next;
        }
        // Keep track of node before tail using prev
        Node prev = tail.back;

        // Create new node with value val
        Node newNode = new Node(val, tail, prev);

        // Join the new node into the doubly linked list
        prev.next = newNode;
        tail.back = newNode;

        // Return the updated linked list
        return head;
    }

    // Function to insert before the Kth node
    private static Node insertBeforeKthElement(Node head, int k, int val) {

        if (k == 1) {
            // K = 1 means node has to be insert before the head
            return insertBeforeHead(head, val);
        }

        // temp will keep track of the node
        Node temp = head;

        // count so that the Kth element can be reached
        int count = 0;
        while (temp != null) {
            count++;
            // On reaching the Kth position, break out of the loop
            if (count == k)
                break;
            // keep moving temp forward till count != K
            temp = temp.next;
        }
        // track the node before the Kth node
        Node prev = temp.back;

        // create new node with data as val
        Node newNode = new Node(val, temp, prev);

        // join the new node in between prev and temp
        prev.next = newNode;
        temp.back = newNode;

        // Set newNode's pointers to prev and temp
        newNode.next = temp;
        newNode.back = prev;

        // Return the head of the updated doubly linked list
        return head;
    }

    // Function to insert a new node before a given node
    private static void insertBeforeNode(Node node, int val) {
        // Get the node before the given node
        Node prev = node.back;

        // Create a new node with the given val
        Node newNode = new Node(val, node, prev);

        // Connect the newNode to the doubly linked list
        prev.next = newNode;
        node.back = newNode;

    }

    // Function to insert a new node with value 'k' at the end of the doubly linked
    // list
    private static Node insertAtTail(Node head, int k) {
        // Create a new node with data 'k'
        Node newNode = new Node(k);

        // If the doubly linked list is empty, set 'head' to the new node
        if (head == null) {
            return newNode;
        }

        // Traverse to the end of the doubly linked list
        Node current = head;
        while (current.next != null) {
            current = current.next;
        }

        // Connect the new node to the last node in the list
        current.next = newNode;
        newNode.back = current;

        return head;
    }

    // Function to print the elements of the doubly linked list
    private static void print(Node head) {
        while (head != null) {
            System.out.print(head.data + " "); // Print the data in the current node
            head = head.next; // Move to the next node
        }
        System.out.println();
    }

    public static void main(String[] args) {
        int[] arr = { 12, 5, 6, 8, 4 };
        Node head = convertArr2DLL(arr); // Convert the array to a doubly linked list
        print(head); // Print the doubly linked list

        System.out.println("Doubly Linked List after deleting the head: ");
        head = deleteTail(head);
        print(head);
    }
}