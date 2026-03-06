/*
 * LeetCode 100 Medium + Hard Problems - Complete Solutions
 * Organized by category for practice and revision
 */

 #include <algorithm>
 #include <array>
 #include <climits>
 #include <cmath>
 #include <functional>
 #include <numeric>
 #include <queue>
 #include <sstream>
 #include <stack>
 #include <string>
 #include <tuple>
 #include <unordered_map>
 #include <unordered_set>
 #include <utility>
 #include <vector>
 using namespace std;

// LeetCode-provided interface for problem 1095 (Find in Mountain Array).
// Forward-declared here so this file can compile locally.
class MountainArray {
public:
    int get(int index);
    int length();
};
 
 // ============================================================================
 // A) ARRAYS / TWO POINTERS / SLIDING WINDOW (20 problems)
 // ============================================================================
 
 // 1. 3Sum (15)
 class Solution1 {
 public:
     vector<vector<int>> threeSum(vector<int>& nums) {
         sort(nums.begin(), nums.end());
         vector<vector<int>> res;
         int n = nums.size();
         
         for (int i = 0; i < n - 2; i++) {
             if (i > 0 && nums[i] == nums[i-1]) continue;
             
             int l = i + 1, r = n - 1;
             while (l < r) {
                 int sum = nums[i] + nums[l] + nums[r];
                 if (sum == 0) {
                     res.push_back({nums[i], nums[l], nums[r]});
                     while (l < r && nums[l] == nums[l+1]) l++;
                     while (l < r && nums[r] == nums[r-1]) r--;
                     l++; r--;
                 } else if (sum < 0) l++;
                 else r--;
             }
         }
         return res;
     }
 };
 
 // 2. 3Sum Closest (16)
 class Solution2 {
 public:
     int threeSumClosest(vector<int>& nums, int target) {
         sort(nums.begin(), nums.end());
         int n = nums.size();
         int closest = nums[0] + nums[1] + nums[2];
         
         for (int i = 0; i < n - 2; i++) {
             int l = i + 1, r = n - 1;
             while (l < r) {
                 int sum = nums[i] + nums[l] + nums[r];
                 if (abs(sum - target) < abs(closest - target)) {
                     closest = sum;
                 }
                 if (sum < target) l++;
                 else r--;
             }
         }
         return closest;
     }
 };
 
 // 3. Container With Most Water (11)
 class Solution3 {
 public:
     int maxArea(vector<int>& height) {
         int l = 0, r = height.size() - 1;
         int maxWater = 0;
         
         while (l < r) {
             int area = min(height[l], height[r]) * (r - l);
             maxWater = max(maxWater, area);
             if (height[l] < height[r]) l++;
             else r--;
         }
         return maxWater;
     }
 };
 
 // 4. Sort Colors (75) - Dutch National Flag
 class Solution4 {
 public:
     void sortColors(vector<int>& nums) {
         int low = 0, mid = 0, high = nums.size() - 1;
         
         while (mid <= high) {
             if (nums[mid] == 0) {
                 swap(nums[low], nums[mid]);
                 low++; mid++;
             } else if (nums[mid] == 1) {
                 mid++;
             } else {
                 swap(nums[mid], nums[high]);
                 high--;
             }
         }
     }
 };
 
 // 5. Subarray Sum Equals K (560)
 class Solution5 {
 public:
     int subarraySum(vector<int>& nums, int k) {
         unordered_map<int, int> count;
         count[0] = 1;
         int sum = 0, ans = 0;
         
         for (int num : nums) {
             sum += num;
             ans += count[sum - k];
             count[sum]++;
         }
         return ans;
     }
 };
 
 // 6. Continuous Subarray Sum (523)
 class Solution6 {
 public:
     bool checkSubarraySum(vector<int>& nums, int k) {
         unordered_map<int, int> modIndex;
         modIndex[0] = -1;
         int sum = 0;
         
         for (int i = 0; i < nums.size(); i++) {
             sum += nums[i];
             int mod = sum % k;
             if (modIndex.count(mod)) {
                 if (i - modIndex[mod] >= 2) return true;
             } else {
                 modIndex[mod] = i;
             }
         }
         return false;
     }
 };
 
 // 7. Minimum Size Subarray Sum (209)
 class Solution7 {
 public:
     int minSubArrayLen(int target, vector<int>& nums) {
         int l = 0, sum = 0, minLen = INT_MAX;
         
         for (int r = 0; r < nums.size(); r++) {
             sum += nums[r];
             while (sum >= target) {
                 minLen = min(minLen, r - l + 1);
                 sum -= nums[l++];
             }
         }
         return minLen == INT_MAX ? 0 : minLen;
     }
 };
 
 // 8. Minimum Window Substring (76)
 class Solution8 {
 public:
     string minWindow(string s, string t) {
         unordered_map<char, int> need;
         for (char c : t) need[c]++;
         
         int l = 0, r = 0, valid = 0;
         int start = 0, len = INT_MAX;
         unordered_map<char, int> window;
         
         while (r < s.size()) {
             char c = s[r++];
             if (need.count(c)) {
                 window[c]++;
                 if (window[c] == need[c]) valid++;
             }
             
             while (valid == need.size()) {
                 if (r - l < len) {
                     start = l;
                     len = r - l;
                 }
                 char d = s[l++];
                 if (need.count(d)) {
                     if (window[d] == need[d]) valid--;
                     window[d]--;
                 }
             }
         }
         return len == INT_MAX ? "" : s.substr(start, len);
     }
 };
 
 // 9. Longest Substring Without Repeating Characters (3)
 class Solution9 {
 public:
     int lengthOfLongestSubstring(string s) {
         unordered_map<char, int> lastPos;
         int l = 0, maxLen = 0;
         
         for (int r = 0; r < s.size(); r++) {
             if (lastPos.count(s[r])) {
                 l = max(l, lastPos[s[r]] + 1);
             }
             lastPos[s[r]] = r;
             maxLen = max(maxLen, r - l + 1);
         }
         return maxLen;
     }
 };
 
 // 10. Longest Repeating Character Replacement (424)
 class Solution10 {
 public:
     int characterReplacement(string s, int k) {
         vector<int> count(26, 0);
         int l = 0, maxFreq = 0, maxLen = 0;
         
         for (int r = 0; r < s.size(); r++) {
             count[s[r] - 'A']++;
             maxFreq = max(maxFreq, count[s[r] - 'A']);
             
             if (r - l + 1 - maxFreq > k) {
                 count[s[l] - 'A']--;
                 l++;
             }
             maxLen = max(maxLen, r - l + 1);
         }
         return maxLen;
     }
 };
 
 // 11. Permutation in String (567)
 class Solution11 {
 public:
     bool checkInclusion(string s1, string s2) {
         if (s1.size() > s2.size()) return false;
         
         vector<int> count1(26, 0), count2(26, 0);
         for (int i = 0; i < s1.size(); i++) {
             count1[s1[i] - 'a']++;
             count2[s2[i] - 'a']++;
         }
         
         int matches = 0;
         for (int i = 0; i < 26; i++) {
             if (count1[i] == count2[i]) matches++;
         }
         
         for (int i = s1.size(); i < s2.size(); i++) {
             if (matches == 26) return true;
             
             int r = s2[i] - 'a';
             count2[r]++;
             if (count1[r] == count2[r]) matches++;
             else if (count1[r] + 1 == count2[r]) matches--;
             
             int l = s2[i - s1.size()] - 'a';
             count2[l]--;
             if (count1[l] == count2[l]) matches++;
             else if (count1[l] - 1 == count2[l]) matches--;
         }
         return matches == 26;
     }
 };
 
 // 12. Find All Anagrams in a String (438)
 class Solution12 {
 public:
     vector<int> findAnagrams(string s, string p) {
         if (p.size() > s.size()) return {};
         
         vector<int> res;
         vector<int> countP(26, 0), countS(26, 0);
         
         for (int i = 0; i < p.size(); i++) {
             countP[p[i] - 'a']++;
             countS[s[i] - 'a']++;
         }
         
         if (countP == countS) res.push_back(0);
         
         for (int i = p.size(); i < s.size(); i++) {
             countS[s[i] - 'a']++;
             countS[s[i - p.size()] - 'a']--;
             if (countP == countS) res.push_back(i - p.size() + 1);
         }
         return res;
     }
 };
 
 // 13. Max Consecutive Ones III (1004)
 class Solution13 {
 public:
     int longestOnes(vector<int>& nums, int k) {
         int l = 0, zeros = 0, maxLen = 0;
         
         for (int r = 0; r < nums.size(); r++) {
             if (nums[r] == 0) zeros++;
             while (zeros > k) {
                 if (nums[l] == 0) zeros--;
                 l++;
             }
             maxLen = max(maxLen, r - l + 1);
         }
         return maxLen;
     }
 };
 
 // 14. Fruit Into Baskets (904)
 class Solution14 {
 public:
     int totalFruit(vector<int>& fruits) {
         unordered_map<int, int> count;
         int l = 0, maxLen = 0;
         
         for (int r = 0; r < fruits.size(); r++) {
             count[fruits[r]]++;
             while (count.size() > 2) {
                 count[fruits[l]]--;
                 if (count[fruits[l]] == 0) count.erase(fruits[l]);
                 l++;
             }
             maxLen = max(maxLen, r - l + 1);
         }
         return maxLen;
     }
 };
 
 // 15. Trapping Rain Water (42)
 class Solution15 {
 public:
     int trap(vector<int>& height) {
         int l = 0, r = height.size() - 1;
         int lMax = 0, rMax = 0, water = 0;
         
         while (l < r) {
             if (height[l] < height[r]) {
                 if (height[l] >= lMax) lMax = height[l];
                 else water += lMax - height[l];
                 l++;
             } else {
                 if (height[r] >= rMax) rMax = height[r];
                 else water += rMax - height[r];
                 r--;
             }
         }
         return water;
     }
 };
 
 // 16. Next Permutation (31)
 class Solution16 {
 public:
     void nextPermutation(vector<int>& nums) {
         int n = nums.size();
         int i = n - 2;
         while (i >= 0 && nums[i] >= nums[i+1]) i--;
         
         if (i >= 0) {
             int j = n - 1;
             while (nums[j] <= nums[i]) j--;
             swap(nums[i], nums[j]);
         }
         reverse(nums.begin() + i + 1, nums.end());
     }
 };
 
 // 17. Rotate Array (189)
 class Solution17 {
 public:
     void rotate(vector<int>& nums, int k) {
         int n = nums.size();
         k %= n;
         reverse(nums.begin(), nums.end());
         reverse(nums.begin(), nums.begin() + k);
         reverse(nums.begin() + k, nums.end());
     }
 };
 
 // 18. Find the Duplicate Number (287) - Floyd Cycle
 class Solution18 {
 public:
     int findDuplicate(vector<int>& nums) {
         int slow = nums[0], fast = nums[0];
         
         do {
             slow = nums[slow];
             fast = nums[nums[fast]];
         } while (slow != fast);
         
         slow = nums[0];
         while (slow != fast) {
             slow = nums[slow];
             fast = nums[fast];
         }
         return slow;
     }
 };
 
 // 19. First Missing Positive (41)
 class Solution19 {
 public:
     int firstMissingPositive(vector<int>& nums) {
         int n = nums.size();
         for (int i = 0; i < n; i++) {
             while (nums[i] > 0 && nums[i] <= n && nums[nums[i] - 1] != nums[i]) {
                 swap(nums[i], nums[nums[i] - 1]);
             }
         }
         
         for (int i = 0; i < n; i++) {
             if (nums[i] != i + 1) return i + 1;
         }
         return n + 1;
     }
 };
 
 // 20. Jump Game (55)
 class Solution20 {
 public:
     bool canJump(vector<int>& nums) {
         int farthest = 0;
         for (int i = 0; i < nums.size(); i++) {
             if (i > farthest) return false;
             farthest = max(farthest, i + nums[i]);
         }
         return true;
     }
 };
 
 // ============================================================================
 // B) HASHING / FREQUENCY / PREFIX TRICKS (10 problems)
 // ============================================================================
 
 // 21. Group Anagrams (49)
 class Solution21 {
 public:
     vector<vector<string>> groupAnagrams(vector<string>& strs) {
         unordered_map<string, vector<string>> groups;
         
         for (string& s : strs) {
             string key = s;
             sort(key.begin(), key.end());
             groups[key].push_back(s);
         }
         
         vector<vector<string>> res;
         for (auto& [k, v] : groups) {
             res.push_back(v);
         }
         return res;
     }
 };
 
 // 22. Top K Frequent Elements (347)
 class Solution22 {
 public:
     vector<int> topKFrequent(vector<int>& nums, int k) {
         unordered_map<int, int> freq;
         for (int num : nums) freq[num]++;
         
         vector<vector<int>> buckets(nums.size() + 1);
         for (auto& [num, count] : freq) {
             buckets[count].push_back(num);
         }
         
         vector<int> res;
         for (int i = buckets.size() - 1; i >= 0 && res.size() < k; i--) {
             for (int num : buckets[i]) {
                 res.push_back(num);
                 if (res.size() == k) break;
             }
         }
         return res;
     }
 };
 
 // 23. Longest Consecutive Sequence (128)
 class Solution23 {
 public:
     int longestConsecutive(vector<int>& nums) {
         unordered_set<int> s(nums.begin(), nums.end());
         int maxLen = 0;
         
         for (int num : s) {
             if (!s.count(num - 1)) {
                 int len = 1;
                 while (s.count(num + len)) len++;
                 maxLen = max(maxLen, len);
             }
         }
         return maxLen;
     }
 };
 
 // 24. Valid Sudoku (36)
 class Solution24 {
 public:
     bool isValidSudoku(vector<vector<char>>& board) {
         vector<unordered_set<char>> rows(9), cols(9), boxes(9);
         
         for (int i = 0; i < 9; i++) {
             for (int j = 0; j < 9; j++) {
                 if (board[i][j] == '.') continue;
                 
                 char c = board[i][j];
                 int box = (i / 3) * 3 + j / 3;
                 
                 if (rows[i].count(c) || cols[j].count(c) || boxes[box].count(c)) {
                     return false;
                 }
                 rows[i].insert(c);
                 cols[j].insert(c);
                 boxes[box].insert(c);
             }
         }
         return true;
     }
 };
 
 // 25. Insert Delete GetRandom O(1) (380)
 class RandomizedSet {
     vector<int> nums;
     unordered_map<int, int> valToIdx;
     
 public:
     RandomizedSet() {}
     
     bool insert(int val) {
         if (valToIdx.count(val)) return false;
         valToIdx[val] = nums.size();
         nums.push_back(val);
         return true;
     }
     
     bool remove(int val) {
         if (!valToIdx.count(val)) return false;
         int idx = valToIdx[val];
         valToIdx[nums.back()] = idx;
         swap(nums[idx], nums.back());
         nums.pop_back();
         valToIdx.erase(val);
         return true;
     }
     
     int getRandom() {
         return nums[rand() % nums.size()];
     }
 };
 
 // 26. Random Pick with Weight (528)
 class Solution26 {
     vector<int> prefix;
     
 public:
     Solution26(vector<int>& w) {
         prefix.push_back(w[0]);
         for (int i = 1; i < w.size(); i++) {
             prefix.push_back(prefix.back() + w[i]);
         }
     }
     
     int pickIndex() {
         int r = rand() % prefix.back() + 1;
         return lower_bound(prefix.begin(), prefix.end(), r) - prefix.begin();
     }
 };
 
 // 27. Find All Duplicates in an Array (442)
 class Solution27 {
 public:
     vector<int> findDuplicates(vector<int>& nums) {
         vector<int> res;
         for (int num : nums) {
             int idx = abs(num) - 1;
             if (nums[idx] < 0) res.push_back(abs(num));
             else nums[idx] = -nums[idx];
         }
         return res;
     }
 };
 
 // 28. Subarray Sums Divisible by K (974)
 class Solution28 {
 public:
     int subarraysDivByK(vector<int>& nums, int k) {
         unordered_map<int, int> count;
         count[0] = 1;
         int sum = 0, ans = 0;
         
         for (int num : nums) {
             sum += num;
             int mod = ((sum % k) + k) % k;
             ans += count[mod];
             count[mod]++;
         }
         return ans;
     }
 };
 
 // 29. Longest Duplicate Substring (1044) - Binary Search + Rolling Hash
 class Solution29 {
     const int base = 26;
     const int mod = 1e9 + 7;
     
     string check(string& s, int len) {
         unordered_map<long long, vector<int>> seen;
         long long hash = 0, power = 1;
         
         for (int i = 0; i < len; i++) {
             hash = (hash * base + (s[i] - 'a')) % mod;
             if (i < len - 1) power = (power * base) % mod;
         }
         seen[hash].push_back(0);
         
         for (int i = len; i < s.size(); i++) {
             hash = ((hash - (s[i - len] - 'a') * power) % mod + mod) % mod;
             hash = (hash * base + (s[i] - 'a')) % mod;
             
             if (seen.count(hash)) {
                 string curr = s.substr(i - len + 1, len);
                 for (int start : seen[hash]) {
                     if (s.substr(start, len) == curr) {
                         return curr;
                     }
                 }
             }
             seen[hash].push_back(i - len + 1);
         }
         return "";
     }
     
 public:
     string longestDupSubstring(string s) {
         int l = 1, r = s.size() - 1;
         string res = "";
         
         while (l <= r) {
             int mid = l + (r - l) / 2;
             string candidate = check(s, mid);
             if (!candidate.empty()) {
                 res = candidate;
                 l = mid + 1;
             } else {
                 r = mid - 1;
             }
         }
         return res;
     }
 };
 
 // 30. Substring with Concatenation of All Words (30)
 class Solution30 {
 public:
     vector<int> findSubstring(string s, vector<string>& words) {
         int wordLen = words[0].size();
         int totalLen = wordLen * words.size();
         unordered_map<string, int> wordCount;
         for (string& w : words) wordCount[w]++;
         
         vector<int> res;
         for (int offset = 0; offset < wordLen; offset++) {
             unordered_map<string, int> seen;
             int l = offset, valid = 0;
             
             for (int r = offset; r + wordLen <= s.size(); r += wordLen) {
                 string word = s.substr(r, wordLen);
                 
                 if (!wordCount.count(word)) {
                     seen.clear();
                     valid = 0;
                     l = r + wordLen;
                     continue;
                 }
                 
                 seen[word]++;
                 if (seen[word] <= wordCount[word]) valid++;
                 
                 while (seen[word] > wordCount[word]) {
                     string leftWord = s.substr(l, wordLen);
                     seen[leftWord]--;
                     if (seen[leftWord] < wordCount[leftWord]) valid--;
                     l += wordLen;
                 }
                 
                 if (valid == words.size()) {
                     res.push_back(l);
                 }
             }
         }
         return res;
     }
 };
 
 // ============================================================================
 // C) STACKS / MONOTONIC STACK / INTERVALS (10 problems)
 // ============================================================================
 
 // 31. Daily Temperatures (739)
 class Solution31 {
 public:
     vector<int> dailyTemperatures(vector<int>& temperatures) {
         stack<int> st;
         vector<int> res(temperatures.size(), 0);
         
         for (int i = 0; i < temperatures.size(); i++) {
             while (!st.empty() && temperatures[i] > temperatures[st.top()]) {
                 res[st.top()] = i - st.top();
                 st.pop();
             }
             st.push(i);
         }
         return res;
     }
 };
 
 // 32. Next Greater Element II (503)
 class Solution32 {
 public:
     vector<int> nextGreaterElements(vector<int>& nums) {
         int n = nums.size();
         vector<int> res(n, -1);
         stack<int> st;
         
         for (int i = 0; i < 2 * n; i++) {
             int idx = i % n;
             while (!st.empty() && nums[idx] > nums[st.top()]) {
                 res[st.top()] = nums[idx];
                 st.pop();
             }
             if (i < n) st.push(idx);
         }
         return res;
     }
 };
 
 // 33. Evaluate Reverse Polish Notation (150)
 class Solution33 {
 public:
     int evalRPN(vector<string>& tokens) {
         stack<int> st;
         
         for (string& token : tokens) {
             if (token == "+" || token == "-" || token == "*" || token == "/") {
                 int b = st.top(); st.pop();
                 int a = st.top(); st.pop();
                 if (token == "+") st.push(a + b);
                 else if (token == "-") st.push(a - b);
                 else if (token == "*") st.push(a * b);
                 else st.push(a / b);
             } else {
                 st.push(stoi(token));
             }
         }
         return st.top();
     }
 };
 
 // 34. Decode String (394)
 class Solution34 {
 public:
     string decodeString(string s) {
         stack<pair<string, int>> st;
         string curr = "";
         int num = 0;
         
         for (char c : s) {
             if (isdigit(c)) {
                 num = num * 10 + (c - '0');
             } else if (c == '[') {
                 st.push({curr, num});
                 curr = "";
                 num = 0;
             } else if (c == ']') {
                 auto [prev, k] = st.top(); st.pop();
                 string temp = curr;
                 curr = prev;
                 while (k--) curr += temp;
             } else {
                 curr += c;
             }
         }
         return curr;
     }
 };
 
 // 35. Largest Rectangle in Histogram (84)
 class Solution35 {
 public:
     int largestRectangleArea(vector<int>& heights) {
         stack<int> st;
         int maxArea = 0;
         
         for (int i = 0; i <= heights.size(); i++) {
             int h = (i == heights.size()) ? 0 : heights[i];
             while (!st.empty() && heights[st.top()] > h) {
                 int height = heights[st.top()];
                 st.pop();
                 int width = st.empty() ? i : i - st.top() - 1;
                 maxArea = max(maxArea, height * width);
             }
             st.push(i);
         }
         return maxArea;
     }
 };
 
 // 36. Maximal Rectangle (85)
 class Solution36 {
 public:
     int maximalRectangle(vector<vector<char>>& matrix) {
         if (matrix.empty()) return 0;
         int rows = matrix.size(), cols = matrix[0].size();
         vector<int> heights(cols, 0);
         int maxArea = 0;
         
         for (int i = 0; i < rows; i++) {
             for (int j = 0; j < cols; j++) {
                 heights[j] = (matrix[i][j] == '1') ? heights[j] + 1 : 0;
             }
             maxArea = max(maxArea, largestRectangleArea(heights));
         }
         return maxArea;
     }
     
 private:
     int largestRectangleArea(vector<int>& heights) {
         stack<int> st;
         int maxArea = 0;
         for (int i = 0; i <= heights.size(); i++) {
             int h = (i == heights.size()) ? 0 : heights[i];
             while (!st.empty() && heights[st.top()] > h) {
                 int height = heights[st.top()];
                 st.pop();
                 int width = st.empty() ? i : i - st.top() - 1;
                 maxArea = max(maxArea, height * width);
             }
             st.push(i);
         }
         return maxArea;
     }
 };
 
 // 37. Merge Intervals (56)
 class Solution37 {
 public:
     vector<vector<int>> merge(vector<vector<int>>& intervals) {
         sort(intervals.begin(), intervals.end());
         vector<vector<int>> res;
         
         for (auto& interval : intervals) {
             if (res.empty() || res.back()[1] < interval[0]) {
                 res.push_back(interval);
             } else {
                 res.back()[1] = max(res.back()[1], interval[1]);
             }
         }
         return res;
     }
 };
 
 // 38. Insert Interval (57)
 class Solution38 {
 public:
     vector<vector<int>> insert(vector<vector<int>>& intervals, vector<int>& newInterval) {
         vector<vector<int>> res;
         int i = 0;
         
         while (i < intervals.size() && intervals[i][1] < newInterval[0]) {
             res.push_back(intervals[i++]);
         }
         
         while (i < intervals.size() && intervals[i][0] <= newInterval[1]) {
             newInterval[0] = min(newInterval[0], intervals[i][0]);
             newInterval[1] = max(newInterval[1], intervals[i][1]);
             i++;
         }
         res.push_back(newInterval);
         
         while (i < intervals.size()) {
             res.push_back(intervals[i++]);
         }
         return res;
     }
 };
 
 // 39. Non-overlapping Intervals (435)
 class Solution39 {
 public:
     int eraseOverlapIntervals(vector<vector<int>>& intervals) {
         sort(intervals.begin(), intervals.end(), [](auto& a, auto& b) {
             return a[1] < b[1];
         });
         
         int end = intervals[0][1];
         int count = 0;
         
         for (int i = 1; i < intervals.size(); i++) {
             if (intervals[i][0] < end) {
                 count++;
             } else {
                 end = intervals[i][1];
             }
         }
         return count;
     }
 };
 
 // 40. Car Fleet (853)
 class Solution40 {
 public:
     int carFleet(int target, vector<int>& position, vector<int>& speed) {
         int n = position.size();
         vector<pair<int, int>> cars;
         for (int i = 0; i < n; i++) {
             cars.push_back({position[i], speed[i]});
         }
         sort(cars.rbegin(), cars.rend());
         
         stack<double> times;
         for (auto& [pos, spd] : cars) {
             double time = (double)(target - pos) / spd;
             if (times.empty() || time > times.top()) {
                 times.push(time);
             }
         }
         return times.size();
     }
 };
 
 // ============================================================================
 // D) BINARY SEARCH / DIVIDE & CONQUER (10 problems)
 // ============================================================================
 
 // 41. Search in Rotated Sorted Array (33)
 class Solution41 {
 public:
     int search(vector<int>& nums, int target) {
         int l = 0, r = nums.size() - 1;
         
         while (l <= r) {
             int mid = l + (r - l) / 2;
             if (nums[mid] == target) return mid;
             
             if (nums[l] <= nums[mid]) {
                 if (target >= nums[l] && target < nums[mid]) r = mid - 1;
                 else l = mid + 1;
             } else {
                 if (target > nums[mid] && target <= nums[r]) l = mid + 1;
                 else r = mid - 1;
             }
         }
         return -1;
     }
 };
 
 // 42. Find Minimum in Rotated Sorted Array (153)
 class Solution42 {
 public:
     int findMin(vector<int>& nums) {
         int l = 0, r = nums.size() - 1;
         
         while (l < r) {
             int mid = l + (r - l) / 2;
             if (nums[mid] > nums[r]) l = mid + 1;
             else r = mid;
         }
         return nums[l];
     }
 };
 
 // 43. Kth Largest Element in an Array (215) - Quickselect
 class Solution43 {
 public:
     int findKthLargest(vector<int>& nums, int k) {
         return quickSelect(nums, 0, nums.size() - 1, nums.size() - k);
     }
     
 private:
     int quickSelect(vector<int>& nums, int l, int r, int k) {
         if (l == r) return nums[l];
         
         int pivotIdx = partition(nums, l, r);
         if (pivotIdx == k) return nums[pivotIdx];
         else if (pivotIdx < k) return quickSelect(nums, pivotIdx + 1, r, k);
         else return quickSelect(nums, l, pivotIdx - 1, k);
     }
     
     int partition(vector<int>& nums, int l, int r) {
         int pivot = nums[r];
         int i = l;
         for (int j = l; j < r; j++) {
             if (nums[j] < pivot) {
                 swap(nums[i++], nums[j]);
             }
         }
         swap(nums[i], nums[r]);
         return i;
     }
 };
 
 // 44. Find Peak Element (162)
 class Solution44 {
 public:
     int findPeakElement(vector<int>& nums) {
         int l = 0, r = nums.size() - 1;
         
         while (l < r) {
             int mid = l + (r - l) / 2;
             if (nums[mid] < nums[mid + 1]) l = mid + 1;
             else r = mid;
         }
         return l;
     }
 };
 
 // 45. Search a 2D Matrix (74)
 class Solution45 {
 public:
     bool searchMatrix(vector<vector<int>>& matrix, int target) {
         int rows = matrix.size(), cols = matrix[0].size();
         int l = 0, r = rows * cols - 1;
         
         while (l <= r) {
             int mid = l + (r - l) / 2;
             int val = matrix[mid / cols][mid % cols];
             if (val == target) return true;
             else if (val < target) l = mid + 1;
             else r = mid - 1;
         }
         return false;
     }
 };
 
 // 46. Find K Closest Elements (658)
 class Solution46 {
 public:
     vector<int> findClosestElements(vector<int>& arr, int k, int x) {
         int l = 0, r = arr.size() - k;
         
         while (l < r) {
             int mid = l + (r - l) / 2;
             if (x - arr[mid] > arr[mid + k] - x) {
                 l = mid + 1;
             } else {
                 r = mid;
             }
         }
         return vector<int>(arr.begin() + l, arr.begin() + l + k);
     }
 };
 
 // 47. Median of Two Sorted Arrays (4)
 class Solution47 {
 public:
     double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
         if (nums1.size() > nums2.size()) swap(nums1, nums2);
         int m = nums1.size(), n = nums2.size();
         int l = 0, r = m;
         
         while (l <= r) {
             int partX = l + (r - l) / 2;
             int partY = (m + n + 1) / 2 - partX;
             
             int maxLeftX = (partX == 0) ? INT_MIN : nums1[partX - 1];
             int minRightX = (partX == m) ? INT_MAX : nums1[partX];
             int maxLeftY = (partY == 0) ? INT_MIN : nums2[partY - 1];
             int minRightY = (partY == n) ? INT_MAX : nums2[partY];
             
             if (maxLeftX <= minRightY && maxLeftY <= minRightX) {
                 if ((m + n) % 2 == 0) {
                     return (max(maxLeftX, maxLeftY) + min(minRightX, minRightY)) / 2.0;
                 } else {
                     return max(maxLeftX, maxLeftY);
                 }
             } else if (maxLeftX > minRightY) {
                 r = partX - 1;
             } else {
                 l = partX + 1;
             }
         }
         return 0.0;
     }
 };
 
 // 48. Split Array Largest Sum (410)
 class Solution48 {
 public:
     int splitArray(vector<int>& nums, int k) {
         int l = *max_element(nums.begin(), nums.end());
         int r = accumulate(nums.begin(), nums.end(), 0);
         
         while (l < r) {
             int mid = l + (r - l) / 2;
             if (canSplit(nums, k, mid)) r = mid;
             else l = mid + 1;
         }
         return l;
     }
     
 private:
     bool canSplit(vector<int>& nums, int k, int maxSum) {
         int sum = 0, count = 1;
         for (int num : nums) {
             if (sum + num > maxSum) {
                 count++;
                 sum = num;
             } else {
                 sum += num;
             }
         }
         return count <= k;
     }
 };
 
 // 49. Find in Mountain Array (1095)
 class Solution49 {
 public:
     int findInMountainArray(int target, MountainArray &mountainArr) {
         int n = mountainArr.length();
         int peak = findPeak(mountainArr, n);
         
         int left = binarySearch(mountainArr, 0, peak, target, true);
         if (left != -1) return left;
         
         return binarySearch(mountainArr, peak, n - 1, target, false);
     }
     
 private:
     int findPeak(MountainArray& arr, int n) {
         int l = 0, r = n - 1;
         while (l < r) {
             int mid = l + (r - l) / 2;
             if (arr.get(mid) < arr.get(mid + 1)) l = mid + 1;
             else r = mid;
         }
         return l;
     }
     
     int binarySearch(MountainArray& arr, int l, int r, int target, bool ascending) {
         while (l <= r) {
             int mid = l + (r - l) / 2;
             int val = arr.get(mid);
             if (val == target) return mid;
             if (ascending) {
                 if (val < target) l = mid + 1;
                 else r = mid - 1;
             } else {
                 if (val > target) l = mid + 1;
                 else r = mid - 1;
             }
         }
         return -1;
     }
 };
 
 // Note: MountainArray interface is assumed to exist with get() and length() methods
 
 // 50. Capacity To Ship Packages Within D Days (1011)
 class Solution50 {
 public:
     int shipWithinDays(vector<int>& weights, int days) {
         int l = *max_element(weights.begin(), weights.end());
         int r = accumulate(weights.begin(), weights.end(), 0);
         
         while (l < r) {
             int mid = l + (r - l) / 2;
             if (canShip(weights, days, mid)) r = mid;
             else l = mid + 1;
         }
         return l;
     }
     
 private:
     bool canShip(vector<int>& weights, int days, int capacity) {
         int sum = 0, count = 1;
         for (int w : weights) {
             if (sum + w > capacity) {
                 count++;
                 sum = w;
             } else {
                 sum += w;
             }
         }
         return count <= days;
     }
 };
 
 // ============================================================================
 // E) LINKED LIST (5 problems)
 // ============================================================================
 
 // Definition for singly-linked list.
 struct ListNode {
     int val;
     ListNode *next;
     ListNode() : val(0), next(nullptr) {}
     ListNode(int x) : val(x), next(nullptr) {}
     ListNode(int x, ListNode *next) : val(x), next(next) {}
 };
 
 // 51. Add Two Numbers (2)
 class Solution51 {
 public:
     ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
         ListNode* dummy = new ListNode();
         ListNode* curr = dummy;
         int carry = 0;
         
         while (l1 || l2 || carry) {
             int sum = carry;
             if (l1) { sum += l1->val; l1 = l1->next; }
             if (l2) { sum += l2->val; l2 = l2->next; }
             
             curr->next = new ListNode(sum % 10);
             curr = curr->next;
             carry = sum / 10;
         }
         return dummy->next;
     }
 };
 
 // 52. Remove Nth Node From End (19)
 class Solution52 {
 public:
     ListNode* removeNthFromEnd(ListNode* head, int n) {
         ListNode* dummy = new ListNode(0, head);
         ListNode* fast = dummy;
         ListNode* slow = dummy;
         
         for (int i = 0; i <= n; i++) {
             fast = fast->next;
         }
         
         while (fast) {
             fast = fast->next;
             slow = slow->next;
         }
         
         slow->next = slow->next->next;
         return dummy->next;
     }
 };
 
 // 53. Reorder List (143)
 class Solution53 {
 public:
     void reorderList(ListNode* head) {
         if (!head || !head->next) return;
         
         ListNode* slow = head;
         ListNode* fast = head;
         while (fast->next && fast->next->next) {
             slow = slow->next;
             fast = fast->next->next;
         }
         
         ListNode* second = slow->next;
         slow->next = nullptr;
         second = reverse(second);
         
         ListNode* first = head;
         while (second) {
             ListNode* temp1 = first->next;
             ListNode* temp2 = second->next;
             first->next = second;
             second->next = temp1;
             first = temp1;
             second = temp2;
         }
     }
     
 private:
     ListNode* reverse(ListNode* head) {
         ListNode* prev = nullptr;
         while (head) {
             ListNode* next = head->next;
             head->next = prev;
             prev = head;
             head = next;
         }
         return prev;
     }
 };
 
 // 54. Merge k Sorted Lists (23)
 class Solution54 {
 public:
     ListNode* mergeKLists(vector<ListNode*>& lists) {
         auto cmp = [](ListNode* a, ListNode* b) { return a->val > b->val; };
         priority_queue<ListNode*, vector<ListNode*>, decltype(cmp)> pq(cmp);
         
         for (ListNode* list : lists) {
             if (list) pq.push(list);
         }
         
         ListNode* dummy = new ListNode();
         ListNode* curr = dummy;
         
         while (!pq.empty()) {
             ListNode* node = pq.top(); pq.pop();
             curr->next = node;
             curr = curr->next;
             if (node->next) pq.push(node->next);
         }
         return dummy->next;
     }
 };
 
 // 55. Reverse Nodes in k-Group (25)
 class Solution55 {
 public:
     ListNode* reverseKGroup(ListNode* head, int k) {
         ListNode* curr = head;
         int count = 0;
         while (curr && count < k) {
             curr = curr->next;
             count++;
         }
         
         if (count == k) {
             curr = reverseKGroup(curr, k);
             while (count-- > 0) {
                 ListNode* temp = head->next;
                 head->next = curr;
                 curr = head;
                 head = temp;
             }
             head = curr;
         }
         return head;
     }
 };
 
 // ============================================================================
 // F) TREES / BST (15 problems)
 // ============================================================================
 
 // Definition for a binary tree node.
 struct TreeNode {
     int val;
     TreeNode *left;
     TreeNode *right;
     TreeNode() : val(0), left(nullptr), right(nullptr) {}
     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 };
 
 // 56. Binary Tree Level Order Traversal (102)
 class Solution56 {
 public:
     vector<vector<int>> levelOrder(TreeNode* root) {
         if (!root) return {};
         
         vector<vector<int>> res;
         queue<TreeNode*> q;
         q.push(root);
         
         while (!q.empty()) {
             int size = q.size();
             vector<int> level;
             for (int i = 0; i < size; i++) {
                 TreeNode* node = q.front(); q.pop();
                 level.push_back(node->val);
                 if (node->left) q.push(node->left);
                 if (node->right) q.push(node->right);
             }
             res.push_back(level);
         }
         return res;
     }
 };
 
 // 57. Validate Binary Search Tree (98)
 class Solution57 {
 public:
     bool isValidBST(TreeNode* root) {
         return validate(root, LONG_MIN, LONG_MAX);
     }
     
 private:
     bool validate(TreeNode* node, long minVal, long maxVal) {
         if (!node) return true;
         if (node->val <= minVal || node->val >= maxVal) return false;
         return validate(node->left, minVal, node->val) && 
                validate(node->right, node->val, maxVal);
     }
 };
 
 // 58. Kth Smallest in BST (230)
 class Solution58 {
 public:
     int kthSmallest(TreeNode* root, int k) {
         stack<TreeNode*> st;
         TreeNode* curr = root;
         int count = 0;
         
         while (curr || !st.empty()) {
             while (curr) {
                 st.push(curr);
                 curr = curr->left;
             }
             curr = st.top(); st.pop();
             if (++count == k) return curr->val;
             curr = curr->right;
         }
         return -1;
     }
 };
 
 // 59. Construct Binary Tree from Preorder and Inorder (105)
 class Solution59 {
 public:
     TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
         unordered_map<int, int> inMap;
         for (int i = 0; i < inorder.size(); i++) {
             inMap[inorder[i]] = i;
         }
         int preIdx = 0;
         return build(preorder, inorder, 0, inorder.size() - 1, preIdx, inMap);
     }
     
 private:
     TreeNode* build(vector<int>& pre, vector<int>& in, int inStart, int inEnd, 
                     int& preIdx, unordered_map<int, int>& inMap) {
         if (inStart > inEnd) return nullptr;
         
         TreeNode* root = new TreeNode(pre[preIdx++]);
         int rootIdx = inMap[root->val];
         
         root->left = build(pre, in, inStart, rootIdx - 1, preIdx, inMap);
         root->right = build(pre, in, rootIdx + 1, inEnd, preIdx, inMap);
         return root;
     }
 };
 
 // 60. Lowest Common Ancestor of a Binary Tree (236)
 class Solution60 {
 public:
     TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
         if (!root || root == p || root == q) return root;
         
         TreeNode* left = lowestCommonAncestor(root->left, p, q);
         TreeNode* right = lowestCommonAncestor(root->right, p, q);
         
         if (left && right) return root;
         return left ? left : right;
     }
 };
 
 // 61. Binary Tree Right Side View (199)
 class Solution61 {
 public:
     vector<int> rightSideView(TreeNode* root) {
         if (!root) return {};
         
         vector<int> res;
         queue<TreeNode*> q;
         q.push(root);
         
         while (!q.empty()) {
             int size = q.size();
             for (int i = 0; i < size; i++) {
                 TreeNode* node = q.front(); q.pop();
                 if (i == size - 1) res.push_back(node->val);
                 if (node->left) q.push(node->left);
                 if (node->right) q.push(node->right);
             }
         }
         return res;
     }
 };
 
 // 62. Path Sum III (437)
 class Solution62 {
 public:
     int pathSum(TreeNode* root, int targetSum) {
         unordered_map<long long, int> prefixSum;
         prefixSum[0] = 1;
         return dfs(root, 0, targetSum, prefixSum);
     }
     
 private:
     int dfs(TreeNode* node, long long currSum, int target, unordered_map<long long, int>& prefixSum) {
         if (!node) return 0;
         
         currSum += node->val;
         int count = prefixSum[currSum - target];
         prefixSum[currSum]++;
         
         count += dfs(node->left, currSum, target, prefixSum);
         count += dfs(node->right, currSum, target, prefixSum);
         
         prefixSum[currSum]--;
         return count;
     }
 };
 
 // 63. Binary Tree Maximum Path Sum (124)
 class Solution63 {
     int maxSum = INT_MIN;
     
 public:
     int maxPathSum(TreeNode* root) {
         dfs(root);
         return maxSum;
     }
     
 private:
     int dfs(TreeNode* node) {
         if (!node) return 0;
         
         int left = max(0, dfs(node->left));
         int right = max(0, dfs(node->right));
         
         maxSum = max(maxSum, node->val + left + right);
         return node->val + max(left, right);
     }
 };
 
 // 64. Serialize and Deserialize Binary Tree (297)
 class Codec {
 public:
     string serialize(TreeNode* root) {
         if (!root) return "#";
         return to_string(root->val) + "," + serialize(root->left) + "," + serialize(root->right);
     }
     
     TreeNode* deserialize(string data) {
         stringstream ss(data);
         return deserializeHelper(ss);
     }
     
 private:
     TreeNode* deserializeHelper(stringstream& ss) {
         string val;
         getline(ss, val, ',');
         if (val == "#") return nullptr;
         
         TreeNode* node = new TreeNode(stoi(val));
         node->left = deserializeHelper(ss);
         node->right = deserializeHelper(ss);
         return node;
     }
 };
 
 // 65. Count of Smaller Numbers After Self (315) - Merge Sort approach
 class Solution65 {
 public:
     vector<int> countSmaller(vector<int>& nums) {
         int n = nums.size();
         vector<int> res(n, 0);
         vector<pair<int, int>> pairs;
         for (int i = 0; i < n; i++) {
             pairs.push_back({nums[i], i});
         }
         mergeSort(pairs, 0, n - 1, res);
         return res;
     }
     
 private:
     void mergeSort(vector<pair<int, int>>& pairs, int l, int r, vector<int>& res) {
         if (l >= r) return;
         
         int mid = l + (r - l) / 2;
         mergeSort(pairs, l, mid, res);
         mergeSort(pairs, mid + 1, r, res);
         merge(pairs, l, mid, r, res);
     }
     
     void merge(vector<pair<int, int>>& pairs, int l, int mid, int r, vector<int>& res) {
         vector<pair<int, int>> temp;
         int i = l, j = mid + 1;
         int rightCount = 0;
         
         while (i <= mid && j <= r) {
             if (pairs[i].first > pairs[j].first) {
                 rightCount++;
                 temp.push_back(pairs[j++]);
             } else {
                 res[pairs[i].second] += rightCount;
                 temp.push_back(pairs[i++]);
             }
         }
         
         while (i <= mid) {
             res[pairs[i].second] += rightCount;
             temp.push_back(pairs[i++]);
         }
         while (j <= r) {
             temp.push_back(pairs[j++]);
         }
         
         for (int k = 0; k < temp.size(); k++) {
             pairs[l + k] = temp[k];
         }
     }
 };
 
 // 66. Implement Trie (Prefix Tree) (208)
 class Trie {
     struct TrieNode {
         vector<TrieNode*> children;
         bool isEnd;
         TrieNode() : children(26, nullptr), isEnd(false) {}
     };
     
     TrieNode* root;
     
 public:
     Trie() {
         root = new TrieNode();
     }
     
     void insert(string word) {
         TrieNode* curr = root;
         for (char c : word) {
             if (!curr->children[c - 'a']) {
                 curr->children[c - 'a'] = new TrieNode();
             }
             curr = curr->children[c - 'a'];
         }
         curr->isEnd = true;
     }
     
     bool search(string word) {
         TrieNode* curr = root;
         for (char c : word) {
             if (!curr->children[c - 'a']) return false;
             curr = curr->children[c - 'a'];
         }
         return curr->isEnd;
     }
     
     bool startsWith(string prefix) {
         TrieNode* curr = root;
         for (char c : prefix) {
             if (!curr->children[c - 'a']) return false;
             curr = curr->children[c - 'a'];
         }
         return true;
     }
 };

