/*
 * 50 Easy String Problems - Implementations
 * Matches the order and numbering in STRING_50_EASY.md
 */

#include <algorithm>
#include <cctype>
#include <numeric>
#include <sstream>
#include <string>
#include <unordered_map>
#include <unordered_set>
#include <utility>
#include <vector>

using namespace std;

// -----------------------------------------------------------------------------
// A) Basics / Traversal / Indexing (1–10)
// -----------------------------------------------------------------------------

// 1. Reverse String (344) - in-place
class SolutionE1 {
public:
    void reverseString(vector<char>& s) {
        int l = 0, r = (int)s.size() - 1;
        while (l < r) {
            swap(s[l], s[r]);
            ++l;
            --r;
        }
    }
};

// 2. Reverse Words in a String III (557)
class SolutionE2 {
public:
    string reverseWords(string s) {
        int n = (int)s.size();
        int start = 0;
        for (int i = 0; i <= n; ++i) {
            if (i == n || s[i] == ' ') {
                reverse(s.begin() + start, s.begin() + i);
                start = i + 1;
            }
        }
        return s;
    }
};

// 3. Valid Palindrome (125)
class SolutionE3 {
public:
    bool isPalindrome(string s) {
        int l = 0, r = (int)s.size() - 1;
        while (l < r) {
            while (l < r && !isalnum((unsigned char)s[l])) ++l;
            while (l < r && !isalnum((unsigned char)s[r])) --r;
            if (l < r) {
                if (tolower((unsigned char)s[l]) != tolower((unsigned char)s[r])) return false;
                ++l;
                --r;
            }
        }
        return true;
    }
};

// 4. Valid Palindrome II (680)
class SolutionE4 {
public:
    bool validPalindrome(string s) {
        int l = 0, r = (int)s.size() - 1;
        while (l < r) {
            if (s[l] == s[r]) {
                ++l;
                --r;
            } else {
                return isPal(s, l + 1, r) || isPal(s, l, r - 1);
            }
        }
        return true;
    }

private:
    bool isPal(const string& s, int l, int r) {
        while (l < r) {
            if (s[l++] != s[r--]) return false;
        }
        return true;
    }
};

// 5. Detect Capital (520)
class SolutionE5 {
public:
    bool detectCapitalUse(string word) {
        int upper = 0;
        for (char c : word) {
            if (isupper((unsigned char)c)) upper++;
        }
        if (upper == 0 || upper == (int)word.size()) return true;
        if (upper == 1 && isupper((unsigned char)word[0])) return true;
        return false;
    }
};

// 6. To Lower Case (709)
class SolutionE6 {
public:
    string toLowerCase(string s) {
        for (char& c : s) {
            if (c >= 'A' && c <= 'Z') c = char(c - 'A' + 'a');
        }
        return s;
    }
};

// 7. Length of Last Word (58)
class SolutionE7 {
public:
    int lengthOfLastWord(string s) {
        int i = (int)s.size() - 1;
        while (i >= 0 && s[i] == ' ') --i;
        int len = 0;
        while (i >= 0 && s[i] != ' ') {
            ++len;
            --i;
        }
        return len;
    }
};

// 8. Implement strStr() (28)
class SolutionE8 {
public:
    int strStr(string haystack, string needle) {
        int n = (int)haystack.size(), m = (int)needle.size();
        if (m == 0) return 0;
        if (m > n) return -1;
        for (int i = 0; i + m <= n; ++i) {
            int j = 0;
            while (j < m && haystack[i + j] == needle[j]) ++j;
            if (j == m) return i;
        }
        return -1;
    }
};

// 9. Excel Sheet Column Title (168)
class SolutionE9 {
public:
    string convertToTitle(int columnNumber) {
        string res;
        while (columnNumber > 0) {
            columnNumber--;
            int rem = columnNumber % 26;
            res.push_back(char('A' + rem));
            columnNumber /= 26;
        }
        reverse(res.begin(), res.end());
        return res;
    }
};

// 10. Excel Sheet Column Number (171)
class SolutionE10 {
public:
    int titleToNumber(string columnTitle) {
        long long ans = 0;
        for (char c : columnTitle) {
            ans = ans * 26 + (c - 'A' + 1);
        }
        return (int)ans;
    }
};

// -----------------------------------------------------------------------------
// B) Frequency / Hashing (11–22)
// -----------------------------------------------------------------------------

