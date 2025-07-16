#include <iostream>
#include <string>
using namespace std;

int main()
{
    // Initializing string
    string str11 = "geeksforgeeks is for geeks";

    // Displaying string
    cout << "The initial string is : ";
    cout << str11 << endl;

    // both size() and length() return length of string and
    // they work as synonyms
    int len = str11.length(); // Same as "len = str6.size();"

    cout << "Length of string 6 is : " << len << "\n";
    ;

    // Resizing string using resize()
    str11.resize(13);

    // Displaying string
    cout << "The string after resize operation is : ";
    cout << str11 << endl;

    // Displaying capacity of string
    cout << "The capacity of string is : ";
    cout << str11.capacity() << endl;

    //Displaying length of the string
    cout << "The length of the string is :" << str11.length() << endl;

    // Decreasing the capacity of string
    // using shrink_to_fit()
    str11.shrink_to_fit();

    // Displaying string
    cout << "The new capacity after shrinking is : ";
    cout << str11.capacity() << endl;

    // clearing string
    str11.clear();

    // Checking if string is empty
    (str11.empty() == 1) ? cout << "String is empty" << endl : cout << "String is not empty" << endl;
}