// https://leetcode.com/problems/find-words-containing-character/description/

class Solution {
public:
    vector<int> findWordsContaining(vector<string>& words, char x) {
        vector<int> ans;
        for(int i=0;i<words.size();i++) {
           if(contain(x, words[i])) ans.push_back(i);
        }

        return ans;
    }

    bool contain(char ch, string str) {
      for(int i=0;i<str.size();i++) {
          if(str[i] == ch) return true;
      }

      return false;
    }
};