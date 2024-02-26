// https://www.geeksforgeeks.org/expand-the-string-according-to-the-given-conditions/

// C++ implementation of the approach
#include<bits/stdc++.h>
using namespace std;

// Function to expand and print the given string
void expandString(string strin)
{
	string temp = "";
	int j;

	for (int i = 0; i < strin.length(); i++) 
	{
		if (strin[i] >= 0) 
		{

			// Subtract '0' to convert char to int
			int num = strin[i] - '0';
			if (strin[i + 1] == '(') 
			{

				// Characters within brackets
				for (j = i + 1; strin[j] != ')'; j++)
				{
					if ((strin[j] >= 'a' && strin[j] <= 'z') || 
						(strin[j] >= 'A' && strin[j] <= 'Z'))
					{
						temp += strin[j];
					}
				}

				// Expanding
				for (int k = 1; k <= num; k++) 
				{
					cout << (temp);
				}

				// Reset the variables
				num = 0;
				temp = "";

				if (j < strin.length()) 
				{
					i = j;
				}
			}
		}
	}
}

// Driver code
int main()
{
	string strin = "3(ab)4(cd)";
	expandString(strin);
}
