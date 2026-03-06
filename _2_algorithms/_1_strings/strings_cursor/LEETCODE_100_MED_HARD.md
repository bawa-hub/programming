# LeetCode 100 (Medium + Hard) — Practice + Revision Sheet

Goal: a **high-signal checklist** of 100 Medium/Hard problems with **implementable solution outlines** (not full code) so you can quickly revise patterns and re-solve.

How to use:
- Pick 5/day: 3 medium + 2 hard.
- For each: write the approach + complexity from memory, then implement in C++.
- Mark ✅ when you can solve in < 25 min (M) / < 45 min (H) without looking.

Legend:
- **Patt** = primary pattern
- **Idea** = minimal algorithmic plan (the “core trick”)
- **TC/SC** = time/space complexity (typical)

---

## A) Arrays / Two Pointers / Sliding Window (20)

1. **(M) 3Sum (15)**  
   - **Patt**: sort + 2 pointers  
   - **Idea**: sort; for each `i`, do `l/r` sweep to find `-a[i]`, skip duplicates.  
   - **TC/SC**: $O(n^2)$ / $O(1)$ extra (output aside)

2. **(M) 3Sum Closest (16)**  
   - **Patt**: sort + 2 pointers  
   - **Idea**: like 3Sum, but track best diff to target.  
   - **TC/SC**: $O(n^2)$ / $O(1)$

3. **(M) Container With Most Water (11)**  
   - **Patt**: 2 pointers  
   - **Idea**: `l/r`; area = min(h[l],h[r])*(r-l); move smaller height inward.  
   - **TC/SC**: $O(n)$ / $O(1)$

4. **(M) Sort Colors (75)**  
   - **Patt**: Dutch National Flag  
   - **Idea**: 3 pointers: `low/mid/high`; swap 0s left, 2s right.  
   - **TC/SC**: $O(n)$ / $O(1)$

5. **(M) Subarray Sum Equals K (560)**  
   - **Patt**: prefix sum + hashmap count  
   - **Idea**: running sum `s`; add `count[s-k]` to answer; then `count[s]++`.  
   - **TC/SC**: $O(n)$ / $O(n)$

6. **(M) Continuous Subarray Sum (523)**  
   - **Patt**: prefix mod + earliest index  
   - **Idea**: store first index for `sum%k`; if seen before and gap>=2 => true.  
   - **TC/SC**: $O(n)$ / $O(min(n,k))$

7. **(M) Minimum Size Subarray Sum (209)**  
   - **Patt**: variable window (positive nums)  
   - **Idea**: expand `r` until sum>=target, shrink `l` to minimize length.  
   - **TC/SC**: $O(n)$ / $O(1)$

8. **(H) Minimum Window Substring (76)**  
   - **Patt**: sliding window + frequency deficit  
   - **Idea**: expand until window covers `t` counts; shrink while valid; track best.  
   - **TC/SC**: $O(n)$ / $O(Σ)$

9. **(M) Longest Substring Without Repeating Characters (3)**  
   - **Patt**: sliding window with last position  
   - **Idea**: `l=max(l,last[c]+1)`; update `last[c]=r`; maximize `r-l+1`.  
   - **TC/SC**: $O(n)$ / $O(Σ)$

10. **(M) Longest Repeating Character Replacement (424)**  
   - **Patt**: sliding window  
   - **Idea**: window valid if `(len - maxFreq) <= k`; expand and shrink.  
   - **TC/SC**: $O(n)$ / $O(Σ)$

11. **(M) Permutation in String (567)**  
   - **Patt**: fixed window + freq  
   - **Idea**: maintain count diff array for 26; slide window length = |s1|.  
   - **TC/SC**: $O(n)$ / $O(1)$

12. **(M) Find All Anagrams in a String (438)**  
   - **Patt**: fixed window  
   - **Idea**: same as above; collect start indices when counts match.  
   - **TC/SC**: $O(n)$ / $O(1)$

13. **(M) Max Consecutive Ones III (1004)**  
   - **Patt**: sliding window (at most k zeros)  
   - **Idea**: keep `zeros`; while `zeros>k` move `l`.  
   - **TC/SC**: $O(n)$ / $O(1)$

14. **(M) Fruit Into Baskets (904)**  
   - **Patt**: sliding window with at most 2 types  
   - **Idea**: hashmap counts; shrink until map size<=2.  
   - **TC/SC**: $O(n)$ / $O(1..n)$

