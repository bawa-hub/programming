// https://leetcode.com/problems/all-paths-from-source-to-target/

class Solution
{
public:
    vector<vector<int>> allPathsSourceTarget(vector<vector<int>> &graph)
    {

        vector<vector<int>> ans;
        vector<int> curr;
        dfs(0, graph.size() - 1, graph, curr, ans);
        return ans;
    }

    void dfs(int src, int dest, vector<vector<int>> &graph, vector<int> curr, vector<vector<int>> &ans)
    {
        curr.push_back(src);
        if (src == dest)
        {
            ans.push_back(curr);
            return;
        }

        for (auto child : graph[src])
        {
            dfs(child, dest, graph, curr, ans);
        }
    }
};