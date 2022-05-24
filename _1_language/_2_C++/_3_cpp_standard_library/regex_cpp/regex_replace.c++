/**
 * regex_replace() This function is used to replace the pattern matching to the 
 * regular expression with a string.
 * **/

// C++ program to demonstrate working of regex_replace()
#include <iostream>
#include <string>
#include <regex>
#include <iterator>
using namespace std;

int main()
{
    string s = "I am looking for GeeksForGeek \n";

    // matches words beginning by "Geek"
    regex r("Geek[a-zA-z]+");

    // regex_replace() for replacing the match with 'geek'
    cout << std::regex_replace(s, r, "geek");

    string result;

    // regex_replace( ) for replacing the match with 'geek'
    regex_replace(back_inserter(result), s.begin(), s.end(),
                  r, "geek");

    cout << result;

    return 0;
}
