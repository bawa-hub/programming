#include <iostream>
#include <string>
using namespace std;

int main()
{
    // initialization by raw string
    string str1("first string");

    // initialization by another string
    string str2(str1);

    // initialization by character with number of occurrence
    string str3(5, '#');

    // initialization by part of another string
    string str4(str1, 6, 6); // from 6th index (second parameter)
                             // 6 characters (third parameter)

    // initialization by part of another string : iteartor version
    string str5(str2.begin(), str2.begin() + 5);

    // assignment operator
    string str6 = str4;

    // C style string initialization
    // Size has to be predefined in character array
    char str7[80] = "GeeksforGeeks";

    cout << "-----------String Declarations--------------"
         << "\n\n";
    cout << str1 << endl;
    cout << str2 << endl;
    cout << str3 << endl;
    cout << str4 << endl;
    cout << str5 << endl;
    cout << str6 << endl;
    cout << str7 << endl;
    cout << "\n\n";
}