#include <iostream>
using namespace std;

void printAtoZ() {
    for(char i = 'a'; i<='z'; i++) cout << i << " ";
    cout << endl;
}

int main() {
   printAtoZ();

   return 0;
}