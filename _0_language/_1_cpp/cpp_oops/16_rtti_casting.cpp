// 16_rtti_casting.cpp
// RTTI and casting: dynamic_cast, static_cast, reinterpret_cast, const_cast

#include <iostream>
#include <typeinfo>

class Base {
public:
    virtual ~Base() = default; // polymorphic base
};

class Derived : public Base {
public:
    void hello() const {
        std::cout << "Derived::hello()\n";
    }
};

void testDynamicCast(Base* b) {
    if (auto* d = dynamic_cast<Derived*>(b)) {
        std::cout << "dynamic_cast succeeded: ";
        d->hello();
    } else {
        std::cout << "dynamic_cast failed (b is not a Derived)\n";
    }
}

int main() {
    Derived d;
    Base* b1 = &d;           // upcast (safe, implicit)
    Base* b2 = new Base;     // plain Base

    std::cout << "typeid(*b1).name() = " << typeid(*b1).name() << '\n';
    std::cout << "typeid(*b2).name() = " << typeid(*b2).name() << '\n';

    testDynamicCast(b1);     // should succeed
    testDynamicCast(b2);     // should fail

    // static_cast example (no runtime check)
    Derived* dFromB1 = static_cast<Derived*>(b1); // OK here, we really have Derived
    dFromB1->hello();

    // reinterpret_cast example (dangerous, just demonstration)
    std::uintptr_t addr = reinterpret_cast<std::uintptr_t>(b1);
    Base* bFromInt = reinterpret_cast<Base*>(addr);
    std::cout << "Reinterpreted pointer still points to something, but this is low-level.\n";
    (void)bFromInt;

    // const_cast example
    const int x = 42;
    const int* px = &x;
    int* py = const_cast<int*>(px); // UB to modify x, but reading is fine
    std::cout << "Value via py (do NOT modify if originally const): " << *py << '\n';

    delete b2;
    return 0;
}