// 11. Valid Anagram (242)
class SolutionE11 {
public:
    bool isAnagram(string s, string t) {
        if (s.size() != t.size()) return false;
        int cnt[26] = {0};
        for (char c : s) cnt[c - 'a']++;
        for (char c : t) {
            if (--cnt[c - 'a'] < 0) return false;
        }
        return true;
    }
};

// 12. Ransom Note (383)
class SolutionE12 {
public:
    bool canConstruct(string ransomNote, string magazine) {
        int cnt[26] = {0};
        for (char c : magazine) cnt[c - 'a']++;
        for (char c : ransomNote) {
            if (--cnt[c - 'a'] < 0) return false;
        }
        return true;
    }
};

// 13. First Unique Character in a String (387)
class SolutionE13 {
public:
    int firstUniqChar(string s) {
        int cnt[26] = {0};
        for (char c : s) cnt[c - 'a']++;
        for (int i = 0; i < (int)s.size(); ++i) {
            if (cnt[s[i] - 'a'] == 1) return i;
        }
        return -1;
    }
};

// 14. Find the Difference (389)
class SolutionE14 {
public:
    char findTheDifference(string s, string t) {
        char x = 0;
        for (char c : s) x ^= c;
        for (char c : t) x ^= c;
        return x;
    }
};

// 15. Isomorphic Strings (205)
class SolutionE15 {
public:
    bool isIsomorphic(string s, string t) {
        if (s.size() != t.size()) return false;
        char map1[256] = {0};
        char map2[256] = {0};
        for (int i = 0; i < (int)s.size(); ++i) {
            unsigned char a = (unsigned char)s[i];
            unsigned char b = (unsigned char)t[i];
            if (map1[a] == 0 && map2[b] == 0) {
                map1[a] = b;
                map2[b] = a;
            } else {
                if (map1[a] != b || map2[b] != a) return false;
            }
        }
        return true;
    }
};

// 16. Word Pattern (290)
class SolutionE16 {
public:
    bool wordPattern(string pattern, string s) {
        vector<string> words;
        string word;
        stringstream ss(s);
        while (ss >> word) words.push_back(word);
        if (words.size() != pattern.size()) return false;

        unordered_map<char, string> p2w;
        unordered_map<string, char> w2p;
        for (int i = 0; i < (int)pattern.size(); ++i) {
            char p = pattern[i];
            const string& w = words[i];
            if (!p2w.count(p) && !w2p.count(w)) {
                p2w[p] = w;
                w2p[w] = p;
            } else {
                if (p2w[p] != w || w2p[w] != p) return false;
            }
        }
        return true;
    }
};

// 17. Check if One String Swap Can Make Strings Equal (1790)
class SolutionE17 {
public:
    bool areAlmostEqual(string s1, string s2) {
        if (s1.size() != s2.size()) return false;
        vector<int> diff;
        for (int i = 0; i < (int)s1.size(); ++i) {
            if (s1[i] != s2[i]) diff.push_back(i);
            if (diff.size() > 2) return false;
        }
        if (diff.empty()) return true;
        if (diff.size() != 2) return false;
        swap(s1[diff[0]], s1[diff[1]]);
        return s1 == s2;
    }
};

// 18. Determine if String Halves Are Alike (1704)
class SolutionE18 {
public:
    bool halvesAreAlike(string s) {
        auto isVowel = [](char c) {
            c = tolower((unsigned char)c);
            return c=='a'||c=='e'||c=='i'||c=='o'||c=='u';
        };
        int n = (int)s.size(), half = n / 2;
        int a = 0, b = 0;
        for (int i = 0; i < half; ++i) if (isVowel(s[i])) a++;
        for (int i = half; i < n; ++i) if (isVowel(s[i])) b++;
        return a == b;
    }
};

// 19. Check if All Characters Have Equal Number of Occurrences (1941)
class SolutionE19 {
public:
    bool areOccurrencesEqual(string s) {
        int cnt[26] = {0};
        for (char c : s) cnt[c - 'a']++;
        int val = 0;
        for (int x : cnt) {
            if (x == 0) continue;
            if (val == 0) val = x;
            else if (x != val) return false;
        }
        return true;
    }
};

// 20. Find Words That Can Be Formed by Characters (1160)
class SolutionE20 {
public:
    int countCharacters(vector<string>& words, string chars) {
        int base[26] = {0};
        for (char c : chars) base[c - 'a']++;
        int total = 0;
        for (auto& w : words) {
            int cur[26];
            copy(begin(base), end(base), cur);
            bool ok = true;
            for (char c : w) {
                if (--cur[c - 'a'] < 0) { ok = false; break; }
            }
            if (ok) total += (int)w.size();
        }
        return total;
    }
};

