#include <iostream>
#include <string>
using namespace std;
int main()
{
    string str("The Geeks for Geeks");
    string str1("The Geeks for Geeks");
    string str2("string");

    // find() returns position to first
    // occurrence of substring "Geeks"
    // Prints 4
    cout << "First occurrence of \"Geeks\" starts from : ";
    cout << str.find("Geeks") << endl;

    // Prints position of first occurrence of
    // any character of "reef" (Prints 2)
    cout << "First occurrence of character from \"reef\" is at : ";
    cout << str.find_first_of("reef") << endl;

    // Prints position of last occurrence of
    // any character of "reef" (Prints 16)
    cout << "Last occurrence of character from \"reef\" is at : ";
    cout << str.find_last_of("reef") << endl;

    // rfind() returns position to last
    // occurrence of substring "Geeks"
    // Prints 14
    cout << "Last occurrence of \"Geeks\" starts from : ";
    cout << str.rfind("Geeks") << endl;

    // substr(a, b) function returns a substring of b length
    // starting from index a
    cout << str2.substr(7, 3) << endl;

    // if second argument is not passed, string till end is
    // taken as substring
    cout << str2.substr(7) << endl;

    // erase(a, b) deletes b characters at index a
    str2.erase(7, 4);
    cout << str2 << endl;

    // iterator version of erase
    str2.erase(str2.begin() + 5, str2.end() - 3);
    cout << str2 << endl;

    str2 = "This is a examples";

    // replace(a, b, str) replaces b characters from a index by str
    str2.replace(2, 7, "ese are test");

    cout << str2 << endl;

    // Comparing strings using compare()
    if (str.compare(str1) == 0)
        cout << "String 1 and String 2 are equal"
             << "\n";
    else
        cout << "Strings are unequal"
             << "\n";
    ;

    // copy()
    // Declaring character array
    char chr[80];

    // using copy() to copy elements into char array
    // copies "geeksforgeeks"
    str1.copy(chr, 13, 0);

    // Diplaying char array
    cout << "The new copied character array is : ";
    cout << chr << endl
         << endl;

    // find returns index where pattern is found.
    // If pattern is not there it returns predefined
    // constant npos whose value is -1

    if (str.find(str1) != string::npos)
        cout << "str4 found in str6 at " << str.find(str1)
             << " pos" << endl;
    else
        cout << "str4 not found in str6" << endl;

    // c_str returns null terminated char array version of string
    const char *charstr = str.c_str();
    printf("%s\n", charstr);
}