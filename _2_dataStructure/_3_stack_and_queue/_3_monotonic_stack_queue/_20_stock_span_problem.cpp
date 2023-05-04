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
        while (!st.empty() && st.top().second <= price)
        {
            st.pop();
        }
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