// 04_static_members.cpp
// Static data members and static member functions

#include <iostream>

class InstanceCounter {
    static int count; // declaration

public:
    InstanceCounter() {
        ++count;
    }

    ~InstanceCounter() {
        --count;
    }

    static int getCount() {
        return count;
    }
};

// Definition of static data member
int InstanceCounter::count = 0;

int main() {
    std::cout << "Initial count: " << InstanceCounter::getCount() << '\n';
    InstanceCounter a;
    std::cout << "After creating a: " << InstanceCounter::getCount() << '\n';
    {
        InstanceCounter b, c;
        std::cout << "Inside block: " << InstanceCounter::getCount() << '\n';
    }
    std::cout << "After block ends: " << InstanceCounter::getCount() << '\n';
    return 0;
}

