#include <iostream>
#include <iomanip>
using namespace std;

double total(int quantity, int price)
{
    double t = quantity * price;
    if (quantity > 1000)
    {
        t = t - (0.1 * t);
        return t;
    }
    else
        return t;
}

int main()
{
    int t;
    cin >> t;
    while (t--)
    {
        int q, p;
        cin >> q >> p;
        cout << fixed;
        cout << setprecision(6);
        cout << total(q, p) << "\n";
    }
    return 0;
}