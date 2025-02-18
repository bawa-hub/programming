// https://practice.geeksforgeeks.org/problems/kth-smallest-element5635/1
// https://takeuforward.org/data-structure/kth-largest-smallest-element-in-an-array/

int kthSmallest(int arr[], int l, int r, int k) {
        int len = r-l+1;
        priority_queue<int> pq;
        for(int i=0;i<len;i++) {
            pq.push(arr[i]);
            if(pq.size() > k) pq.pop();
        }
        return pq.top();
    }

     int findKthLargest(vector<int>& nums, int k) {
        int len = nums.size();
        priority_queue<int> pq;
        for(int i=0;i<nums.size();i++) {
            pq.push(nums[i]);
            if(pq.size()>len - k + 1) pq.pop();
        }

        return pq.top();
    }