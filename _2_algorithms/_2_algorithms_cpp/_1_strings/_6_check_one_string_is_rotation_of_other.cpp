
class Solution
{
public:
    bool rotateString(string s, string goal)
    {
        if (s.size() != goal.size())
            return false;
        string check = s + s;
        return check.find(goal) != string::npos;
    }
};
//  Time Complexity: O(N^2), where N is the length of string(s).
//  Space Complexity: O(N), the space used building s+s.

class Solution
{
private:
    bool rotateString(string A, string B, int rotation)
    {
        for (int i = 0; i < A.length(); i++)
        {
            if (A[i] != B[(i + rotation) % B.length()])
            {
                return false;
            }
        }
        return true;
    }

public:
    bool rotateString(string s, string goal)
    {
        if (s.length() != goal.length())
        {
            return false;
        }
        if (s.length() == 0)
        {
            return true;
        }
        for (int i = 0; i < s.length(); i++)
        {
            if (rotateString(s, goal, i))
            {
                return true;
            }
        }
        return false;
    }
};
// Time Complexity: O(N^2), where N is the length of string(s). For each rotation string(goal), we check up to N
// elements in string(s) and string(goal).

// Space Complexity: O(1), Constant space. We only use pointers to elements of s and goal.