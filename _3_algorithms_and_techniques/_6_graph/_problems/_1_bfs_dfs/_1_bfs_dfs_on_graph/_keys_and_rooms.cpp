// https://leetcode.com/problems/keys-and-rooms/

// dfs
class Solution
{
public:
    bool canVisitAllRooms(vector<vector<int>> &rooms)
    {
        int V = rooms.size();
        vector<int> vis(V, 0);
        int cnt = 0;
        dfs(0, rooms, vis, cnt);
        if (cnt == V)
            return true;
        return false;
    }

    void dfs(int src, vector<vector<int>> &rooms, vector<int> &vis, int &cnt)
    {
        vis[src] = 1;
        cnt++;
        for (auto child : rooms[src])
        {
            if (!vis[child])
                dfs(child, rooms, vis, cnt);
        }
    }
};