15. **(H) Trapping Rain Water (42)**  
   - **Patt**: two pointers or monotonic stack  
   - **Idea**: 2p with `lMax/rMax`; move side with smaller max; add trapped.  
   - **TC/SC**: $O(n)$ / $O(1)$

16. **(M) Next Permutation (31)**  
   - **Patt**: permutation trick  
   - **Idea**: find first decreasing from right; swap with just larger; reverse suffix.  
   - **TC/SC**: $O(n)$ / $O(1)$

17. **(M) Rotate Array (189)**  
   - **Patt**: reverse segments  
   - **Idea**: reverse all; reverse first k; reverse rest.  
   - **TC/SC**: $O(n)$ / $O(1)$

18. **(M) Find the Duplicate Number (287)**  
   - **Patt**: Floyd cycle detection  
   - **Idea**: treat `nums[i]` as next pointer; find cycle entry.  
   - **TC/SC**: $O(n)$ / $O(1)$

19. **(H) First Missing Positive (41)**  
   - **Patt**: in-place hashing  
   - **Idea**: place value `v` in index `v-1` via swaps; then scan first mismatch.  
   - **TC/SC**: $O(n)$ / $O(1)$

20. **(M) Jump Game (55)**  
   - **Patt**: greedy reachability  
   - **Idea**: keep farthest reachable; if `i>far` fail; update `far=max(far,i+nums[i])`.  
   - **TC/SC**: $O(n)$ / $O(1)$

---

## B) Hashing / Frequency / Prefix Tricks (10)

21. **(M) Group Anagrams (49)**  
   - **Patt**: hashing signature  
   - **Idea**: key = 26-count vector (or sorted string); map key -> list.  
   - **TC/SC**: $O(nk)$ / $O(nk)$

22. **(M) Top K Frequent Elements (347)**  
   - **Patt**: bucket / heap  
   - **Idea**: count freq; bucket by freq 1..n; gather from high to low.  
   - **TC/SC**: $O(n)$ / $O(n)$

23. **(M) Longest Consecutive Sequence (128)**  
   - **Patt**: hash set  
   - **Idea**: for each x if (x-1 not in set) expand x,x+1,...  
   - **TC/SC**: $O(n)$ avg / $O(n)$

24. **(M) Valid Sudoku (36)**  
   - **Patt**: hashing sets  
   - **Idea**: 9 row sets + 9 col sets + 9 box sets; detect duplicates.  
   - **TC/SC**: $O(1)$ / $O(1)$

25. **(M) Insert Delete GetRandom O(1) (380)**  
   - **Patt**: hashmap + dynamic array  
   - **Idea**: vector store values; map val->idx; delete by swap-with-last.  
   - **TC/SC**: $O(1)$ avg / $O(n)$

26. **(M) Random Pick with Weight (528)**  
   - **Patt**: prefix sums + binary search  
   - **Idea**: build prefix; pick `r in [1,sum]`; lower_bound prefix>=r.  
   - **TC/SC**: build $O(n)$; query $O(log n)$ / $O(n)$

27. **(M) Find All Duplicates in an Array (442)**  
   - **Patt**: sign marking  
   - **Idea**: for each v, index=abs(v)-1; if nums[idx]<0 => duplicate else negate.  
   - **TC/SC**: $O(n)$ / $O(1)$

28. **(M) Subarray Sums Divisible by K (974)**  
   - **Patt**: prefix mod counts  
   - **Idea**: answer += count[mod]; then count[mod]++. Normalize negative mod.  
   - **TC/SC**: $O(n)$ / $O(k)$

29. **(H) Longest Duplicate Substring (1044)**  
   - **Patt**: binary search length + rolling hash (Rabin-Karp)  
   - **Idea**: check(L): rolling hash all substrings length L, detect collisions carefully.  
   - **TC/SC**: ~ $O(n log n)$ / $O(n)$

30. **(H) Substring with Concatenation of All Words (30)**  
   - **Patt**: sliding window by word length  
   - **Idea**: for offset 0..wlen-1: move in steps of wlen, maintain word counts.  
   - **TC/SC**: $O(n)$-ish / $O(m)$

---

## C) Stacks / Monotonic Stack / Intervals (10)

31. **(M) Daily Temperatures (739)**  
   - **Patt**: monotonic decreasing stack  
   - **Idea**: stack indices; while curr>temp[st.top] pop and set answer.  
   - **TC/SC**: $O(n)$ / $O(n)$

32. **(M) Next Greater Element II (503)**  
   - **Patt**: monotonic stack + circular array  
   - **Idea**: iterate 2n; use modulo index; resolve next greater via stack.  
   - **TC/SC**: $O(n)$ / $O(n)$

