// 06_operator_overloading.cpp
// Operator overloading example with a Complex number type

#include <iostream>

class Complex {
    double re{};
    double im{};

public:
    Complex() = default;
    Complex(double real, double imag = 0.0) : re(real), im(imag) {}

    double real() const { return re; }
    double imag() const { return im; }

    // Addition
    Complex operator+(const Complex& other) const {
        return Complex(re + other.re, im + other.im);
    }

    // Compound assignment
    Complex& operator+=(const Complex& other) {
        re += other.re;
        im += other.im;
        return *this;
    }

    // Equality
    bool operator==(const Complex& other) const {
        return re == other.re && im == other.im;
    }

    bool operator!=(const Complex& other) const {
        return !(*this == other);
    }

    // Output stream
    friend std::ostream& operator<<(std::ostream& os, const Complex& c) {
        os << c.re;
        if (c.im >= 0) os << " + " << c.im << "i";
        else os << " - " << -c.im << "i";
        return os;
    }
};

int main() {
    Complex a(1.0, 2.0);
    Complex b(3.0, -1.0);

    Complex c = a + b;
    a += b;

    std::cout << "a = " << a << '\n';
    std::cout << "b = " << b << '\n';
    std::cout << "c = " << c << '\n';

    std::cout << "a == c? " << std::boolalpha << (a == c) << '\n';
    std::cout << "b != c? " << std::boolalpha << (b != c) << '\n';

    return 0;
}

