// https://www.interviewbit.com/problems/nearest-smaller-element/

#include <bits/stdc++.h>
using namespace std;

vector<int> prevSmaller(vector<int> &A) {
    vector<int> res(A.size(), -1);
    stack<int> st;
    
    for(int i=0;i<A.size();i++) {
        while(!st.empty() && st.top() >= A[i]) {
            st.pop();
        }
        
        if(!st.empty()) res[i] = st.top();
        st.push(A[i]);
    }
    
    return res;
}

int main() {
    vector<int> v = {4, 5, 2, 10, 8};
    vector<int> ans = prevSmaller(v);
    for(int i=0;i<ans.size();i++) cout << ans[i] << " ";
    cout << endl;
}

// Input 1:
//     A = [4, 5, 2, 10, 8]
// Output 1:
//     G = [-1, 4, -1, 2, 2]