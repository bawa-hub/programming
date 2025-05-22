// https://leetcode.com/problems/intersection-of-two-linked-lists/

#include <iostream>
using namespace std;

// brute force
class node
{
public:
    int num;
    node *next;
    node(int val)
    {
        num = val;
        next = NULL;
    }
};

// // utility function to insert node at the end of the linked list
void insertNode(node *&head, int val)
{
    node *newNode = new node(val);

    if (head == NULL)
    {
        head = newNode;
        return;
    }

    node *temp = head;
    while (temp->next != NULL)
        temp = temp->next;

    temp->next = newNode;
    return;
}

// // utility function to check presence of intersection
node *intersectionPresent(node *head1, node *head2)
{
    while (head2 != NULL)
    {
        node *temp = head1;
        while (temp != NULL)
        {
            // if both nodes are same
            if (temp == head2)
                return head2;
            temp = temp->next;
        }
        head2 = head2->next;
    }
    // intersection is not present between the lists return null
    return NULL;
}

// // utility function to print linked list created
void printList(node *head)
{
    while (head->next != NULL)
    {
        cout << head->num << "->";
        head = head->next;
    }
    cout << head->num << endl;
}

int main()
{
    // creation of both lists
    node *head = NULL;
    insertNode(head, 1);
    insertNode(head, 3);
    insertNode(head, 1);
    insertNode(head, 2);
    insertNode(head, 4);
    node *head1 = head;
    head = head->next->next->next;
    node *headSec = NULL;
    insertNode(headSec, 3);
    node *head2 = headSec;
    headSec->next = head;
    // printing of the lists
    cout << "List1: ";
    printList(head1);
    cout << "List2: ";
    printList(head2);
    // checking if intersection is present
    node *answerNode = intersectionPresent(head1, head2);
    if (answerNode == NULL)
        cout << "No intersection\n";
    else
        cout << "The intersection point is " << answerNode->num << endl;
    return 0;
}

// Time Complexity: O(m*n)
// Reason: For each node in list 2 entire lists 1 are iterated.
// Space Complexity: O(1)
// Reason: No extra space is used.

// hashing
class node
{
public:
    int num;
    node *next;
    node(int val)
    {
        num = val;
        next = NULL;
    }
};
// // utility function to insert node at the end of the linked list
void insertNode(node *&head, int val)
{
    node *newNode = new node(val);

    if (head == NULL)
    {
        head = newNode;
        return;
    }

    node *temp = head;
    while (temp->next != NULL)
        temp = temp->next;

    temp->next = newNode;
    return;
}

// // utility function to check presence of intersection
node *intersectionPresent(node *head1, node *head2)
{
    unordered_set<node *> st;
    while (head1 != NULL)
    {
        st.insert(head1);
        head1 = head1->next;
    }
    while (head2 != NULL)
    {
        if (st.find(head2) != st.end())
            return head2;
        head2 = head2->next;
    }
    return NULL;
}

// // utility function to print linked list created
void printList(node *head)
{
    while (head->next != NULL)
    {
        cout << head->num << "->";
        head = head->next;
    }
    cout << head->num << endl;
}

int main()
{
    // creation of both lists
    node *head = NULL;
    insertNode(head, 1);
    insertNode(head, 3);
    insertNode(head, 1);
    insertNode(head, 2);
    insertNode(head, 4);
    node *head1 = head;
    head = head->next->next->next;
    node *headSec = NULL;
    insertNode(headSec, 3);
    node *head2 = headSec;
    headSec->next = head;
    // printing of the lists
    cout << "List1: ";
    printList(head1);
    cout << "List2: ";
    printList(head2);
    // checking if intersection is present
    node *answerNode = intersectionPresent(head1, head2);
    if (answerNode == NULL)
        cout << "No intersection\n";
    else
        cout << "The intersection point is " << answerNode->num << endl;
    return 0;
}

// Time Complexity: O(n+m)
// Reason: Iterating through list 1 first takes O(n), then iterating through list 2 takes O(m).
// Space Complexity: O(n)
// Reason: Storing list 1 node addresses in unordered_set.

