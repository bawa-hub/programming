// https://leetcode.com/problems/online-stock-span/
// https://practice.geeksforgeeks.org/problems/stock-span-problem-1587115621/1
// https://www.geeksforgeeks.org/the-stock-span-problem/

class StockSpanner
{
    stack<pair<int, int>> st;
    int index = -1;

public:
    StockSpanner()
    {
    }

    int next(int price)
    {
        index += 1;
        while (!st.empty() && st.top().second <= price) st.pop();

        if (st.empty())
        {
            st.push({index, price});
            return index + 1;
        }
        
        int result = st.top().first;
        st.push({index, price});
        return index - result;
    }
};

//   Time complexity: O(N)
//  Space complexity:O(N)


// next greater element to left approach (by me)
vector<int> span(vector<int> arr) {
    int n = arr.size();
    stack<pair<int, int> > st;
    vector<int> v;

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
    vector<int> v {100, 80, 60, 70, 60, 75, 85};
    vector<int> res = span(v);
    for(int i=0;i<res.size();i++) {
        cout << res[i] << " ";
    }
}