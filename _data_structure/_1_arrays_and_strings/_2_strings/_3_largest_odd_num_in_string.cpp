// https://leetcode.com/problems/largest-odd-number-in-string/

class Solution
{
public:
    string largestOddNumber(string num)
    {
        int index = -1;
        for (int i = num.length() - 1; i >= 0; i--)
        {
            int n = num[i];
            if (n % 2 != 0)
            {
                index = i;
                break;
            }
        }
        string res = "";
        for (int i = 0; i <= index; i++)
        {
            res += num[i];
        }
        return res;
    }
};