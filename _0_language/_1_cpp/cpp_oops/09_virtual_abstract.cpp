// 09_virtual_abstract.cpp
// Virtual functions, override/final, abstract classes

#include <iostream>
#include <memory>

class Device {
public:
    virtual ~Device() = default;

    virtual void start() = 0;     // pure virtual => abstract class
    virtual void stop() {         // virtual with default implementation
        std::cout << "Device stopping (default)\n";
    }
};

class Printer final : public Device { // final: cannot be further derived
public:
    void start() override {
        std::cout << "Printer starting\n";
    }

    void stop() final override { // final: cannot be overridden further
        std::cout << "Printer stopping\n";
    }
};

void runDevice(Device& d) {
    d.start();
    d.stop();
}

int main() {
    Printer p;
    runDevice(p);

    std::unique_ptr<Device> dev = std::make_unique<Printer>();
    runDevice(*dev);

    return 0;
}