// 67. Word Search II (212)
class Solution67 {
    struct Node {
        array<Node*, 26> next{};
        string word = "";
        Node() { next.fill(nullptr); }
    };

    void insert(Node* root, const string& w) {
        Node* cur = root;
        for (char ch : w) {
            int i = ch - 'a';
            if (!cur->next[i]) cur->next[i] = new Node();
            cur = cur->next[i];
        }
        cur->word = w;
    }

    void dfs(vector<vector<char>>& board, int r, int c, Node* node, vector<string>& out) {
        char ch = board[r][c];
        if (ch == '#') return;
        Node* nxt = node->next[ch - 'a'];
        if (!nxt) return;

        if (!nxt->word.empty()) {
            out.push_back(nxt->word);
            nxt->word.clear(); // avoid duplicates
        }

        board[r][c] = '#';
        static int dr[4] = {1, -1, 0, 0};
        static int dc[4] = {0, 0, 1, -1};
        int R = (int)board.size(), C = (int)board[0].size();
        for (int k = 0; k < 4; k++) {
            int nr = r + dr[k], nc = c + dc[k];
            if (nr >= 0 && nr < R && nc >= 0 && nc < C) dfs(board, nr, nc, nxt, out);
        }
        board[r][c] = ch;
    }

public:
    vector<string> findWords(vector<vector<char>>& board, vector<string>& words) {
        Node* root = new Node();
        for (auto& w : words) insert(root, w);

        vector<string> out;
        int R = (int)board.size();
        if (R == 0) return out;
        int C = (int)board[0].size();
        for (int r = 0; r < R; r++) {
            for (int c = 0; c < C; c++) {
                dfs(board, r, c, root, out);
            }
        }
        return out;
    }
};

