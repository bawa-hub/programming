class Node {
    int data;
    Node next;

    Node(int data1, Node next1) {
        this.data = data1;
        this.next = next1;
    }

    Node(int data1) {
        this.data = data1;
        this.next = null;
    }
};

public class LinkedList {
    // Function to calculate the length of a linked list
    private static int lengthofaLL(Node head) {
        int cnt = 0;
        Node temp = head;
        // Traverse the linked list and count nodes
        while (temp != null) {
            temp = temp.next;
            cnt++;// increment cnt for every node traversed
        }
        return cnt;
    }

    // Function to check if a given element is present in the linked list
    public static int checkifPresent(Node head, int desiredElement) {
        Node temp = head;

        // Traverse the linked list
        while (temp != null) {
            // Check if the current node's data is equal to the desired element
            if (temp.data == desiredElement)
                return 1; // Return 1 if the element is found

            // Move to the next node
            temp = temp.next;
        }

        return 0; // Return 0 if the element is not found in the linked list
    }

    private static Node convertarr2LL(int[] arr) {
        Node head = new Node(arr[0]);
        Node mover = head;
        for (int i = 1; i < arr.length; i++) {
            Node temp = new Node(arr[i]);
            mover.next = temp;
            mover = mover.next;
        }
        return head;
    }

    // Method to print the linked list
    private static void printLL(Node head) {
        while (head != null) {
            System.out.print(head.data + " ");
            head = head.next;
        }
    }

    // Method to remove the head of the linked list
    private static Node removesHead(Node head) {
        // Check if the linked list is empty
        if (head == null)
            return null;
        // Move the head to the next node
        head = head.next;
        // Return the updated head of the linked list
        return head;
        // There’s no need to delete the earlier head since it gets automatically
        // deleted.
    }

    // Function to delete the tail of the linked list
    private static Node deleteTail(Node head) {
        // Check if the linked list is empty or has only one node
        if (head == null || head.next == null)
            return null;
        // Create a temporary pointer for traversal
        Node temp = head;
        // Traverse the list until the second-to-last node
        while (temp.next.next != null) {
            temp = temp.next;
        }
        // Nullify the connection from the second-to-last node to delete the last node
        temp.next = null;
        // Return the updated head of the linked list
        return head;
    }

    // Function to delete the k-th node in a linked list
    private static Node deleteKthNode(Node head, int k) {
        // Check if the list is empty
        if (head == null)
            return head;

        // If k is 1, delete the first node
        if (k == 1) {
            Node temp = head;
            head = head.next;
            temp = null;
            return head;
        }

        // Traverse the list to find the k-th node
        Node temp = head;
        Node prev = null;
        int cnt = 0;

        while (temp != null) {
            cnt++;
            // If the k-th node is found
            if (cnt == k) {
                // Adjust the pointers to skip the k-th node
                prev.next = prev.next.next;
                // Delete the k-th node
                temp = null;
                break;
            }
            // Move to the next node
            prev = temp;
            temp = temp.next;
        }

        return head;
    }

    private static Node deleteNodeWithValue(Node head, int val) {
        if (head == null) {
            return null;
        }
        if (head.data == val) {
            Node curr = head;
            head = head.next;
            return head;
        }

        Node curr = head;
        Node prev = null;
        while (curr != null) {
            if (curr.data == val) {
                prev.next = prev.next.next;
                break;
            }
            prev = curr;
            curr = curr.next;
        }
        return head;
    }

    // Function to insert a new node at the head of the linked list
    public static Node insertHead(Node head, int val) {
        Node temp = new Node(val, head);
        return temp;
    }

    // Function to insert a new node at the tail of the linked list
    public static Node insertTail(Node head, int val) {
        if (head == null)
            return new Node(val);

        Node temp = head;
        while (temp.next != null) {
            temp = temp.next;
        }

        Node newNode = new Node(val);
        temp.next = newNode;

        return head;
    }

    // Function to insert a new node at position k in the linked list
    public static Node insertAtK(Node head, int val, int k) {
        // If the linked list is empty and k is 1, insert the new node as the head
        if (head == null) {
            if (k == 1)
                return new Node(val);
            else
                return head;
        }

        // If k is 1, insert the new node at the beginning of the linked list
        if (k == 1)
            return new Node(val, head);

        int cnt = 0;
        Node temp = head;

        // Traverse the linked list to find the node at position k-1
        while (temp != null) {
            cnt++;
            if (cnt == k - 1) {
                // Insert the new node after the node at position k-1
                Node newNode = new Node(val, temp.next);
                temp.next = newNode;
                break;
            }
            temp = temp.next;
        }

        return head;
    }

    // Function to insert a new node with data 'el' after the node with data 'val'
    public static Node insertAtValueK(Node head, int el, int val) {
        if (head == null) {
            return null;
        }

        // Insert at the beginning if the value matches the head's data
        if (head.data == val) {
            return new Node(el, head);
        }

        Node temp = head;
        while (temp.next != null) {
            // Insert at the current position if the next node has the desired value
            if (temp.next.data == val) {
                Node newNode = new Node(el, temp.next);
                temp.next = newNode;
                break;
            }
            temp = temp.next;
        }
        return head;
    }

    public static void main(String[] args) {
        int[] arr = { 2, 5, 8, 7 };
        Node head = convertarr2LL(arr);
        System.out.println(head.data);
    }
}