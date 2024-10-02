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
}

public class LinkedList {

    private static Node convertArr2LL(int[] arr) {
        Node head = new Node(arr[0]);
        Node curr = head;
        for (int i = 1; i < arr.length; i++) {
          curr.next = new Node(arr[i]);
          curr = curr.next;
        }
      
        return head;
    }

    private static void traverse(Node head) {
        Node curr = head;
        while (curr!= null) {
            System.out.println(curr.data + " ");
            curr = curr.next;
        }
    }

    private static int lengthOfLinkedList(Node head) {
        Node curr = head;
        int count = 0;
        while (curr!= null) {
            count++;
            curr = curr.next;
        }
        return count;
    }

    private static int searchNode(Node head, int val) {
        Node curr = head;
        while (curr!= null) {
            if (curr.data == val) {
                return 1;
            }
            curr = curr.next;
        }
        return 0;
    }

    private static Node deleteHeadNode(Node head) {
        if (head == null) return head;
        head = head.next;
        return head;
    }

    private static Node deleteLastNode(Node head) {
        if(head == null || head.next == null) return null;
        Node temp = head;
        while(temp.next.next != null) {
          temp = temp.next;
        }

        temp.next = null;
        return head;
    }

    private static Node deleteKthNode(Node head, int k) {
        if (head == null || head.next == null) return null;
       if(k == 1) {
        head = head.next;
        return head;
       }
       int cnt = 0;
       Node curr = head;
       Node prev = null;
       while(curr!= null) {
        cnt++;
        if(cnt == k) {
          prev.next = prev.next.next;
          return head;
        }
        prev = curr;
        curr = curr.next;
       }
       return head;
    }

    private static Node deleteNodeWithValue(Node head, int val) {
        if (head == null)
            return null;
        if (head.data == val) {
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

    private static Node insertAtStart(Node head, int val) {
        Node newNode = new Node(val, head);
        return newNode;
    }

    private static Node insertAtLast(Node head,int val) {
        Node newNode = new Node(val);
        if(head == null) return newNode;
        Node curr = head;
        while (curr.next!= null) {
            curr = curr.next;
        }
        curr.next = newNode;
        return head;
    }

    private static Node insertAtPosition(Node head,int val,int pos) {
        if(head == null) {
            if(pos == 1) return new Node(val); 
            else return head;
        }
    
        if(pos==1) return new Node(val, head);
    
        Node curr = head;
        int cnt = 0;
        while(curr != null) {
            cnt++;
            if(cnt == pos-1) {
                Node newNode = new Node(val, curr.next);
                curr.next = newNode;
                return head;
            }
            curr = curr.next;
        }
        return head;
    }
    
    private static Node insertBeforeNode(Node head, int data, int val) {
        if (head == null)
            return null;
        if (head.data == val)
            return new Node(data, head);

        Node curr = head;
        while (curr.next != null) {
            if (curr.next.data == val) {
                Node newNode = new Node(data, curr.next);
                curr.next = newNode;
                return head;
            }
            curr = curr.next;
        }
        return head;
    }

    public static void main(String[] args) {
        int[] arr = {3, 5, 8, 7};
        Node head = convertArr2LL(arr);

        // System.out.println(lengthOfLinkedList(head));
        // System.out.println(searchNode(head, 8));

        // head = deleteHeadNode(head);
        // head = deleteLastNode(head);
        // head = deleteKthNode(head, 5);
        // head = deleteNodeWithValue(head, 3);

        // head = insertAtStart(head, 1);
        // head = insertAtLast(head, 9);
        // head = insertAtPosition(head, 100,3);
        // head = insertBeforeNode(head, 100, 5);

        traverse(head);
    }
}
