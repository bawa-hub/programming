// https://leetcode.com/problems/jewels-and-stones/description/

class Solution {
public:
    int numJewelsInStones(string jewels, string stones) {
        vector<char> j(256, 0);
        for(int i=0;i<jewels.size();i++) {
            j[jewels[i]]++;
        }

        int res = 0;

        for(int k=0;k<stones.size();k++) {
            if(j[stones[k]]!=0) {
               res++;
            }
        }

        return res;
    }
};