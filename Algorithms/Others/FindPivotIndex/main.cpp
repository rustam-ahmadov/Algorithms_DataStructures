#include <iostream>
#include <vector>


//O(n)
int pivotIndex(std::vector<int> &nums)
{
    int n = nums.size();

    int sum = 0;
    for (int i = 0; i < n; i++)
    {
        sum += nums[i];
    }

    int left_sum = 0;
    for (int i = 0; i < n; i++)
    {
        sum -= nums[i];
        if (left_sum == sum)
            return i;

        left_sum += nums[i];
        
    }
    return -1;
}

int main()
{

    std::vector<int> nums = {1, 7, 3, 6, 5, 6};
    std::cout << pivotIndex(nums);
    return 0;
}