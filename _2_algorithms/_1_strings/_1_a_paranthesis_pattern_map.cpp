#include<iostream>
#include<stack>
using namespace std;

/** 

There are 5 core patterns.
Parentheses Problems
│
├── Pattern 1: Validity Checking
├── Pattern 2: Remove Invalid Parentheses
├── Pattern 3: Longest Valid Substring
├── Pattern 4: Depth / Score Problems
└── Pattern 5: Index Tracking Problems

*/

/** Pattern 1 — Validity Checking

Core Idea
Check if parentheses are balanced.
Use a stack.

Rule:

(  -> push
)  -> pop

If stack becomes invalid → string is invalid. 




Example Problem
Problem: Valid Parentheses
Use stack:

stack = []

( -> push
( -> push
) -> pop
) -> pop

Stack empty → valid.

Problems in this Pattern:
Valid Parentheses
Minimum Add to Make Parentheses Valid
Minimum Remove to Make Valid Parentheses

*/

// Template
bool isValid(string s) {
    stack<char> st;

    for(char c : s) {

        if(c == '(')
            st.push(c);

        else {
            if(st.empty())
                return false;

            st.pop();
        }
    }

    return st.empty();
}

/**

Pattern 2 — Remove Invalid Parentheses

Core Idea
We must delete the minimum characters to make the string valid.

Two approaches:

Greedy + stack
or
BFS (hard problems)

Example
"(a(b)c)"
Remove invalid parentheses → valid result.

Key Idea

First pass:
remove extra ')'

Second pass:
remove extra '('

Famous Problems:
Minimum Remove to Make Valid Parentheses
Remove Invalid Parentheses (hard)

Why This Works

First pass ensures:
No extra ')'

Second pass ensures:
No extra '('

Final string becomes valid parentheses.
*/

// Template (Greedy Two Pass)
string removeInvalid(string s) {

    string temp;
    int balance = 0;

    // Pass 1 — Remove extra )
    for(char c : s) {

        if(c == '(') {
            balance++;
            temp += c;
        }

        else if(c == ')') {

            if(balance == 0)
                continue;

            balance--;
            temp += c;
        }

        else
            temp += c;
    }

    // Pass 2 — Remove extra (
        string result;

    for(int i = temp.size()-1; i >= 0; i--) {

        if(temp[i] == '(' && balance > 0) {
            balance--;
            continue;
        }

        result += temp[i];
    }

    reverse(result.begin(), result.end());

    return result;
}

/** 
Pattern 3 — Longest Valid Parentheses

Core Idea
Find the longest valid substring.

Important insight:
We must track indices.
Use stack of indices.

Example
)()())

Longest valid paranthesis:
()()

Length:
4

*/

// Template
void t()
{
    string s = ")()())";

    stack<int> st;
    st.push(-1);

    int maxLen = 0;

    for (int i = 0; i < s.size(); i++)
    {

        if (s[i] == '(')
            st.push(i);

        else
        {

            st.pop();

            if (st.empty())
                st.push(i);
            else
                maxLen = max(maxLen, i - st.top());
        }
    }
}

/** 

Pattern 4 — Depth / Score Problems

These problems care about nesting level.
Use a depth counter.

depth++
depth--

Example Problems:
Remove Outermost Parentheses
Maximum Nesting Depth
Score of Parentheses

Template:

for(char c : s) {

    if(c == '(')
        depth++;

    else
        depth--;
}

The trick is what you do at each depth.

Example:
depth == 0 → primitive finished
depth == 1 → outer parentheses
*/

/*

Pattern 5 — Index Tracking Problems

These problems require matching pairs and positions.

Examples:
Minimum Remove to Make Valid Parentheses
Reverse Substrings Between Each Pair

Use stack storing:
index

Example:
(a(bc)d)

Reverse inside parentheses.
Stack helps track where substring starts.

Template:

stack<int> st;
for(int i = 0; i < s.size(); i++) {

    if(s[i] == '(') {
        st.push(i);
    }

    else if(s[i] == ')') {

        if(!st.empty()) {
            int start = st.top();
            st.pop();

            // process substring start..i
        }
    }
}

Example Use Case

Problem:
Reverse Substrings Between Each Pair of Parentheses

Example:
(a(bc)d)

Steps:
push index of '('
when ')' appears → reverse substring


Another Example (Tracking invalid indices)
Minimum Remove to Make Valid Parentheses

Template:

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
*/