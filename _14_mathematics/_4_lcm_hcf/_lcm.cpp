// https://leetcode.com/problems/smallest-even-multiple/

class Solution {
public:
    int smallestEvenMultiple(int n) {
        if(n==1 || n==2) return 2;
        
        int i=1;
        while(true) {
            if((n*i)%2==0) return n*i;    
            i++;
        }
    return 0;
    }
};