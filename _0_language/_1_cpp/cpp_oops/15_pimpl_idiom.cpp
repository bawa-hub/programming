// 15_pimpl_idiom.cpp
// pImpl (Pointer-to-Implementation) idiom in a single translation unit
// In a real project, `Widget` would be in a header and `Impl` in a .cpp.

#include <iostream>
#include <memory>
#include <string>

class Widget {
public:
    Widget();
    ~Widget();

    void setName(const std::string& name);
    void greet() const;

private:
    struct Impl;                     // forward declaration
    std::unique_ptr<Impl> pImpl;     // pointer to implementation
};

// Definition of the implementation struct
struct Widget::Impl {
    std::string name{"default"};

    void greet() const {
        std::cout << "Hello from Widget(" << name << ")\n";
    }
};

// Widget function definitions
Widget::Widget() : pImpl(std::make_unique<Impl>()) {}

Widget::~Widget() = default; // unique_ptr cleans up Impl

void Widget::setName(const std::string& name) {
    pImpl->name = name;
}

void Widget::greet() const {
    pImpl->greet();
}

int main() {
    Widget w;
    w.greet();
    w.setName("pImpl Example");
    w.greet();
    return 0;
}

