// https://leetcode.com/problems/assign-cookies/

#include <vector>
#include <algorithm>
using namespace std;

                            
#include <vector>   
#include <algorithm>
#include <iostream>

using namespace std;

// Function to find the maximum
// number of content children
int findContentChildren(vector<int>& greed, vector<int>& cookieSize) {
    // Get the size of
    // the greed array
    int n = greed.size();

    // Get the size of
    // the cookieSize array
    int m = cookieSize.size();

    // Sort the greed factors in ascending
    // order to try and satisfy the
    // least greedy children first
    sort(greed.begin(), greed.end());

    // Sort the cookie sizes in ascending
    // order to use the smallest cookies first
    sort(cookieSize.begin(), cookieSize.end());

    // Initialize a pointer for the
    // cookieSize array, starting
    // from the first cookie
    int l = 0;

    // Initialize a pointer for the
    // greed array, starting from
    // the first child
    int r = 0;

    // Iterate while there are
    // cookies and children
    // left to consider
    while (l < m && r < n) {
        // If the current cookie can
        // satisfy the current child's greed
        if (greed[r] <= cookieSize[l]) {
            // Move to the next child,
            // as the current child is satisfied
            r++;
        }
        // Always move to the next cookie
        // whether the current child
        // was satisfied or not
        l++;
    }

    // The value of r at the end of
    // the loop represents the number
    // of children that were satisfied
    return r;
}


int main() {
    vector<int> greed = {1, 5, 3, 3, 4};
    vector<int> cookieSize = {4, 2, 1, 2, 1, 3};
    
    cout << "Array Representing Greed: ";
    for(int i = 0; i < greed.size(); i++){
        cout << greed[i] << " ";
    }
    cout << endl;
    cout << "Array Representing Cookie Size: ";
    for(int i = 0; i < cookieSize.size(); i++){
        cout << cookieSize[i] << " ";
    }
    
    int ans = findContentChildren(greed, cookieSize);
    
    cout << endl << "No. of kids assigned cookies "<< ans;
    cout << endl;

    return 0;
}
                            
//  Time Complexity: O(N logN + M logM + M) where N is the length of the greedy array, M is the length of the cookies array. To sort the greedy and cookies array, the complexity is O(N logN) and O(M logM).
//  After sorting, we iterate over the arrays at most M times as M is the total number of cookies.
// Since each child and each cookie is considered at most once, the time complexity of this part is linear in terms of the size of the cookie array, which is O(M).
// Space Complexity: O(1) as the algorithm uses only a constant amount of extra space regardless of the size of the input array. It does not require any additional data structures that scale with the input size.