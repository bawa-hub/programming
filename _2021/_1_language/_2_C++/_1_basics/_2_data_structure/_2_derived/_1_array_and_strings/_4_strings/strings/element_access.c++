#include <iostream>
#include <string>
using namespace std;

int main()
{

    string str6 = "string";
    // a particular character can be accessed using at /
    // [] operator
    char ch = str6.at(2); // Same as "ch = str6[2];"

    cout << "third character of string 6 is : " << ch << endl;

    // front return first character and back returns last character
    // of string
    char ch_f = str6.front(); // Same as "ch_f = str6[0];"
    char ch_b = str6.back();  // Same as "ch_b = str6[str6.length() - 1];"

    cout << "First char of string 6 is : " << ch_f << ", Last char of string 6 is : "
         << ch_b << endl;
}