// 68. Delete Node in a BST (450)
class Solution68 {
public:
    TreeNode* deleteNode(TreeNode* root, int key) {
        if (!root) return nullptr;
        if (key < root->val) {
            root->left = deleteNode(root->left, key);
        } else if (key > root->val) {
            root->right = deleteNode(root->right, key);
        } else {
            if (!root->left) return root->right;
            if (!root->right) return root->left;
            TreeNode* succ = root->right;
            while (succ->left) succ = succ->left;
            root->val = succ->val;
            root->right = deleteNode(root->right, succ->val);
        }
        return root;
    }
};

// 69. Recover Binary Search Tree (99)
class Solution69 {
public:
    void recoverTree(TreeNode* root) {
        TreeNode *first = nullptr, *second = nullptr, *prev = nullptr;
        inorder(root, prev, first, second);
        if (first && second) swap(first->val, second->val);
    }

private:
    void inorder(TreeNode* node, TreeNode*& prev, TreeNode*& first, TreeNode*& second) {
        if (!node) return;
        inorder(node->left, prev, first, second);
        if (prev && prev->val > node->val) {
            if (!first) first = prev;
            second = node;
        }
        prev = node;
        inorder(node->right, prev, first, second);
    }
};

// 70. All Nodes Distance K in Binary Tree (863)
class Solution70 {
public:
    vector<int> distanceK(TreeNode* root, TreeNode* target, int k) {
        unordered_map<TreeNode*, TreeNode*> parent;
        buildParent(root, nullptr, parent);
        queue<TreeNode*> q;
        unordered_set<TreeNode*> vis;
        q.push(target);
        vis.insert(target);

        int dist = 0;
        while (!q.empty() && dist < k) {
            int sz = (int)q.size();
            while (sz--) {
                TreeNode* cur = q.front(); q.pop();
                for (TreeNode* nxt : {cur->left, cur->right, parent[cur]}) {
                    if (nxt && !vis.count(nxt)) {
                        vis.insert(nxt);
                        q.push(nxt);
                    }
                }
            }
            dist++;
        }

        vector<int> out;
        while (!q.empty()) {
            out.push_back(q.front()->val);
            q.pop();
        }
        return out;
    }

private:
    void buildParent(TreeNode* node, TreeNode* par, unordered_map<TreeNode*, TreeNode*>& parent) {
        if (!node) return;
        parent[node] = par;
        buildParent(node->left, node, parent);
        buildParent(node->right, node, parent);
    }
};

