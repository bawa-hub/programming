// push(), pop(), size(), front(), back(), empty(), emplace(), swap()
#include <iostream> // std::cout
#include <deque>    // std::deque
#include <list>     // std::list
#include <queue>    // std::queue

int main()
{
    std::deque<int> mydeck(3, 100); // deque with 3 elements
    std::list<int> mylist(2, 200);  // list with 2 elements

    std::queue<int> first;          // empty queue
    std::queue<int> second(mydeck); // queue initialized to copy of deque

    std::queue<int, std::list<int>> third; // empty queue with list as underlying container
    std::queue<int, std::list<int>> fourth(mylist);

    std::cout << "size of first: " << first.size() << '\n';
    std::cout << "size of second: " << second.size() << '\n';
    std::cout << "size of third: " << third.size() << '\n';
    std::cout << "size of fourth: " << fourth.size() << '\n';

    // empty()
    std::queue<int> myqueue;
    int sum(0);

    for (int i = 1; i <= 10; i++)
        myqueue.push(i);

    while (!myqueue.empty())
    {
        sum += myqueue.front();
        myqueue.pop();
    }

    std::cout << "total: " << sum << '\n';

    // size()
    std::queue<int> myints;
    std::cout << "0. size: " << myints.size() << '\n';

    for (int i = 0; i < 5; i++)
        myints.push(i);
    std::cout << "1. size: " << myints.size() << '\n';

    myints.pop();
    std::cout << "2. size: " << myints.size() << '\n';
    std::cout << "3. front: " << myints.front() << '\n';
    std::cout << "3. back: " << myints.back() << '\n';

    // emplace()
    std::queue<std::string> myqueue1;

    myqueue1.emplace("First sentence");
    myqueue1.emplace("Second sentence");

    std::cout << "myqueue contains:\n";
    while (!myqueue1.empty())
    {
        std::cout << myqueue1.front() << '\n';
        myqueue1.pop();
    }

    // swap()
    std::queue<int> foo, bar;
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