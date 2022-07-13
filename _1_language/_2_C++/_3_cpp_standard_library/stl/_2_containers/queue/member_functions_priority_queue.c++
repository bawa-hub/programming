// constructing priority queues
#include <iostream>   // std::cout
#include <queue>      // std::priority_queue
#include <vector>     // std::vector
#include <functional> // std::greater

class mycomparison
{
    bool reverse;

public:
    mycomparison(const bool &revparam = false)
    {
        reverse = revparam;
    }
    bool operator()(const int &lhs, const int &rhs) const
    {
        if (reverse)
            return (lhs > rhs);
        else
            return (lhs < rhs);
    }
};

int main()
{
    // constructor()
    int myints[] = {10, 60, 50, 20};

    std::priority_queue<int> first;
    std::priority_queue<int> second(myints, myints + 4);
    std::priority_queue<int, std::vector<int>, std::greater<int>>
        third(myints, myints + 4);
    // using mycomparison:
    typedef std::priority_queue<int, std::vector<int>, mycomparison> mypq_type;

    mypq_type fourth;                    // less-than comparison
    mypq_type fifth(mycomparison(true)); // greater-than comparison

    // empty(), push(), top(), pop()
    std::priority_queue<int> mypq;
    int sum(0);

    for (int i = 1; i <= 10; i++)
        mypq.push(i);

    std::cout << "size: " << mypq.size() << '\n';

    while (!mypq.empty())
    {
        sum += mypq.top();
        mypq.pop();
    }

    std::cout << "total: " << sum << '\n';

    // emplace()
    std::priority_queue<std::string> mypq1;

    mypq1.emplace("orange");
    mypq1.emplace("strawberry");
    mypq1.emplace("apple");
    mypq1.emplace("pear");

    std::cout << "mypq contains:";
    while (!mypq1.empty())
    {
        std::cout << ' ' << mypq1.top();
        mypq1.pop();
    }
    std::cout << '\n';

    // swap()
    std::priority_queue<int> foo, bar;
    foo.push(15);
    foo.push(30);
    foo.push(10);
    bar.push(101);
    bar.push(202);

    foo.swap(bar);

    std::cout << "size of foo: " << foo.size() << '\n';
    std::cout << "size of bar: " << bar.size() << '\n';

    return 0;
}