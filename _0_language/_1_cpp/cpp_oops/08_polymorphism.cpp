// 08_polymorphism.cpp
// Static and dynamic polymorphism examples

#include <iostream>
#include <string>

// ---------- Static polymorphism via overloading ----------

void print(int x) {
    std::cout << "int: " << x << '\n';
}

void print(double x) {
    std::cout << "double: " << x << '\n';
}

void print(const std::string& s) {
    std::cout << "string: " << s << '\n';
}

// ---------- Dynamic polymorphism via virtual functions ----------

class Shape {
public:
    virtual ~Shape() = default;
    virtual double area() const = 0; // pure virtual
};

class Circle : public Shape {
    double r{};
public:
    explicit Circle(double radius) : r(radius) {}
    double area() const override {
        return 3.141592653589793 * r * r;
    }
};

class Rectangle : public Shape {
    double w{};
    double h{};
public:
    Rectangle(double width, double height) : w(width), h(height) {}
    double area() const override {
        return w * h;
    }
};

void printArea(const Shape& s) {
    std::cout << "Area = " << s.area() << '\n';
}

int main() {
    // Static polymorphism
    print(42);
    print(3.14);
    print(std::string("hello"));

    // Dynamic polymorphism
    Circle c(2.0);
    Rectangle r(3.0, 4.0);

    printArea(c);
    printArea(r);

    return 0;
}

