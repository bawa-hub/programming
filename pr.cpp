class Solution
{
public:
    int maxProfit(vector<int> &prices)
    {
        int es = 0;
        int eb = -100000;
        int ps = 0;
        int pb = 0;
        for (int i = 0; i < prices.size(); i++)
        {
            pb = eb;
            eb = max(eb, ps - prices[i]);
            ps = es;
            es = max(es, pb + prices[i]);
        }
        return es;
    }
};