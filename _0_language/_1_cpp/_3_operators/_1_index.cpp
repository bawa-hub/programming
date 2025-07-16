// https://www.tutorialspoint.com/cprogramming/c_operators_precedence.htm
// https://en.cppreference.com/w/cpp/language/operator_precedence
#include <bits/stdc++.h>
using namespace std;

int main()
{
    int a = 1;
    a += 1;
    cout << a << endl;
    // ++a first increment then use
    // a++ first use then increment
    cout << a++ << endl;
    cout << ++a << endl;
}