// difference of length
class node
{
public:
    int num;
    node *next;
    node(int val)
    {
        num = val;
        next = NULL;
    }
};

// // utility function to insert node at the end of the linked list
void insertNode(node *&head, int val)
{
    node *newNode = new node(val);

    if (head == NULL)
    {
        head = newNode;
        return;
    }

    node *temp = head;
    while (temp->next != NULL)
        temp = temp->next;

    temp->next = newNode;
    return;
}


// brute force 
node* intersectionPresent(node* head1,node* head2) {
    while(head2 != NULL) {
        node* temp = head1;
        while(temp != NULL) {
            //if both nodes are same
            if(temp == head2) return head2;
            temp = temp->next;
        }
        head2 = head2->next;
    }
    //intersection is not present between the lists return null
    return NULL;
}
// Time Complexity: O(m*n)
// Reason: For each node in list 2 entire lists 1 are iterated. 
// Space Complexity: O(1)
// Reason: No extra space is used.

// using hashing
node* intersectionPresent(node* head1,node* head2) {
     unordered_set<node*> st;
    while(head1 != NULL) {
       st.insert(head1);
       head1 = head1->next;
    }
    while(head2 != NULL) {
        if(st.find(head2) != st.end()) return head2;
        head2 = head2->next;
    }
    return NULL;

}
// Time Complexity: O(n+m)
// Reason: Iterating through list 1 first takes O(n), then iterating through list 2 takes O(m). 
// Space Complexity: O(n)
// Reason: Storing list 1 node addresses in unordered_set.

// using difference in length
int getDifference(node *head1, node *head2)
{
    int len1 = 0, len2 = 0;
    while (head1 != NULL || head2 != NULL)
    {
        if (head1 != NULL)
        {
            ++len1;
            head1 = head1->next;
        }
        if (head2 != NULL)
        {
            ++len2;
            head2 = head2->next;
        }
    }
    return len1 - len2; // if difference is neg-> length of list2 > length of list1 else vice-versa
}

node *intersectionPresent(node *head1, node *head2)
{
    int diff = getDifference(head1, head2);
    if (diff < 0)
        while (diff++ != 0)
            head2 = head2->next;
    else
        while (diff-- != 0)
            head1 = head1->next;
    while (head1 != NULL)
    {
        if (head1 == head2)
            return head1;
        head2 = head2->next;
        head1 = head1->next;
    }
    return head1;
}
// Time Complexity:
// O(2max(length of list1,length of list2))+O(abs(length of list1-length of list2))+O(min(length of list1,length of list2))
// Reason: Finding the length of both lists takes max(length of list1, length of list2) because it is found simultaneously for both of them. Moving the head pointer ahead by a difference of them. The next one is for searching.
// Space Complexity: O(1)
// Reason: No extra space is used.

// optimized
node *intersectionPresent(node *head1, node *head2)
{
    node *d1 = head1;
    node *d2 = head2;

    while (d1 != d2)
    {
        d1 = d1 == NULL ? head2 : d1->next;
        d2 = d2 == NULL ? head1 : d2->next;
    }

    return d1;
}
// Time Complexity: O(2*max(length of list1,length of list2))
// Reason: Uses the same concept of the difference of lengths of two lists.
// Space Complexity: O(1)
// Reason: No extra data structure is used

// utility function to print linked list created
void printList(node *head)
{
    while (head->next != NULL)
    {
        cout << head->num << "->";
        head = head->next;
    }
    cout << head->num << endl;
}

int main()
{
    // creation of both lists
    node *head = NULL;
    insertNode(head, 1);
    insertNode(head, 3);
    insertNode(head, 1);
    insertNode(head, 2);
    insertNode(head, 4);
    node *head1 = head;
    head = head->next->next->next;
    node *headSec = NULL;
    insertNode(headSec, 3);
    node *head2 = headSec;
    headSec->next = head;
    // printing of the lists
    cout << "List1: ";
    printList(head1);
    cout << "List2: ";
    printList(head2);
    // checking if intersection is present
    node *answerNode = intersectionPresent(head1, head2);
    if (answerNode == NULL)
        cout << "No intersection\n";
    else
        cout << "The intersection point is " << answerNode->num << endl;
    return 0;
}