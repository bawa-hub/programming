#include <iostream>
#include <list>
using namespace std;

// Queue implementation in C++ using `std::list`
int main()
{
    list<string> q;

    q.push_back("A"); // Insert `A` into the queue
    q.push_back("B"); // Insert `B` into the queue
    q.push_back("C"); // Insert `C` into the queue
    q.push_back("D"); // Insert `D` into the queue

    // Returns the total number of elements present in the queue
    cout << "The queue size is " << q.size() << endl;

    // Prints the front of the queue (`A`)
    cout << "The front element is " << q.front() << endl;

    // Prints the rear of the queue (`D`)
    cout << "The rear element is " << q.back() << endl;

    q.pop_front(); // removing the front element (`A`)
    q.pop_front(); // removing the next front element (`B`)

    cout << "The queue size is " << q.size() << endl;

    // check if the queue is empty
    if (q.empty())
    {
        cout << "The queue is empty\n";
    }
    else
    {
        cout << "The queue is not empty\n";
    }

    return 0;
}