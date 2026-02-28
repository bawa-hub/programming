🧠 What is a Permutation?
    A permutation is just a rearrangement of things, where order matters.

📦 Example:

Say you have 3 toys: 🧸 🪀 🚗
Let's label them as: 1, 2, and 3.
All possible arrangements (aka permutations) are:

[1, 2, 3]
[1, 3, 2]
[2, 1, 3]
[2, 3, 1]
[3, 1, 2]
[3, 2, 1]

So there are 3! = 6 permutations.

🌱 Now Think Like a Tree (Recursive Thinking)

Let’s build permutations step by step like a decision tree.
Start with nothing: []

Now at each step, you pick 1 of the unused numbers.


🌳 Tree for [1, 2, 3]:
            []
        /    |    \
      1      2      3
     / \    / \    / \
   2   3  1   3  1   2
   |   |  |   |  |   |
   3   2  3   1  2   1

Each path from root to leaf is a complete permutation.

🧠 How Backtracking Helps

Backtracking means:
̌
    Pick a number (if it’s not already used)

    Go deeper (recursively)

    When done, remove the number (undo choice)

    Try next number

🔁 Step-by-Step for [1, 2, 3]:

Start: []

    Try 1 → [1]

        Try 2 → [1,2]

            Try 3 → [1,2,3] ✅ Store it!

            Go back: remove 3 → [1,2]

        Try 3 → [1,3], then 2 → [1,3,2] ✅

    Backtrack all: try 2 first now…

You explore all paths, and that’s why it’s called backtracking.    

🧱 C++ Template (Simplified):

void backtrack(vector<int>& nums, vector<bool>& used, vector<int>& path) {

    if (path.size() == nums.size()) {
        result.push_back(path);
        return;
    }

    for (int i = 0; i < nums.size(); i++) {
        if (used[i]) continue;

        // pick
        used[i] = true;
        path.push_back(nums[i]);

        // explore
        backtrack(nums, used, path);

        // unpick (backtrack)
        path.pop_back();
        used[i] = false;
    }
}

used[i] says if nums[i] has been added already
path builds one permutation at a time
You “try”, “go deeper”, and “undo” (classic backtracking)

We’ll simulate exactly what the recursive function does at every level with the values of path, used[], and what happens at each step.

⚙️ Initial Setup

nums = [1, 2, 3]
used = [false, false, false]
path = []

🟢 Level 0: Start

We try picking each of the 3 numbers one by one.
🔵 Level 1: Pick 1

path = [1]
used = [true, false, false]

Now go deeper...
🔵 Level 2: Pick 2

path = [1, 2]
used = [true, true, false]

Go deeper...
🔵 Level 3: Pick 3

path = [1, 2, 3]
used = [true, true, true]

All numbers used ✅ — store [1, 2, 3] in result.

Now backtrack!
🔙 Backtrack from [1, 2, 3] → Remove 3

path = [1, 2]
used = [true, true, false]

Try next option at this level → no more choices → backtrack more
🔙 Remove 2

path = [1]
used = [true, false, false]

Now try picking 3 instead
🔵 Level 2: Pick 3

path = [1, 3]
used = [true, false, true]

Go deeper…
🔵 Level 3: Pick 2

path = [1, 3, 2]
used = [true, true, true]

✅ Save [1, 3, 2]

Backtrack all the way to root.
🔙 Remove 3 → Remove 1

path = []
used = [false, false, false]

Now try picking 2 at Level 1
🔵 Level 1: Pick 2

path = [2]
used = [false, true, false]

🔵 Level 2: Pick 1

path = [2, 1]
used = [true, true, false]

🔵 Level 3: Pick 3

path = [2, 1, 3]
used = [true, true, true]

✅ Save it, backtrack!
🔙 Remove 3 → 1

Now try 3 before 1 at this level.
🔵 Path = [2, 3, 1] ✅

✅ Save it, backtrack all the way to root.
🔵 Pick 3 at Level 1

path = [3]
used = [false, false, true]

Then try [1] → [3,1] → [3,1,2] ✅

Backtrack

Try [2] → [3,2] → [3,2,1] ✅
✅ Final Result List

[
 [1, 2, 3],
 [1, 3, 2],
 [2, 1, 3],
 [2, 3, 1],
 [3, 1, 2],
 [3, 2, 1]
]

🔁 Summary of All Steps (like a move list)

1 → 2 → 3 ✅
1 → 3 → 2 ✅
2 → 1 → 3 ✅
2 → 3 → 1 ✅
3 → 1 → 2 ✅
3 → 2 → 1 ✅

Each path is built one element at a time, and we backtrack to try all possibilities.

