// https://leetcode.com/problems/first-bad-version/description/

// The API isBadVersion is defined for you.
// bool isBadVersion(int version);

class Solution {
public:
    int firstBadVersion(int n) {
  int lo = 0, hi = n;
  while (lo + 1 < hi) {
    int mid = lo + (hi - lo) / 2;
    if (isBadVersion(mid)) {
      hi = mid;
    } else {
      lo = mid;
    }
  }

  return hi;
}
};