// https://leetcode.com/problems/find-the-highest-altitude/description/


class Solution {
public:
    int largestAltitude(vector<int>& gain) {
        int maxi = 0;
        int alt = 0;
        for(int i=0;i<gain.size();i++) {
          alt += gain[i];
          maxi=max(maxi, alt);
        }
        return maxi;
    }
};