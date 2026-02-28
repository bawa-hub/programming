// https://practice.geeksforgeeks.org/problems/sum-of-elements-between-k1th-and-k2th-smallest-elements3133/1


class Solution{
    public:
    long long sumBetweenTwoKth( long long A[], long long N, long long K1, long long K2)
    {
        long long k1th = kthSmallest(A, N, K1);
        long long k2th = kthSmallest(A, N, K2);
        
        long long sum = 0;
        for(int i=0;i<N;i++) {
            if(A[i]>k1th && A[i]<k2th) sum+=A[i];
        }
        
        return sum;
    }
    
    long long kthSmallest(long long A[], long long N, long long k) {
        priority_queue<long long> pq;
        for(int i=0;i<N;i++) {
            pq.push(A[i]);
            if(pq.size()>k) pq.pop();
        }
        return pq.top();
    }
};