// 21. Jewels and Stones (771)
class SolutionE21 {
public:
    int numJewelsInStones(string jewels, string stones) {
        bool isJewel[128] = {false};
        for (char c : jewels) isJewel[(unsigned char)c] = true;
        int cnt = 0;
        for (char c : stones) if (isJewel[(unsigned char)c]) cnt++;
        return cnt;
    }
};

// 22. Unique Morse Code Words (804)
class SolutionE22 {
public:
    int uniqueMorseRepresentations(vector<string>& words) {
        static const string codes[26] = {
            ".-","-...","-.-.","-..",".","..-.","--.","....","..",
            ".---","-.-",".-..","--","-.","---",".--.","--.-",".-.",
            "...","-","..-","...-",".--","-..-","-.--","--.."
        };
        unordered_set<string> seen;
        for (auto& w : words) {
            string t;
            for (char c : w) t += codes[c - 'a'];
            seen.insert(t);
        }
        return (int)seen.size();
    }
};

// -----------------------------------------------------------------------------
// C) Two Pointers / Simple Greedy (23–32)
// -----------------------------------------------------------------------------

// 23. Backspace String Compare (844)
class SolutionE23 {
public:
    bool backspaceCompare(string s, string t) {
        return build(s) == build(t);
    }
private:
    string build(const string& s) {
        string out;
        for (char c : s) {
            if (c == '#') {
                if (!out.empty()) out.pop_back();
            } else {
                out.push_back(c);
            }
        }
        return out;
    }
};

// 24. Merge Strings Alternately (1768)
class SolutionE24 {
public:
    string mergeAlternately(string word1, string word2) {
        string res;
        int i = 0, j = 0, n1 = (int)word1.size(), n2 = (int)word2.size();
        while (i < n1 || j < n2) {
            if (i < n1) res.push_back(word1[i++]);
            if (j < n2) res.push_back(word2[j++]);
        }
        return res;
    }
};

// 25. Reverse Prefix of Word (2000)
class SolutionE25 {
public:
    string reversePrefix(string word, char ch) {
        int idx = -1;
        for (int i = 0; i < (int)word.size(); ++i) {
            if (word[i] == ch) { idx = i; break; }
        }
        if (idx == -1) return word;
        reverse(word.begin(), word.begin() + idx + 1);
        return word;
    }
};

// 26. Rotated Digits (788)
class SolutionE26 {
public:
    int rotatedDigits(int n) {
        auto good = [](int x) {
            bool diff = false;
            while (x > 0) {
                int d = x % 10;
                if (d==3||d==4||d==7) return false;
                if (d==2||d==5||d==6||d==9) diff = true;
                x /= 10;
            }
            return diff;
        };
        int cnt = 0;
        for (int i = 1; i <= n; ++i) if (good(i)) cnt++;
        return cnt;
    }
};

// 27. Long Pressed Name (925)
class SolutionE27 {
public:
    bool isLongPressedName(string name, string typed) {
        int i = 0, j = 0, n = (int)name.size(), m = (int)typed.size();
        while (j < m) {
            if (i < n && name[i] == typed[j]) {
                ++i; ++j;
            } else if (j > 0 && typed[j] == typed[j - 1]) {
                ++j;
            } else return false;
        }
        return i == n;
    }
};

// 28. Check If Two String Arrays are Equivalent (1662)
class SolutionE28 {
public:
    bool arrayStringsAreEqual(vector<string>& word1, vector<string>& word2) {
        int i = 0, j = 0, p = 0, q = 0;
        while (i < (int)word1.size() && j < (int)word2.size()) {
            if (word1[i][p] != word2[j][q]) return false;
            if (++p == (int)word1[i].size()) { p = 0; ++i; }
            if (++q == (int)word2[j].size()) { q = 0; ++j; }
        }
        return i == (int)word1.size() && j == (int)word2.size();
    }
};

// 29. Valid Parentheses (20)
class SolutionE29 {
public:
    bool isValid(string s) {
        string st;
        for (char c : s) {
            if (c=='('||c=='['||c=='{') st.push_back(c);
            else {
                if (st.empty()) return false;
                char o = st.back(); st.pop_back();
                if ((c==')' && o!='(') || (c==']' && o!='[') || (c=='}' && o!='{')) return false;
            }
        }
        return st.empty();
    }
};

