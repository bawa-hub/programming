// https://leetcode.com/problems/maximum-running-time-of-n-computers

class Solution {
public:
    typedef long long ll;

    bool possible(vector<int> batteries, ll mid, int n) {
        ll target = n * mid;
        ll sum = 0;
        for(int i=0;i<batteries.size();i++) {
            sum += min((ll)batteries[i], mid);
            if(sum >= target) return true;
        }

        return false;
    }


    long long maxRunTime(int n, vector<int>& batteries) {
        ll l = *min_element(batteries.begin(), batteries.end());

        ll sum = 0;
        for(auto &m : batteries) {
            sum += m;
        }

        ll r = sum/n;

        ll result = 0;

        while(l<=r) {
            ll mid = l + (r-l)/2;
            if(possible(batteries, mid, n)) {
                result = mid;
                l = mid + 1;
            } else {
                r = mid - 1;
            }
        }

        return result;
    }

    // TC: m * log(k)

    // greedy approach
        long long maxRunTime(int n, vector<int>& arr) {
        sort(arr.begin(), arr.end());
        long long total = accumulate(arr.begin(), arr.end(), 0LL);

        for (int i = arr.size() - 1; i >= 0; i--) {
            if (arr[i] <= total / n) break;
            total -= arr[i];
            n--;
        }

        return total / n;
    }
};