// 07_inheritance.cpp
// Basic inheritance example

#include <iostream>
#include <string>

class Animal {
protected:
    std::string name;

public:
    Animal(const std::string& n) : name(n) {}

    void eat() const {
        std::cout << name << " is eating.\n";
    }
};

class Dog : public Animal {
public:
    Dog(const std::string& n) : Animal(n) {}

    void bark() const {
        std::cout << name << " says: Woof!\n";
    }
};

class Cat : public Animal {
public:
    Cat(const std::string& n) : Animal(n) {}

    void meow() const {
        std::cout << name << " says: Meow!\n";
    }
};

int main() {
    Dog d("Buddy");
    d.eat();  // inherited
    d.bark();

    Cat c("Kitty");
    c.eat();
    c.meow();

    return 0;
}

