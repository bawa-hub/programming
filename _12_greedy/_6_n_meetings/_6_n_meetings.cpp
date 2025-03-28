// Q. https://practice.geeksforgeeks.org/problems/n-meetings-in-one-room-1587115620/1

// There is one meeting room in a firm.
// There are N meetings in the form of (start[i], end[i]) where start[i] is start time of meeting i and end[i] is finish time of meeting i.
// What is the maximum number of meetings that can be accommodated in the meeting room when only one meeting can be held in the meeting room at a particular time?

// Note: Start time of one chosen meeting can't be equal to the end time of the other chosen meeting.

#include <bits/stdc++.h>
using namespace std;

struct meeting
{
    int start;
    int end;
    int pos;
};

class Solution
{
public:
    bool static comparator(struct meeting m1, meeting m2)
    {
        if (m1.end < m2.end)
            return true;
        else if (m1.end > m2.end)
            return false;
        else if (m1.pos < m2.pos)
            return true;
        return false;
    }
    void maxMeetings(int s[], int e[], int n)
    {
        struct meeting meet[n];
        for (int i = 0; i < n; i++)
        {
            meet[i].start = s[i], meet[i].end = e[i], meet[i].pos = i + 1;
        }

        sort(meet, meet + n, comparator);

        vector<int> answer;

        int limit = meet[0].end;
        answer.push_back(meet[0].pos);

        for (int i = 1; i < n; i++)
        {
            if (meet[i].start > limit)
            {
                limit = meet[i].end;
                answer.push_back(meet[i].pos);
            }
        }
        cout << "The order in which the meetings will be performed is " << endl;
        for (int i = 0; i < answer.size(); i++)
        {
            cout << answer[i] << " ";
        }
    }
};
int main()
{
    Solution obj;
    int n = 6;
    int start[] = {1, 3, 0, 5, 8, 5};
    int end[] = {2, 4, 5, 7, 9, 9};
    obj.maxMeetings(start, end, n);
    return 0;
}

// Time Complexity: O(n) to iterate through every position and insert them in a data structure. O(n log n)  to sort the data structure in ascending order of end time. O(n)  to iterate through the positions and check which meeting can be performed.
// Overall : O(n) +O(n log n) + O(n) ~O(n log n)
// Space Complexity: O(n)  since we used an additional data structure for storing the start time, end time, and meeting no.