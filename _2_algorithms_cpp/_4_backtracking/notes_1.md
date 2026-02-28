Permutation vs Combination

    Permutation → Order matters (e.g., arranging people in a line).
    Combination → Order does not matter (e.g., choosing a team).Example:

    Permutation: AB and BA are different.
    Combination: AB and BA are the same.

formula
  
  combination(nCr) = n!/r!(n-r)! [selecting r item out of n]
  permutation(nPr) = n!/(n-r)!   [selecting r item out of n and arranging them]

1️⃣ Understand the Basic Framework of Backtracking

def backtrack(current_state, choices):
    if solution_found(current_state):
        save_solution(current_state)
        return

    for choice in choices:
        if is_valid(choice, current_state):
            make_choice(choice, current_state)
            backtrack(current_state, new_choices)
            undo_choice(choice, current_state)  # BACKTRACK

2️⃣ Identify If It’s a Permutation or Combination Problem

    Permutation (Order Matters)
        Use a visited array to track used elements.
        If choosing r out of n, stop recursion when len(current_permutation) == r.

    Combination (Order Does Not Matter)
        Avoid reusing elements already picked.
        Use index-based recursion to prevent duplicate selections.

3️⃣ Steps to Solve Any Backtracking Problem on Permutations & Combinations

✅ Step 1: Choose - Pick a number/letter and add it to the current list.
✅ Step 2: Explore - Recursively continue to pick from remaining choices.
✅ Step 3: Unchoose (Backtrack) - Remove the last added choice to explore other possibilities.

4️⃣ Solving Permutations (With and Without Duplicates)

🔹 Example 1: All Permutations of Distinct Numbers
💡 Problem: Given nums = [1, 2, 3], find all permutations.

def permute(nums):
    res = []  # Store all permutations

    def backtrack(path, used):
        if len(path) == len(nums):  # Base case: A full permutation is formed
            res.append(path[:])  # Store a copy
            return

        for i in range(len(nums)):
            if not used[i]:  # Ensure each number is used only once per permutation
                used[i] = True
                path.append(nums[i])
                backtrack(path, used)
                path.pop()  # BACKTRACK
                used[i] = False

    backtrack([], [False] * len(nums))
    return res

print(permute([1, 2, 3]))

✅ Trick: Use a used array to prevent using the same element twice in one permutation.

🔹 Example 2: Permutations with Duplicates (Avoiding Repeats)
💡 Problem: Given nums = [1, 1, 2], find all unique permutations.

def permuteUnique(nums):
    res = []
    nums.sort()  # Sort to handle duplicates

    def backtrack(path, used):
        if len(path) == len(nums):
            res.append(path[:])
            return
        
        for i in range(len(nums)):
            if used[i]: 
                continue  # Skip already used numbers
            if i > 0 and nums[i] == nums[i - 1] and not used[i - 1]:
                continue  # Skip duplicate numbers

            used[i] = True
            path.append(nums[i])
            backtrack(path, used)
            path.pop()
            used[i] = False  # BACKTRACK

    backtrack([], [False] * len(nums))
    return res

print(permuteUnique([1, 1, 2]))

✅ Trick: Sort the array and use if i > 0 and nums[i] == nums[i - 1] and not used[i - 1] to skip duplicates.

5️⃣ Solving Combinations (With and Without Duplicates)

🔹 Example 3: All Combinations (Pick r Out of n)
💡 Problem: Given n = 4, k = 2, find all combinations of size k.

def combine(n, k):
    res = []

    def backtrack(start, path):
        if len(path) == k:  # Base case: Found a valid combination
            res.append(path[:])
            return

        for i in range(start, n + 1):  # Start from `start` to avoid duplicates
            path.append(i)
            backtrack(i + 1, path)  # Move to next index
            path.pop()  # BACKTRACK

    backtrack(1, [])
    return res

print(combine(4, 2))

✅ Trick: Use start index to avoid duplicates and prevent picking numbers out of order.

🔹 Example 4: Combination Sum (Pick Numbers to Reach a Target)
💡 Problem: Given candidates = [2,3,6,7], find all ways to sum to target = 7.

def combinationSum(candidates, target):
    res = []

    def backtrack(start, path, remaining):
        if remaining == 0:  # Base case: Target reached
            res.append(path[:])
            return
        if remaining < 0:
            return  # Stop if sum exceeds target

        for i in range(start, len(candidates)):  # Start from current index to allow reuse
            path.append(candidates[i])
            backtrack(i, path, remaining - candidates[i])  # `i`, not `i+1`, allows reuse
            path.pop()  # BACKTRACK

    backtrack(0, [], target)
    return res

print(combinationSum([2, 3, 6, 7], 7))

✅ Trick: Use backtrack(i, path, remaining - candidates[i]) to allow reusing numbers.

6️⃣ Common Mistakes to Avoid

❌ Not handling duplicates properly (sort first and use skipping conditions).
❌ Not using a used[] array for permutations (prevents using the same element twice).
❌ Not correctly updating the start index in combinations (prevents duplicates).
❌ Forgetting to backtrack (pop() and used[i] = False) (causes incorrect results).