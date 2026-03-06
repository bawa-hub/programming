#include <iostream>
#include <stack>
using namespace std;

// Below is the FAANG-style Parentheses Mastery Set (10 Problems).
// If you truly understand these with patterns, you can solve almost every parentheses question in interviews.


// 1️⃣ Valid Parentheses (Foundation)
// Pattern: Validity Checking (Stack)
// https://leetcode.com/problems/valid-parentheses/description/

// Idea
// Match opening and closing parentheses.

// Algorithm
// push opening bracket
// if closing → check stack top

bool isValid(string s) {
    stack<char> st;

    for(char c : s) {

        if(c=='(' || c=='{' || c=='[')
            st.push(c);

        else {

            if(st.empty()) return false;

            char top = st.top();
            st.pop();

            if((c==')' && top!='(') ||
               (c=='}' && top!='{') ||
               (c==']' && top!='['))
                return false;
        }
    }

    return st.empty();
}
// Time: O(n)

// 2️⃣ Minimum Add to Make Parentheses Valid
// Pattern: Greedy + Counter
// https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/description/

// Idea
// Track unmatched parentheses.

// balance = open brackets
// If extra ) → need insertion.

int minAddToMakeValid(string s) {

    int balance = 0;
    int add = 0;

    for(char c : s) {

        if(c=='(')
            balance++;

        else {

            if(balance==0)
                add++;
            else
                balance--;
        }
    }

    return add + balance;
}

// 3️⃣ Remove Outermost Parentheses
// Pattern: Depth Tracking
// https://leetcode.com/problems/remove-outermost-parentheses/description/

// Idea
// Skip parentheses at depth = 1.

string removeOuterParentheses(string s) {

    int depth = 0;
    string res;

    for(char c : s) {

        if(c=='(') {

            if(depth>0)
                res+=c;

            depth++;
        }

        else {

            depth--;

            if(depth>0)
                res+=c;
        }
    }

    return res;
}

// 4️⃣ Maximum Nesting Depth
// Pattern: Depth Counter
// https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/description/

// Idea
// Track max depth.

int maxDepth(string s) {

    int depth = 0;
    int ans = 0;

    for(char c : s) {

        if(c=='(')
            depth++;

        ans = max(ans, depth);

        if(c==')')
            depth--;
    }

    return ans;
}

// 5️⃣ Minimum Remove to Make Valid Parentheses
// Pattern: Stack + Index
// https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/description/

// Idea
// Store invalid indices.

string minRemoveToMakeValid(string s) {

    stack<int> st;
    set<int> remove;

    for(int i=0;i<s.size();i++) {

        if(s[i]=='(')
            st.push(i);

        else if(s[i]==')') {

            if(st.empty())
                remove.insert(i);
            else
                st.pop();
        }
    }

    while(!st.empty()) {
        remove.insert(st.top());
        st.pop();
    }

    string res;

    for(int i=0;i<s.size();i++)
        if(!remove.count(i))
            res+=s[i];

    return res;
}

// 6️⃣ Score of Parentheses
// Pattern: Depth / Stack
// https://leetcode.com/problems/score-of-parentheses/description/

// Idea
// Rules:
// () = 1
// AB = A + B
// (A) = 2 * A

int scoreOfParentheses(string s) {

    stack<int> st;
    st.push(0);

    for(char c : s) {

        if(c=='(')
            st.push(0);

        else {

            int v = st.top();
            st.pop();

            int score = max(2*v,1);

            st.top() += score;
        }
    }

    return st.top();
}

// 7️⃣ Generate Parentheses
// Pattern: Backtracking
// https://leetcode.com/problems/generate-parentheses/description/

// Idea

// Generate all valid combinations.
// Rules:
// open < n
// close < open

void solve(int open,int close,int n,string cur,vector<string>&ans){

    if(cur.size()==2*n){
        ans.push_back(cur);
        return;
    }

    if(open<n)
        solve(open+1,close,n,cur+"(",ans);

    if(close<open)
        solve(open,close+1,n,cur+")",ans);
}

// 8️⃣ Longest Valid Parentheses
// Pattern: Stack + Index
// https://leetcode.com/problems/longest-valid-parentheses/description/

// Idea
// Track last invalid index.

int longestValidParentheses(string s) {

    stack<int> st;
    st.push(-1);

    int ans=0;

    for(int i=0;i<s.size();i++){

        if(s[i]=='(')
            st.push(i);

        else{

            st.pop();

            if(st.empty())
                st.push(i);

            else
                ans=max(ans,i-st.top());
        }
    }

    return ans;
}

// 9️⃣ Reverse Substrings Between Parentheses
// Pattern: Stack + String Manipulation
// https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/description/

// Idea:
// push current string
// reverse when closing bracket appears

// 🔟 Remove Invalid Parentheses (Hard)
// Pattern: BFS / Backtracking
// https://leetcode.com/problems/remove-invalid-parentheses/description/

// Idea:
// Try removing parentheses level by level
// Stop when valid string found