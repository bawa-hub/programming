// 10_multiple_inheritance.cpp
// Multiple inheritance and virtual inheritance (diamond problem)

#include <iostream>
#include <string>

class Named {
protected:
    std::string name;
public:
    explicit Named(std::string n) : name(std::move(n)) {}
};

class Identified {
protected:
    int id{};
public:
    explicit Identified(int id) : id(id) {}
};

// Multiple inheritance of two independent bases
class Entity : public Named, public Identified {
public:
    Entity(std::string name, int id)
        : Named(std::move(name)), Identified(id) {}

    void print() const {
        std::cout << "Entity{id=" << id << ", name=" << name << "}\n";
    }
};

// Diamond inheritance example

class A {
public:
    int value{};
};

class B : virtual public A {};
class C : virtual public A {};

class D : public B, public C {
public:
    void setValue(int v) {
        value = v; // only one A subobject due to virtual inheritance
    }
};

int main() {
    Entity e("Player", 1);
    e.print();

    D d;
    d.setValue(42);
    std::cout << "D.value = " << d.value << '\n';

    return 0;
}

