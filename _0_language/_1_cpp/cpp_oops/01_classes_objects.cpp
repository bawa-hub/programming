// 01_classes_objects.cpp
// Basic example of classes and objects in C++

#include <iostream>
#include <string>

class Player {
public:
    std::string name;
    int health{};

    void takeDamage(int dmg) {
        health -= dmg;
        if (health < 0) health = 0;
    }

    void print() const {
        std::cout << "Player{name=" << name << ", health=" << health << "}\n";
    }
};

int main() {
    // Object on stack
    Player p1;
    p1.name = "Vikram";
    p1.health = 100;
    p1.takeDamage(30);
    p1.print();

    // Object on heap
    Player* p2 = new Player;
    p2->name = "AI";
    p2->health = 80;
    p2->takeDamage(10);
    p2->print();

    delete p2; // manual cleanup for heap-allocated object

    return 0;
}

