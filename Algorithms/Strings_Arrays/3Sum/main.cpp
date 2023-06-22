#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

void print(std::vector<vector<int>> &nums)
{
    for (int i = 0; i < nums.size(); i++)
    {
        std::cout << '[';

        for (int j = 0; j < nums[i].size(); j++)
        {
            std::cout << nums[i][j];
            if (j < nums[i].size() - 1)
                std::cout << ',';
        }
        std::cout << ']';
        if (i < nums.size() - 1)
            std::cout << ',';
    }
}

void print(std::vector<int> &nums)
{

    for (int j = 0; j < nums.size(); j++)
        std::cout << nums[j] << ' ';
    std::cout << '\n';
}

vector<vector<int>> threeSum(vector<int> &nums)
{
    std::sort(std::begin(nums), std::end(nums));

    const int n = nums.size();
    vector<vector<int>> result;
    if (n < 3)
        return result;

    int i = 0;
    int s = 0;
    while (nums[i] < 1 && n - i >= 3)
    {
        if (i > 0 && nums[i] == nums[i - 1])
        {
            i++;
            continue;
        }
        int sum = nums[i] * -1;
        for (int l = i + 1, r = n - 1; l < r;)
        {
            if (l > i + 1 && nums[l] == nums[l - 1])
            {
                l++;
                continue;
            }
            int l_r_sum = nums[l] + nums[r];
            if (l_r_sum == sum)
            {
                result.push_back({nums[i], nums[l], nums[r]});
                l++;
                r--;
            }
            else if (l_r_sum < sum)
                l++;
            else
                r--;
        }
        i++;
    }
    return result;
}

int main()
{
    vector<int> nums = {-2, 0, 0, 2, 2};

    vector<vector<int>> result = threeSum(nums);
    print(result);

    return 0;
}