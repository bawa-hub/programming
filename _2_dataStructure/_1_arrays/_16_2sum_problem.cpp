// https://leetcode.com/problems/two-sum/

// brute force
// vector<int> twoSum(vector<int> &nums, int target)
// {
//     vector<int> res;
//     for (int i = 0; i < nums.size(); ++i)
//     {
//         for (int j = i + 1; j < nums.size(); ++j)
//         {
//             if (nums[i] + nums[j] == target)
//             {
//                 res.emplace_back(i);
//                 res.emplace_back(j);
//                 break;
//             }
//         }
//         if (res.size() == 2)
//             break;
//     }
//     return res;
// }

// TC: O(n^2)
// SC: O(1)

// two pointer approach
// vector<int> twoSum(vector<int> &nums, int target)
// {

//     vector<int> res, store;
//     store = nums;

//     sort(store.begin(), store.end());

//     int left = 0, right = nums.size() - 1;
//     int n1, n2;

//     while (left < right)
//     {
//         if (store[left] + store[right] == target)
//         {

//             n1 = store[left];
//             n2 = store[right];

//             break;
//         }
//         else if (store[left] + store[right] > target)
//             right--;
//         else
//             left++;
//     }

//     for (int i = 0; i < nums.size(); ++i)
//     {

//         if (nums[i] == n1)
//             res.emplace_back(i);
//         else if (nums[i] == n2)
//             res.emplace_back(i);
//     }

//     return res;
// }

// TC: O(nlogn)
// SC: O(n)

// hashing (most efficient)
vector<int> twoSum(vector<int> &nums, int target)
{

    vector<int> res;
    unordered_map<int, int> mp;

    for (int i = 0; i < nums.size(); ++i)
    {

        if (mp.find(target - nums[i]) != mp.end())
        {

            res.emplace_back(i);
            res.emplace_back(mp[target - nums[i]]);
            return res;
        }

        mp[nums[i]] = i;
    }

    return res;
}

// TC: O(n)
// SC: O(n)