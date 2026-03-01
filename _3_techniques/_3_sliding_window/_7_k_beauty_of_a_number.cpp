// https://leetcode.com/problems/find-the-k-beauty-of-a-number/

class Solution {
public:
    int divisorSubstrings(int num, int k) {
        string s = to_string(num);

       int i=0,j=0;
       int n = s.size();
       int cnt = 0;
        while(j<n) {
            if(j-i+1==k) {
               int curr = stoi(s.substr(i, k));
               if(curr !=0 && num%curr==0) cnt++;
               i++;
            }
            j++;
        }
        return cnt;
    }
};