// 18_default_delete_special_members.cpp
// Defaulted and deleted special member functions

#include <iostream>

class NonCopyable {
public:
    NonCopyable() = default;
    ~NonCopyable() = default;

    NonCopyable(const NonCopyable&) = delete;
    NonCopyable& operator=(const NonCopyable&) = delete;
};

class OnlyMovable {
public:
    OnlyMovable() = default;
    ~OnlyMovable() = default;

    OnlyMovable(const OnlyMovable&) = delete;
    OnlyMovable& operator=(const OnlyMovable&) = delete;

    OnlyMovable(OnlyMovable&&) = default;
    OnlyMovable& operator=(OnlyMovable&&) = default;
};

class NoDefaultCtor {
    int value;
public:
    NoDefaultCtor() = delete;          // must supply a value
    explicit NoDefaultCtor(int v) : value(v) {}

    int get() const { return value; }
};

int main() {
    NonCopyable a;
    // NonCopyable b = a;          // ERROR if uncommented: copy deleted

    OnlyMovable m1;
    OnlyMovable m2 = std::move(m1);   // OK: move ctor

    // NoDefaultCtor nd;           // ERROR: default ctor deleted
    NoDefaultCtor nd2(42);
    std::cout << "nd2.get() = " << nd2.get() << '\n';

    (void)m2;
    (void)a;
    return 0;
}

