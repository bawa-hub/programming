// https://leetcode.com/problems/string-to-integer-atoi/description/
// https://practice.geeksforgeeks.org/problems/implement-atoi/1
// https://www.geeksforgeeks.org/write-your-own-atoi/
// https://www.geeksforgeeks.org/recursive-implementation-of-atoi/

// iterative
int myAtoi(string str)
{
    int sign = 1, base = 0, i = 0;
    while (str[i] == ' ')
    {
        i++;
    }
    if (str[i] == '-' || str[i] == '+')
    {
        sign = 1 - 2 * (str[i++] == '-');
    }

    while (str[i] >= '0' && str[i] <= '9')
    {
        if (base > INT_MAX / 10 || (base == INT_MAX / 10 && str[i] - '0' > 7))
        {
            if (sign == 1)
                return INT_MAX;
            else
                return INT_MIN;
        }
        base = 10 * base + (str[i++] - '0');
    }
    return base * sign;
}
// TC: O(n)
// SC: O(1)