// https://leetcode.com/problems/minimum-recolors-to-get-k-consecutive-black-blocks/

class Solution {
public:
    int minimumRecolors(string blocks, int k) {
        int i=0,j=0,len = blocks.size(),mini = 1000,cnt = 0;

        while(j<len) {
            if(blocks[j]=='W') cnt++;
            if(j-i+1==k) {
              mini = min(mini, cnt);
              if(blocks[i]=='W') cnt--;
              i++;
            }
            j++;
        }

        return mini;

    }
};