// 30. Remove All Adjacent Duplicates In String (1047)
class SolutionE30 {
public:
    string removeDuplicates(string s) {
        string st;
        for (char c : s) {
            if (!st.empty() && st.back() == c) st.pop_back();
            else st.push_back(c);
        }
        return st;
    }
};

// 31. Make The String Great (1544)
class SolutionE31 {
public:
    string makeGood(string s) {
        string st;
        for (char c : s) {
            if (!st.empty() && abs(st.back() - c) == 32) st.pop_back();
            else st.push_back(c);
        }
        return st;
    }
};

// 32. Check If a Word Occurs As a Prefix of Any Word in a Sentence (1455)
class SolutionE32 {
public:
    int isPrefixOfWord(string sentence, string searchWord) {
        string w;
        stringstream ss(sentence);
        int idx = 1;
        while (ss >> w) {
            if (w.rfind(searchWord, 0) == 0) return idx;
            idx++;
        }
        return -1;
    }
};

// -----------------------------------------------------------------------------
// D) String Building / Parsing (33–42)
// -----------------------------------------------------------------------------

// 33. Add Strings (415)
class SolutionE33 {
public:
    string addStrings(string num1, string num2) {
        int i = (int)num1.size() - 1, j = (int)num2.size() - 1;
        int carry = 0;
        string res;
        while (i >= 0 || j >= 0 || carry) {
            int sum = carry;
            if (i >= 0) sum += num1[i--] - '0';
            if (j >= 0) sum += num2[j--] - '0';
            res.push_back(char('0' + (sum % 10)));
            carry = sum / 10;
        }
        reverse(res.begin(), res.end());
        return res;
    }
};

// 34. Add Binary (67)
class SolutionE34 {
public:
    string addBinary(string a, string b) {
        int i = (int)a.size() - 1, j = (int)b.size() - 1;
        int carry = 0;
        string res;
        while (i >= 0 || j >= 0 || carry) {
            int sum = carry;
            if (i >= 0) sum += a[i--] - '0';
            if (j >= 0) sum += b[j--] - '0';
            res.push_back(char('0' + (sum & 1)));
            carry = sum >> 1;
        }
        reverse(res.begin(), res.end());
        return res;
    }
};

// 35. Roman to Integer (13)
class SolutionE35 {
public:
    int romanToInt(string s) {
        unordered_map<char, int> val{
            {'I',1},{'V',5},{'X',10},{'L',50},{'C',100},{'D',500},{'M',1000}
        };
        int n = (int)s.size();
        int ans = 0;
        for (int i = 0; i < n; ++i) {
            int v = val[s[i]];
            if (i + 1 < n && v < val[s[i + 1]]) ans -= v;
            else ans += v;
        }
        return ans;
    }
};

// 36. Integer to Roman (12)
class SolutionE36 {
public:
    string intToRoman(int num) {
        vector<pair<int,string>> mp = {
            {1000,"M"},{900,"CM"},{500,"D"},{400,"CD"},
            {100,"C"},{90,"XC"},{50,"L"},{40,"XL"},
            {10,"X"},{9,"IX"},{5,"V"},{4,"IV"},{1,"I"}
        };
        string res;
        for (auto& [v, sym] : mp) {
            while (num >= v) {
                num -= v;
                res += sym;
            }
        }
        return res;
    }
};

// 37. Valid Word Abbreviation (408)
class SolutionE37 {
public:
    bool validWordAbbreviation(string word, string abbr) {
        int i = 0, j = 0, n = (int)word.size(), m = (int)abbr.size();
        while (i < n && j < m) {
            if (isdigit((unsigned char)abbr[j])) {
                if (abbr[j] == '0') return false; // leading zero
                int num = 0;
                while (j < m && isdigit((unsigned char)abbr[j])) {
                    num = num * 10 + (abbr[j] - '0');
                    j++;
                }
                i += num;
            } else {
                if (word[i] != abbr[j]) return false;
                i++; j++;
            }
        }
        return i == n && j == m;
    }
};

// 38. Number of Segments in a String (434)
class SolutionE38 {
public:
    int countSegments(string s) {
        int cnt = 0;
        int n = (int)s.size();
        for (int i = 0; i < n; ++i) {
            if (s[i] != ' ' && (i == 0 || s[i-1] == ' ')) cnt++;
        }
        return cnt;
    }
};

