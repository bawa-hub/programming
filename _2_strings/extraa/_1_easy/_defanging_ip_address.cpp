// https://leetcode.com/problems/defanging-an-ip-address/description/

class Solution {
public:
    string defangIPaddr(string address) {
        int n = address.size();
        string res = "";

        for(int i=0;i<n;i++) {
            if(address[i]=='.') {
                res+="[.]";
            } else res+=address[i];
        }

        return res;
    }
};