

class Solution
{
public:
    int beautySum(string s)
    {
        int ans = 0;
        int n = s.size();
        for (int i = 0; i < n; i++)
        {
            int cnt[26] = {};
            int max_f = INT_MIN;
            int min_f = INT_MAX;
            for (int j = i; j < n; j++)
            {
                int ind = s[j] - 'a';
                cnt[ind]++;
                max_f = max(max_f, cnt[ind]);
                min_f = cnt[ind];
                for (int k = 0; k < 26; k++)
                {
                    if (cnt[k] >= 1)
                        min_f = min(min_f, cnt[k]);
                }
                ans += (max_f - min_f);
            }
        }
        return ans;
    }

    // using unordered_map
    int beautySum(string s)
    {
        int l, m;
        unordered_map<char, int> umap;
        int res = 0;
        for (int i = 0; i < s.length(); i++)
        {
            umap.clear();
            for (int j = i; j < s.length(); j++)
            {
                umap[s[j]]++;
                if (j - i > 1)
                {
                    m = 0;
                    l = INT_MAX;
                    for (auto i : umap)
                    {
                        m = max(i.second, m);
                        l = min(i.second, l);
                    }
                    res += m - l;
                }
            }
        }
        return res;
    }
};