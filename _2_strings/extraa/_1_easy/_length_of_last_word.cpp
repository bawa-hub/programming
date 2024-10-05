// https://leetcode.com/problems/length-of-last-word/description/

class Solution {
public:
    int lengthOfLastWord(string s) {
        int n=s.size();
        int i=n-1;
        while(isblank(s[i])) {
            i--;
        }
        int cnt=0;
        while(i>=0) {
            if(isblank(s[i])) break;
            else cnt++;
            i--;
        }
        return cnt;
    }
};