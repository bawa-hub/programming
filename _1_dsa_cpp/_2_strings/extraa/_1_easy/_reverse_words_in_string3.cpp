// https://leetcode.com/problems/reverse-words-in-a-string-iii/description/

class Solution {
public:
    string reverseWords(string s) {
        vector<string> arr;

         string str = "";
        for(int i=0;i<s.size();i++) {
            if(isspace(s[i])) {
                arr.push_back(str);
                arr.push_back(" ");
                str = "";
            }
            else if(i == s.size()-1) {
                str+=s[i];
                arr.push_back(str);
            }
            else str+=s[i];
        }

        string res = "";

        for(int i=0;i<arr.size();i++) {
            if(arr[i]==" ") {
              res+=" ";
            } else {
              res+= rev(arr[i]);
            }
        }

        return res;
    }

    string rev(string &s) {
        int i=0,j=s.size()-1;

        while(i<=j) {
            swap(s[i], s[j]);
            i++;j--;
        }

        return s;
    }
};