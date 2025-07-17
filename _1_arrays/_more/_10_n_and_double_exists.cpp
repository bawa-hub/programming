// https://leetcode.com/problems/check-if-n-and-its-double-exist/description/

class Solution {
public:
    bool checkIfExist(vector<int>& arr) {
        unordered_set<int> st;

        for(int i=0;i<arr.size();i++) {
            int num = arr[i];
            if(st.find(2*num) != st.end() || (num%2 == 0 && st.find(num/2) != st.end())) {
                return true;
            } else {
                st.insert(num);
            }
        }

        return false;
    }
};