class Node {
    public String data;
    public Node next;
    public Node back;
    
    // Default constructor
    public Node() {
        this.data = "";
        this.next = null;
        this.back = null;
    }
    
    // Constructor with data
    public Node(String x) {
        this.data = x;
        this.next = null;
        this.back = null;
    }
    
    // Constructor with data, next, and back pointers
    public Node(String x, Node next, Node back) {
        this.data = x;
        this.next = next;
        this.back = back;
    }
}

class BrowserHistory {
    private Node currentPage;
    
    // Constructor with the homepage URL
    public BrowserHistory(String homepage) {
        currentPage = new Node(homepage);
    }
    
    // Method to visit a new URL
    public void visit(String url) {
        Node newNode = new Node(url);
        currentPage.next = newNode;
        newNode.back = currentPage;
        currentPage = newNode;
    }
    
    // Method to go back a certain number of steps
    public String back(int steps) {
        while (steps > 0) {
            if (currentPage.back != null) {
                currentPage = currentPage.back;
            } else {
                break;
            }
            steps--;
        }
        return currentPage.data;
    }
    
    // Method to go forward a certain number of steps
    public String forward(int steps) {
        while (steps > 0) {
            if (currentPage.next != null) {
                currentPage = currentPage.next;
            } else {
                break;
            }
            steps--;
        }
        return currentPage.data;
    }
}
