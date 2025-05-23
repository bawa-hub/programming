// https://leetcode.com/problems/combinations/
// https://leetcode.com/problems/combinations/solutions/27006/a-template-to-those-combination-problems/
// https://leetcode.com/problems/combinations/editorial/


// https://www.geeksforgeeks.org/permutations-and-combinations/
// Combination Formula is used to choose ‘r’ components out of a total number of ‘n’ components, and is given by:
// nCr = n!/r!(n-r)!

// chatgpt
class Solution {
public:
    vector<vector<int>> result;
    
    void backtrack(int start, int n, int k, vector<int>& comb) {
        // ✅ Base Case: combination of size k
        if (comb.size() == k) {
            result.push_back(comb);
            return;
        }

        for (int i = start; i <= n; ++i) {
            comb.push_back(i);                    // choose
            backtrack(i + 1, n, k, comb);         // explore
            comb.pop_back();                      // un-choose (backtrack)
        }
    }

    vector<vector<int>> combine(int n, int k) {
        vector<int> comb;
        backtrack(1, n, k, comb);
        return result;
    }
};


// recursive
class Solution
{
public:
    vector<vector<int>> combine(int n, int k)
    {
        vector<vector<int>> res;
        vector<int> comb;
        backtrack(res, 1, n, k, comb);
        return res;
    }

    void backtrack(vector<vector<int>> &res, int cur, int n, int k, vector<int> &comb)
    {
        if (k == 0)
        {
            res.push_back(comb);
            return;
        }

        // pick
        comb.push_back(cur);
        backtrack(res, cur + 1, n, k - 1, comb);
        comb.pop_back();

        // not pick
        // If cur>n-k, there are not enough numbers left, we have to select the current element
        if (cur <= n - k)
        {
            backtrack(res, cur + 1, n, k, comb);
        }
    }
};

// iterative
class Solution
{
public:
    vector<vector<int>> combine(int n, int k)
    {
        vector<vector<int>> result;
        int i = 0;
        vector<int> p(k, 0);
        while (i >= 0)
        {
            p[i]++;
            if (p[i] > n)
                --i;
            else if (i == k - 1)
                result.push_back(p);
            else
            {
                ++i;
                p[i] = p[i - 1];
            }
        }
        return result;
    }
};