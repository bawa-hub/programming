#include <iostream>

using namespace std;
struct node
{ // single node
    int data;
    node *next;
    node *previous;
    node(int d)
    { // constructor
        data = d;
        previous = NULL;
        next = NULL;
    }
};
node *delete_last(node *head)
{
    node *current = head;
    while (current->next != NULL)
    { // iterating to the last node
        current = current->next;
    }
    node *temp = current;           // storing the last node
    current->previous->next = NULL; // changing the next pointer of
    neighboring node delete temp;   // deleting the last node
    return head;
}
// Time Complexity: O(size of the DLL)
// Space Complexity: O(1)

void print_dll(node *head)
{
    node *current = head;
    while (current != NULL)
    {
        cout << current->data << " ";
        current = current->next;
    }
    cout << endl;
}
int main()
{
    // creating nodes for the DLL
    node *head = new node(10);
    node *node1 = new node(20);
    node *node2 = new node(30);

    // 	constructing the DLL
    head->next = node1;
    node1->previous = head;
    node1->next = node2;
    node2->previous = node1;
    cout << "Original Doubly Linked List" << endl;
    print_dll(head);
    head = delete_last(head);
    cout << "After deleting the Last node" << endl;
    print_dll(head); // utility function
    return 0;
}