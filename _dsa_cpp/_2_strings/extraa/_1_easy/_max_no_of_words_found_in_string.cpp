// https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/description/

class Solution {
public:
    int mostWordsFound(vector<string>& sentences) {
        int n = sentences.size();
        int maxi = INT_MIN;
        for(int i=0;i<n;i++) {
           int spaces = 0;
           for(int j=0;j<sentences[i].size();j++) {
               if(isblank(sentences[i][j])) spaces++;
           }
           maxi = max(maxi, spaces+1);
        }

        return maxi;
    }
};