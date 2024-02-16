#include <bits/stdc++.h>
using namespace std;



vector<vector<int>> convert(vector<int> nums) {
    unordered_map<int, int> mp;;

        for(int i=0;i<nums.size();i++) mp[nums[i]]++;

        vector<vector<int>> res;

        while(true) {
            vector<int> temp;
            vector<int> keys;
        for (auto& pair : mp) {
            temp.push_back(pair.first);
           pair.second--;
        if (pair.second == 0) {
            keys.push_back(pair.first);
        }
    }

    for(int key : keys) {
        mp.erase(key);
    }
        if(mp.size() == 0) break;
        }

        return res;
}

int main() {

    vector<int> nums = {1,3,4,1,2,3,1};
    return convert(nums);

}