// ============================================================================
// G) GRAPHS / BFS / DFS / TOPO / UNION-FIND (15 problems)
// ============================================================================

// Definition for undirected graph node (LeetCode 133).
class Node {
public:
    int val;
    vector<Node*> neighbors;
    Node() : val(0) {}
    Node(int _val) : val(_val) {}
    Node(int _val, vector<Node*> _neighbors) : val(_val), neighbors(std::move(_neighbors)) {}
};

// 71. Number of Islands (200)
class Solution71 {
public:
    int numIslands(vector<vector<char>>& grid) {
        int R = (int)grid.size();
        if (R == 0) return 0;
        int C = (int)grid[0].size();
        int ans = 0;
        for (int r = 0; r < R; r++) {
            for (int c = 0; c < C; c++) {
                if (grid[r][c] == '1') {
                    ans++;
                    flood(grid, r, c);
                }
            }
        }
        return ans;
    }
private:
    void flood(vector<vector<char>>& g, int r, int c) {
        int R = (int)g.size(), C = (int)g[0].size();
        if (r < 0 || r >= R || c < 0 || c >= C || g[r][c] != '1') return;
        g[r][c] = '0';
        flood(g, r + 1, c);
        flood(g, r - 1, c);
        flood(g, r, c + 1);
        flood(g, r, c - 1);
    }
};

