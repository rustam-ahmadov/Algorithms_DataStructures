#include <iostream>
#include <vector>


//O(n)
std::vector<int> runningSum(std::vector<int> &nums)
{
    int n = nums.size();
    std::vector<int> running_sum;
    int sum = 0;
    for (int i = 0; i < n; i++)
    {
        sum+=nums[i];
        running_sum.push_back(sum);
    }
    return running_sum;
}

int main()
{

    return 0;
}