#include <iostream>
using namespace std;

void print(int num)
{
    if (num < 1)
        return;
    cout << num << endl;
    print(num - 1);
}

int main()
{
    print(50);
}