// 72. Rotting Oranges (994)
class Solution72 {
public:
    int orangesRotting(vector<vector<int>>& grid) {
        int R = (int)grid.size();
        if (R == 0) return 0;
        int C = (int)grid[0].size();
        queue<pair<int,int>> q;
        int fresh = 0;
        for (int r = 0; r < R; r++) {
            for (int c = 0; c < C; c++) {
                if (grid[r][c] == 2) q.push({r,c});
                else if (grid[r][c] == 1) fresh++;
            }
        }
        if (fresh == 0) return 0;
        int minutes = -1;
        int dr[4] = {1,-1,0,0};
        int dc[4] = {0,0,1,-1};
        while (!q.empty()) {
            int sz = (int)q.size();
            minutes++;
            while (sz--) {
                auto [r,c] = q.front(); q.pop();
                for (int k = 0; k < 4; k++) {
                    int nr = r + dr[k], nc = c + dc[k];
                    if (nr>=0 && nr<R && nc>=0 && nc<C && grid[nr][nc]==1) {
                        grid[nr][nc] = 2;
                        fresh--;
                        q.push({nr,nc});
                    }
                }
            }
        }
        return fresh == 0 ? minutes : -1;
    }
};

// 73. Course Schedule (207)
class Solution73 {
public:
    bool canFinish(int numCourses, vector<vector<int>>& prerequisites) {
        vector<vector<int>> g(numCourses);
        vector<int> indeg(numCourses, 0);
        for (auto& p : prerequisites) {
            g[p[1]].push_back(p[0]);
            indeg[p[0]]++;
        }
        queue<int> q;
        for (int i = 0; i < numCourses; i++) if (indeg[i] == 0) q.push(i);
        int seen = 0;
        while (!q.empty()) {
            int u = q.front(); q.pop();
            seen++;
            for (int v : g[u]) if (--indeg[v] == 0) q.push(v);
        }
        return seen == numCourses;
    }
};

