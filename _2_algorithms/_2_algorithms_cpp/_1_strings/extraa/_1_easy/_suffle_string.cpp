// https://leetcode.com/problems/shuffle-string/description/

class Solution {

public:

    string restoreString(string s, vector<int>& indices) {

        string res(s.size(), '#');

        int j=0;

        for(int i=0;i<indices.size();i++) {

            res[indices[i]] = s[j++];

        }



        return res;

    }

};