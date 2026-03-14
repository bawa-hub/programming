// 19_exception_safety_raii.cpp
// Exception safety levels and RAII usage

#include <iostream>
#include <stdexcept>

class Resource {
public:
    Resource() {
        std::cout << "Resource acquired\n";
    }

    ~Resource() {
        std::cout << "Resource released\n";
    }
};

class Container {
    int* data;
    std::size_t size;

public:
    Container() : data(nullptr), size(0) {}

    explicit Container(std::size_t n)
        : data(new int[n]{}), size(n) {}

    ~Container() {
        delete[] data;
    }

    std::size_t getSize() const { return size; }

    // Strong exception safety: copy-and-swap
    void resize(std::size_t newSize) {
        if (newSize == size) return;

        int* newData = new int[newSize]{}; // may throw bad_alloc
        // copy existing elements (up to min(size, newSize))
        std::size_t n = (newSize < size) ? newSize : size;
        for (std::size_t i = 0; i < n; ++i) {
            newData[i] = data[i];
        }

        // Commit: no throws after this point
        delete[] data;
        data = newData;
        size = newSize;
    }
};

void mayThrow(bool reallyThrow) {
    Resource r; // RAII: will always be released
    std::cout << "Inside mayThrow\n";
    if (reallyThrow) {
        throw std::runtime_error("Something went wrong");
    }
    std::cout << "Exiting mayThrow normally\n";
}

int main() {
    std::cout << "--- Exception + RAII demo ---\n";
    try {
        mayThrow(true);
    } catch (const std::exception& ex) {
        std::cout << "Caught exception: " << ex.what() << '\n';
    }

    std::cout << "\n--- Strong exception safety resize demo ---\n";
    Container c(5);
    std::cout << "Initial size = " << c.getSize() << '\n';
    c.resize(10);
    std::cout << "After resize size = " << c.getSize() << '\n';

    return 0;
}

