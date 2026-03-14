// 13_smart_pointers_polymorphism.cpp
// Smart pointers with polymorphism

#include <iostream>
#include <memory>
#include <string>
#include <vector>

class Shape {
public:
    virtual ~Shape() = default;
    virtual double area() const = 0;
    virtual std::string name() const = 0;
};

class Circle : public Shape {
    double r{};
public:
    explicit Circle(double radius) : r(radius) {}
    double area() const override { return 3.141592653589793 * r * r; }
    std::string name() const override { return "Circle"; }
};

class Rectangle : public Shape {
    double w{};
    double h{};
public:
    Rectangle(double width, double height) : w(width), h(height) {}
    double area() const override { return w * h; }
    std::string name() const override { return "Rectangle"; }
};

int main() {
    std::vector<std::unique_ptr<Shape>> shapes;
    shapes.push_back(std::make_unique<Circle>(2.0));
    shapes.push_back(std::make_unique<Rectangle>(3.0, 4.0));

    for (const auto& s : shapes) {
        std::cout << s->name() << " area = " << s->area() << '\n';
    }

    // shared_ptr example
    std::shared_ptr<Shape> sharedCircle = std::make_shared<Circle>(5.0);
    std::shared_ptr<Shape> sharedCircle2 = sharedCircle; // reference count++

    std::cout << sharedCircle->name() << " area = " << sharedCircle->area() << '\n';
    std::cout << "Use count of sharedCircle: " << sharedCircle.use_count() << '\n';

    return 0;
}

