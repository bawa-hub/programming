// https://takeuforward.org/data-structure/sieve-of-eratosthenes

#include<bits/stdc++.h>
using namespace std;

class Solution {
public:
    std::vector<int> primesInRange(std::vector<std::vector<int>>& queries) {
        if (queries.empty()) {
            return {};
        }

        // Find the maximum value in the queries 
        // to determine the sieve range
        int maxVal = 0;
        for (const auto& query : queries) {
            maxVal = std::max(maxVal, query[1]);
        }

        // Step 1: Use the Sieve of Eratosthenes 
        // to find all primes up to maxVal
        std::vector<bool> isPrime(maxVal + 1, true);
        isPrime[0] = isPrime[1] = false;  // 0 and 1 are not primes
        for (int p = 2; p * p <= maxVal; ++p) {
            if (isPrime[p]) {
                for (int i = p * p; i <= maxVal; i += p) {
                    isPrime[i] = false;
                }
            }
        }

        // Step 2: Create a prefix sum array 
        // to count primes up to each number
        std::vector<int> primeCount(maxVal + 1, 0);
        for (int i = 1; i <= maxVal; ++i) {
            primeCount[i] = primeCount[i - 1];
            if (isPrime[i]) {
                primeCount[i]++;
            }
        }

        // Step 3: Process each query to find the number of primes 
        // in the given range
        std::vector<int> result;
        for (const auto& query : queries) {
            int start = query[0];
            int end = query[1];
            if (start == 0) {
                result.push_back(primeCount[end]);
            } else {
                result.push_back(primeCount[end] - primeCount[start - 1]);
            }
        }

        return result;
    }
};

int main() {
    // Example usage
    
    
    // Create an instance of the Solution class
    Solution solution;
    
    // Call the function and store the result
    auto result = solution.primesInRange(queries);
    
    // Output the result
    cout << "The number of primes in the given ranges are: ";
    for (int count : result) {
        cout << count << " ";
    }
    cout << endl;

    return 0;
}

// Time Complexity: O(N log (log N)) for the Sieve of Eratosthenes, where n is the maximum value in the queries. The prefix sum computation takes O(n), and each query is processed in O(1) time. Thus, the overall time complexity is dominated by the sieve, resulting in O(N log (log N)).
// Space Complexity: O(N) for storing the prime status array and the prefix sum array, where n is the maximum value in the queries. The space complexity is primarily due to the storage of these arrays, with each requiring space proportional to the maximum value. 