// 74. Course Schedule II (210)
class Solution74 {
public:
    vector<int> findOrder(int numCourses, vector<vector<int>>& prerequisites) {
        vector<vector<int>> g(numCourses);
        vector<int> indeg(numCourses, 0);
        for (auto& p : prerequisites) {
            g[p[1]].push_back(p[0]);
            indeg[p[0]]++;
        }
        queue<int> q;
        for (int i = 0; i < numCourses; i++) if (indeg[i] == 0) q.push(i);
        vector<int> order;
        while (!q.empty()) {
            int u = q.front(); q.pop();
            order.push_back(u);
            for (int v : g[u]) if (--indeg[v] == 0) q.push(v);
        }
        if ((int)order.size() != numCourses) return {};
        return order;
    }
};

// 75. Clone Graph (133)
class Solution75 {
public:
    Node* cloneGraph(Node* node) {
        if (!node) return nullptr;
        unordered_map<Node*, Node*> mp;
        return clone(node, mp);
    }
private:
    Node* clone(Node* node, unordered_map<Node*, Node*>& mp) {
        if (mp.count(node)) return mp[node];
        Node* cp = new Node(node->val);
        mp[node] = cp;
        for (Node* nei : node->neighbors) cp->neighbors.push_back(clone(nei, mp));
        return cp;
    }
};

// 76. Accounts Merge (721) - DSU
class Solution76 {
    struct DSU {
        vector<int> p, r;
        DSU(int n): p(n), r(n,0) { iota(p.begin(), p.end(), 0); }
        int find(int x){ return p[x]==x?x:p[x]=find(p[x]); }
        void unite(int a,int b){
            a=find(a); b=find(b);
            if(a==b) return;
            if(r[a]<r[b]) swap(a,b);
            p[b]=a;
            if(r[a]==r[b]) r[a]++;
        }
    };
public:
    vector<vector<string>> accountsMerge(vector<vector<string>>& accounts) {
        unordered_map<string,int> id;
        unordered_map<string,string> owner;
        int idx = 0;
        for (auto& acc : accounts) {
            for (int i = 1; i < (int)acc.size(); i++) {
                if (!id.count(acc[i])) id[acc[i]] = idx++;
                owner[acc[i]] = acc[0];
            }
        }
        DSU dsu(idx);
        for (auto& acc : accounts) {
            for (int i = 2; i < (int)acc.size(); i++) {
                dsu.unite(id[acc[1]], id[acc[i]]);
            }
        }
        unordered_map<int, vector<string>> groups;
        for (auto& [email, eid] : id) groups[dsu.find(eid)].push_back(email);

        vector<vector<string>> res;
        for (auto& [root, emails] : groups) {
            sort(emails.begin(), emails.end());
            vector<string> merged;
            merged.push_back(owner[emails[0]]);
            merged.insert(merged.end(), emails.begin(), emails.end());
            res.push_back(std::move(merged));
        }
        return res;
    }
};

// 77. Network Delay Time (743)
class Solution77 {
public:
    int networkDelayTime(vector<vector<int>>& times, int n, int k) {
        vector<vector<pair<int,int>>> g(n+1);
        for (auto& e : times) g[e[0]].push_back({e[1], e[2]});
        const int INF = 1e9;
        vector<int> dist(n+1, INF);
        priority_queue<pair<int,int>, vector<pair<int,int>>, greater<pair<int,int>>> pq;
        dist[k] = 0;
        pq.push({0, k});
        while (!pq.empty()) {
            auto [d,u] = pq.top(); pq.pop();
            if (d != dist[u]) continue;
            for (auto [v,w] : g[u]) {
                if (dist[v] > d + w) {
                    dist[v] = d + w;
                    pq.push({dist[v], v});
                }
            }
        }
        int ans = 0;
        for (int i = 1; i <= n; i++) {
            if (dist[i] == INF) return -1;
            ans = max(ans, dist[i]);
        }
        return ans;
    }
};

