// https://practice.geeksforgeeks.org/problems/allocate-minimum-number-of-pages0937/1
// https://www.geeksforgeeks.org/allocate-minimum-number-pages/
// https://www.codingninjas.com/studio/problems/allocate-books_1090540




#include <bits/stdc++.h>
using namespace std;

// brute force
int countStudents(vector<int> &arr, int pages) {
    int n = arr.size(); //size of array.
    int students = 1;
    long long pagesStudent = 0;
    for (int i = 0; i < n; i++) {
        if (pagesStudent + arr[i] <= pages) {
            //add pages to current student
            pagesStudent += arr[i];
        }
        else {
            //add pages to next student
            students++;
            pagesStudent = arr[i];
        }
    }
    return students;
}

int findPages(vector<int>& arr, int n, int m) {
    //book allocation impossible:
    if (m > n) return -1;

    int low = *max_element(arr.begin(), arr.end());
    int high = accumulate(arr.begin(), arr.end(), 0);

    for (int pages = low; pages <= high; pages++) {
        if (countStudents(arr, pages) == m) {
            return pages;
        }
    }
    return low;
}
// Time Complexity: O(N * (sum(arr[])-max(arr[])+1)), where N = size of the array, sum(arr[]) = sum of all array elements, max(arr[]) = maximum of all array elements.
// Reason: We are using a loop from max(arr[]) to sum(arr[]) to check all possible numbers of pages. Inside the loop, we are calling the countStudents() function for each number. Now, inside the countStudents() function, we are using a loop that runs for N times.

// Space Complexity:  O(1) as we are not using any extra space to solve this problem.

// optimized
int countStudents(vector<int> &arr, int pages) {
    int n = arr.size(); //size of array.
    int students = 1;
    long long pagesStudent = 0;
    for (int i = 0; i < n; i++) {
        if (pagesStudent + arr[i] <= pages) {
            //add pages to current student
            pagesStudent += arr[i];
        }
        else {
            //add pages to next student
            students++;
            pagesStudent = arr[i];
        }
    }
    return students;
}

int findPages(vector<int>& arr, int n, int m) {
    //book allocation impossible:
    if (m > n) return -1;

    int low = *max_element(arr.begin(), arr.end());
    int high = accumulate(arr.begin(), arr.end(), 0);
    while (low <= high) {
        int mid = (low + high) / 2;
        int students = countStudents(arr, mid);
        if (students > m) {
            low = mid + 1;
        }
        else {
            high = mid - 1;
        }
    }
    return low;
}
// Time Complexity: O(N * log(sum(arr[])-max(arr[])+1)), where N = size of the array, sum(arr[]) = sum of all array elements, max(arr[]) = maximum of all array elements.
// Reason: We are applying binary search on [max(arr[]), sum(arr[])]. Inside the loop, we are calling the countStudents() function for the value of ‘mid’. Now, inside the countStudents() function, we are using a loop that runs for N times.

// Space Complexity:  O(1) as we are not using any extra space to solve this problem.

int main()
{
    vector<int> arr = {25, 46, 28, 49, 24};
    int n = 5;
    int m = 4;
    int ans = findPages(arr, n, m);
    cout << "The answer is: " << ans << "\n";
    return 0;
}



// Time Complexity: O(N*log (M - max)), where N is the number of different books , max denotes the maximum number of pages from all the books and M denotes the sum of number of pages of all different books
// Auxiliary Space: O(1)