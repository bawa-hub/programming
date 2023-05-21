#include <iostream>
using namespace std;

int main()
{
    int i;
    float f;
    double d;
    char c;
    string s;
    string line;
    int arr[3];

    cout << "Enter an integer: ";
    cin >> i;
    cout << i << endl;

    cout << "Enter a float: ";
    cin >> f;
    cout << f << endl;

    cout << "Enter double: ";
    cin >> d;
    cout << d << endl;

    cout << "Enter a character: ";
    cin >> c;
    cout << c << endl;
    cout << "ascii value" << (int)c << endl;

    cout << "Enter a string: ";
    cin >> s;
    cout << s << endl;

    cout << "Enter 3 space seperated integers: ";
    for (int i = 0; i < 3; i++)
    {
        cin >> arr[i];
    }
    for (int i = 0; i < 3; i++)
    {
        cout << arr[i] << " ";
    }
    cout << endl;

    return 0;
}