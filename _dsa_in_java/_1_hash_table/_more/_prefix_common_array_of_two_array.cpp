// https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/description/

class Solution {
public:
    vector<int> findThePrefixCommonArray(vector<int>& A, vector<int>& B) {
        
        unordered_map<int, bool> ma, mb;
        
        int s = A.size();
        vector<int> res;
         int cnt = 0;
        
        for(int i=0;i<s;i++) {
            ma[A[i]] = true;
            mb[B[i]] = true;
            
            if(A[i]==B[i]) cnt++;
            else {
                if(mb.find(A[i])!=mb.end()) cnt++;
                if(ma.find(B[i])!=ma.end()) cnt++;
            }
            res.push_back(cnt);
        }
        
        return res;
        
        
    }
};