#include <iostream>
#include <vector>
void print(std::vector<int> &arr)
{
    for (int i = 0; i < arr.size(); i++)
        std::cout << arr[i] << ' ';
    std::cout << '\n';
}

int maxProduct(std::vector<int> &nums)
{
    int n = nums.size();
    if (n == 1)
        return nums[0];

    // -2, 3, 4, 5
    int max = INT_MIN, curMax = 1;
    for (int i = 0; i < n; i++)
    {
        curMax *= nums[i];
        max = max < curMax ? curMax : max;
        if (!curMax)
            curMax = 1;
        std::cout << "Max: " << max << '\n';
    }

    curMax = 1;

    for (int i = n - 1; i >= 0; i--)
    {
        curMax *= nums[i];
        max = max < curMax ? curMax : max;
        if (!curMax)
            curMax = 1;
    }
    return max;
}

int main()
{
    std::vector<int> nums{-2, 3, -4};
    int max = maxProduct(nums);

    return 0;
}