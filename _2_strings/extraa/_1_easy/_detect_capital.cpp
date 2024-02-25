// https://leetcode.com/problems/detect-capital/description/

class Solution {
public:
    bool detectCapitalUse(string word) {
        int cnt = 0;
        bool res = false;
        for(int i=0;i<word.size();i++) {
            if(word[i]>=65&&word[i]<=90) cnt++;
        }

        if(cnt==word.size()||cnt==0||(cnt==1 && (word[0]>=65&&word[0]<=90)) )  res = true;

        return res;
    }
};