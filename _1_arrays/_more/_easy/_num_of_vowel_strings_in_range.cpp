// https://leetcode.com/problems/count-the-number-of-vowel-strings-in-range/description/

class Solution {
public:
    int vowelStrings(vector<string>& words, int left, int right) {
        int cnt = 0;
        for(int i=left;i<=right;i++) {
            if(isVowel(words[i])) cnt++;
         }
        return cnt;
    }
    
    bool isVowel(string s) {
        int i=0;
        int j = s.length() -1;
        if((s[i]=='a'||s[i]=='e'||s[i]=='i'||s[i]=='o'||s[i]=='u')&&(s[j]=='a'||s[j]=='e'||s[j]=='i'||s[j]=='o'||s[j]=='u')) {
            return true;
        }
        return false;
    }
};