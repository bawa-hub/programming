// https://practice.geeksforgeeks.org/problems/first-negative-integer-in-every-window-of-size-k3345/1

class Solution {
  public:
    vector<int> FirstNegativeInteger(vector<int>& arr, int k) {
        int i=0,j=0;
        queue<int> q;
        vector<int> res;
        
        while(j<arr.size()) {
            if(arr[j]<0) q.push(arr[j]);
            if(j-i+1==k) {
                if(q.empty()) res.push_back(0);
                else res.push_back(q.front());
                
                if(arr[i++]<0) q.pop();
             }
            j++;
        }
        
        return res;
    }
};