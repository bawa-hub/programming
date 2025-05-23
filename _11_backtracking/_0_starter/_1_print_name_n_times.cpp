#include <iostream>
#include <string>
using namespace std;

void print(string name, int n)
{
    if (n == 0)
        return;
    cout << name << endl;
    print(name, n - 1);
}

int main()
{
    string name = "Bawa";
    int n = 10;
    print(name, n);
}