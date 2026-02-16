// https://leetcode.com/problems/lemonade-change/description/

class Solution {
public:
    bool lemonadeChange(vector<int>& bills) {
        int fv = 0, tn = 0, tw = 0;
        for(int i=0;i<bills.size();i++) {
           if(bills[i] == 5) fv++;
           else if(bills[i] == 10) {
            if(fv>0) {
                fv--;tn++;
            } else return false;
           } else { 
             if(tn > 0 && fv > 0) {
                tn--;fv--;
             } else if(fv >=3) {
                fv -= 3;
             } else return false;
           }
        }

        return true;
    }
};