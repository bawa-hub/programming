// https://en.cppreference.com/w/c/language/type
// https://en.cppreference.com/w/cpp/language/type

// integer
// character
// boolean
// floating point
// double floating point
// void
// wide character

// ranges
// roughly
// -10^9 < int < 10^9
// -10^12 < long int < 10^12
// -10^18 < long long int < 10^18

#include <iostream>
using namespace std;

int main()
{
    char c = 'c';
    int a = 3;
    float b = 1.4;
    double d = 4.56;
    bool bo = false;
    cout << c << " " << a << " " << b << " " << d << " " << bo << endl;

    // overflow
    int a1 = 100000;
    int b1 = 100000;
    int c1 = a1 * b1;
    long long d1 = a1 * 1LL * b1;
    cout << c1 << endl; // unexpected answer
    cout << d1 << endl;

    int mx = INT_MAX;
    cout << mx + 1 << endl; // overflow

    // double precision
    double p = 100000;
    double q = 100000;
    double r = p * q;
    cout << r << endl;
    cout << fixed << r << endl;
    cout << fixed << setprecision(0) << r << endl;

    r = 1e24;
    cout << fixed << r << endl; // precision error
}