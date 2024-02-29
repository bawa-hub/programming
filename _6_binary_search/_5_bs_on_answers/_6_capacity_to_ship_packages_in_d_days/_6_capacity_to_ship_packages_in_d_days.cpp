// https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/


// brute force
int findDays(vector<int> &weights, int cap) {
    int days = 1; //First day.
    int load = 0;
    int n = weights.size(); //size of array.
    for (int i = 0; i < n; i++) {
        if (load + weights[i] > cap) {
            days += 1; //move to next day
            load = weights[i]; //load the weight.
        }
        else {
            //load the weight on the same day.
            load += weights[i];
        }
    }
    return days;
}

int leastWeightCapacity(vector<int> &weights, int d) {
    //Find the maximum and the summation:
    int maxi = *max_element(weights.begin(), weights.end());
    int sum = accumulate(weights.begin(), weights.end(), 0);

    for (int i = maxi; i <= sum; i++) {
        if (findDays(weights, i) <= d) {
            return i;
        }
    }
    //dummy return statement:
    return -1;
}
// Time Complexity: O(N * (sum(weights[]) – max(weights[]) + 1)), where sum(weights[]) = summation of all the weights, max(weights[]) = maximum of all the weights, N = size of the weights array.
// Reason: We are using a loop from max(weights[]) to sum(weights[]) to check all possible weights. Inside the loop, we are calling findDays() function for each weight. Now, inside the findDays() function, we are using a loop that runs for N times.
// Space Complexity: O(1) as we are not using any extra space to solve this problem.

int main()
{
    vector<int> weights = {5, 4, 5, 2, 3, 4, 5, 6};
    int d = 5;
    int ans = leastWeightCapacity(weights, d);
    cout << "The minimum capacity should be: " << ans << "\n";
    return 0;
}

// binary search
class Solution
{
public:
    int findDays(vector<int> &weights, int cap)
    {
        int days = 1, load = 0;
        for (int i = 0; i < weights.size(); i++)
        {
            if (weights[i] + load > cap)
            {
                days += 1;
                load = weights[i];
            }
            else
                load += weights[i];
        }
        return days;
    }

    int shipWithinDays(vector<int> &weights, int days)
    {
        int l = *max_element(weights.begin(), weights.end());
        int r = accumulate(weights.begin(), weights.end(), 0);
        while (l <= r)
        {
            int mid = l + (r - l) / 2;
            if (findDays(weights, mid) <= days)
                r = mid - 1;
            else
                l = mid + 1;
        }
        return l;
    }
};



// Time Complexity: O(N * log(sum(weights[]) – max(weights[]) + 1)), where sum(weights[]) = summation of all the weights, max(weights[]) = maximum of all the weights, N = size of the weights array.
// Reason: We are applying binary search on the range [max(weights[]), sum(weights[])]. For every possible answer ‘mid’, we are calling findDays() function. Now, inside the findDays() function, we are using a loop that runs for N times.
// Space Complexity: O(1) as we are not using any extra space to solve this problem.