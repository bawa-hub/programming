//  push(), pop(), top(), size(), empty(), emplace(), swap()
#include <iostream> // std::cout
#include <stack>    // std::stack
#include <vector>   // std::vector
#include <deque>    // std::deque

int main()
{

    // constructor()
    std::deque<int> mydeque(3, 100);   // deque with 3 elements
    std::vector<int> myvector(2, 200); // vector with 2 elements

    std::stack<int> first;           // empty stack
    std::stack<int> second(mydeque); // stack initialized to copy of deque

    std::stack<int, std::vector<int>> third; // empty stack using vector
    std::stack<int, std::vector<int>> fourth(myvector);

    std::cout << "size of first: " << first.size() << '\n';
    std::cout << "size of second: " << second.size() << '\n';
    std::cout << "size of third: " << third.size() << '\n';
    std::cout << "size of fourth: " << fourth.size() << '\n';

    // empty()
    std::stack<int> mystack;
    int sum(0);

    for (int i = 1; i <= 10; i++)
        mystack.push(i);

    while (!mystack.empty())
    {
        sum += mystack.top();
        mystack.pop();
    }

    std::cout << "total: " << sum << '\n';

    // size()
    std::stack<int> myints;
    std::cout << "0. size: " << myints.size() << '\n';

    for (int i = 0; i < 5; i++)
        myints.push(i);
    std::cout << "1. size: " << myints.size() << '\n';

    myints.pop();
    std::cout << "2. size: " << myints.size() << '\n';

    // top()
    std::stack<int> mystack1;

    mystack1.push(10);
    mystack1.push(20);

    mystack1.top() -= 5;

    std::cout << "mystack.top() is now " << mystack1.top() << '\n';

    // emplace()
    std::stack<std::string> mystack2;

    mystack2.emplace("First sentence");
    mystack2.emplace("Second sentence");

    std::cout << "mystack contains:\n";
    while (!mystack2.empty())
    {
        std::cout << mystack2.top() << '\n';
        mystack2.pop();
    }

    // swap()
    std::stack<int> foo, bar;
    foo.push(10);
    foo.push(20);
    foo.push(30);
    bar.push(111);
    bar.push(222);

    foo.swap(bar);

    std::cout << "size of foo: " << foo.size() << '\n';
    std::cout << "size of bar: " << bar.size() << '\n';

    return 0;
}