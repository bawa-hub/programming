/** Under the hood, the std::sort algorithm typically uses a variation of the QuickSort or IntroSort algorithm, both of which are comparison-based sorting algorithms. Here's a simplified explanation of how the comparator works within the sorting algorithm:

    Function Pointer or Functor:
        When you pass a comparator function to std::sort, it's either a function pointer or a functor (a class with operator() overloaded).
        In C++, functions are treated as first-class citizens, meaning you can pass them around as arguments to other functions.

    Comparator Invocation:
        The sorting algorithm invokes the comparator function multiple times during the sorting process.
        The comparator is called to compare pairs of elements to determine their relative order.

    Comparisons:
        The comparator function returns true if the first element should precede the second one in the sorted sequence, and false otherwise.
        Based on these comparisons, the sorting algorithm rearranges the elements to achieve the desired order.

    Algorithm Execution:
        The sorting algorithm (e.g., QuickSort) uses these comparisons to partition the elements and recursively sort subarrays.
        The exact behavior of the sorting algorithm depends on the specific implementation and the characteristics of the data being sorted.

    Efficiency:
        Efficient sorting algorithms, like QuickSort and IntroSort, exploit the comparator's results to minimize the number of comparisons needed to sort the elements.
        By using the comparator's information intelligently, these algorithms can achieve optimal or near-optimal performance.

    Stability:
        If the comparator defines a strict weak ordering (i.e., if it satisfies the requirements for comparison functions), the sorting algorithm maintains the stability of the sort.
        Stability means that equal elements retain their relative order in the sorted sequence.

Overall, the comparator function plays a crucial role in guiding the sorting algorithm's behavior and determining the final order of the elements. The sorting algorithm relies on the comparator's logic to make decisions about how to arrange the elements efficiently and correctly. **/

// https://en.cppreference.com/w/cpp/algorithm/sort.html

#include <iostream>
#include <algorithm>
#include <vector>
using namespace std;

// Custom comparison function
bool customCompare(int a, int b) {
    // Define your custom sorting criteria here
    // For example, let's sort integers in ascending order
    return a < b;
}

int main() {
    vector<int> numbers = {5, 2, 8, 1, 9};

    // default sort in ascending order
    // sort(numbers.begin(), numbers.end());

    // Sorting using custom comparison function
    sort(numbers.begin(), numbers.end(), customCompare);

    // Display sorted numbers
    cout << "Sorted numbers in ascending order: ";
    for (int num : numbers) {
        cout << num << " ";
    }
    cout << endl;

     // Sorting using lambda expression
    sort(numbers.begin(), numbers.end(), [](int a, int b) {
        // Define your custom sorting criteria here
        // For example, let's sort integers in descending order
        return a > b;
    });

    cout << "Sorted numbers in descending order: ";
    for (int num : numbers) {
        cout << num << " ";
    }
    cout << endl;

    return 0;
}
