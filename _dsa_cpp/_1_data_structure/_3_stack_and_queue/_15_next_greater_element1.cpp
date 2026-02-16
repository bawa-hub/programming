// https://leetcode.com/problems/next-greater-element-i/

// brute force
// using nested for loop

class Solution
{
public:
    vector<int> nextGreaterElement(vector<int> &nums1, vector<int> &nums2)
    {
        unordered_map<int, int> mpx(nums2.size());
        stack<int> st;

        for (int i = nums2.size() - 1; i >= 0; i--)
        {
            while (!st.empty() && st.top() <= nums2[i])
            {
                st.pop();
            }

            if (!st.empty())
            {
                mpx[nums2[i]] = st.top();
            }
            else
            {
                mpx[nums2[i]] = -1;
            }

            st.push(nums2[i]);
        }
        vector<int> ans;
        for (int i = 0; i < nums1.size(); i++)
            ans.push_back(mpx[nums1[i]]);

        return ans;
    }
};
// Time Complexity: O(N)
// Space Complexity: O(N)

int main() {
  Solution obj;
  vector < int > v {5,7,1,2,6,0};
  vector < int > res = obj.nextGreaterElements(v);
  cout << "The next greater elements are" << endl;
  for (int i = 0; i < res.size(); i++) {
    cout << res[i] << " ";
  }
}
