// https://leetcode.com/problems/palindrome-linked-list/

#include <bits/stdc++.h>
using namespace std;

// using extra ds
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

void insertNode(node *head, int val)
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

bool isPalindrome(node *head)
{
    vector<int> arr;
    while (head != NULL)
    {
        arr.push_back(head->num);
        head = head->next;
    }
    for (int i = 0; i < arr.size() / 2; i++)
        if (arr[i] != arr[arr.size() - i - 1])
            return false;
    return true;
}

int main()
{
    node *head = NULL;
    insertNode(head, 1);
    insertNode(head, 2);
    insertNode(head, 3);
    insertNode(head, 2);
    insertNode(head, 1);
    isPalindrome(head) ? cout << "True" : cout << "False";
    return 0;
}

// Time Complexity: O(N)
// Reason: Iterating through the list to store elements in the array.
// Space Complexity: O(N)
// Reason: Using an array to store list elements for further computations.

// optimized
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

void insertNode(node *head, int val)
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

node *reverse(node *ptr)
{
    node *pre = NULL;
    node *nex = NULL;
    while (ptr != NULL)
    {
        nex = ptr->next;
        ptr->next = pre;
        pre = ptr;
        ptr = nex;
    }
    return pre;
}

bool isPalindrome(node *head)
{
    if (head == NULL || head->next == NULL)
        return true;

    node *slow = head;
    node *fast = head;

    while (fast->next != NULL && fast->next->next != NULL)
    {
        slow = slow->next;
        fast = fast->next->next;
    }

    slow->next = reverse(slow->next);
    slow = slow->next;
    node *dummy = head;

    while (slow != NULL)
    {
        if (dummy->num != slow->num)
            return false;
        dummy = dummy->next;
        slow = slow->next;
    }
    return true;
}

int main()
{
    node *head = NULL;
    insertNode(head, 1);
    insertNode(head, 2);
    insertNode(head, 3);
    insertNode(head, 2);
    insertNode(head, 1);
    isPalindrome(head) ? cout << "True" : cout << "False";
    return 0;
}

// Time Complexity: O(N/2)+O(N/2)+O(N/2)
// Reason: O(N/2) for finding the middle element, reversing the list from the middle element, and traversing again to find palindrome respectively.
// Space Complexity: O(1)
// Reason: No extra data structures are used.