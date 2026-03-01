// https://leetcode.com/problems/minimum-window-substring/
// https://leetcode.com/problems/minimum-window-substring/solutions/26808/here-is-a-10-line-template-that-can-solve-most-substring-problems/

       string minWindow(string s, string t) {

        vector<int> map(128,0);
        for(auto c: t) map[c]++;
        int counter=t.size(), i=0, j=0, mini=INT_MAX, head=0;

        while(j<s.size()){
            if(map[s[j]]-- > 0) counter--; 

            while(counter==0){ 
                if(j-i+1<mini)  mini=j- (head=i) +1; // head is used to extract the substring
                if(map[s[i]]==0) counter++; 
                map[s[i]]++;
                i++;
            } 

             j++; 
        }

        return mini==INT_MAX? "":s.substr(head, mini);
    }