// C++ code to implement Fibonacci series
#include <iostream>
#include <chrono>
using namespace std;
using namespace std::chrono;

// Function for fibonacci

int fib(int n)
{
    // base condition
    if (n == 0)
        return 0;

    // Self work
    if (n == 1 || n == 2)
        return 1;

    // Recursion function
    else
        return (fib(n - 1) + fib(n - 2));
}

// Driver Code
int main()
{

    // Get starting timepoint
    auto start = high_resolution_clock::now();

    // for loop to print the fiboancci series.
    for (int i = 0; i < 45; i++)
    {
        cout << fib(i) << " ";
    }

    // Get ending timepoint
    auto stop = high_resolution_clock::now();

    auto duration = duration_cast<microseconds>(stop - start);

    cout << "Time taken by function: "
         << duration.count() << " microseconds" << endl;

    return 0;
}