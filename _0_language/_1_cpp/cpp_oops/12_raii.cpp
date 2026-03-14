// 12_raii.cpp
// RAII: Resource Acquisition Is Initialization

#include <fstream>
#include <iostream>
#include <mutex>
#include <string>

class FileWriter {
    std::ofstream out;

public:
    explicit FileWriter(const std::string& path)
        : out(path, std::ios::out) {
        if (!out) {
            throw std::runtime_error("Failed to open file");
        }
    }

    ~FileWriter() {
        // Destructor ensures file is closed even if exceptions occur
        if (out.is_open()) {
            out.close();
        }
    }

    void writeLine(const std::string& line) {
        out << line << '\n';
    }
};

class MutexGuard {
    std::mutex& m;

public:
    explicit MutexGuard(std::mutex& m) : m(m) {
        m.lock();
    }

    ~MutexGuard() {
        m.unlock();
    }
};

int main() {
    // RAII with file
    try {
        FileWriter writer("raii_example.txt");
        writer.writeLine("Hello from RAII!");
    } catch (const std::exception& ex) {
        std::cerr << "Error: " << ex.what() << '\n';
    }

    // RAII with mutex
    std::mutex m;
    {
        MutexGuard guard(m); // m.lock()
        std::cout << "In critical section\n";
    }                         // guard destroyed -> m.unlock()

    std::cout << "After critical section\n";
    return 0;
}

