#include <bits/stdc++.h>

using namespace std;

#include <bits/stdc++.h>

using namespace std;

// Solution 1: Using two Stacks where push operation is O(N)
struct Queue
{
    stack<int> input, output;

    // Push elements in queue
    void Push(int data)
    {
        // Pop out all elements from the stack input
        while (!input.empty())
        {
            output.push(input.top());
            input.pop();
        }
        // Insert the desired element in the stack input
        cout << "The element pushed is " << data << endl;
        input.push(data);
        // Pop out elements from the stack output and push them into the stack input
        while (!output.empty())
        {
            input.push(output.top());
            output.pop();
        }
    }
    // Pop the element from the Queue
    int Pop()
    {
        if (input.empty())
        {
            cout << "Stack is empty";
            exit(0);
        }
        int val = input.top();
        input.pop();
        return val;
    }
    // Return the Topmost element from the Queue
    int Top()
    {
        if (input.empty())
        {
            cout << "Stack is empty";
            exit(0);
        }
        return input.top();
    }
    // Return the size of the Queue
    int size()
    {
        return input.size();
    }
};
// Time Complexity: O(N )
// Space Complexity: O(2N)

int main()
{
    Queue q;
    q.Push(3);
    q.Push(4);
    cout << "The element poped is " << q.Pop() << endl;
    q.Push(5);
    cout << "The top of the queue is " << q.Top() << endl;
    cout << "The size of the queue is " << q.size() << endl;
}

// Solution 2: Using two Stacks where push operation is O(1)
class MyQueue
{
public:
    stack<int> input, output;
    /** Initialize your data structure here. */
    MyQueue()
    {
    }

    /** Push element x to the back of queue. */
    void push(int x)
    {
        cout << "The element pushed is " << x << endl;
        input.push(x);
    }

    /** Removes the element from in front of queue and returns that element. */
    int pop()
    {
        // shift input to output
        if (output.empty())
            while (input.size())
                output.push(input.top()), input.pop();

        int x = output.top();
        output.pop();
        return x;
    }

    /** Get the front element. */
    int top()
    {
        // shift input to output
        if (output.empty())
            while (input.size())
                output.push(input.top()), input.pop();
        return output.top();
    }

    bool empty()
    {
        return input.empty() && output.empty();
    }

    int size()
    {
        return (output.size() + input.size());
    }
};
int main()
{
    MyQueue q;
    q.push(3);
    q.push(4);
    cout << "The element poped is " << q.pop() << endl;
    q.push(5);
    cout << "The top of the queue is " << q.top() << endl;
    cout << "The size of the queue is " << q.size() << endl;
}

// Time Complexity: O(1 )
// Space Complexity: O(2N)