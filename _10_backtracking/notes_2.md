🔥 Backtracking Framework

1️⃣ Identify Key Elements

    Choices: What are the possible options at each step?
    Constraints: What conditions must be met?
    Goal: When should we stop and record a solution?

2️⃣ General Backtracking Template

void backtrack(parameters) {
    // 🛑 BASE CASE: Stop when a valid solution is found
    if (goal reached) {
        store solution;
        return;
    }

    // 🔄 ITERATE OVER CHOICES
    for (each possible choice) {
        if (choice is valid) {  // ✅ Check constraints
            make choice;         // ✅ Modify state
            backtrack(new parameters); // 🚀 Recursive call
            undo choice;         // 🔙 BACKTRACK to previous state
        }
    }
}

🔥 Step-by-Step Approach

1️⃣ Define a Recursive Function

    This function explores all possible solutions.

2️⃣ Define Base Case (Stopping Condition)

    When a valid solution is reached, store it and return.

3️⃣ Loop Over Possible Choices

    Try every possibility that hasn't been used yet.

4️⃣ Check Constraints

    If a choice is invalid, skip it.

5️⃣ Make the Choice (Modify State)

    Add the choice to the current path.

6️⃣ Recur with the New State

    Call backtrack() with the updated parameters.

7️⃣ Undo the Choice (Backtrack)

    Remove the last choice to explore a new path.


🔥 Framework Applied to Different Problems

1️⃣ Permutations (All Orderings)

void backtrack(vector<int>& nums, vector<bool>& used, vector<int>& path, vector<vector<int>>& res) {
    if (path.size() == nums.size()) {  // 🎯 Base case: full permutation
        res.push_back(path);
        return;
    }

    for (int i = 0; i < nums.size(); i++) {
        if (!used[i]) { // ✅ Check constraints
            used[i] = true;
            path.push_back(nums[i]);
            backtrack(nums, used, path, res);
            path.pop_back();  // 🔙 BACKTRACK
            used[i] = false;
        }
    }
}

✅ Trick: Use a used[] array to track visited numbers.

2️⃣ Combinations (Choose k Elements)

void backtrack(int start, int n, int k, vector<int>& path, vector<vector<int>>& res) {
    if (path.size() == k) {  // 🎯 Base case: exact `k` elements
        res.push_back(path);
        return;
    }

    for (int i = start; i <= n; i++) {
        path.push_back(i);
        backtrack(i + 1, n, k, path, res);
        path.pop_back();  // 🔙 BACKTRACK
    }
}

✅ Trick: Use start to prevent duplicates.

3️⃣ Subsets (Power Set)

void backtrack(vector<int>& nums, int start, vector<int>& path, vector<vector<int>>& res) {
    res.push_back(path); // Store every subset

    for (int i = start; i < nums.size(); i++) {
        path.push_back(nums[i]);
        backtrack(nums, i + 1, path, res);
        path.pop_back();  // 🔙 BACKTRACK
    }
}

✅ Trick: Always add path to res to store all subsets.

4️⃣ Combination Sum (Choose Numbers to Reach Target)

void backtrack(vector<int>& candidates, int start, int target, vector<int>& path, vector<vector<int>>& res) {
    if (target == 0) {  // 🎯 Base case: found a valid sum
        res.push_back(path);
        return;
    }
    if (target < 0) return;  // Stop if sum exceeds target

    for (int i = start; i < candidates.size(); i++) {
        path.push_back(candidates[i]);
        backtrack(candidates, i, target - candidates[i], path, res); // `i` allows reuse
        path.pop_back();  // 🔙 BACKTRACK
    }
}

✅ Trick: Reuse numbers by not increasing i.

5️⃣ N-Queens Problem (Place N Queens on an N x N Chessboard)
bool isSafe(vector<string>& board, int row, int col, int n) {
    for (int i = 0; i < row; i++) {
        if (board[i][col] == 'Q') return false; // Check column
        if (col - (row - i) >= 0 && board[i][col - (row - i)] == 'Q') return false; // Check left diagonal
        if (col + (row - i) < n && board[i][col + (row - i)] == 'Q') return false; // Check right diagonal
    }
    return true;
}

void backtrack(vector<string>& board, int row, int n, vector<vector<string>>& res) {
    if (row == n) { // 🎯 Base case: placed `n` queens
        res.push_back(board);
        return;
    }

    for (int col = 0; col < n; col++) {
        if (isSafe(board, row, col, n)) { // ✅ Check constraints
            board[row][col] = 'Q';
            backtrack(board, row + 1, n, res);
            board[row][col] = '.';  // 🔙 BACKTRACK
        }
    }
}
✅ Trick: Use a helper function (isSafe()) to check constraints.

🔥 Pro Tips

1️⃣ Recognizing a Backtracking Problem

✅ Use backtracking when:

    The problem requires generating all possible solutions.
    You need to make choices, then undo them if they fail.
    Constraints must be followed at each step.

2️⃣ Optimizing Backtracking

🚀 Prune the Search Space Early

    Sort input (if applicable) to detect duplicate solutions.
    Use hash sets or bit masking to speed up checking constraints.
    Use start index in combinations to avoid unnecessary recursion.

🚀 Memoization + Backtracking

    If the problem has overlapping subproblems, store solutions in a cache to avoid recomputation.

🚀 Iterative vs Recursive

    Recursive backtracking is easier to write.
    Iterative (stack-based) is used when recursion depth is too large.

🔥 Summary: One-Page Cheat Sheet

Problem Type	Base Case	   Recursion	Trick
Permutation	    path.size() == n	Loop over unused elements	Track used[]
Combination	path.size() == k	Start from i+1	Prevent duplicates
Subset	Always add path	Start from i+1	Store all paths
Combination Sum	target == 0	Continue at i (reuse)	Stop if target < 0
N-Queens	row == n	Place Q, check isSafe()	Use helper function

