// 03_ctors_dtors_this.cpp
// Constructors, destructor, and the `this` pointer

#include <iostream>
#include <string>

class Logger {
    std::string name;
    bool enabled{};

public:
    // Default constructor
    Logger() : name("default"), enabled(true) {
        std::cout << "Logger(" << name << ") constructed (default)\n";
    }

    // Parameterized constructor
    Logger(const std::string& name, bool enabled = true)
        : name(name), enabled(enabled) {
        std::cout << "Logger(" << this->name << ") constructed (param)\n";
    }

    // Copy constructor
    Logger(const Logger& other)
        : name(other.name), enabled(other.enabled) {
        std::cout << "Logger(" << name << ") copied\n";
    }

    // Destructor
    ~Logger() {
        std::cout << "Logger(" << name << ") destroyed\n";
    }

    Logger& setEnabled(bool flag) {
        this->enabled = flag;   // using this to emphasize member access
        return *this;           // enables chaining
    }

    void log(const std::string& msg) const {
        if (enabled) {
            std::cout << "[" << name << "] " << msg << '\n';
        }
    }
};

int main() {
    Logger a;                       // default
    a.log("Hello from default logger");

    Logger b("main");               // parameterized
    b.log("Hello from main logger");

    Logger c = b;                   // copy
    c.setEnabled(false).log("You should not see this");

    // Temporary object
    Logger("temp").log("Temporary logger");

    return 0;
}

