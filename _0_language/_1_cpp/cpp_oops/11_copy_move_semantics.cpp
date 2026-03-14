// 11_copy_move_semantics.cpp
// Copy constructor/assignment and move constructor/assignment

#include <algorithm>
#include <iostream>

class Buffer {
    std::size_t size{};
    int* data{};

public:
    Buffer() = default;

    explicit Buffer(std::size_t n)
        : size(n), data(n ? new int[n] : nullptr) {
        std::cout << "Buffer(size=" << size << ") constructed\n";
    }

    // Destructor
    ~Buffer() {
        std::cout << "Buffer(size=" << size << ") destroyed\n";
        delete[] data;
    }

    // Copy constructor
    Buffer(const Buffer& other)
        : size(other.size), data(other.size ? new int[other.size] : nullptr) {
        std::cout << "Buffer copied\n";
        std::copy(other.data, other.data + size, data);
    }

    // Copy assignment
    Buffer& operator=(const Buffer& other) {
        std::cout << "Buffer copy-assigned\n";
        if (this != &other) {
            delete[] data;
            size = other.size;
            data = other.size ? new int[other.size] : nullptr;
            std::copy(other.data, other.data + size, data);
        }
        return *this;
    }

    // Move constructor
    Buffer(Buffer&& other) noexcept
        : size(other.size), data(other.data) {
        std::cout << "Buffer moved (ctor)\n";
        other.size = 0;
        other.data = nullptr;
    }

    // Move assignment
    Buffer& operator=(Buffer&& other) noexcept {
        std::cout << "Buffer move-assigned\n";
        if (this != &other) {
            delete[] data;
            size = other.size;
            data = other.data;
            other.size = 0;
            other.data = nullptr;
        }
        return *this;
    }

    std::size_t getSize() const { return size; }
};

Buffer makeBuffer(std::size_t n) {
    Buffer tmp(n);
    return tmp; // NRVO or move
}

int main() {
    Buffer a(10);
    Buffer b = a;              // copy constructor
    Buffer c;
    c = a;                     // copy assignment

    Buffer d = makeBuffer(5);  // move constructor (typically)
    Buffer e;
    e = makeBuffer(3);         // move assignment (typically)

    std::cout << "d.size = " << d.getSize() << '\n';
    std::cout << "e.size = " << e.getSize() << '\n';

    return 0;
}

