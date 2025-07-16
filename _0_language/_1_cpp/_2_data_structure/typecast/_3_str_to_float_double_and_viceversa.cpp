/**
 *
    std::stof() - convert string to float
    std::stod() - convert string to double
    std::stold() - convert string to long double.

 *
 */

#include <iostream>
#include <string>
#include <cstdlib> // cstdlib is needed for atoi()
#include <sstream> // for using stringstream
using namespace std;

int main()
{

    // C++ string to float and double
    string str = "123.4567";
    // convert string to float
    float num_float = stof(str);
    // convert string to double
    double num_double = stod(str);
    cout << "num_float = " << num_float << endl;
    cout << "num_double = " << num_double << endl;

    // C++ char Array to double
    char str1[] = "123.4567";
    double num_double1 = atof(str1);
    cout << "num_double = " << num_double1 << std::endl;

    // float and double to string Using to_string()
    float num_float2 = 123.4567F;
    double num_double2 = 123.4567;
    string str2 = to_string(num_float2);
    string str3 = to_string(num_double2);
    cout << "Float to String = " << str2 << endl;
    cout << "Double to String = " << str3 << endl;

    // float and double to string Using stringstream
    float num_float3 = 123.4567F;
    double num_double3 = 123.4567;
    stringstream ss1;
    stringstream ss2;
    ss1 << num_float3;
    ss2 << num_double3;
    string str4 = ss1.str();
    string str5 = ss2.str();
    cout << "Float to String = " << str4 << endl;
    cout << "Double to String = " << str5 << endl;

    return 0;
}