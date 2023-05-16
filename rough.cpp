#include <iostream>
#include <vector>
#include <stack>

using namespace std;

vector<int> span(vector<int> arr) {
    int n = arr.size();
    stack<pair<int, int> > st;
    vector<int> v(n);

    for(int i=0;i<n;i++) {
        while(!st.empty() && st.top().first <= arr[i]) st.pop();
        if(!st.empty()) v.push_back(st.top().second);
        else v.push_back(-1);
        st.push(make_pair(arr[i], i));
    }

    for(int i=0;i<n;i++) {
        v[i] = i-v[i];
    }

    return v;

}

int main()
{
    vector<int> v {100, 80, 60,70,75,85};

    vector<int> res = span(v);

    for(int i=0;i<res.size();i++) {
        cout << res[i] << " ";
    }
}