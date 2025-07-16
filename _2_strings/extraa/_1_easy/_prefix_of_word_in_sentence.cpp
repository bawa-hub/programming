// https://leetcode.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/description/


class Solution {
public:
    int isPrefixOfWord(string sentence, string searchWord) {
        vector<string> words = split(sentence);
        for(int i=0;i<words.size();i++) {
            string temp = words[i];
            cout << temp << endl;
            int j = 0;
            if(temp.size() < searchWord.size()) continue;
            while((temp[j] == searchWord[j]) && j < searchWord.size()) j++;
            if(j == searchWord.size()) return i + 1;
        }

        return -1;
    }

    vector<string> split(string sentence) {
        vector<string> result; 
        for(int i=0;i<sentence.size();i++) {
            string temp = "";
            while(sentence[i] != ' ' && i < sentence.size()) {
                temp += sentence[i++];
            }
            result.push_back(temp);
            continue;
        }

        return result;
    }
};