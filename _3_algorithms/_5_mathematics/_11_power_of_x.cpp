using namespace std;

long long power(int N, int R)
{
    int MOD = 1e9 + 7;
    // step 1
    if (R == 0)
        return 1;
    // step 2
    long long ans = power(N, R - 1);
    // step 3 . note we need to do mod as the result may
    // overflow
    ans = (ans % MOD * N % MOD) % MOD;
    return ans;
}