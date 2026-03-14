// 14_templates_and_oop.cpp
// Combining templates (static polymorphism) with OOP (dynamic polymorphism)

#include <iostream>
#include <memory>
#include <string>
#include <vector>

// Generic function (template)
template <typename T>
T add(T a, T b) {
    return a + b;
}

// Base interface
class Operation {
public:
    virtual ~Operation() = default;
    virtual int apply(int x, int y) const = 0;
    virtual std::string name() const = 0;
};

class AddOperation : public Operation {
public:
    int apply(int x, int y) const override { return x + y; }
    std::string name() const override { return "Add"; }
};

class MultiplyOperation : public Operation {
public:
    int apply(int x, int y) const override { return x * y; }
    std::string name() const override { return "Multiply"; }
};

// Template function operating on any Operation-like type
template <typename Op>
void runAndPrint(const Op& op, int a, int b) {
    std::cout << op.name() << "(" << a << ", " << b << ") = "
              << op.apply(a, b) << '\n';
}

int main() {
    // Template add
    std::cout << "add<int>(1, 2) = " << add(1, 2) << '\n';
    std::cout << "add<double>(1.5, 2.3) = " << add(1.5, 2.3) << '\n';

    // Dynamic polymorphism with Operation hierarchy
    std::vector<std::unique_ptr<Operation>> ops;
    ops.push_back(std::make_unique<AddOperation>());
    ops.push_back(std::make_unique<MultiplyOperation>());

    int a = 3;
    int b = 4;
    for (const auto& op : ops) {
        std::cout << op->name() << "(" << a << ", " << b << ") = "
                  << op->apply(a, b) << '\n';
    }

    // Static polymorphism (template) using concrete types
    AddOperation addOp;
    MultiplyOperation mulOp;
    runAndPrint(addOp, 5, 6);
    runAndPrint(mulOp, 5, 6);

    return 0;
}