33. **(M) Evaluate Reverse Polish Notation (150)**  
   - **Patt**: stack  
   - **Idea**: push numbers; on operator pop 2, apply, push result.  
   - **TC/SC**: $O(n)$ / $O(n)$

34. **(M) Decode String (394)**  
   - **Patt**: stack (counts + prev string)  
   - **Idea**: push (prev, k); build curr; on ']' pop and expand.  
   - **TC/SC**: $O(n)$ / $O(n)$

35. **(H) Largest Rectangle in Histogram (84)**  
   - **Patt**: monotonic increasing stack  
   - **Idea**: push indices; on drop, pop height, compute width with new top.  
   - **TC/SC**: $O(n)$ / $O(n)$

36. **(H) Maximal Rectangle (85)**  
   - **Patt**: histogram per row  
   - **Idea**: build heights for each row; run histogram algo each row.  
   - **TC/SC**: $O(R*C)$ / $O(C)$

37. **(M) Merge Intervals (56)**  
   - **Patt**: sort intervals  
   - **Idea**: sort by start; merge if overlap else push new.  
   - **TC/SC**: $O(n log n)$ / $O(n)$

38. **(M) Insert Interval (57)**  
   - **Patt**: interval sweep  
   - **Idea**: add all ending before new; merge overlaps; add rest.  
   - **TC/SC**: $O(n)$ / $O(n)$

39. **(M) Non-overlapping Intervals (435)**  
   - **Patt**: greedy by end time  
   - **Idea**: sort by end; keep interval with earliest end; count removals.  
   - **TC/SC**: $O(n log n)$ / $O(1)$

40. **(M) Car Fleet (853)**  
   - **Patt**: sort + stack of times  
   - **Idea**: sort by position desc; compute time; if time<=stack top merge else new fleet.  
   - **TC/SC**: $O(n log n)$ / $O(n)$

---

## D) Binary Search / Divide & Conquer (10)

41. **(M) Search in Rotated Sorted Array (33)**  
   - **Patt**: binary search with rotated logic  
   - **Idea**: identify sorted half each step; decide where target lies.  
   - **TC/SC**: $O(log n)$ / $O(1)$

42. **(M) Find Minimum in Rotated Sorted Array (153)**  
   - **Patt**: binary search  
   - **Idea**: compare mid with right; shrink toward pivot.  
   - **TC/SC**: $O(log n)$ / $O(1)$

43. **(M) Kth Largest Element in an Array (215)**  
   - **Patt**: quickselect  
   - **Idea**: partition around pivot; recurse into side containing kth.  
   - **TC/SC**: avg $O(n)$ / $O(1)$

44. **(M) Find Peak Element (162)**  
   - **Patt**: binary search on slope  
   - **Idea**: if nums[mid] < nums[mid+1] go right else left.  
   - **TC/SC**: $O(log n)$ / $O(1)$

45. **(M) Search a 2D Matrix (74)**  
   - **Patt**: binary search in flattened array  
   - **Idea**: treat as 1D index -> (r=i/cols,c=i%cols).  
   - **TC/SC**: $O(log(RC))$ / $O(1)$

46. **(M) Find K Closest Elements (658)**  
   - **Patt**: binary search window  
   - **Idea**: binary search left bound in [0,n-k]; compare x-a[mid] vs a[mid+k]-x.  
   - **TC/SC**: $O(log(n-k)+k)$ / $O(1)$

47. **(H) Median of Two Sorted Arrays (4)**  
   - **Patt**: binary search partition  
   - **Idea**: partition A and B so left size = (m+n+1)/2 and maxLeft<=minRight.  
   - **TC/SC**: $O(log min(m,n))$ / $O(1)$

48. **(H) Split Array Largest Sum (410)**  
   - **Patt**: binary search answer  
   - **Idea**: check(mid): greedy count partitions with max sum <= mid; shrink range.  
   - **TC/SC**: $O(n log(sum))$ / $O(1)$

49. **(H) Find in Mountain Array (1095)**  
   - **Patt**: 3 binary searches  
   - **Idea**: find peak; binary search ascending left; binary search descending right.  
   - **TC/SC**: $O(log n)$ / $O(1)$

50. **(H) Capacity To Ship Packages Within D Days (1011)**  
   - **Patt**: binary search answer  
   - **Idea**: check(cap): greedy days count; adjust cap range.  
   - **TC/SC**: $O(n log(sum))$ / $O(1)$

