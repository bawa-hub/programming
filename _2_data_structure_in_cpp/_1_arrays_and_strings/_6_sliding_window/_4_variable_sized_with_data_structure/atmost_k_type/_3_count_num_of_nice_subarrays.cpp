// https://leetcode.com/problems/count-number-of-nice-subarrays/

    int numberOfSubarrays(vector<int>& A, int k) {
        return atMost(A, k) - atMost(A, k - 1);
    }


    int atMost(vector<int>& A, int k) {
        int i = 0,j=0,res=0,n = A.size();

        while(j<n) {
            if(A[j]%2!=0) k--;

            while(k<0) {
                if(A[i]%2!=0) k++;
                i++;
            }

            res+=(j-i+1);
            j++;
        }
        return res;
    }

//     Time O(N) for one pass
// Space O(1)

// three pointer
    int numberOfSubarrays(vector<int>& A, int k) {
        int res = 0, i = 0, count = 0, n = A.size();
        for (int j = 0; j < n; j++) {
            if (A[j] & 1)
                --k, count = 0;
            while (k == 0)
                k += A[i++] & 1, ++count;
            res += count;
        }
        return res;
    }
//     Time O(N) for one pass
// Space O(1)