#include <iostream>
#include <string>
using namespace std;

int main()
{
    // Declaring string
    string str10;

    // Taking string input using getline()
    // "geeksforgeek" in givin output
    getline(cin, str10);

    // Displaying string
    cout << "The initial string is : ";
    cout << str10 << endl;

    // Using push_back() to insert a character
    // at end
    // pushes 's' in this case
    str10.push_back('s');

    // Displaying string
    cout << "The string after push_back operation is : ";
    cout << str10 << endl;

    // Using pop_back() to delete a character
    // from end
    // pops 's' in this case
    str10.pop_back();

    // Displaying string
    cout << "The string after pop_back operation is : ";
    cout << str10 << endl;

    string str13 = "monu";
    string str14 = "sonu";

    // Displaying strings before swapping
    cout << "The 1st string before swapping is : ";
    cout << str13 << endl;
    cout << "The 2nd string before swapping is : ";
    cout << str14 << endl;

    // using swap() to swap string content
    str13.swap(str14);

    // Displaying strings after swapping
    cout << "The 1st string after swapping is : ";
    cout << str13 << endl;
    cout << "The 2nd string after swapping is : ";
    cout << str14 << endl;

    // demonstrate working of insert()
    string str8("Geeksfor");

    // Printing the original string
    cout << "Original string is: " << str8 << endl;

    // Inserting "Geeks" at 8th index position
    str8.insert(8, "Geeks");

    // Printing the modified string
    // Prints "GeeksforGeeks"
    cout << "After insertion string becomes: " << str8 << endl;

    // append add the argument string at the end
    str10.append(" extension"); // same as str6 += " extension"

    // another version of append, which appends part of other string
    str10.append(str10, 0, 6); // at 0th position 6 character

    cout << "After appending string 6 becomes: " << str10 << endl;
}