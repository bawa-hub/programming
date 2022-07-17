// https://www.cs.cmu.edu/~pattis/15-1XX/common/handouts/ascii.html
#include <iostream>
using namespace std;

int main()
{
    char c = 'A';
    cout << c << endl;
    // print ascii value
    cout << (int)c << endl;

    char a;
    cout << "Enter a character: ";
    cin >> a;
    cout << a << "with ascii value of " << (int)a << endl;

    // change char to int
    char b = '5';
    cout << b << "with ascii value of " << b - '0' << endl;

    // int to char
    int i = 2;
    char ch = '0' + i;
    cout << "ch: " << ch;
}