// 78. Word Ladder (127)
class Solution78 {
public:
    int ladderLength(string beginWord, string endWord, vector<string>& wordList) {
        unordered_set<string> dict(wordList.begin(), wordList.end());
        if (!dict.count(endWord)) return 0;
        unordered_map<string, vector<string>> buckets;
        int L = (int)beginWord.size();
        for (auto& w : dict) {
            for (int i = 0; i < L; i++) {
                string key = w;
                key[i] = '*';
                buckets[key].push_back(w);
            }
        }
        queue<pair<string,int>> q;
        unordered_set<string> vis;
        q.push({beginWord, 1});
        vis.insert(beginWord);
        while (!q.empty()) {
            auto [w, d] = q.front(); q.pop();
            if (w == endWord) return d;
            for (int i = 0; i < L; i++) {
                string key = w;
                key[i] = '*';
                auto it = buckets.find(key);
                if (it == buckets.end()) continue;
                for (auto& nxt : it->second) {
                    if (!vis.count(nxt)) {
                        vis.insert(nxt);
                        q.push({nxt, d + 1});
                    }
                }
                buckets.erase(key); // important pruning
            }
        }
        return 0;
    }
};

// 79. Word Ladder II (126)
class Solution79 {
public:
    vector<vector<string>> findLadders(string beginWord, string endWord, vector<string>& wordList) {
        unordered_set<string> dict(wordList.begin(), wordList.end());
        vector<vector<string>> res;
        if (!dict.count(endWord)) return res;

        unordered_map<string, vector<string>> parents; // child -> list of parents on shortest paths
        unordered_set<string> curLevel{beginWord};
        bool found = false;

        while (!curLevel.empty() && !found) {
            for (auto& w : curLevel) dict.erase(w);
            unordered_set<string> nextLevel;

            for (auto& w : curLevel) {
                string s = w;
                for (int i = 0; i < (int)s.size(); i++) {
                    char orig = s[i];
                    for (char ch = 'a'; ch <= 'z'; ch++) {
                        if (ch == orig) continue;
                        s[i] = ch;
                        if (dict.count(s)) {
                            nextLevel.insert(s);
                            parents[s].push_back(w);
                            if (s == endWord) found = true;
                        }
                    }
                    s[i] = orig;
                }
            }
            curLevel.swap(nextLevel);
        }

        if (!found) return res;
        vector<string> path{endWord};
        backtrack(endWord, beginWord, parents, path, res);
        return res;
    }

private:
    void backtrack(const string& word, const string& beginWord,
                   unordered_map<string, vector<string>>& parents,
                   vector<string>& path, vector<vector<string>>& res) {
        if (word == beginWord) {
            vector<string> p = path;
            reverse(p.begin(), p.end());
            res.push_back(std::move(p));
            return;
        }
        for (auto& par : parents[word]) {
            path.push_back(par);
            backtrack(par, beginWord, parents, path, res);
            path.pop_back();
        }
    }
};

// 80. Alien Dictionary (269)
class Solution80 {
public:
    string alienOrder(vector<string>& words) {
        unordered_map<char, unordered_set<char>> g;
        unordered_map<char, int> indeg;
        for (auto& w : words) for (char c : w) indeg.emplace(c, 0);

        for (int i = 0; i + 1 < (int)words.size(); i++) {
            string& a = words[i];
            string& b = words[i+1];
            int m = min(a.size(), b.size());
            int j = 0;
            while (j < m && a[j] == b[j]) j++;
            if (j == m) {
                if (a.size() > b.size()) return ""; // invalid prefix order
                continue;
            }
            char u = a[j], v = b[j];
            if (!g[u].count(v)) {
                g[u].insert(v);
                indeg[v]++;
            }
        }

        queue<char> q;
        for (auto& [c, d] : indeg) if (d == 0) q.push(c);
        string order;
        while (!q.empty()) {
            char u = q.front(); q.pop();
            order.push_back(u);
            for (char v : g[u]) {
                if (--indeg[v] == 0) q.push(v);
            }
        }
        if ((int)order.size() != (int)indeg.size()) return "";
        return order;
    }
};

// 81. Minimum Number of Refueling Stops (871)
class Solution81 {
public:
    int minRefuelStops(int target, int startFuel, vector<vector<int>>& stations) {
        priority_queue<int> pq;
        long long fuel = startFuel;
        int i = 0, n = (int)stations.size();
        int stops = 0;
        while (fuel < target) {
            while (i < n && stations[i][0] <= fuel) {
                pq.push(stations[i][1]);
                i++;
            }
            if (pq.empty()) return -1;
            fuel += pq.top(); pq.pop();
            stops++;
        }
        return stops;
    }
};

// 82. Evaluate Division (399)
class Solution82 {
public:
    vector<double> calcEquation(vector<vector<string>>& equations, vector<double>& values,
                                vector<vector<string>>& queries) {
        unordered_map<string, vector<pair<string,double>>> g;
        for (int i = 0; i < (int)equations.size(); i++) {
            auto& e = equations[i];
            double v = values[i];
            g[e[0]].push_back({e[1], v});
            g[e[1]].push_back({e[0], 1.0 / v});
        }
        vector<double> res;
        for (auto& q : queries) {
            if (!g.count(q[0]) || !g.count(q[1])) {
                res.push_back(-1.0);
                continue;
            }
            if (q[0] == q[1]) {
                res.push_back(1.0);
                continue;
            }
            unordered_set<string> vis;
            double ans = -1.0;
            dfs(q[0], q[1], 1.0, g, vis, ans);
            res.push_back(ans);
        }
        return res;
    }

private:
    void dfs(const string& u, const string& t, double acc,
             unordered_map<string, vector<pair<string,double>>>& g,
             unordered_set<string>& vis, double& ans) {
        if (vis.count(u) || ans != -1.0) return;
        vis.insert(u);
        if (u == t) { ans = acc; return; }
        for (auto& [v,w] : g[u]) dfs(v, t, acc * w, g, vis, ans);
    }
};

// 83. Critical Connections in a Network (1192)
class Solution83 {
public:
    vector<vector<int>> criticalConnections(int n, vector<vector<int>>& connections) {
        vector<vector<int>> g(n);
        for (auto& e : connections) {
            g[e[0]].push_back(e[1]);
            g[e[1]].push_back(e[0]);
        }
        vector<int> tin(n, -1), low(n, -1);
        vector<vector<int>> bridges;
        int timer = 0;
        dfs(0, -1, timer, g, tin, low, bridges);
        return bridges;
    }
private:
    void dfs(int u, int p, int& timer, vector<vector<int>>& g,
             vector<int>& tin, vector<int>& low, vector<vector<int>>& bridges) {
        tin[u] = low[u] = timer++;
        for (int v : g[u]) {
            if (v == p) continue;
            if (tin[v] != -1) {
                low[u] = min(low[u], tin[v]);
            } else {
                dfs(v, u, timer, g, tin, low, bridges);
                low[u] = min(low[u], low[v]);
                if (low[v] > tin[u]) bridges.push_back({u, v});
            }
        }
    }
};

// 84. Shortest Path in a Grid with Obstacles Elimination (1293)
class Solution84 {
public:
    int shortestPath(vector<vector<int>>& grid, int k) {
        int R = (int)grid.size(), C = (int)grid[0].size();
        if (R == 1 && C == 1) return 0;
        vector<vector<int>> best(R, vector<int>(C, -1)); // max remaining k seen at cell
        queue<tuple<int,int,int>> q;
        q.push({0,0,k});
        best[0][0] = k;
        int steps = 0;
        int dr[4] = {1,-1,0,0};
        int dc[4] = {0,0,1,-1};
        while (!q.empty()) {
            int sz = (int)q.size();
            steps++;
            while (sz--) {
                auto [r,c,rem] = q.front(); q.pop();
                for (int i = 0; i < 4; i++) {
                    int nr = r + dr[i], nc = c + dc[i];
                    if (nr<0 || nr>=R || nc<0 || nc>=C) continue;
                    int nrem = rem - grid[nr][nc];
                    if (nrem < 0) continue;
                    if (nr == R-1 && nc == C-1) return steps;
                    if (best[nr][nc] >= nrem) continue;
                    best[nr][nc] = nrem;
                    q.push({nr,nc,nrem});
                }
            }
        }
        return -1;
    }
};

// 85. Swim in Rising Water (778)
class Solution85 {
public:
    int swimInWater(vector<vector<int>>& grid) {
        int n = (int)grid.size();
        vector<vector<int>> dist(n, vector<int>(n, INT_MAX));
        priority_queue<tuple<int,int,int>, vector<tuple<int,int,int>>, greater<tuple<int,int,int>>> pq;
        dist[0][0] = grid[0][0];
        pq.push({dist[0][0], 0, 0});
        int dr[4] = {1,-1,0,0};
        int dc[4] = {0,0,1,-1};
        while (!pq.empty()) {
            auto [d,r,c] = pq.top(); pq.pop();
            if (r == n-1 && c == n-1) return d;
            if (d != dist[r][c]) continue;
            for (int i = 0; i < 4; i++) {
                int nr = r + dr[i], nc = c + dc[i];
                if (nr<0 || nr>=n || nc<0 || nc>=n) continue;
                int nd = max(d, grid[nr][nc]);
                if (nd < dist[nr][nc]) {
                    dist[nr][nc] = nd;
                    pq.push({nd, nr, nc});
                }
            }
        }
        return -1;
    }
};

