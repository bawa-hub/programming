// https://stackoverflow.com/questions/7663709/how-can-i-convert-a-stdstring-to-int/#answer-26530289

#include <iostream>
#include <string>
#include <sstream>
// cstdlib is needed for atoi()
#include <cstdlib>
using namespace std;

int main()
{

    // C++ string to int Using stoi()
    string str = "123";
    int num;
    num = stoi(str);
    cout << num + 1 << endl;

    // char Array to int Using atoi()
    char str1[] = "456";
    int num1 = atoi(str1);
    cout << num1 + 1 << endl;

    // C++ int to string Using to_string()
    int num2 = 123;
    string str2 = to_string(num2);
    cout << str2 << endl;

    // C++ int to string Using stringstream
    int num3 = 15;
    stringstream ss;
    ss << num;
    string strNum = ss.str();
    cout << strNum;

    return 0;
}