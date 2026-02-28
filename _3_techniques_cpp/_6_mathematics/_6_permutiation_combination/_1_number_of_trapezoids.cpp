// https://leetcode.com/problems/count-number-of-trapezoids-i

class Solution {
public:
    typedef long long ll;
    int countTrapezoids(vector<vector<int>>& points) {
        unordered_map<int, int> mp;

        const int mod = 1e9 + 7;
        for(int i=0;i<points.size();i++) {
           mp[points[i][1]]++;
        }

        ll ans = 0, sum = 0;
        for(auto it : mp) {
            // no of ways to select 2 item out of item nC2 = (n*n-1)/2
             ll e = ((ll)it.second * (it.second - 1))/2;
             ans += e * sum;
             sum += e;
        }

        return ans % mod;

    }
};

// x1, x2, x3, x4, .........xn 
// x1 + (x2*x1) + (x3*x2 + x3* x1) + ...........
// x1 + (x2*x1) + x3(x2+x1) + .............