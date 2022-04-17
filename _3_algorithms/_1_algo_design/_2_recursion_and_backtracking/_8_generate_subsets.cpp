
// https://www.youtube.com/watch?v=u0e29JIdxZU&list=PLauivoElc3gjpEVTdncOKYN8fAiMm9a5g&index=4
#include <iostream>
#include <vector>
using namespace std;

// total subsets of n elements = 2^n

vector<vector<int>> subsets;

void generate(vector<int> &subset, int i, vector<int> nums)
{
    if (i == nums.size())
    {
        subsets.push_back(subset);
        return;
    }

    // ith element not in subset
    generate(subset, i + 1, nums);

    // ith element in subset
    subset.push_back(nums[i]);
    generate(subset, i + 1, nums);
    subset.pop_back();
}

int main()
{
    int n;
    cin >> n;
    vector<int> nums(n);
    for (int i = 0; i < n; i++)
    {
        cin >> nums[i];
    }
    vector<int> empty;
    generate(empty, 0, nums);
    for (auto subset : subsets)
    {
        cout << "[ ";
        for (auto ele : subset)
        {
            cout << ele << " ";
        }
        cout << "]";
        cout << endl;
    }
}