---

## E) Linked List (5)

51. **(M) Add Two Numbers (2)**  
   - **Patt**: simulation  
   - **Idea**: digit-wise add with carry; build new list.  
   - **TC/SC**: $O(n)$ / $O(1)$ extra

52. **(M) Remove Nth Node From End (19)**  
   - **Patt**: 2 pointers  
   - **Idea**: advance fast n; then move both until fast hits end; remove slow->next.  
   - **TC/SC**: $O(n)$ / $O(1)$

53. **(M) Reorder List (143)**  
   - **Patt**: split + reverse + merge  
   - **Idea**: find mid; reverse second half; weave alternately.  
   - **TC/SC**: $O(n)$ / $O(1)$

54. **(H) Merge k Sorted Lists (23)**  
   - **Patt**: heap or divide&conquer  
   - **Idea**: min-heap of heads; pop/push next; or merge pairs.  
   - **TC/SC**: $O(N log k)$ / $O(k)$

55. **(H) Reverse Nodes in k-Group (25)**  
   - **Patt**: pointer manipulation  
   - **Idea**: check k nodes exist; reverse segment; connect; repeat.  
   - **TC/SC**: $O(n)$ / $O(1)$

---

## F) Trees / BST (15)

56. **(M) Binary Tree Level Order Traversal (102)**  
   - **Patt**: BFS  
   - **Idea**: queue; process by level size.  
   - **TC/SC**: $O(n)$ / $O(n)$

57. **(M) Validate Binary Search Tree (98)**  
   - **Patt**: DFS with bounds  
   - **Idea**: each node must satisfy (low, high); recurse with updated bounds.  
   - **TC/SC**: $O(n)$ / $O(h)$

58. **(M) Kth Smallest in BST (230)**  
   - **Patt**: inorder traversal  
   - **Idea**: iterative stack inorder; count nodes until k.  
   - **TC/SC**: $O(h+k)$ / $O(h)$

59. **(M) Construct Binary Tree from Preorder and Inorder (105)**  
   - **Patt**: recursion + hashmap index  
   - **Idea**: root=pre[pi]; split inorder by rootIndex; recurse left/right by sizes.  
   - **TC/SC**: $O(n)$ / $O(n)$

60. **(M) Lowest Common Ancestor of a Binary Tree (236)**  
   - **Patt**: DFS  
   - **Idea**: return node if matches p/q; if both sides non-null => current is LCA.  
   - **TC/SC**: $O(n)$ / $O(h)$

61. **(M) Binary Tree Right Side View (199)**  
   - **Patt**: BFS/DFS  
   - **Idea**: level order; take last node each level (or DFS prioritize right).  
   - **TC/SC**: $O(n)$ / $O(n)$

62. **(M) Path Sum III (437)**  
   - **Patt**: prefix sums on tree  
   - **Idea**: DFS with running sum; map of prefix counts; add count[sum-target].  
   - **TC/SC**: $O(n)$ / $O(h)$

63. **(H) Binary Tree Maximum Path Sum (124)**  
   - **Patt**: DFS DP  
   - **Idea**: for each node compute best downward gain; update global with left+node+right.  
   - **TC/SC**: $O(n)$ / $O(h)$

64. **(H) Serialize and Deserialize Binary Tree (297)**  
   - **Patt**: BFS or preorder + null markers  
   - **Idea**: serialize with separators and `#`; deserialize by reading tokens recursively/queue.  
   - **TC/SC**: $O(n)$ / $O(n)$

65. **(H) Count of Smaller Numbers After Self (315)**  
   - **Patt**: BIT/Fenwick or merge sort counting  
   - **Idea**: coordinate compress; traverse from right; query prefix counts.  
   - **TC/SC**: $O(n log n)$ / $O(n)$

66. **(M) Implement Trie (Prefix Tree) (208)**  
   - **Patt**: Trie  
   - **Idea**: node with children[26]/map; insert/search/startsWith.  
   - **TC/SC**: $O(L)$ / $O(total chars)$

67. **(H) Word Search II (212)**  
   - **Patt**: Trie + DFS pruning  
   - **Idea**: build trie of words; DFS board; follow trie edges; mark visited; collect terminal.  
   - **TC/SC**: depends / trie+recursion

68. **(M) Delete Node in a BST (450)**  
   - **Patt**: BST recursion  
   - **Idea**: find node; if 2 children swap with inorder successor; delete successor.  
   - **TC/SC**: $O(h)$ / $O(h)$

