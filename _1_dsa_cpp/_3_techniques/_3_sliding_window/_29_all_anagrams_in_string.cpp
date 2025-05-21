// https://leetcode.com/problems/find-all-anagrams-in-a-string/description/

#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

class Solution {
public:
    vector<int> findAnagrams(string text, string pat) {
        vector<int> res = search(pat, text);
        return res;
    }

private:
    vector<int> search(string pat, string text) {
        vector<int> list;
        int i = 0,j = 0,k = pat.length(),n = text.length(),ans = 0;
        unordered_map<char, int> map;

        for (int l = 0; l < pat.length(); l++) {
            map[pat[l]] = map[pat[l]] + 1;
        }

        int count = map.size();

        while (j < n) {

            if (map.count(text[j]) > 0) {
                map[text[j]]--;
                if (map[text[j]] == 0) {
                    count--;
                }
            }

            if (j - i + 1 == k) {
                if (count == 0) {
                    list.push_back(i);
                }
                if (map.count(text[i]) > 0) {
                    map[text[i]] = map[text[i]] + 1;
                    if (map[text[i]] == 1) {
                        count++;
                    }
                }
                i++;
            }

            j++;
        }
        return list;
    }
};