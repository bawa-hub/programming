#include <iostream>
#include <iomanip>
using namespace std;

int main()
{
    int x;
    float y;
    cin >> x;
    cin >> y;
    cout << fixed << setprecision(2);
    if (x % 5 == 0 && y >= x + 0.50)
        y = y - x - 0.50;
    cout << y << endl;

    return 0;
}
