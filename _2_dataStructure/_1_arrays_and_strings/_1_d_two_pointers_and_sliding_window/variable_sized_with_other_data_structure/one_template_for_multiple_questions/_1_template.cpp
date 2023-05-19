// https://leetcode.com/problems/subarrays-with-k-different-integers

#include <vector>
#include <unordered_map>
using namespace std;

class Solution {
public:
    int subarraysWithKDistinct(vector<int>& A, int K) {
        return subarraysWithAtMostKDistinct(A, K) - subarraysWithAtMostKDistinct(A, K - 1);
    }

private:
    int subarraysWithAtMostKDistinct(vector<int>& s, int k) {
        unordered_map<int, int> lookup;
        int l = 0, r = 0, counter = 0, res = 0;
        while (r < s.size()) {
            lookup[s[r]]++;
            if (lookup[s[r]] == 1) {
                counter++;
            }
            r++;
            while (l < r && counter > k) {
                lookup[s[l]]--;
                if (lookup[s[l]] == 0) {
                    counter--;
                }
                l++;
            }
            res += r - l;
        }
        return res;
    }
};

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

class Solution {
public:
    int lengthOfLongestSubstring(std::string s) {
        std::unordered_map<char, int> lookup;
        int l = 0, r = 0, counter = 0, res = 0;

        while (r < s.length()) {
            lookup[s[r]]++;
            if (lookup[s[r]] == 1) {
                counter++;
            }
            r++;

            while (l < r && counter < (r - l)) {
                lookup[s[l]]--;
                if (lookup[s[l]] == 0) {
                    counter--;
                }
                l++;
            }

            res = std::max(res, r - l);
        }

        return res;
    }
};

// https://leetcode.com/problems/longest-substring-with-at-most-two-distinct-characters/

class Solution {
public:
    int lengthOfLongestSubstringTwoDistinct(std::string s) {
        std::unordered_map<char, int> lookup;
        int l = 0, r = 0, counter = 0, res = 0;

        while (r < s.length()) {
            lookup[s[r]]++;
            if (lookup[s[r]] == 1) {
                counter++;
            }
            r++;

            while (l < r && counter > 2) {
                lookup[s[l]]--;
                if (lookup[s[l]] == 0) {
                    counter--;
                }
                l++;
            }

            res = std::max(res, r - l);
        }

        return res;
    }
};

// https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/
class Solution {
public:
    int lengthOfLongestSubstringKDistinct(std::string s, int k) {
        std::unordered_map<char, int> lookup;
        int l = 0, r = 0, counter = 0, res = 0;

        while (r < s.length()) {
            lookup[s[r]]++;
            if (lookup[s[r]] == 1) {
                counter++;
            }
            r++;

            while (l < r && counter > k) {
                lookup[s[l]]--;
                if (lookup[s[l]] == 0) {
                    counter--;
                }
                l++;
            }

            res = std::max(res, r - l);
        }

        return res;
    }
};