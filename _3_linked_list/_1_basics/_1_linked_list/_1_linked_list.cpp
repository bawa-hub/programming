#include <bits/stdc++.h>
using namespace std;

class Node {
    public:
    int data;
    Node *next;

    Node(int x) {
        data = x;
        next = NULL;
    }

    Node(int x, Node *next) {
        data = x;
        this->next = next;
    }
};

Node* convertArr2LL(vector<int> &arr) {
  Node* head = new Node(arr[0]);
  Node* curr = head;
  for (int i = 1; i < arr.size(); i++) {
    curr->next = new Node(arr[i]);
    curr = curr->next;
  }

  return head;
}

void traverse(Node *head) {
    Node *curr = head;
    while (curr!= nullptr) {
        cout << curr->data << " ";
        curr = curr->next;
    }
    cout << endl;
}

int lengthOfLinkedList(Node *head) {
    Node *curr = head;
    int count = 0;
    while (curr!= nullptr) {
        count++;
        curr = curr->next;
    }
    return count;
}

int searchNode(Node *head, int val) {
    Node *curr = head;
    while (curr!= nullptr) {
        if (curr->data == val) {
            return 1;
        }
        curr = curr->next;
    }
    return 0;
}

Node * insertAtStart(Node *head, int val) {
    Node *newNode = new Node(val, head);
    return newNode;
}

Node* insertAtLast(Node* head,int val) {
    Node *newNode = new Node(val);
    if(head == NULL) return newNode;
    Node *curr = head;
    while (curr->next!= nullptr) {
        curr = curr->next;
    }
    curr->next = newNode;
    return head;
}

Node* insertAtPosition(Node* head,int val,int pos) {
    if(head == nullptr) {
        if(pos == 1) {
            return new Node(val);
        } else {
            return head;
        }
    }

    if(pos==1) return new Node(val, head);

    Node *curr = head;
    int cnt = 0;
    while(curr != nullptr) {
        cnt++;
        if(cnt == pos-1) {
            Node *newNode = new Node(val, curr->next);
            curr->next = newNode;
            return head;
        }
        curr = curr->next;
    }
    return head;
}

Node* insertBeforeNode(Node* head,int data,int val) {
    if(head == nullptr) {
        return nullptr;
    }

    if(head->data == val) return new Node(data, head);

    Node *curr = head;
    while(curr->next != nullptr) {
        if(curr->next->data == val) {
            Node *newNode = new Node(data, curr->next);
            curr->next = newNode;
            return head;
        }
        curr = curr->next;
    }
    return head;
}

Node* deleteHeadNode(Node* head) {
    if (head == nullptr) {
        return nullptr;
    }
    Node *temp = head;
    head = head->next;
    delete temp;
    return head;
}

Node* deleteLastNode(Node* head) {
  if(head == nullptr || head->next == nullptr) {
    return nullptr;
  }
  Node *temp = head;
  while(temp->next->next!= nullptr) {
    temp = temp->next;
  }
  delete temp->next;
  temp->next = nullptr;

  return head;
}

Node* deleteKthNode(Node* head, int k) {
    if (head == nullptr || head->next == nullptr) {
        return nullptr;
    }
   if(k == 1) {
    Node* temp = head;
    head = head->next;
    delete temp;
    return head;
   }
   int cnt = 0;
   Node* curr = head;
   Node* prev = nullptr;
   while(curr!= nullptr) {
    cnt++;
    if(cnt == k) {
      prev->next = prev->next->next;
      delete curr;
      return head;
    }
    prev = curr;
    curr = curr->next;
   }
   return head;
}

Node* deleteNodeWithValue(Node* head, int val) {
    if (head == nullptr) {
        return nullptr;
    }
   if(head->data == val) {
    Node* curr = head;
    head = head->next;
    delete curr;
    return head;
   }

   Node* curr = head;
   Node* prev = nullptr;
   while(curr!= nullptr) {
    if(curr->data == val) {
      prev->next = prev->next->next;
      delete curr;
      break;
    }
    prev = curr;
    curr = curr->next;
   }
   return head;
}

int main() {
    vector<int> arr = {2,5,8,7,6};
    Node* head = convertArr2LL(arr);
    traverse(head);
    head = insertBeforeNode(head, 100,6);
    traverse(head);
    return 0;
}