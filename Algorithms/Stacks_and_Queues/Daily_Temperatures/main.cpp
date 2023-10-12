#include <iostream>
#include <vector>
#include <stack>

using namespace std;

vector<int> dailyTemperatures(vector<int> &temperatures)
{

    const int n = temperatures.size();
    if (n == 1)
        return {0};

    vector<int> ans(n);
    stack<pair<int, int>> s;
    for (int i = 0; i < n - 1; i++)
    {
        if (temperatures[i] < temperatures[i + 1])
        {
            ans[i] = 1;
            while (!s.empty() && temperatures[i + 1] > s.top().first)
            {
                pair p = s.top();
                ans[p.second] = i + 1 - p.second;
                s.pop();
            }
        }
        else
            s.push({temperatures[i], i});
    }
    ans[n - 1] = 0;
    return ans;
}

int main()
{

    vector<int> temp = {73, 74, 75, 71, 69, 72, 76, 73};
    temp = dailyTemperatures(temp);
    for (int i = 0; i < temp.size(); i++)
        std::cout << temp[i] << ' ';
    std::cout << "END";
    return 0;
}