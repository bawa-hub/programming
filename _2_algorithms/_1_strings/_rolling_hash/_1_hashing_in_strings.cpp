// https://cp-algorithms.com/string/string-hashing.html
// https://www.geeksforgeeks.org/string-hashing-using-polynomial-rolling-hash-function/

/*

1. Rolling Hash Implementation (C++)

We treat string as a polynomial.
hash(s)=s0​Bn−1+s1​Bn−2+...+sn

Common constants:
base = 31 or 131
mod = 1e9+7

*/

#include <bits/stdc++.h>
using namespace std;

class RollingHash {
public:
    const long long base = 31;
    const long long mod = 1e9 + 7;

    vector<long long> prefix;
    vector<long long> power;

    RollingHash(string s) {
        int n = s.size();
        prefix.resize(n + 1);
        power.resize(n + 1);

        power[0] = 1;

        for (int i = 1; i <= n; i++) {
            power[i] = (power[i - 1] * base) % mod;
        }

        for (int i = 0; i < n; i++) {
            prefix[i + 1] =
                (prefix[i] * base + (s[i] - 'a' + 1)) % mod;
        }
    }

    long long getHash(int l, int r) {
        long long hash =
            (prefix[r + 1] -
             prefix[l] * power[r - l + 1] % mod +
             mod) % mod;

        return hash;
    }
};

class RollingHash1 {

    static const long long mod1 = 1000000007;
    static const long long mod2 = 1000000009;
    static const long long base = 91138233;

    vector<long long> hash1, hash2;
    vector<long long> power1, power2;

    RollingHash(string s) {

        int n = s.size();

        hash1.resize(n+1);
        hash2.resize(n+1);
        power1.resize(n+1);
        power2.resize(n+1);

        power1[0] = power2[0] = 1;

        for(int i=1;i<=n;i++)
        {
            power1[i] = (power1[i-1] * base) % mod1;
            power2[i] = (power2[i-1] * base) % mod2;
        }

        for(int i=0;i<n;i++)
        {
            hash1[i+1] = (hash1[i]*base + s[i]) % mod1;
            hash2[i+1] = (hash2[i]*base + s[i]) % mod2;
        }
    }

    pair<long long,long long> getHash(int l,int r)
    {
        long long x1 =
        (hash1[r+1] - hash1[l]*power1[r-l+1] % mod1 + mod1) % mod1;

        long long x2 =
        (hash2[r+1] - hash2[l]*power2[r-l+1] % mod2 + mod2) % mod2;

        return {x1,x2};
    }
};





int main() {
    string s = "abcdef";
    RollingHash rh(s);

    cout << rh.getHash(1,3) << endl; // bcd

    s = "banana";

    RollingHash1 rh1(s);

    auto h1 = rh1.getHash(1,3);
    auto h2 = rh1.getHash(3,5);

    if(h1 == h2)
        cout<<"equal substring";

}