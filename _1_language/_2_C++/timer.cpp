#include <iostream>
#include <chrono>
#include <unordered_map>
using namespace std;
using namespace std::chrono;

int fib(int n, unordered_map<int, int> &memo)
{
    if (memo[n])
        return memo[n];
    if (n <= 2)
        return 1;
    memo[n] = fib(n - 1, memo) + fib(n - 2, memo);
    return memo[n];
}

int main()
{
    // Get starting timepoint
    auto start = high_resolution_clock::now();

    unordered_map<int, int> mp{};

    for (int i = 1; i <= 45; i++)
    {
        cout << fib(i, mp) << " ";
    }

    cout << endl;

    // Get ending timepoint
    auto stop = high_resolution_clock::now();

    auto duration = duration_cast<microseconds>(stop - start);

    cout << "Time taken by function: "
         << duration.count() << " microseconds" << endl;
}