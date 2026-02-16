// finding x^n

// Basic Method
int recursivePower(int x, int n)
{
    if (n == 0)
        return 1;
    return x * recursivePower(x, n - 1);
}

int iterativePower(int x, int n)
{
    int result = 1;
    while (n > 0)
    {
        result = result * x;
        n--;
    }
    return result;
}

// Time Complexity = O(n)
// this method is not suitable for very large number as 10^18

// Optimized method
/**
 * Binary Exponentiation
 * 
 * example --
 * 
 * 3^10 = (3^2)^5 = 9^5
 * 9^5 = 9.(9^2)^2 = 9.81^2 = 9.81.81 = 59049
 * 
 * Time Complexity = O(log N)
*/

int recursiveBinaryExponentiation(int x, int n)
{
    if (n == 0)
        return 1;
    else if (n % 2 == 0) //n is even
        return recursiveBinaryExponentiation(x * x, n / 2);
    else //n is odd
        return x * recursiveBinaryExponentiation(x * x, (n - 1) / 2);
}

int iterativeBinaryExponentiation(int x, int n)
{
    int result = 1;
    while (n > 0)
    {
        if (n % 2 == 1)
            result = result * x;
        x = x * x;
        n = n / 2;
    }
    return result;
}

/**
 * Storing answers  that are too large for their datatypes is an issue with above method
 * 
 * Must use modulus(%)
 * 
 * **/

int recursiveModularExponentiation(int x, int n, int M)
{
    if (n == 0)
        return 1;
    else if (n % 2 == 0) //n is even
        return recursiveModularExponentiation((x * x) % M, n / 2, M);
    else //n is odd
        return (x * recursiveModularExponentiation((x * x) % M, (n - 1) / 2, M)) % M;
}

int modularExponentiation(int x, int n, int M)
{
    int result = 1;
    while (n > 0)
    {
        if (n % 2 == 1)
            result = (result * x) % M;
        x = (x * x) % M;
        n = n / 2;
    }
    return result;
}

/**
 * Recursive solution analysis-
 * time complexity = O(log N)
 * memory complexity = O(log N)
 * 
 * Iterative solution analysis-
 * time complexity = O(log N)
 * memory complexity = O(1)
 * **/