// 39. Replace All Digits with Characters (1844)
class SolutionE39 {
public:
    string replaceDigits(string s) {
        for (int i = 1; i < (int)s.size(); i += 2) {
            s[i] = char(s[i - 1] + (s[i] - '0'));
        }
        return s;
    }
};

// 40. Check if the Sentence Is Pangram (1832)
class SolutionE40 {
public:
    bool checkIfPangram(string sentence) {
        bool seen[26] = {false};
        for (char c : sentence) {
            if (c >= 'a' && c <= 'z') seen[c - 'a'] = true;
        }
        for (bool b : seen) if (!b) return false;
        return true;
    }
};

// 41. Maximum Number of Words Found in Sentences (2114)
class SolutionE41 {
public:
    int mostWordsFound(vector<string>& sentences) {
        int best = 0;
        for (auto& s : sentences) {
            int spaces = 0;
            for (char c : s) if (c == ' ') spaces++;
            best = max(best, spaces + 1);
        }
        return best;
    }
};

// 42. Goal Parser Interpretation (1678)
class SolutionE42 {
public:
    string interpret(string command) {
        string res;
        int n = (int)command.size();
        for (int i = 0; i < n; ) {
            if (command[i] == 'G') {
                res.push_back('G');
                ++i;
            } else if (i + 1 < n && command[i] == '(' && command[i + 1] == ')') {
                res.push_back('o');
                i += 2;
            } else {
                res += "al";
                i += 4;
            }
        }
        return res;
    }
};

// -----------------------------------------------------------------------------
// E) Small Pattern Matching / Simple Logic (43–50)
// -----------------------------------------------------------------------------

// 43. Repeated Substring Pattern (459)
class SolutionE43 {
public:
    bool repeatedSubstringPattern(string s) {
        string t = s + s;
        return t.substr(1, t.size() - 2).find(s) != string::npos;
    }
};

// 44. Longest Common Prefix (14)
class SolutionE44 {
public:
    string longestCommonPrefix(vector<string>& strs) {
        if (strs.empty()) return "";
        string prefix = strs[0];
        for (int i = 1; i < (int)strs.size(); ++i) {
            while (!prefix.empty() && strs[i].find(prefix) != 0) {
                prefix.pop_back();
            }
            if (prefix.empty()) break;
        }
        return prefix;
    }
};

// 45. Valid Word Square (422)
class SolutionE45 {
public:
    bool validWordSquare(vector<string>& words) {
        int n = (int)words.size();
        for (int i = 0; i < n; ++i) {
            for (int j = 0; j < (int)words[i].size(); ++j) {
                if (j >= n || i >= (int)words[j].size()) return false;
                if (words[i][j] != words[j][i]) return false;
            }
        }
        return true;
    }
};

// 46. String Matching in an Array (1408)
class SolutionE46 {
public:
    vector<string> stringMatching(vector<string>& words) {
        vector<string> res;
        int n = (int)words.size();
        for (int i = 0; i < n; ++i) {
            for (int j = 0; j < n; ++j) {
                if (i == j) continue;
                if (words[j].find(words[i]) != string::npos) {
                    res.push_back(words[i]);
                    break;
                }
            }
        }
        return res;
    }
};

// 47. Find the Index of the First Occurrence in a String (28) - same as E8
class SolutionE47 {
public:
    int strStr(string haystack, string needle) {
        int n = (int)haystack.size(), m = (int)needle.size();
        if (m == 0) return 0;
        if (m > n) return -1;
        for (int i = 0; i + m <= n; ++i) {
            int j = 0;
            while (j < m && haystack[i + j] == needle[j]) ++j;
            if (j == m) return i;
        }
        return -1;
    }
};

// 48. Check if a String Is an Acronym of Words (2828)
class SolutionE48 {
public:
    bool isAcronym(vector<string>& words, string s) {
        if (words.size() != s.size()) return false;
        for (int i = 0; i < (int)words.size(); ++i) {
            if (words[i].empty() || words[i][0] != s[i]) return false;
        }
        return true;
    }
};

// 49. Shuffle String (1528)
class SolutionE49 {
public:
    string restoreString(string s, vector<int>& indices) {
        string res(s.size(), ' ');
        for (int i = 0; i < (int)indices.size(); ++i) {
            res[indices[i]] = s[i];
        }
        return res;
    }
};

// 50. Defanging an IP Address (1108)
class SolutionE50 {
public:
    string defangIPaddr(string address) {
        string res;
        for (char c : address) {
            if (c == '.') res += "[.]";
            else res.push_back(c);
        }
        return res;
    }
};