// ============================================================================
// H) DYNAMIC PROGRAMMING (15 problems)
// ============================================================================

// 86. House Robber (198)
class Solution86 {
public:
    int rob(vector<int>& nums) {
        int prev2 = 0, prev1 = 0;
        for (int x : nums) {
            int cur = max(prev1, prev2 + x);
            prev2 = prev1;
            prev1 = cur;
        }
        return prev1;
    }
};

// 87. Coin Change (322)
class Solution87 {
public:
    int coinChange(vector<int>& coins, int amount) {
        const int INF = 1e9;
        vector<int> dp(amount + 1, INF);
        dp[0] = 0;
        for (int a = 1; a <= amount; a++) {
            for (int c : coins) {
                if (a - c >= 0 && dp[a - c] != INF) dp[a] = min(dp[a], dp[a - c] + 1);
            }
        }
        return dp[amount] == INF ? -1 : dp[amount];
    }
};

// 88. Longest Increasing Subsequence (300)
class Solution88 {
public:
    int lengthOfLIS(vector<int>& nums) {
        vector<int> tails;
        for (int x : nums) {
            auto it = lower_bound(tails.begin(), tails.end(), x);
            if (it == tails.end()) tails.push_back(x);
            else *it = x;
        }
        return (int)tails.size();
    }
};

// 89. Partition Equal Subset Sum (416)
class Solution89 {
public:
    bool canPartition(vector<int>& nums) {
        int sum = accumulate(nums.begin(), nums.end(), 0);
        if (sum % 2) return false;
        int target = sum / 2;
        vector<char> dp(target + 1, 0);
        dp[0] = 1;
        for (int x : nums) {
            for (int s = target; s >= x; s--) dp[s] = dp[s] || dp[s - x];
        }
        return dp[target];
    }
};

// 90. Unique Paths (62)
class Solution90 {
public:
    int uniquePaths(int m, int n) {
        vector<long long> dp(n, 1);
        for (int i = 1; i < m; i++) {
            for (int j = 1; j < n; j++) dp[j] += dp[j - 1];
        }
        return (int)dp[n - 1];
    }
};

// 91. Longest Palindromic Substring (5)
class Solution91 {
public:
    string longestPalindrome(string s) {
        int n = (int)s.size();
        int bestL = 0, bestR = -1;
        for (int i = 0; i < n; i++) {
            expand(s, i, i, bestL, bestR);
            expand(s, i, i + 1, bestL, bestR);
        }
        return s.substr(bestL, bestR - bestL + 1);
    }
private:
    void expand(const string& s, int l, int r, int& bestL, int& bestR) {
        int n = (int)s.size();
        while (l >= 0 && r < n && s[l] == s[r]) { l--; r++; }
        l++; r--;
        if (r - l > bestR - bestL) { bestL = l; bestR = r; }
    }
};

// 92. Palindromic Substrings (647)
class Solution92 {
public:
    int countSubstrings(string s) {
        int n = (int)s.size();
        int ans = 0;
        for (int i = 0; i < n; i++) {
            ans += expand(s, i, i);
            ans += expand(s, i, i + 1);
        }
        return ans;
    }
private:
    int expand(const string& s, int l, int r) {
        int n = (int)s.size();
        int cnt = 0;
        while (l >= 0 && r < n && s[l] == s[r]) { cnt++; l--; r++; }
        return cnt;
    }
};

// 93. Decode Ways (91)
class Solution93 {
public:
    int numDecodings(string s) {
        int n = (int)s.size();
        if (n == 0 || s[0] == '0') return 0;
        long long prev2 = 1, prev1 = 1;
        for (int i = 1; i < n; i++) {
            long long cur = 0;
            if (s[i] != '0') cur += prev1;
            int two = (s[i-1]-'0')*10 + (s[i]-'0');
            if (two >= 10 && two <= 26) cur += prev2;
            prev2 = prev1;
            prev1 = cur;
        }
        return (int)prev1;
    }
};

// 94. Word Break (139)
class Solution94 {
public:
    bool wordBreak(string s, vector<string>& wordDict) {
        unordered_set<string> dict(wordDict.begin(), wordDict.end());
        int n = (int)s.size();
        vector<char> dp(n + 1, 0);
        dp[0] = 1;
        for (int i = 1; i <= n; i++) {
            for (int j = 0; j < i; j++) {
                if (!dp[j]) continue;
                if (dict.count(s.substr(j, i - j))) { dp[i] = 1; break; }
            }
        }
        return dp[n];
    }
};

// 95. Word Break II (140)
class Solution95 {
public:
    vector<string> wordBreak(string s, vector<string>& wordDict) {
        unordered_set<string> dict(wordDict.begin(), wordDict.end());
        unordered_map<int, vector<string>> memo;
        return dfs(0, s, dict, memo);
    }
private:
    vector<string> dfs(int i, const string& s, unordered_set<string>& dict,
                       unordered_map<int, vector<string>>& memo) {
        if (memo.count(i)) return memo[i];
        int n = (int)s.size();
        vector<string> res;
        if (i == n) { res.push_back(""); return memo[i] = res; }
        for (int j = i + 1; j <= n; j++) {
            string w = s.substr(i, j - i);
            if (!dict.count(w)) continue;
            auto tails = dfs(j, s, dict, memo);
            for (auto& t : tails) {
                if (t.empty()) res.push_back(w);
                else res.push_back(w + " " + t);
            }
        }
        return memo[i] = res;
    }
};

// 96. Edit Distance (72)
class Solution96 {
public:
    int minDistance(string word1, string word2) {
        int n = (int)word1.size(), m = (int)word2.size();
        vector<int> dp(m + 1);
        iota(dp.begin(), dp.end(), 0);
        for (int i = 1; i <= n; i++) {
            vector<int> ndp(m + 1);
            ndp[0] = i;
            for (int j = 1; j <= m; j++) {
                if (word1[i-1] == word2[j-1]) ndp[j] = dp[j-1];
                else ndp[j] = 1 + min({dp[j], ndp[j-1], dp[j-1]});
            }
            dp.swap(ndp);
        }
        return dp[m];
    }
};

// 97. Regular Expression Matching (10)
class Solution97 {
public:
    bool isMatch(string s, string p) {
        int n = (int)s.size(), m = (int)p.size();
        vector<vector<char>> dp(n + 1, vector<char>(m + 1, 0));
        dp[0][0] = 1;
        for (int j = 2; j <= m; j++) {
            if (p[j-1] == '*') dp[0][j] = dp[0][j-2];
        }
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= m; j++) {
                if (p[j-1] == '*') {
                    dp[i][j] = dp[i][j-2]; // use 0 of preceding
                    char prev = p[j-2];
                    if (prev == '.' || prev == s[i-1]) dp[i][j] = dp[i][j] || dp[i-1][j];
                } else if (p[j-1] == '.' || p[j-1] == s[i-1]) {
                    dp[i][j] = dp[i-1][j-1];
                }
            }
        }
        return dp[n][m];
    }
};

// 98. Distinct Subsequences (115)
class Solution98 {
public:
    int numDistinct(string s, string t) {
        int n = (int)s.size(), m = (int)t.size();
        vector<long long> dp(m + 1, 0);
        dp[0] = 1;
        for (int i = 1; i <= n; i++) {
            for (int j = m; j >= 1; j--) {
                if (s[i-1] == t[j-1]) dp[j] += dp[j-1];
            }
        }
        return (int)dp[m];
    }
};

// 99. Burst Balloons (312)
class Solution99 {
public:
    int maxCoins(vector<int>& nums) {
        int n = (int)nums.size();
        vector<int> a(n + 2, 1);
        for (int i = 0; i < n; i++) a[i + 1] = nums[i];
        int N = n + 2;
        vector<vector<int>> dp(N, vector<int>(N, 0));
        for (int len = 2; len < N; len++) {
            for (int l = 0; l + len < N; l++) {
                int r = l + len;
                for (int k = l + 1; k < r; k++) {
                    dp[l][r] = max(dp[l][r], dp[l][k] + dp[k][r] + a[l] * a[k] * a[r]);
                }
            }
        }
        return dp[0][N - 1];
    }
};

// 100. Longest Increasing Path in a Matrix (329)
class Solution100 {
public:
    int longestIncreasingPath(vector<vector<int>>& matrix) {
        int R = (int)matrix.size();
        if (R == 0) return 0;
        int C = (int)matrix[0].size();
        vector<vector<int>> memo(R, vector<int>(C, 0));
        int ans = 0;
        for (int r = 0; r < R; r++) {
            for (int c = 0; c < C; c++) ans = max(ans, dfs(matrix, r, c, memo));
        }
        return ans;
    }
private:
    int dfs(vector<vector<int>>& a, int r, int c, vector<vector<int>>& memo) {
        if (memo[r][c]) return memo[r][c];
        int R = (int)a.size(), C = (int)a[0].size();
        int best = 1;
        int dr[4] = {1,-1,0,0};
        int dc[4] = {0,0,1,-1};
        for (int i = 0; i < 4; i++) {
            int nr = r + dr[i], nc = c + dc[i];
            if (nr<0 || nr>=R || nc<0 || nc>=C) continue;
            if (a[nr][nc] > a[r][c]) best = max(best, 1 + dfs(a, nr, nc, memo));
        }
        return memo[r][c] = best;
    }
};
