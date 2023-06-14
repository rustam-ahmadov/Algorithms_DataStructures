#include <iostream>
#include <vector>

int maxProfit1(std::vector<int> &prices)
{
    int profit = 0;
    const int n = prices.size();

    int maxProfitPerCycle = 0;
    for (int i = 0; i < n - 1; i++)
    {
        for (int j = i + 1; j < n; j++)
            if (maxProfitPerCycle < prices[j] - prices[i])
                maxProfitPerCycle = prices[j] - prices[i];
        if (profit < maxProfitPerCycle)
            profit = maxProfitPerCycle;
    }
    return profit;
}

int maxProfit2(std::vector<int> &prices)
{
    int maxProfit = 0;
    const int n = prices.size();

    int maxProfitPerCycle = 0;
    for (int i = 0; i < n - 1; i++)
    {
        for (int j = i + 1; j < n; j++)
        {
            if (prices[j] <= prices[i])
                break;
            int profit = prices[j] - prices[i];
            maxProfitPerCycle = maxProfitPerCycle < profit ? profit : maxProfitPerCycle;
        }
        maxProfit = maxProfit < maxProfitPerCycle ? maxProfitPerCycle : maxProfit;
    }
    return maxProfit;
}
int main()
{
    std::vector<int> prices{7, 4, 11, 1, 3, 7};
    const int profit = maxProfit2(prices);
    std::cout << profit;
    return 0;
}