🌲 Visual Tree of Permutations for [1, 2, 3]

                          []
               ┌──────────┼──────────┐
              [1]        [2]        [3]
           ┌──┴──┐     ┌──┴──┐     ┌──┴──┐
         [1,2] [1,3]  [2,1] [2,3] [3,1] [3,2]
         ┌──┘   └──┐  ┌──┘   └──┐  ┌──┘   └──┐
      [1,2,3] [1,3,2] [2,1,3] [2,3,1] [3,1,2] [3,2,1]

Each branch is a recursive call choosing one number.
Each leaf node is a complete permutation

🧱 Call Stack Simulation (Line-by-line)

Let’s use this simplified version of the C++ code:

void backtrack(vector<int>& nums, vector<bool>& used, vector<int>& path) {

    if (path.size() == nums.size()) {
        result.push_back(path);
        return;
    }

    for (int i = 0; i < nums.size(); i++) {
        if (used[i]) continue;

        used[i] = true;
        path.push_back(nums[i]);

        backtrack(nums, used, path);   // 🔁 RECURSION HAPPENS HERE

        path.pop_back();
        used[i] = false;
    }
}

🎯 Step-by-step for [1,2,3]:

📞 Call 1: path = []

Loop runs i = 0 → 1 → 2

✅ i = 0 → choose 1 → path = [1]
📞 Call 2: path = [1]

Loop i = 0 → 1 → 2

    i = 0 is used → skip

    i = 1 → choose 2 → path = [1,2]

📞 Call 3: path = [1,2]

Loop i = 0 → 1 → 2

    i = 0 and i = 1 used → skip

    i = 2 → choose 3 → path = [1,2,3]

✅ Base case: path.size() == nums.size()

Save [1,2,3] to result.

⬅️ Return to previous call → backtrack: pop 3, mark used[2]=false
🧪 This continues:

    Next try 3 before 2 at same depth

    Each recursive call is pushed on the call stack

    When it hits base case or finishes loop, it pops and backtracks

Each push/pop is like:

Call: path = [1,2,3] → Save
Return: path = [1,2]
Backtrack: remove 3

🧠 Visualize Call Stack Like Boxes

Top  ➤ [1,2,3] → SAVE
         ↑
       [1,2]
         ↑
        [1]
         ↑
         []

Each level is a new function call with:

    Different path
    Updated used[]


We’ll simulate the exact behavior step-by-step for nums = [1, 2, 3] with all key variables shown during each recursive call.

#include <iostream>
#include <vector>
using namespace std;

vector<vector<int>> result;

void backtrack(vector<int>& nums, vector<bool>& used, vector<int>& path) {

    if (path.size() == nums.size()) {
        result.push_back(path);
        return;
    }

    for (int i = 0; i < nums.size(); i++) {
        if (used[i]) continue;

        // pick
        used[i] = true;
        path.push_back(nums[i]);

        // explore
        backtrack(nums, used, path);

        // unpick (backtrack)
        path.pop_back();
        used[i] = false;
    }
}

vector<vector<int>> permute(vector<int>& nums) {

    vector<bool> used(nums.size(), false);
    vector<int> path;
    backtrack(nums, used, path);
    return result;
}

▶️ Call 1: permute([1, 2, 3])

    used = [false, false, false]

    path = []

▶️ Call 2: backtrack([], [F, F, F])

Try i = 0 → 1 → 2

    i=0 is not used
    → pick 1
    → path = [1]
    → used = [T, F, F]
    → go deeper

▶️ Call 3: backtrack([1], [T, F, F])

Try i = 0 → 1 → 2

    i=0 used → skip

    i=1 → pick 2
    → path = [1, 2]
    → used = [T, T, F]
    → go deeper

▶️ Call 4: backtrack([1, 2], [T, T, F])

Try i = 0 → 1 → 2

    i=2 → pick 3
    → path = [1, 2, 3]
    → used = [T, T, T]

✅ path.size() == nums.size()
👉 save: [1, 2, 3] to result
⬅️ Backtrack from Call 4

    pop 3, mark used[2]=false

    back to path [1, 2]

Loop ends → backtrack again
⬅️ Backtrack from Call 3

    pop 2, mark used[1]=false

    try i=2 → pick 3

→ path = [1, 3]
→ used = [T, F, T]
→ go deeper
▶️ Call 5: backtrack([1, 3], [T, F, T])

Try i=0 → i=1

    i=1 → pick 2
    → path = [1, 3, 2]
    → used = [T, T, T]

✅ save [1, 3, 2]

⬅️ backtrack: pop 2, unmark
⬅️ pop 3, unmark

This goes on... in total we’ll generate:

[1, 2, 3]
[1, 3, 2]
[2, 1, 3]
[2, 3, 1]
[3, 1, 2]
[3, 2, 1]

🧠 Key Learning Points

    Every recursive call adds a level to the call stack

    When base case hits → result is saved

    Then we undo last move (backtrack) and try other options

    used[i] is like a visited marker


