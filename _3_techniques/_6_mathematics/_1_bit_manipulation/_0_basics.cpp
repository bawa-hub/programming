// https://takeuforward.org/data-structure/introduction-to-bit-manipulation-theory
#include <iostream>
#include <string>
#include <algorithm>
using namespace std;

// bit manipulation must know tricks:
// 1. binary number conversion
// 2. 1's and 2's complement
// 3. Operators -  AND, OR, XOR, NOT, SHIFT
// 4. swap two numbers
// 5. check if ith bit is set or not
// 6, interact with ith bit
// 7. set, clear, toggle the ith bit
// 8. remove last set bit
// 9. count number of set bits
// 10. check if number is power of 2 or not

// convert from decimal to binary
string convertDecToBinary(int num) {
    string res = "";
    while (num != 1) {
        if(num%2==1) res += '1';
        else res += '0';
        num /= 2;
    }
    res += "1";

    reverse(res.begin(),res.end());
    return res;
}
// TC : O(logn)
// SC : O(logn)

int convertBinaryToDec(string s) {
    int p2 = 1;
    int num = 0;

    for(int i=s.size()-1;i>=0;i--) {
        if(s[i]=='1') num += p2;
        p2 *= 2;
    }

    return num;
}
// TC : O(len(s))
// SC : O(1)

void swapWithoutThirdVairable(int &a, int &b) {
    a = a^b;
    b = a^b; // (a^b)^a --> xor of same number is 0; => b = a
    a = a ^ b; // (a^b)^a  => b
}





int main() {
//    int num;
//    cin >> num;
//    cout << convertDecToBinary(num) << endl;

//    string s;
//    cin >> s;
//    cout << convertBinaryToDec(s) << endl;

    // int a,b;
    // cin >> a >> b;
    // swapWithoutThirdVairable(a,b);
    // cout << a << " " << b << endl;
}
