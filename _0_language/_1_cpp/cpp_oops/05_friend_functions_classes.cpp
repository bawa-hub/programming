// 05_friend_functions_classes.cpp
// Friend functions and friend classes

#include <iostream>
#include <string>

class Box;

// Friend function example
void printBox(const Box& b);

class Box {
private:
    double width{};
    double height{};
    double depth{};

public:
    Box(double w, double h, double d) : width(w), height(h), depth(d) {}

    // Declare free function as friend
    friend void printBox(const Box& b);

    // Declare friend class
    friend class BoxScaler;
};

void printBox(const Box& b) {
    std::cout << "Box(" << b.width << " x " << b.height << " x " << b.depth << ")\n";
}

class BoxScaler {
public:
    void scale(Box& b, double factor) const {
        b.width *= factor;
        b.height *= factor;
        b.depth *= factor;
    }
};

int main() {
    Box b(2.0, 3.0, 4.0);
    printBox(b);

    BoxScaler scaler;
    scaler.scale(b, 1.5);
    printBox(b);

    return 0;
}

