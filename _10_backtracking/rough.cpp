#include <iostream>

using namespace std;


void permutation(int nums[], vector<int> &ds, int n, vector<vector<int>> &res) {
    if(ds.size() == n) res.push_back(ds);
   for(int i=0;i<n;i++) {

   }
}
int main() {
    int nums[] = {1,2,3};
    int n = sizeof(nums)/sizeof(nums[0]);
   vector<vector<int>> res;
    vector<int> ds;
    permutation(nums, ds, n, res);
}