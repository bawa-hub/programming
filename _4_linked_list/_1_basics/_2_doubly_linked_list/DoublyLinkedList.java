public class DoublyLinkedList {
    public static class Node {
        public int data;       
        public Node next;      
        public Node back;     
    
        public Node(int data, Node next, Node back) {
            this.data = data;
            this.next = next;
            this.back = back;
        }
    
        public Node(int data) {
            this.data = data;
            this.next = null;
            this.back = null;
        }
    }

    private static Node convertArr2DLL(int[] arr) {
        Node head = new Node(arr[0]);
        Node prev = head;

        for (int i = 1; i < arr.length; i++) {
            Node temp = new Node(arr[i], null, prev);
            prev.next = temp;
            prev = temp;
        }
        return head;
    }

    private static void traverse(Node head) {
        while (head != null) {
            System.out.print(head.data + " ");
            head = head.next;
        }
        System.out.println();
    }

    private static Node deleteHead(Node head) {
        if (head == null || head.next == null) return null;
    
        Node prev = head;      
        head = head.next;    
        head.back = null;   
        prev.next = null;  
    
        return head;          
    }

    private static Node deleteTail(Node head) {
        if (head == null || head.next == null) return null;
        
        Node tail = head;
        while (tail.next != null) {
            tail = tail.next; 
        }
        
        Node newTail = tail.back;
        newTail.next = null;
        tail.back = null;
        
        return head;
    }

    private static Node deleteKthElement(Node head, int k){
        if(head==null) return null;

        int count = 0;
        Node kNode = head;
        while(kNode!=null){
            count++;
            if(count==k) break;
            kNode = kNode.next;
        }
        Node prev = kNode.back;
        Node front = kNode.next;
        
        if(prev==null && front == null) return null;
        else if (prev==null) return deleteHead(head);
        else if(front == null) return deleteTail(head);
        
        prev.next = front;
        front.back = prev;
        
        kNode.next = null;
        kNode.back = null;
        
        return head;
    }

    private static void deleteGivenNode(Node temp){
        Node prev = temp.back;
        Node front = temp.next;
        
        if(front==null){
            prev.next = null;
            temp.back = null;
            return;
        }
        
        prev.next = front;
        front.back = prev;
        
        temp.next = null;
        temp.back = null;
        
        return;
    }

    private static Node insertBeforeHead(Node head, int val){
        Node newHead = new Node(val , head, null);
        head.back = newHead;
        return newHead;
    }

    private static Node insertBeforeTail(Node head, int val){
        if(head.next==null) return insertBeforeHead(head, val);
        
        Node tail = head;
        while(tail.next!=null){
            tail = tail.next;
        }
        Node prev = tail.back;
        Node newNode = new Node(val, tail, prev);
        
        prev.next = newNode;
        tail.back = newNode;
        
        return head;
    }

    private static Node insertBeforeKthElement(Node head, int k, int val){
       if(k==1) return insertBeforeHead(head, val);
        
        Node temp = head;
        
        int count = 0;
        while(temp!=null){
            count ++;
            if(count == k) break;
            temp = temp.next;
        }
        Node prev = temp.back;
        
        Node newNode = new Node(val, temp, prev);
        
        prev.next = newNode;
        temp.back = newNode;
        
        newNode.next = temp;
        newNode.back = prev;
        
        return head;
    }

    private static void insertBeforeNode(Node node, int val){
        Node prev = node.back;
        
        Node newNode = new Node(val, node, prev);
        
        prev.next = newNode;
        node.back = newNode;
        
        return;
    }

    private static Node insertAtTail(Node head, int k) {
        Node newNode = new Node(k);
    
        if (head == null)  return newNode;
        Node tail = head;
        while (tail.next != null) {
            tail = tail.next;
        }
    
        tail.next = newNode;
        newNode.back = tail;
    
        return head;
    }

    public static void main(String[] args) {
        int[] arr = {12, 5, 6, 8, 4};
        Node head = convertArr2DLL(arr);

        // head = deleteHead(head);
        // head = deleteTail(head);
        // head = deleteKthElement(head, 2);
        // deleteGivenNode(head.next);
        

        traverse(head);
    }
}
