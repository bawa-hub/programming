// https://www.geeksforgeeks.org/problems/count-occurences-of-anagrams5839/1

class Solution {
  public:
int search(string &pat, string &txt) {
    int cnt = 0;
    unordered_map<char, int> mp;

    // Step 1: Count frequency of characters in pattern `pat`
    for (int i = 0; i < pat.size(); i++) {
        mp[pat[i]]++;
    }

    int char_count = mp.size(); // Number of unique characters in `pat`

    int i = 0, j = 0;
    while (j < txt.size()) {
        // Reduce the frequency of the current character in the window
        // if (mp.find(txt[j]) != mp.end()) {
            mp[txt[j]]--;
            if (mp[txt[j]] == 0) {
                char_count--; // Found one required character fully in the window
            }
        // }

        // When window size reaches the size of `pat`, check if it's a valid match
        if (j - i + 1 == pat.size()) {
            if (char_count == 0) {
                cnt++; // All characters of `pat` are present in the current window
            }

            // Move the window to the right by incrementing `i`
            if (mp.find(txt[i]) != mp.end()) {
                mp[txt[i]]++;
                if (mp[txt[i]] == 1) {
                    char_count++; // Restoring the count of characters at the left of the window
                }
            }
            i++; // Increment left pointer to shrink the window
        }

        j++; // Increment right pointer to expand the window
    }

    return cnt;
}