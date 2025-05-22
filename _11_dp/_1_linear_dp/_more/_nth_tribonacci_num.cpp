// https://leetcode.com/problems/n-th-tribonacci-number/description/class Solution {
public:
    int tribonacci(int n) {
        if(n==0 || n==1)return n;
        if(n==2)return 1;
        int a=0,b=1,c=1,sum=0;
        for(int i=3;i<=n;i++)
        {
            sum=a+b+c; //sum of three
            a=b;   //move forward to next value
            b=c;  //mov for...
            c=sum;  //move for..
        }return sum;
    }
};