// 20_composition_vs_inheritance.cpp
// Composition vs inheritance and simple SOLID-style design

#include <iostream>
#include <string>

// Bad-style inheritance (for illustration only):
// A `LoggingVector` that publicly inherits from std::string or std::vector
// would be an example of "is-not-really-a" misuse. We avoid that here.

// Better: use composition.

class Logger {
public:
    void log(const std::string& msg) const {
        std::cout << "[LOG] " << msg << '\n';
    }
};

class FileStorage {
public:
    void save(const std::string& data) const {
        std::cout << "Saving data: " << data << '\n';
    }
};

// Composition: UserService "has-a" Logger and FileStorage
class UserService {
    Logger logger;          // composed
    FileStorage storage;    // composed

public:
    void createUser(const std::string& name) {
        logger.log("Creating user: " + name);
        storage.save("User:" + name);
    }
};

// Interface-style inheritance for real "is-a"
class Animal {
public:
    virtual ~Animal() = default;
    virtual void speak() const = 0;
};

class Dog : public Animal {
public:
    void speak() const override {
        std::cout << "Woof!\n";
    }
};

class Cat : public Animal {
public:
    void speak() const override {
        std::cout << "Meow!\n";
    }
};

int main() {
    std::cout << "--- Composition demo ---\n";
    UserService service;
    service.createUser("Vikram");

    std::cout << "\n--- Inheritance (interface) demo ---\n";
    Animal* a1 = new Dog;
    Animal* a2 = new Cat;
    a1->speak();
    a2->speak();
    delete a1;
    delete a2;

    return 0;
}

