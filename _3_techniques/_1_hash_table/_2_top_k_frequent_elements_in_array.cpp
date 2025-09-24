// https://practice.geeksforgeeks.org/problems/top-k-frequent-elements-in-array/1
// https://leetcode.com/problems/top-k-frequent-elements/

#include <bits/stdc++.h>
using namespace std;

// using max heap
class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        unordered_map<int, int> mp;
        for(int i=0;i<nums.size();i++) {
              mp[nums[i]]++;
        }

     priority_queue<pair<int, int>> pq;
        for(auto i: mp) {
             pq.push({i.second, i.first});
        }

        vector<int> res;
        while(k-->0) {
            res.push_back(pq.top().second);
            pq.pop();
        }

        return res;
    }
};

//  min-heap optimization approach (keep heap size = k)
class Solution
{
public:
    vector<int> topKFrequent(vector<int> &nums, int k)
    {
        unordered_map<int, int> mp;
        for (int i = 0; i < nums.size(); i++)
            mp[nums[i]]++;

        //    min heap
        priority_queue<pair<int, int>, vector<pair<int, int>>, greater<pair<int, int>>> pq;

        for (auto i : mp)
        {
            pq.push({i.second, i.first});
            if (pq.size() > k)
                pq.pop();
        }

        vector<int> res;
        while (pq.size() > 0)
        {
            res.push_back(pq.top().second);
            pq.pop();
        }
        return res;
    }
};


// optimal O(n) bucket sort

// vector<int> topKFrequent(vector<int>& nums, int k) {
//     unordered_map<int, int> freq;
//     for (int num : nums) {
//         freq[num]++;
//     }

//     // bucket[i] contains numbers that appear exactly i times
//     vector<vector<int>> bucket(nums.size() + 1);
//     for (auto& p : freq) {
//         bucket[p.second].push_back(p.first);
//     }

//     vector<int> res;
//     for (int i = nums.size(); i >= 0 && res.size() < k; i--) {
//         for (int num : bucket[i]) {
//             res.push_back(num);
//             if (res.size() == k) break;
//         }
//     }

//     return res;
// }

// int main() {
//     vector<int> nums = {1, 1, 1, 2, 2, 3};
//     int k = 2;
//     vector<int> res = topKFrequent(nums, k);
//     for (int x : res) cout << x << " ";
//     return 0;
// }
