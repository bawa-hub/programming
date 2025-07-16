#include <iostream>
#include <string>
using namespace std;

int main()
{
    // Initializing string`
    string str12 = "geeksforgeeks";

    // Declaring iterator
    std::string::iterator it;

    // Declaring reverse iterator
    std::string::reverse_iterator it1;

    // Displaying string
    cout << "The string using forward iterators is : ";
    for (it = str12.begin(); it != str12.end(); it++)
        cout << *it;
    cout << endl;

    // Displaying reverse string
    cout << "The reverse string using reverse iterators is : ";
    for (it1 = str12.rbegin(); it1 != str12.rend(); it1++)
        cout << *it1;
    cout << endl;
}