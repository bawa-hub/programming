// https://leetcode.com/problems/kth-largest-element-in-an-array/
// https://www.geeksforgeeks.org/k-largestor-smallest-elements-in-an-array/
// https://takeuforward.org/data-structure/kth-largest-smallest-element-in-an-array/


class Solution {
public:
    int findKthLargest(vector<int>& arr, int k) {
        int left = 0, right = arr.size() - 1, kth;
    while (1) {
        int idx = partition(arr, left, right);
        if (idx == k - 1) {
            kth = arr[idx];
            break;
        }
        if (idx < k - 1) {
            left = idx + 1;
        } else {
            right = idx - 1;
        }
    }
    return kth;
    }


int partition(vector<int>& arr, int left, int right) {
    int pivot = arr[left] ;
    int l = left + 1 ;
    int r = right;
    while (l <= r) {
        if (arr[l] < pivot && arr[r] > pivot) {
            swap(arr[l], arr[r]);
            l++ ;
            r-- ;
        }
        if (arr[l] >= pivot) {
            l++;
        }
        if (arr[r] <= pivot) {
            r--;
        }
    }
    swap(arr[left], arr[r]);
    return r;
}
};