69. **(H) Recover Binary Search Tree (99)**  
   - **Patt**: inorder anomaly detection  
   - **Idea**: inorder traversal; find two nodes where order breaks; swap values.  
   - **TC/SC**: $O(n)$ / $O(h)$ (or $O(1)$ Morris)

70. **(H) All Nodes Distance K in Binary Tree (863)**  
   - **Patt**: graph conversion + BFS  
   - **Idea**: parent pointers via DFS; BFS from target with visited until distance K.  
   - **TC/SC**: $O(n)$ / $O(n)$

---

## G) Graphs / BFS / DFS / Topo / Union-Find (15)

71. **(M) Number of Islands (200)**  
   - **Patt**: DFS/BFS flood fill  
   - **Idea**: iterate cells; when land, BFS/DFS mark visited.  
   - **TC/SC**: $O(RC)$ / $O(RC)$

72. **(M) Rotting Oranges (994)**  
   - **Patt**: multi-source BFS  
   - **Idea**: push all rotten; BFS by layers; count fresh; track minutes.  
   - **TC/SC**: $O(RC)$ / $O(RC)$

73. **(M) Course Schedule (207)**  
   - **Patt**: cycle detection / topo  
   - **Idea**: Kahn (indegree queue) or DFS states (0/1/2).  
   - **TC/SC**: $O(V+E)$ / $O(V+E)$

74. **(M) Course Schedule II (210)**  
   - **Patt**: topo sort  
   - **Idea**: Kahn; output order; if size!=V return empty.  
   - **TC/SC**: $O(V+E)$ / $O(V+E)$

75. **(M) Clone Graph (133)**  
   - **Patt**: DFS/BFS + hashmap old->new  
   - **Idea**: create node copies on demand; connect neighbors.  
   - **TC/SC**: $O(V+E)$ / $O(V)$

76. **(M) Accounts Merge (721)**  
   - **Patt**: DSU / Union-Find  
   - **Idea**: union emails of same account; group by root; sort emails.  
   - **TC/SC**: ~$O(N α(N))$ / $O(N)$

77. **(M) Network Delay Time (743)**  
   - **Patt**: Dijkstra  
   - **Idea**: adjacency list; min-heap distances; relax edges.  
   - **TC/SC**: $O(E log V)$ / $O(V+E)$

78. **(H) Word Ladder (127)**  
   - **Patt**: BFS  
   - **Idea**: precompute wildcard patterns (`h*t`) -> words; BFS levels.  
   - **TC/SC**: $O(N*L)$ build + BFS / $O(N*L)$

79. **(H) Word Ladder II (126)**  
   - **Patt**: BFS layers + backtracking paths  
   - **Idea**: BFS to build parent lists per word on shortest depth; then DFS build sequences.  
   - **TC/SC**: heavy / heavy

80. **(H) Alien Dictionary (269)**  
   - **Patt**: build graph + topo  
   - **Idea**: edges from first differing char; invalid if prefix issue; topo sort.  
   - **TC/SC**: $O(total chars)$ / $O(Σ+E)$

81. **(H) Minimum Number of Refueling Stops (871)**  
   - **Patt**: greedy + max-heap  
   - **Idea**: drive until fuel insufficient; add all reachable stations’ fuel to max-heap; refuel best.  
   - **TC/SC**: $O(n log n)$ / $O(n)$

82. **(M) Evaluate Division (399)**  
   - **Patt**: graph weighted edges  
   - **Idea**: build graph a->b weight; query via DFS/BFS multiplying weights.  
   - **TC/SC**: build $O(E)$; query $O(V+E)$ / $O(V+E)$

83. **(H) Critical Connections in a Network (1192)**  
   - **Patt**: Tarjan bridges  
   - **Idea**: DFS with tin/low; if low[v] > tin[u], (u,v) is bridge.  
   - **TC/SC**: $O(V+E)$ / $O(V+E)$

84. **(H) Shortest Path in a Grid with Obstacles Elimination (1293)**  
   - **Patt**: BFS state = (r,c,kRemaining)  
   - **Idea**: visited[r][c] = max remaining; only proceed if better.  
   - **TC/SC**: $O(R*C*K)$ / $O(R*C*K)$

85. **(H) Swim in Rising Water (778)**  
   - **Patt**: Dijkstra / binary search + BFS  
   - **Idea**: min-heap on max-elevation-so-far; pop best, relax neighbors with max(curr,grid).  
   - **TC/SC**: $O(n^2 log n)$ / $O(n^2)$

---

## H) Dynamic Programming (15)

