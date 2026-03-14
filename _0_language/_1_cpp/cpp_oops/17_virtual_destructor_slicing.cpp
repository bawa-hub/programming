// 17_virtual_destructor_slicing.cpp
// Virtual destructors and object slicing

#include <iostream>
#include <memory>

class Base {
public:
    Base() { std::cout << "Base constructed\n"; }
    virtual ~Base() { std::cout << "Base destroyed\n"; } // virtual destructor

    virtual void whoAmI() const {
        std::cout << "I am Base\n";
    }
};

class Derived : public Base {
    int* data;
public:
    Derived() : data(new int[5]{1,2,3,4,5}) {
        std::cout << "Derived constructed with resource\n";
    }

    ~Derived() override {
        std::cout << "Derived destroyed, releasing resource\n";
        delete[] data;
    }

    void whoAmI() const override {
        std::cout << "I am Derived\n";
    }
};

void callWhoAmIByReference(const Base& b) {
    b.whoAmI(); // virtual dispatch
}

int main() {
    std::cout << "--- Virtual destructor demo ---\n";
    {
        Base* b = new Derived(); // constructed as Derived
        b->whoAmI();             // calls Derived::whoAmI
        delete b;                // calls Derived::~Derived then Base::~Base (because virtual)
    }

    std::cout << "\n--- Object slicing demo ---\n";
    Derived d;
    Base sliced = d;             // object slicing: Derived part is sliced off
    callWhoAmIByReference(d);    // prints "I am Derived"
    callWhoAmIByReference(sliced); // prints "I am Base"

    return 0;
}

