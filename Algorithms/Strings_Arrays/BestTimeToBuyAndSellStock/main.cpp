#include <iostream>
#include <vector>

// O(n^2)
int maxProfit(std::vector<int> &prices)
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

// O(n)
int maxProfitN(std::vector<int> &prices)
{
    const int n = prices.size();

    int bestDay = 0;
    int profit = 0;
    int maxProfit = 0;
    for (int i = 1; i < n; i++)
    {
        if (prices[bestDay] > prices[i])
            bestDay = i;
        else
            profit = prices[i] - prices[bestDay];

        if (maxProfit < profit)
            maxProfit = profit;
    }
    return maxProfit;
}

int main()
{
    std::vector<int> prices{7, 4, 11, 1, 3, 7};
    const int profit = maxProfitN(prices);
    std::cout << profit;
    return 0;
}