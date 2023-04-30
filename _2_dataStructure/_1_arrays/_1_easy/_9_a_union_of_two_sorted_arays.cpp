// https://takeuforward.org/data-structure/union-of-two-sorted-arrays/

#include <bits/stdc++.h>

using namespace std;

// using map
// In unordered_map the keys are stored in random order, while in the map the keys are stored in sorted order (ascending order by default)
vector<int> FindUnion(int arr1[], int arr2[], int n, int m)
{
    map<int, int> freq;
    vector<int> Union;
    for (int i = 0; i < n; i++)
        freq[arr1[i]]++;
    for (int i = 0; i < m; i++)
        freq[arr2[i]]++;
    for (auto &it : freq)
        Union.push_back(it.first);
    return Union;
}
// Time Compleixty : O( (m+n)log(m+n) ) . Inserting a key in map takes logN times, where N is no of elements in map. At max map can store m+n elements {when there are no common elements and elements in arr,arr2 are distntict}. So Inserting m+n th element takes log(m+n) time. Upon approximation across insertion of all elements in worst it would take O((m+n)log(m+n) time.
// Using unordered_map also takes the same time, On average insertion in unordered_map takes O(1) time but sorting the union vector takes O((m+n)log(m+n))  time. Because at max union vector can have m+n elements.
// Space Complexity : O(m+n) {If Space of Union Vector is considered}
// O(1) {If Space of union Vector is not considered}

// using set
// In unordered_set the elements are stored in random order, while in a set the keys are stored in sorted order (ascending order by default)
vector<int> FindUnion(int arr1[], int arr2[], int n, int m)
{
    set<int> s;
    vector<int> Union;
    for (int i = 0; i < n; i++)
        s.insert(arr1[i]);
    for (int i = 0; i < m; i++)
        s.insert(arr2[i]);
    for (auto &it : s)
        Union.push_back(it);
    return Union;
}
// Time Compleixty : O( (m+n)log(m+n) ) . Inserting a element in set takes logN time, where N is no of elements in set. At max set can store m+n elements {when there are no common elements and elements in arr,arr2 are distntict}. So Inserting m+n th element takes log(m+n) time. Upon approximation across inserting all elements in worst it would take O((m+n)log(m+n) time.
// Using unordered_set also takes the same time, On average insertion in unordered_set takes O(1) time but sorting the union vector takes O((m+n)log(m+n))  time. Because at max union vector can have m+n elements.
// Space Complexity : O(m+n) {If Space of Union Vector is considered}
// O(1) {If Space of union Vector is not considered}

// using pointer (if arrays are sorted)
vector<int> FindUnion(int arr1[], int arr2[], int n, int m)
{
    int i = 0, j = 0;  // pointers
    vector<int> Union; // Uninon vector
    while (i < n && j < m)
    {
        if (arr1[i] <= arr2[j]) // Case 1 and 2
        {
            if (Union.size() == 0 || Union.back() != arr1[i])
                Union.push_back(arr1[i]);
            i++;
        }
        else // case 3
        {
            if (Union.size() == 0 || Union.back() != arr2[j])
                Union.push_back(arr2[j]);
            j++;
        }
    }
    while (i < n) // IF any element left in arr1
    {
        if (Union.back() != arr1[i])
            Union.push_back(arr1[i]);
        i++;
    }
    while (j < m) // If any elements left in arr2
    {
        if (Union.back() != arr2[j])
            Union.push_back(arr2[j]);
        j++;
    }
    return Union;
}
// Time Complexity: O(m+n), Because at max i runs for n times and j runs for m times. When there are no common elements in arr1 and arr2 and all elements in arr1, arr2 are distinct.
// Space Complexity : O(m+n) {If Space of Union Vector is considered}
// O(1) {If Space of union Vector is not considered}

int main()
{
    int n = 10, m = 7;
    int arr1[] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
    int arr2[] = {2, 3, 4, 4, 5, 11, 12};
    vector<int> Union = FindUnion(arr1, arr2, n, m);
    cout << "Union of arr1 and arr2 is " << endl;
    for (auto &val : Union)
        cout << val << " ";
    return 0;
}