86. **(M) House Robber (198)**  
   - **Patt**: 1D DP  
   - **Idea**: dp[i]=max(dp[i-1], dp[i-2]+a[i]); optimize to 2 vars.  
   - **TC/SC**: $O(n)$ / $O(1)$

87. **(M) Coin Change (322)**  
   - **Patt**: unbounded knapsack  
   - **Idea**: dp[amt]=min(dp[amt], dp[amt-coin]+1).  
   - **TC/SC**: $O(n*amount)$ / $O(amount)$

88. **(M) Longest Increasing Subsequence (300)**  
   - **Patt**: patience sorting  
   - **Idea**: maintain tails; lower_bound for each x; replace/extend.  
   - **TC/SC**: $O(n log n)$ / $O(n)$

89. **(M) Partition Equal Subset Sum (416)**  
   - **Patt**: 0/1 knapsack  
   - **Idea**: target=sum/2; dp boolean; iterate nums, dp backwards.  
   - **TC/SC**: $O(n*target)$ / $O(target)$

90. **(M) Unique Paths (62)**  
   - **Patt**: grid DP  
   - **Idea**: dp[r][c]=dp[r-1][c]+dp[r][c-1].  
   - **TC/SC**: $O(RC)$ / $O(C)$

91. **(M) Longest Palindromic Substring (5)**  
   - **Patt**: expand around center (or DP/Manacher)  
   - **Idea**: for each center i (odd/even), expand while equal, track best.  
   - **TC/SC**: $O(n^2)$ / $O(1)$

92. **(M) Palindromic Substrings (647)**  
   - **Patt**: expand centers  
   - **Idea**: count expansions for all centers.  
   - **TC/SC**: $O(n^2)$ / $O(1)$

93. **(M) Decode Ways (91)**  
   - **Patt**: DP on string  
   - **Idea**: dp[i]=ways; add single-digit if valid; add two-digit if 10..26.  
   - **TC/SC**: $O(n)$ / $O(1)$

94. **(M) Word Break (139)**  
   - **Patt**: DP + dictionary  
   - **Idea**: dp[i]=true if exists j<i with dp[j] and s[j..i) in dict.  
   - **TC/SC**: $O(n^2)$ / $O(n)$

95. **(H) Word Break II (140)**  
   - **Patt**: DFS + memo  
   - **Idea**: recursion from index; memo index -> list sentences; combine with words.  
   - **TC/SC**: exponential output / memo helps

96. **(H) Edit Distance (72)**  
   - **Patt**: 2D DP  
   - **Idea**: dp[i][j]=min(insert,delete,replace) with base dp[0][*], dp[*][0].  
   - **TC/SC**: $O(nm)$ / $O(min(n,m))$

97. **(H) Regular Expression Matching (10)**  
   - **Patt**: DP  
   - **Idea**: dp[i][j] match s[0..i), p[0..j); handle '.' and '*' cases.  
   - **TC/SC**: $O(nm)$ / $O(nm)$

98. **(H) Distinct Subsequences (115)**  
   - **Patt**: DP counting  
   - **Idea**: dp[i][j]=# ways s[0..i) form t[0..j); if char match add dp[i-1][j-1].  
   - **TC/SC**: $O(nm)$ / $O(m)$

99. **(H) Burst Balloons (312)**  
   - **Patt**: interval DP  
   - **Idea**: add 1 borders; dp[l][r]=max over k last burst in (l,r): dp[l][k]+dp[k][r]+a[l]*a[k]*a[r].  
   - **TC/SC**: $O(n^3)$ / $O(n^2)$

100. **(H) Longest Increasing Path in a Matrix (329)**  
   - **Patt**: DFS + memo  
   - **Idea**: dp[r][c]=1+max(dp[nr][nc]) for neighbors with greater value; memoize.  
   - **TC/SC**: $O(RC)$ / $O(RC)$

---

## Extra: quick “pattern reminders”

- **Sliding window**: keep window *valid* via counts/constraints; when valid, try to shrink.
- **Prefix sum + map**: `ans += count[prefix - need]` is the core move.
- **Monotonic stack**: stack maintains increasing/decreasing property to answer “next greater” / “range where this is minimum”.
- **Binary search on answer**: when check(mid) is monotonic (true/false boundary), search mid.
- **Tree DP**: return “best downward contribution”; update global answer separately.
- **Graph BFS**: for shortest in unweighted, BFS; for weighted positive, Dijkstra.
- **DP**: define state precisely, write transitions, then optimize space.

