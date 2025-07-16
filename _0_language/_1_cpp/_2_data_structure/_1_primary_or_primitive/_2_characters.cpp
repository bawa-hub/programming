// https://www.cs.cmu.edu/~pattis/15-1XX/common/handouts/ascii.html
// ascii characters https://www.ascii-code.com/

#include <iostream>
using namespace std;

void printAtoZ() {
    for(char i = 'a'; i<='z'; i++) cout << i << " ";
    cout << endl;
}

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

    // toLower or toUpper
    char c1 = 'a';
    cout << toupper(c1);
    char c2 = 'A';
    cout << tolower(c2);
}