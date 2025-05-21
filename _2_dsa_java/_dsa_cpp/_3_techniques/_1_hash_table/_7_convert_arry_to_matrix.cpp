// https://leetcode.com/problems/convert-an-array-into-a-2d-array-with-conditions/description/

// brute force
vector<vector<int>> findMatrix(vector<int>& nums) {
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
    res.push_back(temp);
        if(mp.size() == 0) break;
        }

        return res;

    }