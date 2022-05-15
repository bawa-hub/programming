#include <iostream>
using namespace std;

int main()
{
    // your code goes here
    int t;
    cin >> t;
    while (t--)
    {
        int h, t, a = 0, b = 0, c = 0;
        float s;
        cin >> h >> s >> t;
        if (h > 50)
        {
            a = 1;
        }
        if (s < 0.7)
        {
            b = 1;
        }
        if (t > 5600)
        {
            c = 1;
        }
        if (a == 1 && b == 1 && c == 1)
        {
            cout << "10" << endl;
        }
        else if (a == 1 && b == 1 && c == 0)
        {
            cout << "9" << endl;
        }
        else if (a == 0 && b == 1 && c == 1)
        {
            cout << "8" << endl;
        }
        else if (a == 1 && b == 0 && c == 1)
        {
            cout << "7" << endl;
        }
        else if (a == 0 && b == 1 && c == 0)
            cout << "6" << endl;
        else if (a == 0 && b == 0 && c == 1)
            cout << "6" << endl;
        else if (a == 1 && b == 0 && c == 0)
            cout << "6" << endl;
        else if (a == 0 && b == 0 && c == 0)
            cout << "5" << endl;
    }
    return 0;
}
