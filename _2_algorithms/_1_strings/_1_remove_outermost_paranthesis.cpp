// https://leetcode.com/problems/remove-outermost-parentheses/

#include <iostream>
using namespace std;

class Solution
{
public:
    string removeOuterParentheses(string s)
    {

        stack<char> st;
        string ans = "";

        for (auto c : s)
        {

            if (!st.empty())
                ans += c;
            
            if (c == '(')
                st.push(c);
            else
            {
                st.pop();
                if (st.empty())
                    ans.pop_back();
            }
        }

        return ans;
    }
    // Time complexity: O(N)
    // Space complexity: O(N)

    string removeOuterParenthesesOptimized(string S) {
        int count = 0;
        string str;
        for (char c : S) {
            if (c == '(') {
                if (count++) {
                    str += '(';
                }
            } else {
                if (--count) {
                    str += ')';
                }
            }
        }
        return str;
    }
    // TC: O(n)
// SC: O(1)
};


// Pattern Behind This Problem: Parentheses / Depth Tracking pattern.

// Key idea:
// Track nesting depth

// This pattern appears in problems like:

// Maximum nesting depth
// Score of parentheses
// Longest valid parentheses
// Remove outermost parentheses

// Only 5 patterns cover almost every parentheses question on LeetCode.
// 20  Valid Parentheses
// 32  Longest Valid Parentheses
// 921 Minimum Add to Make Parentheses Valid
// 1249 Minimum Remove to Make Valid Parentheses
// 1021 Remove Outermost Parentheses
// 856 Score of Parentheses

// What Does “Depth” Mean in Parentheses?
// Depth = how many parentheses are currently open.
// In other words:
// Depth tells you how deeply nested you are inside parentheses.
// Think of it like entering and leaving rooms.
// "(" → you enter a room → depth increases
// ")" → you leave a room → depth decreases

// Nested Example (())
// | char | action | depth |
// | ---- | ------ | ----- |
// | (    | open   | 1     |
// | (    | open   | 2     |
// | )    | close  | 1     |
// | )    | close  | 0     |

// More Complex Example
// (()(()))
// | char | depth |
// | ---- | ----- |
// | (    | 1     |
// | (    | 2     |
// | )    | 1     |
// | (    | 2     |
// | (    | 3     |
// | )    | 2     |
// | )    | 1     |
// | )    | 0     |

// Code Representation
// Depth is usually tracked with a counter.

// int depth = 0;

// for(char c : s) {

//     if(c == '(')
//         depth++;

//     if(c == ')')
//         depth--;
// }

// Why Depth Is Important
// Depth tells you:
// 1️⃣ Which parentheses are outermost
// 2️⃣ How deeply nested you are
// 3️⃣ When a primitive substring ends

// in above problem (()())
// ( depth=1
// ( depth=2
// ) depth=1
// ( depth=2
// ) depth=1
// ) depth=0

// Important moment:
// depth = 0
// This means:
// one primitive block finished

// (()())(())
// Depth becomes 0 twice → meaning two primitive groups.

// Interview Tip
// Whenever you see parentheses questions, think:
// Track depth
// Track stack
// Almost every solution uses one of these two.

// The Two Ways Parentheses Problems Are Solved
// Almost every parentheses problem uses one of these two tools:

// 1️⃣ Depth Counter
// depth++
// depth--

// Used when we only care about nesting level.

// Example problems:
// Remove outer parentheses
// Maximum depth
// Score of parentheses

// 2️⃣ Stack

// Used when we must match positions.

// Example problems:
// Valid parentheses
// Longest valid parentheses
// Remove invalid parentheses

// The Golden Rule
// Whenever you see parentheses problems ask:

// Do I need nesting level?
// or
// Do I need matching pairs?

// If nesting → depth
// If matching → stack