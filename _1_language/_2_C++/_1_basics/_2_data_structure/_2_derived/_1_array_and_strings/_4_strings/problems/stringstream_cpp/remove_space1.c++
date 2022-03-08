// C++ program to remove spaces using stringstream
// and getline()
#include <iostream>
#include <sstream>
#include <string>
using namespace std;

// Function to remove spaces
string removeSpaces(string str)
{
    // Storing the whole string
    // into string stream
    stringstream ss(str);
    string temp;

    // Making the string empty
    str = "";

    // Running loop till end of stream
    // and getting every word
    while (getline(ss, temp, ' '))
    {
        // Concatenating in the string
        // to be returned
        str = str + temp;
    }
    return str;
}
// Driver function
int main()
{
    // Sample Inputs
    string s = "This is a test";
    cout << removeSpaces(s) << endl;

    s = "geeks for geeks";
    cout << removeSpaces(s) << endl;

    s = "geeks quiz is awsome!";
    cout << removeSpaces(s) << endl;

    s = "I love	 to	 code";
    cout << removeSpaces(s) << endl;

    return 0;
}

// Code contributed by saychakr13
